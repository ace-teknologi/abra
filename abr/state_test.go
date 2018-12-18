package abr

import (
	"testing"
)

var validStates = []State{
	{ShortTitle: "ACT"},
	{ShortTitle: "NSW"},
	{ShortTitle: "NT"},
	{ShortTitle: "QLD"},
	{ShortTitle: "SA"},
	{ShortTitle: "TAS"},
	{ShortTitle: "VIC"},
	{ShortTitle: "WA"},
}

var invalidStates = []State{
	{ShortTitle: "derpderpder"},
	{ShortTitle: ""},
	{ShortTitle: "  "},
}

func TestFindStateWithValidStates(t *testing.T) {
	for _, state := range validStates {
		if state, err := FindState(state.ShortTitle); state == nil {
			t.Errorf("Expecting to find State %s, but it wasn't found: %s", state.ShortTitle, err)
		}
	}

	for _, state := range invalidStates {
		if state, _ := FindState(state.ShortTitle); state != nil {
			t.Errorf("Expecting to not find State %s, but did.", state.ShortTitle)
		}
	}

	return
}
