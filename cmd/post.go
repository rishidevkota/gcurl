package cmd

import (
	"encoding/json"
	"fmt"
	"gcurl/client"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post [url]",
	Short: "http post request",
	Args:  cobra.MinimumNArgs(1),
	Example: `
	gcurl post -H '{"headers": {"accept": "application/json"}}' \
	-B '{"body": {"title": "foo", "body": "bar", "userId": 1}}' \
	https://jsonplaceholder.typicode.com/posts
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var req client.Request
		var headersM map[string]string
		var bodyM map[string]interface{}

		if headers != "" {
			err := json.Unmarshal([]byte(headers), &headersM)
			if err != nil {
				return fmt.Errorf("unacceptable header format: %v", err)
			}
			req.Headers = &headersM
		}

		if body != "" {
			err := json.Unmarshal([]byte(headers), &bodyM)
			if err != nil {
				return fmt.Errorf("unacceptable body format: %v", err)
			}
			req.Body = &bodyM
		}

		req.Method = "POST"
		req.URL = args[0]
		if err := client.Do(req); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	postCmd.Flags().StringVarP(&headers, "headers", "H", "", "headers")
	postCmd.Flags().StringVarP(&body, "body", "B", "", "body")
	rootCmd.AddCommand(postCmd)
}
