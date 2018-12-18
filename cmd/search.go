package cmd

import (
	"fmt"

	"github.com/ace-teknologi/go-abn/abr"
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

	searchResults, err := client.SearchByName(searchString, false)
	if err != nil {
		return err
	}

	fmt.Printf("Found %d Business Entities:\n", searchResults.NumberOfRecords)
	for i, r := range searchResults.SearchResultsRecord {
		fmt.Printf("\n%d.\n", i)
		fmt.Printf("  %s %s\n", r.ABN, r.MainName.OrganisationName)
		fmt.Printf("  %s %s\n", r.MainBusinessPhysicalAddress.Postcode, r.MainBusinessPhysicalAddress.StateCode)
		fmt.Printf("  %d/100 %s\n\n", r.MainName.Score, r.MainName.IsCurrentIndicator)
	}

	return nil
}
