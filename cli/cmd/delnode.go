package cmd

import (
	"errors"
	"fmt"

	"github.com/graytonio/openveritas/cli/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(delNodeCmd)
}

var delNodeCmd = &cobra.Command{
	Use:   "del-node <node-name>",
	Short: "Delete Node",
	Args:  cobra.ExactArgs(1),
	RunE:  delNodeCmdRun,
}

func delNodeCmdRun(cmd *cobra.Command, args []string) error {
	node_name := args[0]
	_, err := api.DeleteNode(config.Host, node_name)
	if err != nil {
		return errors.New(err.Message)
	}

	fmt.Printf("Node %s Deleted\n", node_name)
	return nil
}
