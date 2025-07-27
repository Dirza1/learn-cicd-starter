package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"
)

type Test struct {
	Name          string
	Header        http.Header
	ExpectedOut   string
	ExpectedError error
}

var TestHeader = http.Header{}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestGetAPIKey(t *testing.T) {
	tests := []Test{
		{
			Name:          "Correct Header",
			Header:        TestHeader,
			ExpectedOut:   "APIKey",
			ExpectedError: nil,
		},
		{
			Name:          "incorrect Header",
			Header:        TestHeader,
			ExpectedOut:   "",
			ExpectedError: ErrNoAuthHeaderIncluded,
		},
		{
			Name:          "Incorrect Header syntax",
			Header:        TestHeader,
			ExpectedOut:   "",
			ExpectedError: errors.New("malformed authorization header"),
		},
		{
			Name:          "Header to long",
			Header:        TestHeader,
			ExpectedOut:   "",
			ExpectedError: errors.New("malformed authorization header"),
		},
	}

	for _, test := range tests {
		fmt.Printf("Starting test %s with the following parameters:\n Expected Output:%s\nExpected Error: %s\n",
			test.Name,
			test.ExpectedOut,
			test.ExpectedError)
		switch test.Name {
		case "Correct Header":
			TestHeader.Set("Authorization", "ApiKey APIKey")
		case "incorrect Header":
			TestHeader.Set("Authorization", "")
		case "Incorrect Header syntax":
			TestHeader.Set("Authorization", "test")
		case "Header to long":
			TestHeader.Set("Authorization", "test Test Test Test")
		}

		output, err := GetAPIKey(TestHeader)
		fmt.Printf("Output from test:\nOutput: %s\nTest error: %s\n\n\n", output, err)
		if err != nil && test.ExpectedError.Error() != err.Error() {
			t.Fail()
		}
		if output != test.ExpectedOut {
			t.Fail()
		}

	}
}
