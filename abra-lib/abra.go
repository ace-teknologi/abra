// Package abra provides a Go wrapper for the Australian Business Register
package abra

// Abra is the interface implemented by Client
type Abra interface {
	Search(string) (*BusinessEntity, error)
	SearchByABN(string, bool) (*BusinessEntity, error)
	SearchByABNv201408(string, bool) (*BusinessEntity, error)
	SearchByACN(string, bool) (*BusinessEntity, error)
	SearchByASIC(string, bool) (*BusinessEntity, error)
	SearchByASICv201408(string, bool) (*BusinessEntity, error)
	SearchByName(string, *NameSearchParams) (*ResponseSearchResultsList, error)
	SearchByNameAdvancedSimpleProtocol2017(string, *NameSearchParams) (*ResponseSearchResultsList, error)
	SearchByNameBestGuess(string, *NameSearchParams) (*BusinessEntity, error)
}
