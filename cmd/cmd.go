package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ace-teknologi/go-abn/abr"
	"github.com/spf13/cobra"
)

const (
	guidFlagName = "GUID"
)

var GUID string

var rootCmd = &cobra.Command{
	Use:   "goabn",
	Short: "Goabn looks up an ABN or ACN using the ABR",
	Long: `A command line interface to the Australian Business Register.
            More information available at https://github.com/ace-teknologi/go-abn`,
	RunE: func(cmd *cobra.Command, arg []string) error {
		fmt.Printf("[DEBUG] %v", arg)
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&GUID, guidFlagName, "g", "", "Your ABR GUID")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setGUID() error {
	// If GUID wasn't set as a flag, check for an ENV
	flag := rootCmd.Flags().Lookup(guidFlagName)
	if flag.Value.String() == "" {
		g, ok := os.LookupEnv(abr.GUIDEnvName)
		if ok {
			flag.Value.Set(g)
		} else {
			return errors.New(abr.MissingGUIDError)
		}
	}
	return nil
}
