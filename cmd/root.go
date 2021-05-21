package cmd

import "github.com/spf13/cobra"

var body string
var headers string

var rootCmd = &cobra.Command{
	Use:   "gcurl",
	Short: "It is curl but with json support",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
