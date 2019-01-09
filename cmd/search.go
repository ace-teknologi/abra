package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ace-teknologi/go-abn/abr"
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
	searchByNameCmd.Flags().StringVarP(&outputFormat, outputFormatFlag, "f", "", "Output format: json, xml, text")
	searchByNameCmd.Flags().StringVarP(&outputFormatTextTemplatePath, outputFormatTextTemplatePathFlag, "t", "", "Path to text output template")
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

	outputFormat, err := setOutputType(outputFormat)
	if err != nil {
		return err
	}

	if outputFormat == outputTypeTEXT {
		fmt.Printf("Found %d Business Entities:\n", searchResults.NumberOfRecords)

		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		t, err := template.ParseFiles(filepath.Join(cwd, "./cmd/templates/search.txt.gtpl"))
		if err != nil {
			return err
		}

		for i, entity := range searchResults.SearchResultsRecord {
			fmt.Printf("%d.\n", i+1)
			err = t.Execute(os.Stdout, entity)
			if err != nil {
				return err
			}
		}
	} else if outputFormat == outputTypeJSON {
		b, err := json.Marshal(searchResults)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	} else if outputFormat == outputTypeXML {
		b, err := xml.Marshal(searchResults)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	}

	return nil
}
