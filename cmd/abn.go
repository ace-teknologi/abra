package cmd

import (
	"fmt"

	"../abr"
	"github.com/spf13/cobra"
)

const (
	findABNStringFlag = "find-abn"
)

var findABNString string

var findABNCmd = &cobra.Command{
	Use:   "find-abn",
	Short: "Finds an ABN in the ABR",
	Long:  `Finds an ABN in the ABR`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return findABN()
	},
}

func init() {
	rootCmd.AddCommand(findABNCmd)
	findABNCmd.Flags().StringVarP(&findABNString, findABNStringFlag, "s", "", "A nine digit ABN for you to search")
}

func findABN() error {
	// Ensure we have a GUID
	err := setGUID()
	if err != nil {
		return err
	}

	client, err := abr.NewWithGuid(GUID)
	if err != nil {
		return err
	}

	entity, err := client.SearchByABN(findABNString, false)
	if err != nil {
		return err
	}

	fmt.Printf("Found Business Entity: %v\n", entity.Name())

	return nil
}
