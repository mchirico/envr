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


k logs envar-demo --follow

k run --rm -i --tty dev --image=mchirico/ubuntu:latest --restart=Never --pod-running-timeout=6m0s -- bash -il


kubectl cp file argo/pod:/root/file

argo \
--kubeconfig ~/.kube/config \
-n argo submit \
--serviceaccount cicd \
--node-field-selector='name=ip-10-0-0-1.ec2.internal' \
 -p message="goodbye world" \
 --watch https://raw.githubusercontent.com/argoproj/argo-workflows/master/examples/global-parameters.yaml



# This will generate the secret
SECRET=$(kubectl get sa cicd -o=jsonpath='{.secrets[0].name}')
ARGO_TOKEN="Bearer $(kubectl get secret $SECRET -o=jsonpath='{.data.token}' | base64 --decode)"
echo $ARGO_TOKEN


aws ecr describe-repositories

ref: https://jmespath.org/tutorial.html

aws ecr describe-repositories \
--query 'repositories[*].[repositoryUri,repositoryName]'


aws ecr describe-repositories \
--query 'repositories[?repositoryName=='\''repodev'\''].[repositoryUri,repositoryName]'


kubectl create secret docker-registry regcred \
  --docker-server=<aws-account-id>.dkr.ecr.<aws-region>.amazonaws.com \
  --docker-username=AWS \
  --docker-password=$(aws ecr get-login-password) \
  -o yaml




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
