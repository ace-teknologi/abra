package cmd

import (
	"fmt"

	"github.com/aceteknologi/go-abn/abr"
	"github.com/spf13/cobra"
)

const (
	searchStringFlag = "find-abn"
)

var searchString string

var searchCmd = &cobra.Command{
	Use:   "find-abn",
	Short: "Finds an ABN in the ABR",
	Long:  `Finds an ABN in the ABR`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return search()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&searchString, searchStringFlag, "s", "", "A nine digit ABN for you to search")
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

	entity, err := client.SearchByABN(searchString, false)
	if err != nil {
		return err
	}

	fmt.Printf("Found Business Entity: %v\n", entity.Name())

	return nil
}
