/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mchirico/envr/fixtures"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "argo dev template",
	Long: `

Can ssh into the argo pod for dev and testing.
This pulls from:

   mchirico/ubuntu:latest


kubectl exec --stdin --tty argopod -- /bin/bash

`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := fixtures.Read("../fixtures/argo/argoUbunto.yaml")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(r))
	},
}

func init() {
	rootCmd.AddCommand(devCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
