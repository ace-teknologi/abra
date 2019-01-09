package abr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	BaseURL          = "https://www.abn.business.gov.au/abrxmlsearch/ABRXMLSearch.asmx/"
	contentType      = "application/json"
	GUIDEnvName      = "ABR_GUID"
	MissingGUIDError = "The ABR_GUID environment variable must be set"
	userAgent        = "go-abn v0.0.0"
)

var defaultNameSearchParams = NameSearchParams{
	ActiveABNsOnly:   true,
	BusinessName:     true,
	LegalName:        true,
	MaxSearchResults: 50,
	MinimumScore:     0,
	Postcode:         "",
	States:           make([]State, 0),
	TradingName:      true,
	TypicalSearch:    true,
}

// Client provides a client connection to the ABR
type Client struct {
	BaseURL *url.URL
	GUID    string

	httpClient *http.Client
}

// Payload encapsulates a Response from the API
type Payload struct {
	Response *Response `xml:"response,omitempty"`
}

// Response represents the XML responses available via the ABRXMLSearch
type Response struct {
	UsageStatement          string    `xml:"usageStatement,omitempty"`
	DateRegisterLastUpdated abnDate   `xml:"dateRegisterLastUpdated,omitempty"`
	DateTimeRetrieved       time.Time `xml:"dateTimeRetrieved,omitempty"`

	BusinessEntity       *BusinessEntity            `xml:"businessEntity,omitempty"`
	BusinessEntity200506 *BusinessEntity            `xml:"businessEntity200506,omitempty"`
	BusinessEntity200709 *BusinessEntity            `xml:"businessEntity200709,omitempty"`
	BusinessEntity201205 *BusinessEntity            `xml:"businessEntity201205,omitempty"`
	BusinessEntity201408 *BusinessEntity            `xml:"businessEntity201408,omitempty"`
	Exception            *ResponseException         `xml:"exception,omitempty"`
	AbnList              *ResponseABNList           `xml:"abnList,omitempty"`
	SearchResultsList    *ResponseSearchResultsList `xml:"searchResultsList,omitempty"`
}

// ResponseABNList
type ResponseABNList struct {
	NumberOfRecords int32    `xml:"numberOfRecords,omitempty"`
	Abn             []string `xml:"abn,omitempty"`
}

// ResponseException returns the exception
type ResponseException struct {
	ExceptionDescription string `xml:"exceptionDescription,omitempty"`
	ExceptionCode        string `xml:"exceptionCode,omitempty"`
}

// ResponseSearchResultsList
type ResponseSearchResultsList struct {
	NumberOfRecords int32  `xml:"numberOfRecords,omitempty"`
	ExceedsMaximum  string `xml:"exceedsMaximum,omitempty"`

	SearchResultsRecord []*SearchResultsRecord `xml:"searchResultsRecord,omitempty"`
}

// NameSearchParams
type NameSearchParams struct {
	ActiveABNsOnly   bool
	BusinessName     bool
	LegalName        bool
	MaxSearchResults int32
	MinimumScore     int32
	Postcode         string
	States           []State
	TradingName      bool
	TypicalSearch    bool
}

// SearchResultsRecord
type SearchResultsRecord struct {
	ABN                         ABN                          `xml:"ABN,omitempty"`
	BusinessName                *SearchResultName            `xml:"businessName,omitempty"`
	LegalName                   *SearchResultName            `xml:"legalName,omitempty"`
	MainName                    *SearchResultName            `xml:"mainName,omitempty"`
	MainTradingName             *SearchResultName            `xml:"mainTradingName,omitempty"`
	OtherTradingName            *SearchResultName            `xml:"otherTradingName,omitempty"`
	MainBusinessPhysicalAddress *MainBusinessPhysicalAddress `xml:"mainBusinessPhysicalAddress,omitempty"`
}

// MainName
type SearchResultName struct {
	OrganisationName   string `xml:"organisationName,omitempty"`
	FullName           string `xml:"fullName,omitempty"`
	Score              int32  `xml:"score,omitempty"`
	IsCurrentIndicator string `xml:"isCurrentIndicator,omitempty"`
}

