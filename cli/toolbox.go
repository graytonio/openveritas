package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

type PropertyReturn struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	NodeName string `json:"node_name"`
	PropertyName string `json:"property_name"`
	PropertyValue interface{} `json:"property_value"`
}

func propertyToString(prop PropertyReturn) string {
	return fmt.Sprintf("%s:%s\t%v\n", prop.NodeName, prop.PropertyName, prop.PropertyValue)
}

func parseStringToJSONString(str string) string {
	i, err := strconv.Atoi(str)
	if err == nil {
		return  fmt.Sprintf("%d", i)
	}
	f, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return fmt.Sprintf("%f", f)
	}
	b, err := strconv.ParseBool(str)
	if err == nil {
		return fmt.Sprintf("%t", b)
	}
	return fmt.Sprintf(`"%s"`, str)
}

func printHelp(command string) {
	switch command {
	case "get":
		fmt.Println("Usage: veritas get <node_name> [prop_name]")
	case "set":
		fmt.Println("Usage: veritas set <node_name> <prop_name> <prop_value>")
	case "new":
		fmt.Println("Usage veritas new <node_name>")
	default:
		fmt.Println("Usage: veritas <command> [options]\n")
		bold("Commands: ")
		fmt.Println("\tget\tget the properties of a node")
		fmt.Println("\tset\tset the property of a node")
		fmt.Println("\tnew\tcreate a new node\n")
	}
}

func parseBody(resp *http.Response, ptr interface{}) error {
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, ptr)
	if err != nil {
		return err
	}
	return nil
}

var bold = color.New(color.Bold, color.FgWhite).PrintlnFunc()