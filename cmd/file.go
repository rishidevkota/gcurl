package cmd

import (
	"encoding/json"
	"fmt"
	"gcurl/client"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file [file name]",
	Short: "it open json request file and do the request",
	Args:  cobra.MinimumNArgs(1),
	Example: `
	file contain example for get request
	{
		"method": "GET",
		"url": "https://jsonplaceholder.typicode.com/todos/1",
		"headers": {
			"accept": "application/json, text/plain, */*",
		}
	}

	file contain example for post request
	{
		"method": "POST",
		"url": "https://jsonplaceholder.typicode.com/posts",
		"headers": {
			"accept": "application/json, text/plain, */*",
		},
		"body": {
			"title": "foo",
			"body": "bar",
			"userId": 1
		}
	}
	`,
	RunE: func(cmd *cobra.Command, args []string) error {

		b, err := ioutil.ReadFile(args[0])
		if err != nil {
			return err
		}
		var req client.Request
		err = json.Unmarshal(b, &req)
		if err != nil {
			return fmt.Errorf("unacceptable json format: %v", err)
		}

		if err := client.Do(req); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
