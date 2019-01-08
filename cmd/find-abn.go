package cmd

import (
	"fmt"

	"github.com/ace-technologi/go-abn/abr"
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

	fmt.Printf("Found Business Entity from ABN: %s\n\n", findABNString)
	fmt.Printf("  Name: %s\n", entity.Name())
	if len(entity.ASICNumber) > 0 {
		fmt.Printf("  ACN: %s\n", entity.ASICNumber)
	}
	fmt.Printf("  ABNs:\n")
	for _, abn := range entity.ABNs {
		fmt.Printf("  - %s\n", abn.String())
	}
	fmt.Printf("  Addresses:\n")
	for _, address := range entity.PhysicalAddresses {
		fmt.Printf("  - %s %s\n", address.Postcode, address.StateCode)
	}
	fmt.Printf("  Type: %s (%s)\n", entity.EntityType.EntityTypeCode, entity.EntityType.EntityDescription)
	fmt.Printf("  Statuses:\n")
	for _, es := range entity.EntityStatuses {
		fmt.Printf("  - %s (%s-%s)\n", es.EntityStatusCode, es.EffectiveFrom, es.EffectiveTo)
	}
	fmt.Printf("  Also known as:\n")
	for _, name := range entity.BusinessNames {
		fmt.Printf("  - %s\n", name.OrganisationName)
	}
	for _, name := range entity.HumanNames {
		fmt.Printf("  - %s %s %s\n", name.GivenName, name.OtherGivenName, name.FamilyName)
	}
	for _, name := range entity.MainNames {
		fmt.Printf("  - %s\n", name.OrganisationName)
	}
	for _, name := range entity.MainTradingNames {
		fmt.Printf("  - %s\n", name.OrganisationName)
	}
	for _, name := range entity.OtherTradingNames {
		fmt.Printf("  - %s\n", name.OrganisationName)
	}
	fmt.Printf("\n")

	return nil
}
