package models

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNewWebSource(t *testing.T) {
	configSource := ConfigFileSource.New(ConfigFileSource{}, "../../examples/apt-term-config.json")
	configFileString, _ := configSource.LoadConfig()

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
		testWebSource := ConfigWebSource.New(ConfigWebSource{}, testServer.server.URL)
		if testWebSource.GetSourceType() != "web" {
			t.Errorf("Config source type not set properly.")
		}

		if reflect.TypeOf(testWebSource).String() != "*models.ConfigWebSource" {
			t.Errorf("Config source not of the right type.")
		}
	}

}

func TestLoadWebConfig(t *testing.T) {
	configSource := ConfigFileSource.New(ConfigFileSource{}, "../../examples/apt-term-config.json")
	configFileString, _ := configSource.LoadConfig()

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

	for _, testServer := range testTable {
		t.Run(testServer.name, func(t *testing.T) {
			defer testServer.server.Close()
			testWebSource := ConfigWebSource.New(ConfigWebSource{}, testServer.server.URL)
			configFileString, err := testWebSource.LoadConfig()

			if reflect.TypeOf(configFileString).String() != reflect.TypeOf(testServer.expectedOutput).String() {
				t.Errorf("LoadConfig does not return the correct type.")
			}

			if !reflect.DeepEqual(configFileString, testServer.expectedOutput) {
				t.Errorf("LoadConfig did not output the expected data.")
			}

			if !(err == nil && testServer.expectedErr == nil) {
				if reflect.TypeOf(err.Error()).String() != reflect.TypeOf(testServer.expectedErr.Error()).String() {
					t.Errorf("Error not of the same type as expected error.")
				}
			}

			if err != nil && !strings.Contains(err.Error(), testServer.expectedErr.Error()) {
				t.Errorf("Error improperly created when loading config.")
			}
		})

	}
}
