package cmd

import (
	"fmt"

	"../abr"
	"github.com/spf13/cobra"
)

const (
	searchByNameStringFlag = "search"
)

var searchByNameString string

var searchByNameCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches the ABR",
	Long:  `Searches the ABR`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return searchByName()
	},
}

func init() {
	rootCmd.AddCommand(searchByNameCmd)
	searchByNameCmd.Flags().StringVarP(&searchByNameString, searchByNameStringFlag, "s", "", "The name of the company to search for")
}

func searchByName() error {
	// Ensure we have a GUID
	err := setGUID()
	if err != nil {
		return err
	}

	client, err := abr.NewWithGuid(GUID)
	if err != nil {
		return err
	}

	searchResults, err := client.SearchByName(searchByNameString, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Found %d Business Entities:\n", searchResults.NumberOfRecords)
	for i, r := range searchResults.SearchResultsRecord {
		fmt.Printf("\n%d.\t%s %s\n", (i + 1), r.ABN.IdentifierValue, r.FriendlyName())
		fmt.Printf("\t%s %s\n", r.MainBusinessPhysicalAddress.Postcode, r.MainBusinessPhysicalAddress.StateCode)
		fmt.Printf("\t%d/100 %s\n", r.Score(), r.IsCurrentIndicator())
	}

	return nil
}
