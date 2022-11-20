/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mchirico/envr/util"

	"github.com/spf13/cobra"
)

// swapallCmd represents the swapall command
var swapallCmd = &cobra.Command{
	Use:   "swapall",
	Short: "looks at file.yaml and swaps all {VARS} if environment ${VARS} exists",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println(`

			Usage: envr swapall ./fileTemplate.yaml ./target.yaml

            Where fileTemplate.yaml has {ENVARS} and target.yaml gets results
`)

			return
		}

		r, err := util.FindAll(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		e := util.NewEnv(r)
		err = e.Replace(args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(swapallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// swapallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// swapallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
