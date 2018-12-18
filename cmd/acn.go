package cmd

import (
	"fmt"

	"../abr"
	"github.com/spf13/cobra"
)

const (
	findACNStringFlag = "find-acn"
)

var findACNString string

var findACNCmd = &cobra.Command{
	Use:   "find-acn",
	Short: "Finds an ACN in the ABR",
	Long:  `Finds an ACN in the ABR`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return findACN()
	},
}

func init() {
	rootCmd.AddCommand(findACNCmd)
	findACNCmd.Flags().StringVarP(&findACNString, findACNStringFlag, "s", "", "A nine digit ACN for you to search")
}

func findACN() error {
	// Ensure we have a GUID
	err := setGUID()
	if err != nil {
		return err
	}

	client, err := abr.NewWithGuid(GUID)
	if err != nil {
		return err
	}

	entity, err := client.SearchByACN(findACNString, false)
	if err != nil {
		return err
	}

	fmt.Printf("Found Business Entity: %v\n", entity.Name())

	return nil
}
