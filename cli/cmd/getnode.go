package cmd

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/graytonio/openveritas/cli/api"
	"github.com/spf13/cobra"
)

func init() {
	getNodeCmd.Flags().BoolVarP(&detailed, "details", "d", false, "Show more node details")
	rootCmd.AddCommand(getNodeCmd)
}

var getNodeCmd = &cobra.Command{
	Use:   "get-node [<node-name>]",
	Short: "Get Node Details",
	Args:  cobra.RangeArgs(0, 1),
	RunE:  getNodeCmdRun,
}

func getNodeCmdRun(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 1:
		node_name := args[0]
		return getNodeDetails(node_name)
	default:
		return getNodeList()
	}
}

func getNodeList() error {
	nodes, err := api.GetAllNodes(config.Host)
	if err != nil {
		if err.Code == http.StatusNotFound {
			fmt.Println(err.Message)
			return nil
		}
		return errors.New(err.Message)
	}

	printNodeArray(nodes, detailed, config.Basic)
	return nil
}

func getNodeDetails(node_name string) error {
	node, err := api.GetNodeByName(config.Host, node_name)
	if err != nil {
		if err.Code == http.StatusNotFound {
			fmt.Println(err.Message)
			return nil
		}
		return errors.New(err.Message)
	}

	printNode(node, detailed, config.Basic)
	return nil
}
