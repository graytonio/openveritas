package cmd

import (
	"github.com/graytonio/openveritas/cli/api"
	"github.com/spf13/cobra"
)

func init() {
	getPropCmd.Flags().BoolVarP(&detailed, "details", "d", false, "Show more prop details")
	rootCmd.AddCommand(getPropCmd)
}

var getPropCmd = &cobra.Command{
	Use:   "get-prop <node-name> [prop_name]",
	Short: "Get Properties of A Node",
	Args:  cobra.RangeArgs(1, 2),
	RunE:  getPropCmdRun,
}

func getPropCmdRun(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 1:
		node_name := args[0]
		return getAllPropertiesOfNode(node_name)
	default:
		node_name := args[0]
		prop_name := args[1]
		return getNodeProperty(node_name, prop_name)
	}
}

func getAllPropertiesOfNode(node_name string) error {
	props, err := api.GetAllPropertiesOfNode(config.Host, node_name)
	if err != nil {
		return err
	}

	printPropArray(props, config.Basic, detailed)
	return nil
}

func getNodeProperty(node_name string, prop_name string) error {
	prop, err := api.GetPropertyOfNodeByName(config.Host, node_name, prop_name)
	if err != nil {
		return err
	}

	printProp(prop, config.Basic, detailed)
	return nil
}
