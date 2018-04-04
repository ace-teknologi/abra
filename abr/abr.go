// A Go wrapper for the Australian Business Register
package abr

import (
	"encoding/xml"
	"fmt"
	"time"
)

// BusinessEntity represents a legal entity and encapsulates everything that the
// ABR knows about such an entity
type BusinessEntity struct {
	RecordLastUpdatedDate abnDate `xml:"recordLastUpdatedDate,omitempty"`

	// So many ways of naming a business!
	BusinessNames     []*BusinessName `xml:"businessName,omitempty"`
	HumanNames        []*HumanName    `xml:"legalName,omitempty"`
	MainNames         []*BusinessName `xml:"mainName,omitempty"`
	MainTradingNames  []*BusinessName `xml:"mainTradingName,omitempty"`
	OtherTradingNames []*BusinessName `xml:"otherTradingName,omitempty"`

	// Other details
	ABNs                 []*ABN                 `xml:"ABN,omitempty"`
	ASICNumber           string                 `xml:"ASICNumber,omitempty"`
	EntityStatuses       []*EntityStatus        `xml:"entityStatus,omitempty"`
	EntityType           *EntityType            `xml:"entityType,omitempty"`
	GSTRegisteredPeriods []*GSTRegisteredPeriod `xml:"goodsAndServicesTax,omitempty"`
	PhysicalAddresses    []*AddressDetails      `xml:"mainBusinessPhysicalAddress,omitempty"`
}

// BusinessName represents some kind of name that was in use by a BusinessEntity
// at some point in history
type BusinessName struct {
	OrganisationName string  `xml:"organisationName"`
	EffectiveFrom    abnDate `xml:"effectiveFrom"`
	EffectiveTo      abnDate `xml:"effectiveTo"`
}

// HumanName represents the name of a human that was in use by a BusinessEntity
// at some point in history
type HumanName struct {
	GivenName      string `xml:"givenName"`
	OtherGivenName string `xml:"otherGivenName"`
	FamilyName     string `xml:"familyName"`

	EffectiveFrom abnDate `xml:"effectiveFrom"`
	EffectiveTo   abnDate `xml:"effectiveTo"`
}

// ABN represents an actual Australian Business Number
type ABN struct {
	IdentifierValue         string  `xml:"identifierValue,omitempty"`
	IsCurrentIndicator      string  `xml:"isCurrentIndicator,omitempty"`
	ReplacedIdentifierValue string  `xml:"replacedIdentifierValue,omitempty"`
	ReplacedFrom            abnDate `xml:"replacedFrom,omitempty"`
}

// Address details represents an address used by a BusinessEntity for a period
// of time
type AddressDetails struct {
	StateCode string `xml:"stateCode,omitempty"`
	Postcode  string `xml:"postcode,omitempty"`

	EffectiveFrom abnDate `xml:"effectiveFrom,omitempty"`
	EffectiveTo   abnDate `xml:"effectiveTo,omitempty"`
}

// EntityStatus represents the status of an entity (Active or Cancelled)
type EntityStatus struct {
	EntityStatusCode string  `xml:"entityStatusCode,omitempty"`
	EffectiveFrom    abnDate `xml:"effectiveFrom,omitempty"`
	EffectiveTo      abnDate `xml:"effectiveTo,omitempty"`
}

// EntityType represents the type of an entity - individal, pty ltd, etc.
type EntityType struct {
	EntityTypeCode    string `xml:"entityTypeCode,omitempty"`
	EntityDescription string `xml:"entityDescription,omitempty"`
}

// GSTRegisteredPeriod represents a period during which the entity was
// registered for GST
type GSTRegisteredPeriod struct {
	EffectiveFrom abnDate `xml:"effectiveFrom,omitempty"`
	EffectiveTo   abnDate `xml:"effectiveTo,omitempty"`
}

// abnDate represents a yyyy-mm-dd formatted date
type abnDate struct {
	time.Time
}

func (a *abnDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const form = "2006-01-02" // yyyy-mm-dd
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(form, v)
	if err != nil {
		return err
	}
	*a = abnDate{parse}
	return nil
}

// Name returns the most relevant name that is available
func (b *BusinessEntity) Name() string {
	for _, name := range b.MainNames {
		return name.OrganisationName
	}

	for _, name := range b.MainTradingNames {
		return name.OrganisationName
	}

	for _, name := range b.OtherTradingNames {
		return name.OrganisationName
	}

	for _, name := range b.BusinessNames {
		return name.OrganisationName
	}

	for _, name := range b.HumanNames {
		return fmt.Sprintf("%s %s", name.GivenName, name.FamilyName)
	}

	return ""
}

// ABN returns the current ABN
func (b *BusinessEntity) ABN() string {
	for _, abn := range b.ABNs {
		if abn.IsCurrentIndicator == "Y" {
			return abn.IdentifierValue
		}
	}
	return ""
}