// MainBusinessPhysicalAddress
type MainBusinessPhysicalAddress struct {
	StateCode          string `xml:"stateCode,omitempty"`
	Postcode           string `xml:"postcode,omitempty"`
	IsCurrentIndicator string `xml:"isCurrentIndicator,omitempty"`
}

// FriendlyName provides a FriendlyName for a `SearchResultName`
func (r *SearchResultsRecord) Name() string {
	if r.MainName != nil {
		return r.MainName.OrganisationName
	}
	if r.MainTradingName != nil {
		return r.MainTradingName.OrganisationName
	}
	if r.BusinessName != nil {
		return r.BusinessName.OrganisationName
	}
	if r.OtherTradingName != nil {
		return r.OtherTradingName.OrganisationName
	}
	if r.LegalName != nil {
		return r.LegalName.FullName
	}
	return ""
}

func (r *SearchResultsRecord) Score() int32 {
	if r.MainName != nil {
		return r.MainName.Score
	}
	if r.MainTradingName != nil {
		return r.MainTradingName.Score
	}
	if r.BusinessName != nil {
		return r.BusinessName.Score
	}
	if r.OtherTradingName != nil {
		return r.OtherTradingName.Score
	}
	if r.LegalName != nil {
		return r.LegalName.Score
	}
	return 0
}

func (r *SearchResultsRecord) IsCurrentIndicator() string {
	if r.MainName != nil {
		return r.MainName.IsCurrentIndicator
	}
	if r.MainTradingName != nil {
		return r.MainTradingName.IsCurrentIndicator
	}
	if r.BusinessName != nil {
		return r.BusinessName.IsCurrentIndicator
	}
	if r.OtherTradingName != nil {
		return r.OtherTradingName.IsCurrentIndicator
	}
	if r.LegalName != nil {
		return r.LegalName.IsCurrentIndicator
	}
	return ""
}

// NewClient returns a pointer to an initialized instance of a Client
func NewClient() (*Client, error) {
	guid, ok := os.LookupEnv(GUIDEnvName)
	if !ok {
		return nil, errors.New(MissingGUIDError)
	}

	return NewWithGuid(guid)
}

// NewClientWithGuid returns a pointer to an initalized instance of a Client,
// configured with a `guid`
func NewWithGuid(guid string) (*Client, error) {
	rawurl, ok := os.LookupEnv("ABR_ENDPOINT")
	if !ok {
		rawurl = BaseURL
	}

	baseurl, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	client := &Client{
		BaseURL:    baseurl,
		httpClient: &http.Client{},
		GUID:       guid,
	}
	return client, nil
}

// SearchByABN is an alias for SearchByABNv201408
func (c *Client) SearchByABN(abn string, hist bool) (*BusinessEntity, error) {
	return c.SearchByABNv201408(abn, hist)
}

