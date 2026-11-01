package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of pranav-pyq",
	Long:  `Displays version of PRANAV PYQ FETCHER AMRITA installed on the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("PRANAV PYQ FETCHER AMRITA v2.0.0")
	},
}
