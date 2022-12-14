/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mchirico/envr/fixtures"

	"github.com/spf13/cobra"
)

// argoCmd represents the argo command
var argoCmd = &cobra.Command{
	Use:   "argo",
	Short: " print argo template",
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := fixtures.Read("../fixtures/argo/argo.yaml")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(r))
	},
}

func init() {
	rootCmd.AddCommand(argoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// argoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// argoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
