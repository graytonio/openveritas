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
	NodeName string `arg:"" help:"Name of node to query"`
	PropertyName string `arg:"" help:"Name of property to lookup"`
}

type SetPropertyCMD struct {
	NodeName string `arg:"" help:"Name of node to query"`
	PropertyName string `arg:"" help:"Name of property to set"`
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
	NodeName string `arg:"" help:"Name of node to update"`
	NewName string `arg:"" help:"New name for node"`
}

type DeleteNodeCMD struct {
	NodeName string `arg:"" help:"Name of node to delete"`
}

type NewNodeCMD struct {
	NodeName string `arg:"" help:"Name of node to create"`
}

type debugFlag bool
type Context struct {
	Config string
	Logger *log.Logger
}