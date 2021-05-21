package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"gcurl/client"

	"github.com/spf13/cobra"
)

var headers string

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "http get request",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var req client.Request
		var headersM map[string]string

		if headers != "" {
			err := json.Unmarshal([]byte(headers), &headersM)
			if err != nil {
				return errors.New(fmt.Sprintf("unacceptable header format: %v", err))
			}
			req.Headers = &headersM
		}

		req.Method = "GET"
		req.URL = args[0]
		client.Do(req)

		return nil
	},
}

func init() {
	getCmd.Flags().StringVarP(&headers, "headers", "H", "", "headers")
	rootCmd.AddCommand(getCmd)
}
