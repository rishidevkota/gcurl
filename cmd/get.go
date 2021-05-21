package cmd

import (
	"encoding/json"
	"fmt"
	"gcurl/client"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "http get request",
	Args:  cobra.MinimumNArgs(1),
	Example: `
	gcurl get -H '{"headers": {"accept": "application/json"}}' https://jsonplaceholder.typicode.com/todos/1 
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var req client.Request
		var headersM map[string]string

		if headers != "" {
			err := json.Unmarshal([]byte(headers), &headersM)
			if err != nil {
				return fmt.Errorf("unacceptable header format: %v", err)
			}
			req.Headers = &headersM
		}

		req.Method = "GET"
		req.URL = args[0]
		if err := client.Do(req); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	getCmd.Flags().StringVarP(&headers, "headers", "H", "", "headers")
	rootCmd.AddCommand(getCmd)
}
