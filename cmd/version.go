package cmd

import (
	"fmt"

	abra "github.com/ace-teknologi/abra/abra-lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of abra",
	Long:  `All software has versions. This is abra's`,
	RunE: func(cmd *cobra.Command, args []string) error {
		printVersion()
		return nil
	},
}

func printVersion() {
	fmt.Printf("Abra version %s", abra.Version)
}
