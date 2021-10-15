package cmd

import (
	"github.com/aahemm/container-engine/pkg/container"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a simple container",
	Long:  "Runs containers using golang and kernel",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		volume, _ := cmd.Flags().GetString("volume")
		container.RunContainer(name, volume, args)
	},
}

var runcCmd = &cobra.Command{
	Use:   "runc",
	Short: "Runs a process to run container",
	Long:  "Runs containers using golang and kernel",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		volume, _ := cmd.Flags().GetString("volume")
		container.RunInsideContainer(name, volume, args)
	},
}

func init() {
	runCmd.PersistentFlags().StringP("volume", "v", "", "Set a path for volume")
	runCmd.PersistentFlags().StringP("name", "n", "", "Set a name for container")

	runcCmd.PersistentFlags().StringP("volume", "v", "", "Set a path for volume")
	runcCmd.PersistentFlags().StringP("name", "n", "", "Set a name for container")
}
