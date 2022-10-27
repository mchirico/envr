/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mchirico/envr/fixtures"

	"github.com/spf13/cobra"
)

// skaffoldCmd represents the skaffold command
var skaffoldCmd = &cobra.Command{
	Use:   "skaffold",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := fixtures.Read("../fixtures/argo/skaffold.yaml")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(r))
	},
}

func init() {
	rootCmd.AddCommand(skaffoldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// skaffoldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// skaffoldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
