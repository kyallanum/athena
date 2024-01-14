package models

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewWebSource(t *testing.T) {
	configSource := NewWebSource("../../examples/apt-term-config.json")
	configFileString, _ := configSource.Config()

	testTable := []struct {
		name   string
		server *httptest.Server
	}{
		{
			name: "good-test-server-response",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(configFileString)
			})),
		},
	}

	for _, testServer := range testTable {
		t.Run(testServer.name, func(t *testing.T) {
			testWebSource := NewWebSource(testServer.server.URL)
			if testWebSource.SourceType() != "web" {
				t.Errorf("Config source type not set properly.")
			}

			if reflect.TypeOf(testWebSource).String() != "*models.WebSource" {
				t.Errorf("Config source not of the right type.")
			}
		})
	}

}

func TestLoadWebConfig(t *testing.T) {
	configSource := NewFileSource("../../examples/apt-term-config.json")
	configFileString, _ := configSource.Config()

	testTable := []struct {
		name           string
		server         *httptest.Server
		expectedOutput []byte
		expectedErr    error
	}{
		{
			name: "good-test-server-response",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(configFileString)
			})),
			expectedOutput: configFileString,
			expectedErr:    nil,
		},
		{
			name: "404-error-server-response",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			})),
			expectedOutput: nil,
			expectedErr:    fmt.Errorf("received status code 404 when attempting to get file"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()
			testWebSource := NewWebSource(test.server.URL)
			configFileString, err := testWebSource.Config()

			if reflect.TypeOf(configFileString).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("LoadConfig does not return the correct type.")
			}

			if !reflect.DeepEqual(configFileString, test.expectedOutput) {
				t.Errorf("LoadConfig did not output the expected data.")
			}

			if !checkExpectedError(err, test.expectedErr) {
				t.Errorf("Error returned improperly: \n\tExpected: %s\n\tReceived: %s", test.expectedErr.Error(), err.Error())
			}
		})

	}
}
