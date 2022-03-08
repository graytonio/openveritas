package cmd

import (
	"fmt"
	"net/http"

	"github.com/graytonio/openveritas/cli/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setPropCmd)
}

var setPropCmd = &cobra.Command{
	Use:   "set-prop <node-name> <prop-name> <prop-value>",
	Short: "Set a property to a value",
	Args:  cobra.ExactArgs(3),
	RunE:  setPropRun,
}

func setPropRun(cmd *cobra.Command, args []string) error {
	node_name := args[0]
	prop_name := args[1]
	prop_value := args[2]

	status_code, err := api.PutProp(config.Host, node_name, prop_name, prop_value)
	if err != nil {
		return err
	}

	switch status_code {
	case http.StatusOK:
		fmt.Printf("%s Updated to %s on %s", prop_name, prop_value, node_name)
	case http.StatusCreated:
		fmt.Printf("%s Set to %s on %s", prop_name, prop_value, node_name)
	}
	return nil
}
