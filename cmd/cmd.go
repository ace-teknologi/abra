package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	abra "github.com/ace-teknologi/abra/abra-lib"
	"github.com/spf13/cobra"
)

const (
	guidFlagName                     = "GUID"
	outputFormatFlag                 = "output-format"
	outputFormatTextTemplatePathFlag = "text-output-template"
	outputTypeJSON                   = "json"
	outputTypeTEXT                   = "text"
	outputTypeXML                    = "xml"
	defaultFindTemplatePath          = "./cmd/templates/abn.txt.gtpl"
	defaultSearchTemplatePath        = "./cmd/templates/search.txt.gtpl"
)

var GUID string
var outputFormat string
var outputFormatTextTemplatePath string

// ErrInvalidOutputTypeMessage provides an error message if output type is not valid
var ErrInvalidOutputTypeMessage = fmt.Errorf("Invalid output type. Please choose from json, text, xml")

var rootCmd = &cobra.Command{
	Use:   "abra",
	Short: "Abra looks up an ABN or ACN using the ABR",
	Long: `A command line interface to the Australian Business Register.
More information available at https://github.com/ace-teknologi/abra`,
	RunE: func(cmd *cobra.Command, arg []string) error {
		fmt.Println("Maybe try abra help")
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&GUID, guidFlagName, "g", "", "Your ABR GUID")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setGUID() error {
	// If GUID wasn't set as a flag, check for an ENV
	flag := rootCmd.Flags().Lookup(guidFlagName)
	if flag.Value.String() == "" {
		g, ok := os.LookupEnv(abra.GUIDEnvName)
		if ok {
			flag.Value.Set(g)
		} else {
			return errors.New(abra.MissingGUIDError)
		}
	}
	return nil
}

func setOutputType(f string) (string, error) {
	if f == "" {
		return outputTypeTEXT, nil
	}

	if f != outputTypeJSON && f != outputTypeTEXT && f != outputTypeXML {
		return "", ErrInvalidOutputTypeMessage
	}

	return f, nil
}

func setOutputTypeTextTemplate(request string, path string) (*template.Template, error) {
	if request != "search" && request != "find" {
		return nil, fmt.Errorf("Invalid request type. Either `search` or `find` expected")
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	t, err := template.ParseFiles(filepath.Join(cwd, path))
	if err != nil {
		switch request {
		case "find":
			t, err = template.ParseFiles(filepath.Join(cwd, defaultFindTemplatePath))
		case "search":
			t, err = template.ParseFiles(filepath.Join(cwd, defaultSearchTemplatePath))
		}
	}
	if err != nil {
		return nil, err
	}

	return t, nil
}
