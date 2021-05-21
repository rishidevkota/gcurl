package cmd

import (
	"encoding/json"
	"gcurl/client"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "it open json request file and do the request",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		b, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}
		var req client.Request
		err = json.Unmarshal(b, &req)
		if err != nil {
			log.Fatal(err)
		}

		client.Do(req)
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
