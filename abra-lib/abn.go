package abra

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ABN represents an actual Australian Business Number
type ABN struct {
	IdentifierValue         string  `xml:"identifierValue,omitempty"`
	IdentifierStatus        string  `xml:"identifierStatus,omitempty"`
	IsCurrentIndicator      string  `xml:"isCurrentIndicator,omitempty"`
	ReplacedIdentifierValue string  `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom            abnDate `xml:"replacedFrom,omitempty"`
}

// IsValid checks whether your ABN has a valid identifier
// (https://www.abr.business.gov.au/HelpAbnFormat.aspx)
func (a *ABN) IsValid() (bool, error) {
	return ValidateABN(a.IdentifierValue)
}

// String returns the ABN as a formatted string
func (a *ABN) String() string {
	return a.IdentifierValue
}

// ValidateABN tests a string to see if it is a valid ABN
func ValidateABN(abn string) (bool, error) {
	// Strip whitespace
	raw := strings.Replace(abn, " ", "", -1)
	if len(raw) != 11 {
		return false, fmt.Errorf("Invalid ABN: found a %d character string", len(raw))
	}

	var checksum int
	// compute the checksum
	for i, char := range raw {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false, fmt.Errorf("Invalid ABN: found character %s", string(char))
		}
		if i == 0 {
			// For the first digit, subtract 1 then multiply by ten
			checksum += (digit - 1) * 10
		} else {
			// For all other digits, multiply by (2i - 1)
			checksum += (digit * (2*i - 1))
		}
	}

	if math.Mod(float64(checksum), float64(89)) > 0 {
		return false, fmt.Errorf("Invalid checksum: %d", checksum)
	}

	return true, nil
}
