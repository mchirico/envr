/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "common commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
kubectl exec --stdin --tty shell-demo -- /bin/bash

kubectl cp file argo/pod:/root/file

argo --kubeconfig ~/.kube/config -n argo submit --serviceaccount cicd \
 -p message="goodbye world" \
 --watch https://raw.githubusercontent.com/argoproj/argo-workflows/master/examples/global-parameters.yaml


# This will generate the secret
SECRET=$(kubectl get sa cicd -o=jsonpath='{.secrets[0].name}')
ARGO_TOKEN="Bearer $(kubectl get secret $SECRET -o=jsonpath='{.data.token}' | base64 --decode)"
echo $ARGO_TOKEN

`)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
