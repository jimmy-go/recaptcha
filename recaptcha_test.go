// Package recaptcha contains tests for recaptcha package.
//
// The MIT License (MIT)
//
// Copyright (c) 2016 github.com/haisum, Haisum (haisumbhatti@gmail.com)
//
// changes by: github.com/jimmy-go.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package recaptcha

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// T struct for tests.
type T struct {
	Purpose        string
	Input          Credentials
	Expected       error
	ExpectedVerify error
}

// Credentials for input data.
type Credentials struct {
	Key      string
	Secret   string
	Endpoint string
	Solution string
	Body     string
}

func TestTable(t *testing.T) {
	const FailHost = "http://127.0.0.1:8080"
	table := []T{
		{
			Input: Credentials{
				Key:      "",
				Secret:   "1",
				Endpoint: Endpoint,
				Solution: "123",
			},
			Expected:       ErrInvalidKey,
			ExpectedVerify: ErrInvalidKey,
		},
		{
			Input: Credentials{
				Key:      "1",
				Secret:   "",
				Endpoint: Endpoint,
				Solution: "123",
			},
			Expected:       ErrInvalidSecret,
			ExpectedVerify: ErrInvalidSecret,
		},
		{
			Purpose: "Void credentials",
			Input: Credentials{
				Key:      "1",
				Secret:   "1",
				Endpoint: Endpoint,
				Solution: "123",
			},
			Expected:       nil,
			ExpectedVerify: errors.New("recaptcha: response errors: invalid-input-response, invalid-input-secret"),
		},
		{
			Purpose: "Fail endpoint",
			Input: Credentials{
				Key:      "1",
				Secret:   "1",
				Endpoint: FailHost,
				Solution: "123",
			},
			Expected:       nil,
			ExpectedVerify: errors.New("Post http://127.0.0.1:8080: dial tcp 127.0.0.1:8080: getsockopt: connection refused"),
		},
		{
			Purpose: "Custom endpoint",
			Input: Credentials{
				Key:      "1",
				Secret:   "1",
				Endpoint: "http://1.2.3.4/",
				Solution: "123",
				Body:     "response body html",
			},
			Expected:       nil,
			ExpectedVerify: errors.New("recaptcha: response body html"),
		},
		{
			Purpose: "Empty body",
			Input: Credentials{
				Key:      "1",
				Secret:   "1",
				Endpoint: "http://1.2.3.4/",
				Solution: "123",
				Body:     "",
			},
			Expected:       nil,
			ExpectedVerify: ErrEmptyResponse,
		},
		{
			Purpose: "Response with errors",
			Input: Credentials{
				Key:      "1",
				Secret:   "1",
				Solution: "123",
				Body:     `{"error-codes":["one","two"]}`,
			},
			Expected:       nil,
			ExpectedVerify: errors.New("recaptcha: response errors: one, two"),
		},
		{
			Purpose: "Response without errors",
			Input: Credentials{
				Key:      "1",
				Secret:   "1",
				Solution: "123",
				Body:     `{"error-codes":[]}`,
			},
			Expected:       nil,
			ExpectedVerify: nil,
		},
	}
	for i := range table {
		x := table[i]

		verifyURL = x.Input.Endpoint

		re, err := New(x.Input.Key, x.Input.Secret)
		if fmt.Sprintf("%s", err) != fmt.Sprintf("%s", x.Expected) {
			t.Fail()
			t.Logf("new : purpose [%v] expected [%s] actual [%s]", x.Purpose, x.Expected, err)
			continue
		}
		// prevent panic
		if re == nil {
			continue
		}

		// Make a mock server for custom endpoints test.
		if verifyURL != Endpoint && verifyURL != FailHost {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, x.Input.Body)
			}))
			defer ts.Close()
			verifyURL = ts.URL
		}

		err = re.Verify(x.Input.Solution)
		if fmt.Sprintf("%s", err) != fmt.Sprintf("%s", x.ExpectedVerify) {
			t.Fail()
			t.Logf("verify : purpose [%v] expected [%s] actual [%s]", x.Purpose, x.ExpectedVerify, err)
			continue
		}
	}
}

// TestRecaptcha test.
//
// You need to start example/main.go with your
// credentials in order to validate is working.
func TestRecaptcha(t *testing.T) {
	verifyURL = Endpoint

	re, err := New("invalidKey", "invalidSecret")
	if err != nil {
		t.Fail()
		t.Logf("new : err [%s]", err)
	}
	err = re.Verify("testresponse")
	if err == nil {
		t.Fail()
		t.Logf("new : err [%s]", err)
	}
	expected := "recaptcha: response errors: invalid-input-response, invalid-input-secret"
	if err.Error() != expected {
		t.Fail()
		t.Logf("new : err [%s]", err)
	}
}

func TestRecaptchaFail(t *testing.T) {
	verifyURL = Endpoint

	re, err := New("123", "456")
	if err != nil {
		t.Fail()
		t.Logf("new : err [%s]", err)
	}
	err = re.Verify("testresponse")
	if err == nil {
		t.Fail()
		t.Logf("new : err [%s]", err)
	}
}
