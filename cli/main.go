package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"gopkg.in/ini.v1"
)

var fs *flag.FlagSet
var url string = "http://localhost:9295"
var conf string

func init() {
	fs = flag.NewFlagSet("args", flag.ContinueOnError)
	fs.StringVar(&conf, "config", "~/.config/veritasrc", "Location of config file")
	err := fs.Parse(os.Args[1:])

	if err != nil {
		fmt.Printf("Error parsing arguments %v", err.Error())
		os.Exit(1)
	}

	loadConfig()
}

func loadConfig() {
	cfg, err := ini.Load(conf)
	if err != nil {
		fmt.Printf("Failed to read configuration %s", err.Error())
		os.Exit(1)
	}

	url = cfg.Section("server").Key("host").String()
}

func main() {
	args := fs.Args()

	if len(args) == 0 {
		printHelp("")
		os.Exit(1)
	}

	command := args[0]
	var err error
	switch command {
	case "get":
		err = veritasGet(args[1:])
	case "set":
		err = veritasSet(args[1:])
	case "new":
		err = veritasNew(args[1:])
	}

	if err != nil {
		fmt.Println(err.Error() + "\n")
		printHelp(command)
		os.Exit(1)
	}
}

func veritasNew(args []string) error {
	if len(args) == 0 {
		return errors.New("node_name is required")
	}
	node_name := args[0]

	var jsonData = []byte(fmt.Sprintf(`{ "name": "%s" }`, node_name))
	resp, err := http.Post(fmt.Sprintf("%s/node", url), "application/json", bytes.NewBuffer((jsonData)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("Created new node %s\n", node_name)
		return nil
	} else if resp.StatusCode == http.StatusConflict {
		fmt.Printf("Node %s already exists\n", node_name)
		return nil
	}

	return errors.New("Error occured while creating node " + node_name)
}

func veritasSet(args []string) error {
	if len(args) == 0 {
		return errors.New("node_name is required")
	} else if len(args) == 1 {
		return errors.New("prop_name is required")
	} else if len(args) == 2 {
		return errors.New("prop_value is required")
	}

	node_name := args[0]
	prop_name := args[1]
	prop_value := parseStringToJSONString(args[2])
	var jsonData = []byte(fmt.Sprintf(`{ "property_name": "%s", "property_value": %v }`, prop_name, prop_value))
	resp, err := http.Post(fmt.Sprintf("%s/node/%s/prop", url, node_name), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("Created property %s on node %s\n", prop_name, node_name)
		return nil
	}

	return errors.New("Error occured while creating property " + prop_name + " for node " + node_name)
}

func veritasGet(args []string) error {
	if len(args) == 0 {
		//TODO get list of nodes
		return errors.New("node_name is required")
	}
	node_name := args[0]

	// node_name and prop_name
	if len(args) == 2 {
		prop_name := args[1]

		//Request
		resp, err := http.Get(fmt.Sprintf("%s/node/%s/prop/%s", url, node_name, prop_name))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		//Check status code
		if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("Property %s not found for node %s\n", prop_name, node_name)
			return nil
		}

		//Parse Body
		var data PropertyReturn
		err = parseBody(resp, &data)
		if err != nil {
			return err
		}

		fmt.Printf("[0] %s", propertyToString(data))
	} else {
		//Request
		resp, err := http.Get(fmt.Sprintf("%s/node/%s/prop", url, node_name))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		//Check status code
		if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("Node %s not found", node_name)
			return nil
		}

		//Parse Body
		var data []PropertyReturn
		err = parseBody(resp, &data)
		if err != nil {
			return err
		}

		if len(data) == 0 {
			fmt.Printf("Node %s has no properties\n", node_name)
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		for i, s := range data {
			fmt.Fprintf(w, "[%d] %s", i, propertyToString(s))
		}
		w.Flush()
	}

	return nil
}