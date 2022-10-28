/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mchirico/envr/kustomize"
	"github.com/spf13/cobra"
)

// kustomizeCmd represents the kustomize command
var kustomizeCmd = &cobra.Command{
	Use:   "kustomize",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(`
envr kustomize <directory-to-tmp>

`)
			return
		}
		dir := args[0]
		path, err := kustomize.CreateKustomizeTmpDir(dir)
		if err != nil {
			return
		}
		fmt.Println(`
building kustomize setup in: ` + path + `\n`)

		err = kustomize.PopulateKustomizeTmpDir(dir)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(kustomizeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kustomizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kustomizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