// SearchByABNv201408 wraps the API call to query an ABN
func (c *Client) SearchByABNv201408(abn string, hist bool) (*BusinessEntity, error) {
	if ok, err := ValidateABN(abn); !ok {
		return nil, err
	}

	data := url.Values{}
	data.Set("authenticationGuid", c.GUID)
	if hist {
		data.Add("includeHistoricalDetails", "Y")
	} else {
		data.Add("includeHistoricalDetails", "N")
	}
	data.Add("searchString", abn)

	req, err := c.newRequest("POST", "SearchByABNv201408", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	resp := Payload{}

	_, err = c.do(req, &resp)
	if err != nil {
		return nil, err
	}

	// ABR returns a 200 response for some exceptions. Check for these.
	if resp.Response.Exception != nil {
		log.Printf("[ERROR] %v\n", resp.Response.Exception.ExceptionDescription)
		return nil, fmt.Errorf(resp.Response.Exception.ExceptionDescription)
	}

	return resp.Response.BusinessEntity201408, nil
}

// SearchByACN is an alias for SearchByASICv201408
func (c *Client) SearchByACN(abn string, hist bool) (*BusinessEntity, error) {
	return c.SearchByASICv201408(abn, hist)
}

// SearchByASIC is an alias for SearchByASICv201408
func (c *Client) SearchByASIC(abn string, hist bool) (*BusinessEntity, error) {
	return c.SearchByASICv201408(abn, hist)
}

// SearchByASICv201408 wraps the API call to query an ACN
func (c *Client) SearchByASICv201408(acn string, hist bool) (*BusinessEntity, error) {
	if ok, err := ValidateACN(acn); !ok {
		return nil, fmt.Errorf(err)
	}

	data := url.Values{}
	data.Set("authenticationGuid", c.GUID)
	if hist {
		data.Add("includeHistoricalDetails", "Y")
	} else {
		data.Add("includeHistoricalDetails", "N")
	}
	data.Add("searchString", acn)

	req, err := c.newRequest("POST", "SearchByASICv201408", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	resp := Payload{}

	_, err = c.do(req, &resp)
	if err != nil {
		return nil, err
	}

	// ABR returns a 200 response for some exceptions. Check for these.
	if resp.Response.Exception != nil {
		log.Printf("[ERROR] %v\n", resp.Response.Exception.ExceptionDescription)
		return nil, fmt.Errorf(resp.Response.Exception.ExceptionDescription)
	}

	return resp.Response.BusinessEntity201408, nil
}

// SearchByName wraps the API call to query by `name` and other optional details
func (c *Client) SearchByName(name string, params *NameSearchParams) (*ResponseSearchResultsList, error) {
	return c.SearchByNameAdvancedSimpleProtocol2017(name, params)
}

func (c *Client) SearchByNameAdvancedSimpleProtocol2017(name string, params *NameSearchParams) (*ResponseSearchResultsList, error) {
	// Strip whitespace
	name = strings.Trim(name, " ")

	if len(name) == 0 {
		return nil, fmt.Errorf("No `name` to search with provided")
	}

	if params == nil {
		params = &defaultNameSearchParams
	}

	data := url.Values{}
	// Authentication GUID
	data.Set("authenticationGuid", c.GUID)
	// Search value
	data.Add("name", name)
	// Optional Parameters
	// Postcode isolation
	if params.Postcode != "" {
		data.Add("postcode", params.Postcode)
	} else {
		data.Add("postcode", "")
	}
	// LegalName
	if params.LegalName {
		data.Add("legalName", "Y")
	} else {
		data.Add("legalName", "N")
	}
	// TradingName
	if params.TradingName {
		data.Add("tradingName", "Y")
	} else {
		data.Add("tradingName", "N")
	}
	// BusinessName
	if params.BusinessName {
		data.Add("businessName", "Y")
	} else {
		data.Add("businessName", "N")
	}
	// ActiveABNsOnly
	if params.ActiveABNsOnly {
		data.Add("activeABNsOnly", "Y")
	} else {
		data.Add("activeABNsOnly", "N")
	}
	// State isolated searches
	if len(params.States) > 0 {
		for _, state := range States {
			data.Add(state.ShortTitle, "N")
		}

		for _, state := range params.States {
			data.Add(state.ShortTitle, "Y")
		}
	} else {
		for _, state := range States {
			data.Add(state.ShortTitle, "Y")
		}
	}
	// Quality filtering
	if params.TypicalSearch {
		data.Add("searchWidth", "typical")
	} else {
		data.Add("searchWidth", "narrow")
	}
	// MinimumScore
	if params.MinimumScore > 0 {
		data.Add("minimumScore", fmt.Sprintf("%d", params.MinimumScore))
	} else {
		data.Add("minimumScore", "0")
	}
	// MinimumScore
	if params.MaxSearchResults > 0 {
		data.Add("maxSearchResults", fmt.Sprintf("%d", params.MaxSearchResults))
	} else {
		data.Add("maxSearchResults", "100")
	}

	req, err := c.newRequest("POST", "ABRSearchByNameAdvancedSimpleProtocol2017", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	resp := Payload{}

	_, err = c.do(req, &resp)
	if err != nil {
		return nil, err
	}

	// ABR returns a 200 response for some exceptions. Check for these.
	if resp.Response.Exception != nil {
		log.Printf("[ERROR] %v\n", resp.Response.Exception.ExceptionDescription)
		return nil, fmt.Errorf(resp.Response.Exception.ExceptionDescription)
	}

	return resp.Response.SearchResultsList, nil
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	rel, _ := url.Parse(path)
	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Header.Set("Accept", "application/xml")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	b, err := req.GetBody()
	buf := new(bytes.Buffer)
	buf.ReadFrom(b)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[ERROR] received status %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, v)

	return resp, err
}
