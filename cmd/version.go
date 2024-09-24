package cmd

import (
	"fmt"

	"github.com/nuomizi-fw/stargazer/constants"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Stargazer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stargazer version ", constants.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
