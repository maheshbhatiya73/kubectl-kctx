package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload kubeconfig from local ~/.kube/config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Reloading kubeconfig from ~/.kube/config...")

		kubeconfig := os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = os.Getenv("HOME") + "/.kube/config"
		}

		_, err := os.Stat(kubeconfig)
		if os.IsNotExist(err) {
			fmt.Printf("Kubeconfig not found at %s\n", kubeconfig)
			return
		}

		cmdExec := exec.Command("kubectl", "config", "get-contexts")
		cmdExec.Stdout = cmd.OutOrStdout()
		cmdExec.Stderr = cmd.ErrOrStderr()

		if err := cmdExec.Run(); err != nil {
			fmt.Println("Failed to validate kubeconfig:", err)
			return
		}

		fmt.Println("Kubeconfig is available and loaded.")
	},
}

func init() {
	rootCmd.AddCommand(reloadCmd)
}
