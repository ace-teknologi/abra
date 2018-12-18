package cmd

import (
	"fmt"

	"github.com/ace-teknologi/go-abn/abr"
	"github.com/spf13/cobra"
)

const (
	searchStringFlag = "find-acn"
)

var searchString string

var searchCmd = &cobra.Command{
	Use:   "find-acn",
	Short: "Finds an ACN in the ABR",
	Long:  `Finds an ACN in the ABR`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return search()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&searchString, searchStringFlag, "s", "", "A nine digit ACN for you to search")
}

func search() error {
	// Ensure we have a GUID
	err := setGUID()
	if err != nil {
		return err
	}

	client, err := abr.NewWithGuid(GUID)
	if err != nil {
		return err
	}

	entity, err := client.SearchByACN(searchString, false)
	if err != nil {
		return err
	}

	fmt.Printf("Found Business Entity: %v\n", entity.Name())

	return nil
}
