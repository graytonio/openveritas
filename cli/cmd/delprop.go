package cmd

import (
	"fmt"

	"github.com/graytonio/openveritas/cli/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(delPropCmd)
}

var delPropCmd = &cobra.Command{
	Use:   "del-prop <node-name> <prop-name>",
	Short: "Delete Property",
	Args:  cobra.ExactArgs(2),
	RunE:  delPropCmdRun,
}

func delPropCmdRun(cmd *cobra.Command, args []string) error {
	node_name := args[0]
	prop_name := args[1]
	_, err := api.DeleteProp(config.Host, node_name, prop_name)
	if err != nil {
		return err
	}

	fmt.Printf("Property %s Deleted From Node %s", prop_name, node_name)
	return nil
}
