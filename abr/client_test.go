package abr

import (
	"os"
	"testing"
)

func init() {
	guid, ok := os.LookupEnv("TEST_ABR_GUID")
	if !ok {
		panic("You must set TEST_ABR_GUID in order to run tests")
	}
	os.Setenv("ABR_GUID", guid)
}

func TestABRClient(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Error(err)
		return
	}

	if client.BaseURL.String() != BaseURL {
		t.Errorf("Expected endpoint to be %s, got %s", BaseURL, client.BaseURL.String())
	}
}

func TestABRClientNoEnvSet(t *testing.T) {
	guid := os.Getenv("ABR_GUID")
	os.Unsetenv("ABR_GUID")
	defer os.Setenv("ABR_GUID", guid)

	_, err := NewClient()
	if err == nil {
		t.Errorf("Expected an error, none was raised")
	} else if err.Error() != MissingGUIDError {
		t.Error(err)
	}

	return
}

var abnTestCases = []struct {
	abn  string
	name string
}{
	{"99124391073", "COzero Pty Ltd"},
	{"26154482283", "Oneflare Pty Ltd"},
	{"65433405893", "STUART J AULD"},
}

func TestSearchByABNv201408(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Error(err)
		return
	}

	for _, c := range abnTestCases {
		entity, err := client.SearchByABNv201408(c.abn, true)
		if err != nil {
			t.Error(err)
			continue
		}

		if entity.Name() != c.name {
			t.Errorf("Expected %v, got %v", c.name, entity.Name())
		}

		if entity.ABN() != c.abn {
			t.Errorf("Expected %v, got %v", c.abn, entity.ABN())
		}
	}
	return
}
