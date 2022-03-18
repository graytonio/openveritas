package cmd

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/graytonio/openveritas/cli/api"
	"github.com/spf13/cobra"
)

func init() {
	queryPropCmd.Flags().BoolVarP(&detailed, "details", "d", false, "Show more prop details")
	rootCmd.AddCommand(queryPropCmd)
}

var queryPropCmd = &cobra.Command{
	Use:   "query-prop <query-string>",
	Short: "Query Properties",
	Args:  cobra.ExactArgs(1),
	RunE:  queryPropCmdRun,
}

func queryPropCmdRun(cmd *cobra.Command, args []string) error {
	query_string := args[0]
	props, err := api.QueryPropertyByName(config.Host, query_string)
	if err != nil {
		if err.Code == http.StatusNotFound {
			fmt.Println(err.Message)
			return nil
		}
		return errors.New(err.Message)
	}

	printPropArray(props, config.Basic, detailed)
	return nil
}
