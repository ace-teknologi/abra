package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/ace-teknologi/go-abn/abr"
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
	findACNCmd.Flags().StringVarP(&outputFormat, outputFormatFlag, "f", "", "Output format: json, xml, text")
	findACNCmd.Flags().StringVarP(&outputFormatTextTemplatePath, outputFormatTextTemplatePathFlag, "t", "", "Path to text output template")
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

	outputFormat, err := setOutputType(outputFormat)
	if err != nil {
		return err
	}

	switch outputFormat {
	case outputTypeTEXT:
		fmt.Printf("Found Business Entity from ACN: %s\n\n", findACNString)

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
