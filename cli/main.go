package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/alecthomas/kong"
)

var host string = "https://veritas.conheart.com"

func propToString(prop *Property) string {
	return fmt.Sprintf("%s:%s=%v", prop.NodeName, prop.PropertyName, prop.PropertyValue)
}

// Get A Property of a Node
func (r *GetPropertyCMD) Run(ctx *Context) error {
	prop, err := getNodeProperty(ctx.Logger, host, r.NodeName, r.PropertyName)
	if err != nil { return err }
	fmt.Println(propToString(prop))
	return nil
}

// TODO Set Property
func (r *SetPropertyCMD) Run(ctx *Context) error {
	err := setNodeProperty(ctx.Logger, host, r.NodeName, r.PropertyName, r.Value)
	if err != nil { return err }
	fmt.Printf("%s set to %v on node %s", r.PropertyName, r.Value, r.NodeName)
	return nil
}

// Get All Properties of a Node
func (r *GetNodeCMD) Run(ctx *Context) error {
	if r.NodeName != "" {
		props, err := getAllNodeProperties(ctx.Logger, host, r.NodeName)
		if err != nil { return err }
		for _, p := range *props {
			fmt.Println(propToString(&p))
		}	
	} else {
		nodes, err := getAllNodes(ctx.Logger, host)
		if err != nil { return err }
		for _, n := range *nodes {
			fmt.Println(n.NodeName)
		}
	}
	return nil
}

// Update Node Name
func (r *SetNodeCMD) Run(ctx *Context) error {
	err := updateNodeName(ctx.Logger, host, r.NodeName, r.NewName)
	if err != nil { return err }
	fmt.Printf("Updated node %s to %s\n", r.NodeName, r.NewName)
	return nil
}

// TODO Create Node
func (r *NewNodeCMD) Run(ctx *Context) error {
	fmt.Println(r)
	return nil
}

func (d debugFlag) BeforeApply(logger *log.Logger) error {
	logger.SetOutput(os.Stdout)
	return nil
}

var cli struct {
	Config string `help:"Path to config file"`
	Debug debugFlag `help:"Enable debug mode"`

	GetProperty GetPropertyCMD `cmd:"" help:"Get the value of a property on a node"`
	SetProperty SetPropertyCMD `cmd:"" help:"Set the value of a property on a node"`
	
	GetNode GetNodeCMD `cmd:"" help:"Get all properties of a node"`
	SetNode SetNodeCMD `cmd:"" help:"Rename Node"`
	NewNode NewNodeCMD `cmd:"" help:"Create new Node"`
}

func main() {
	logger := log.New(ioutil.Discard, "", log.LstdFlags)
	ctx := kong.Parse(&cli, kong.UsageOnError(), kong.Bind(logger))
	err := ctx.Run(&Context{ Config: cli.Config, Logger: logger })
	ctx.FatalIfErrorf(err)
}