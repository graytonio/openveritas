package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

func appendToHostString(host string, routes ...string) string {
	route := strings.Join(routes, "")
	u, _ := url.Parse(host)
	u.Path = path.Join(u.Path, route)
	return u.String()
}

func parseJSONBody(resp *http.Response, data interface{}) (error) {
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, data)
	if err != nil { return err }
	return nil
}

func parseStringToJSONString(str string) string {

	i, err := strconv.ParseInt(str, 0, 64)
	if err == nil {
		return fmt.Sprintf("%d", i)
	}

	b, err := strconv.ParseBool(str)
	if err == nil {
		return fmt.Sprintf("%t", b)
	}

	f, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return fmt.Sprintf("%f", f)
	}

	return fmt.Sprintf(`"%s"`, str)
}

func getNodeProperty(logger *log.Logger, string, node_name string, property_name string)(*Property, error) {
	path := appendToHostString(host, "/node/", node_name, "/prop/", property_name)
	logger.Printf("GET Request: %s", path)
	resp, err := http.Get(path)
	logger.Println(resp.StatusCode)
	
	if err != nil { 
		logger.Println(err)
		return nil, err 
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("property %s not found for node %s", property_name, node_name)
	}

	var data Property
	err = parseJSONBody(resp, &data)
	if err != nil { 
		logger.Println(err.Error())
		return nil, err 
	}
	return &data, nil
}

func setNodeProperty(logger *log.Logger, host string, node_name string, property_name string, property_value string) error {
	path := appendToHostString(host, "/node/", node_name, "/prop/", property_name)
	logger.Printf("POST Request: %s", path)
	formated_prop_value := parseStringToJSONString(property_value)
	jsonData := []byte(fmt.Sprintf(`{ "property_value": %v }`, formated_prop_value))
	logger.Printf("Payload: %s", string(jsonData))
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(jsonData))
	if err != nil { 
		logger.Println(err.Error())
		return err 
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	logger.Println(resp.StatusCode)
	if err != nil {
		logger.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("error creating property %s for node %s", property_name, node_name)
}

func getAllNodeProperties(logger *log.Logger, host string, node_name string) (*[]Property, error) {
	path := appendToHostString(host, "/node/", node_name, "/prop")
	logger.Printf("GET Request: %s", path)
	resp, err := http.Get(path)
	logger.Println(resp.StatusCode)

	if err != nil { 
		logger.Println(err.Error())
		return nil, err 
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("node %s not found", node_name)
	}

	var data []Property
	err = parseJSONBody(resp, &data)
	if err != nil { 
		logger.Println(err)
		return nil, err 
	}
	return &data, nil
}

// TODO Get all nodes with property
func getAllPropertyNodes(host string, prop_name string) (*[]Property, error) {
	return nil, nil
}

func getAllNodes(logger *log.Logger, host string) (*[]Node, error) {
	path := appendToHostString(host, "/node")
	logger.Printf("GET Request: %s", path)
	resp, err := http.Get(path)
	logger.Println(resp.StatusCode)
	if err != nil { 
		logger.Println(err.Error())
		return nil, err 
	}
	defer resp.Body.Close()

	var data []Node
	err = parseJSONBody(resp, &data)
	if err != nil { 
		logger.Println(err.Error())
		return nil, err 
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no nodes found")
	}

	return &data, nil
}

func updateNodeName(logger *log.Logger, host string, node_name string, new_name string) error {
	path := appendToHostString(host, "/node/", node_name)
	logger.Printf("PUT Request: %s", path)
	jsonData :=  []byte(fmt.Sprintf(`{ "name": "%s" }`, new_name))
	logger.Printf("Payload: %s", string(jsonData))
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(jsonData))
	if err != nil { 
		logger.Println(err)
		return err 
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	logger.Println(resp.StatusCode)
	if err != nil { 
		logger.Println(err)
		return err 
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("error updating node %s", node_name)
}

func createNode(host string, node_name string) error {
	path := appendToHostString(host, "/node")
	jsonData := []byte(fmt.Sprintf(`{ name: "%s" }`, node_name))
	resp, err := http.Post(path, "application/json", bytes.NewBuffer(jsonData))
	if err != nil { return err }

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	return fmt.Errorf("error creating node %s", node_name)
}