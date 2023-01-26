package cmd

import (
	"github.com/nyogjtrc/deciduous/ver"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		ver.Print()
	},
}
