package abra

// ClientAPI provides and interface to enable mocking the service client's
// API operation.
type ClientAPI interface {
	SearchByABN(string, bool) (*BusinessEntity, error)
	SearchByABNv201408(string, bool) (*BusinessEntity, error)
	SearchByACN(string, bool) (*BusinessEntity, error)
	SearchByASIC(string, bool) (*BusinessEntity, error)
	SearchByASICv201408(string, bool) (*BusinessEntity, error)
	SearchByName(string, params *NameSearchParams) (*ResponseSearchResultsList, error)
	SearchByNameAdvancedSimpleProtocol2017(string, params *NameSearchParams) (*ResponseSearchResultsList, error)
}
