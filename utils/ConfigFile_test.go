package utils

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestCreateConfigurationFromFile(t *testing.T) {
	config, err := CreateConfiguration("../examples/apt-term-config.json")
	if err != nil {
		t.Errorf("Error returned during CreateConfiguration when there shouldn't have.")
	}

	if config.Name != "Apt Terminal" {
		t.Errorf("Name improperly returned from CreateConfiguration")
	}

	if reflect.TypeOf(config).String() != "*models.Configuration" {
		t.Errorf("Improper type returned when running CreateConfiguration")
	}
}

func TestCreateConfigurationFromFileBadFile(t *testing.T) {
	_, err := CreateConfiguration("bad_file_name")
	if err == nil {
		t.Errorf("Error not returned when it should have after calling CreateConfiguration")
	}

	if err.Error() != "unable to create configuration object: \n\tunable to get file information for file: bad_file_name. error: stat bad_file_name: no such file or directory" {
		t.Errorf("Error improperly returned when it should have. \nError: %s", err.Error())
	}
}

func TestCreateConfigurationFromWeb(t *testing.T) {
	file, _ := os.Open("../examples/apt-term-config.json")
	scanner := bufio.NewScanner(file)

	configFileBytes := make([]byte, 0)

	for scanner.Scan() {
		configFileBytes = append(configFileBytes, scanner.Bytes()...)
	}

	testTable := []struct {
		name               string
		server             *httptest.Server
		expectedOutputType string
		expectedErr        string
	}{
		{
			name: "good-test-create-web-config",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(configFileBytes)
			})),
			expectedOutputType: "*models.Configuration",
			expectedErr:        "",
		},
		{
			name: "bad-test-create-web-config-404",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			})),
			expectedOutputType: "*models.Configuration",
			expectedErr:        "received status code 404 when attempting to get file",
		},
	}

	for _, testServer := range testTable {
		t.Run(testServer.name, func(t *testing.T) {
			defer testServer.server.Close()
			testWebSource, err := CreateConfiguration(testServer.server.URL)

			if reflect.TypeOf(testWebSource).String() != testServer.expectedOutputType {
				t.Errorf("Create Configuration returned the wrong output type for web configuration")
			}

			if (err != nil) && (testServer.expectedErr != "") {
				if !strings.Contains(err.Error(), testServer.expectedErr) {
					t.Errorf("Create Configuration returned an improper error when using incorrect web URL")
				}
			}

		})
	}
}
