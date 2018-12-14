package abr

import (
	"fmt"
	"strconv"
	"strings"
)

// ACN represents an actual Australian Company Number
type ACN struct {
	IdentifierValue         string  `xml:"identifierValue,omitempty"`
	IsCurrentIndicator      string  `xml:"isCurrentIndicator,omitempty"`
	ReplacedIdentifierValue string  `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom            abnDate `xml:"replacedFrom,omitempty"`
}

// IsValid checks whether your ACN has a valid identifier
// (https://asic.gov.au/for-business/registering-a-company/steps-to-register-a-company/australian-company-numbers/australian-company-number-digit-check/)
func (a *ACN) IsValid() (bool, string) {
	// Strip whitespace
	raw := strings.Replace(a.IdentifierValue, " ", "", -1)
	if len(raw) != 9 {
		return false, fmt.Sprintf("Invalid ACN: found a %d character string", len(raw))
	}

	var checksum int
	var mapVal int
	var checkDigit int
	// compute the mapped values
	for i, char := range raw {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false, fmt.Sprintf("Invalid ACN: found character %s", string(char))
		}
		if i < 8 {
			mapVal += (8 - i) * digit
		}
		if i == 8 {
			checkDigit = digit
		}
	}
	checksum = (10 - (mapVal % 10)) % 10

	if checkDigit != checksum {
		return false, fmt.Sprintf("Invalid checksum: %d", checksum)
	}

	return true, ""
}

// ValidateACN tests a string to see if it is a valid ACN
func ValidateACN(acn string) (bool, string) {
	acnobj := ACN{IdentifierValue: acn}
	return acnobj.IsValid()
}
