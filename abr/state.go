package abr

import (
	"fmt"
)

// State represents an actual Australian Company Number
type State struct {
	Title      string
	ShortTitle string
}

var States = []State{
	State{Title: "Australian Capital Territory", ShortTitle: "ACT"},
	State{Title: "New South Wales", ShortTitle: "NSW"},
	State{Title: "Northern Territory", ShortTitle: "NT"},
	State{Title: "Queensland", ShortTitle: "QLD"},
	State{Title: "South Australia", ShortTitle: "SA"},
	State{Title: "Tasmania", ShortTitle: "TAS"},
	State{Title: "Victoria", ShortTitle: "VIC"},
	State{Title: "Western Australia", ShortTitle: "WA"},
}

// IsValid checks whether your State is valid
func (s *State) IsValid() (bool, error) {
	validStateTitle := false
	validStateShortTitle := false

	for _, state := range States {
		if state.Title == s.Title {
			validStateTitle = true
		}
		if state.ShortTitle == s.ShortTitle {
			validStateShortTitle = true
		}
	}

	if validStateTitle && validStateShortTitle {
		return true, nil
	}

	return false, fmt.Errorf("Invalid state: %s (%s)", s.Title, s.ShortTitle)
}

func FindState(shortTitle string) (*State, error) {
	var state *State

	for _, s := range States {
		if s.ShortTitle == shortTitle {
			state = &s
		}
	}

	if state == nil {
		return nil, fmt.Errorf("Could not find state with ShortTitle: %s", shortTitle)
	}

	return state, nil
}
