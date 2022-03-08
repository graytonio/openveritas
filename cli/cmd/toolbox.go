package cmd

import (
	"fmt"
	"os"

	"github.com/graytonio/openveritas/cli/api"
	"github.com/jedib0t/go-pretty/table"
)

func propToString(prop *api.Property, detailed bool) string {
	if detailed {
		return fmt.Sprintf("%s:%s %v\n\tCreated: %s\n\tUpdate: %s", prop.NodeName, prop.PropertyName, prop.PropertyValue, prop.CreatedAt, prop.UpdatedAt)
	}

	return fmt.Sprintf("%s:%s %v", prop.NodeName, prop.PropertyName, prop.PropertyValue)
}

func propToTableRow(prop *api.Property, detailed bool) table.Row {
	if detailed {
		return table.Row{
			prop.NodeName,
			prop.PropertyName,
			prop.PropertyValue,
			prop.CreatedAt,
			prop.UpdatedAt,
		}
	}

	return table.Row{
		prop.NodeName,
		prop.PropertyName,
		prop.PropertyValue,
	}
}

func printProp(prop *api.Property, basic bool, detailed bool) {
	if basic {
		fmt.Println(propToString(prop, detailed))
		return
	}

	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(os.Stdout)

	if detailed {
		t.AppendHeader(table.Row{"Node", "Property", "Value", "Created", "Updated"})
	} else {
		t.AppendHeader(table.Row{"Node", "Property", "Value"})
	}
	t.AppendRow(propToTableRow(prop, detailed))
	t.Render()
}

func printPropArray(props []api.Property, basic bool, detailed bool) {
	if basic {
		for _, p := range props {
			fmt.Println(propToString(&p, detailed))
		}
		return
	}

	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(os.Stdout)
	if detailed {
		t.AppendHeader(table.Row{"Node", "Property", "Value", "Created", "Updated"})
	} else {
		t.AppendHeader(table.Row{"Node", "Property", "Value"})
	}
	for _, p := range props {
		t.AppendRow(propToTableRow(&p, detailed))
	}
	t.Render()
}

func nodeToString(node *api.Node, detailed bool) string {
	if detailed {
		return fmt.Sprintf("%s\n\tCreated: %s\n\tUpdated: %s", node.NodeName, node.CreatedAt.String(), node.UpdatedAt.String())
	}
	return fmt.Sprintf("%s", node.NodeName)
}

func nodeToTableRow(node *api.Node, detailed bool) table.Row {
	if detailed {
		return table.Row{node.NodeName, node.CreatedAt, node.UpdatedAt}
	}
	return table.Row{node.NodeName}
}

func printNode(node *api.Node, detailed bool, basic bool) {
	if basic {
		fmt.Println(nodeToString(node, detailed))
		return
	}
	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(os.Stdout)

	if detailed {
		t.AppendHeader(table.Row{"Name", "Created", "Updated"})
	} else {
		t.AppendHeader(table.Row{"Name"})
	}
	t.AppendRow(nodeToTableRow(node, detailed))
	t.Render()
}

func printNodeArray(nodes []api.Node, detailed bool, basic bool) {
	if basic {
		for _, node := range nodes {
			fmt.Println(nodeToString(&node, detailed))
		}
		return
	}
	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	t.SetOutputMirror(os.Stdout)
	if detailed {
		t.AppendHeader(table.Row{"Name", "Created", "Updated"})
	} else {
		t.AppendHeader(table.Row{"Name"})
	}
	for _, node := range nodes {
		t.AppendRow(nodeToTableRow(&node, detailed))
	}
	t.Render()
}
