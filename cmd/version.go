package cmd

import (
	"fmt"

	"github.com/aceteknologi/go-abn/abr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-abn",
	Long:  `All software has versions. This is go-abn's`,
	RunE: func(cmd *cobra.Command, args []string) error {
		printVersion()
		return nil
	},
}

func printVersion() {
	fmt.Printf("Go ABN version %s", abr.Version)
}
