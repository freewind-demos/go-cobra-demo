package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  `Show version of this application`,
	Run: func(cmd *cobra.Command, args []string) {
		println("v1.0.0")
	},
}
