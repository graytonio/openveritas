package main

import "log"

type Property struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	NodeName string `json:"node_name"`
	PropertyName string `json:"property_name"`
	PropertyValue interface{} `json:"property_value"`
}

type Node struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	NodeName string `json:"name"`
}

type GetPropertyCMD struct {
	PropertyName string `arg:"" help:"Name of property to lookup"`
	NodeName string `arg:"" optional:"" help:"Name of node to query"`
}

type SetPropertyCMD struct {
	PropertyName string `arg:"" help:"Name of property to set"`
	NodeName string `arg:"" help:"Name of node to query"`
	Value string `arg:"" help:"Value to set property to"`
}

type DeletePropertyCMD struct {
	NodeName string `arg:"" help:"Name of node"`
	PropertyName string `arg:"" help:"Name of property to delete"`
}

type GetNodeCMD struct {
	NodeName string `arg:"" optional:"" help:"Name of node to query"`
}

type SetNodeCMD struct {
	NodeName string `arg:"" help:"Name of node to set"`
	NewName string `arg:"" optional:"" help:"New name for node"`
}

type DeleteNodeCMD struct {
	NodeName string `arg:"" help:"Name of node to delete"`
}

type debugFlag bool
type Config struct {
	host string
}
type Context struct {
	Config Config
	Logger *log.Logger
}