package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	abra "github.com/ace-teknologi/abra/abra-lib"
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
	findABNCmd.Flags().StringVarP(&outputFormat, outputFormatFlag, "f", "", "Output format: json, xml, text")
	findABNCmd.Flags().StringVarP(&outputFormatTextTemplatePath, outputFormatTextTemplatePathFlag, "t", "", "Path to text output template")
}

func findABN() error {
	// Ensure we have a GUID
	err := setGUID()
	if err != nil {
		return err
	}

	client, err := abra.NewWithGuid(GUID)
	if err != nil {
		return err
	}

	entity, err := client.SearchByABN(findABNString, false)
	if err != nil {
		return err
	}

	outputFormat, err := setOutputType(outputFormat)
	if err != nil {
		return err
	}

	switch outputFormat {
	case outputTypeTEXT:
		fmt.Printf("Found Business Entity from ABN: %s\n\n", findABNString)

		t, err := setOutputTypeTextTemplate("find", outputFormatTextTemplatePath)
		if err != nil {
			return err
		}

		err = t.Execute(os.Stdout, entity)
		if err != nil {
			return err
		}
	case outputTypeJSON:
		b, err := json.Marshal(entity)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	case outputTypeXML:
		b, err := xml.Marshal(entity)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	default:
		return ErrInvalidOutputTypeMessage
	}

	return nil
}
