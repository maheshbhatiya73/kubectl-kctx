package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubectl-kctx",
	Short: "Manage and reload your Kubernetes kubeconfig contexts",
	Long:  `kubectl-kctx is a kubectl plugin to manage, reload, and monitor kubeconfig contexts easily.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
