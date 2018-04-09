package cmd

import (
	"fmt"

	"github.com/sjauld/go-abn/abr"
	"github.com/spf13/cobra"
)

const (
	searchStringFlag = "search"
)

var searchString string

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches the ABR",
	Long:  `Searches the ABR`,
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
