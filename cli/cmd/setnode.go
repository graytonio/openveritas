package cmd

import (
	"fmt"
	"net/http"

	"github.com/graytonio/openveritas/cli/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setNodeCmd)
}

var setNodeCmd = &cobra.Command{
	Use:   "set-node <node-name> [<new-name>]",
	Short: "Set a nodes name",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  setNodeCmdRun,
}

func setNodeCmdRun(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 1:
		node_name := args[0]
		return createNode(node_name)
	default:
		node_name := args[0]
		new_name := args[1]
		return updateNode(node_name, new_name)
	}
}

func createNode(node_name string) error {
	status_code, err := api.PutNode(config.Host, node_name, node_name)
	if err != nil {
		return err
	}

	switch status_code {
	case http.StatusCreated:
		fmt.Printf("Node %s Created\n", node_name)
	case http.StatusOK:
		fmt.Printf("Node %s Already Exists\n", node_name)
	}
	return nil
}

func updateNode(node_name string, new_name string) error {
	status_code, err := api.PutNode(config.Host, node_name, new_name)
	if err != nil {
		return err
	}

	switch status_code {
	case http.StatusCreated:
		fmt.Printf("Node %s Created\n", node_name)
	case http.StatusOK:
		fmt.Printf("Node %s Updated to %s\n", node_name, new_name)
	}
	return nil
}
