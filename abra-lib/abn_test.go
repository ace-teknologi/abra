package abra

import (
	"testing"
)

var validABNs = []ABN{
	{IdentifierValue: "99124391073"},
	{IdentifierValue: "26154482283"},
	{IdentifierValue: "65433405893"},
	{IdentifierValue: "99 124 391 073"},
	{IdentifierValue: "99  12 43 9 10 7   3"},
	{IdentifierValue: "  99 124 391 073  "},
}

var invalidABNs = []ABN{
	{IdentifierValue: "derpderpder"},
	{IdentifierValue: "99124391074"},
	{IdentifierValue: "26154482284"},
	{IdentifierValue: "65433405894"},
	{IdentifierValue: ""},
	{IdentifierValue: "1"},
	{IdentifierValue: "one"},
	{IdentifierValue: "991243910731"},
	{IdentifierValue: "26154482283d"},
	{IdentifierValue: "65433405893*"},
}

func TestABNIsValid(t *testing.T) {
	for _, abn := range validABNs {
		if ok, msg := abn.IsValid(); !ok {
			t.Errorf("Expecting ABN %s to be valid, but it isn't: %s", abn.IdentifierValue, msg)
		}
	}

	for _, abn := range invalidABNs {
		if ok, _ := abn.IsValid(); ok {
			t.Errorf("Expecting ABN %s to be invalid, but it isn't!", abn.IdentifierValue)
		}
	}

	return
}
