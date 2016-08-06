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

import "testing"

// T struct for tests.
type T struct {
	Key      string
	Secret   string
	Expected error
}

var (
	table = []T{
		{
			Key:      "",
			Secret:   "1",
			Expected: errInvalidKey,
		},
		{
			Key:      "1",
			Secret:   "",
			Expected: errInvalidSecret,
		},
		{
			Key:      "1",
			Secret:   "1",
			Expected: nil,
		},
	}
)

func TestTable(t *testing.T) {
	for _, x := range table {
		_, err := New(x.Key, x.Secret)
		if err != x.Expected {
			t.Fail()
			t.Logf("TestTable : expected [%s] actual [%s]", x.Expected, err)
		}
	}
}

// TestRecaptcha test.
// FIXME; You need to start example/main.go with your
// credentials in order to validate is working.
func TestRecaptcha(t *testing.T) {
	re, err := New("invalidKey", "invalidSecret")
	if err != nil {
		t.Fail()
		t.Logf("TestRecaptcha : new : err [%s]", err)
	}
	err = re.Verify("testresponse")
	if err == nil {
		t.Fail()
		t.Logf("TestRecaptcha : new : err [%s]", err)
	}
	expected := "recaptcha: response errors: invalid-input-response, invalid-input-secret"
	if err.Error() != expected {
		t.Fail()
		t.Logf("TestRecaptcha : new : err [%s]", err)
	}
}

func TestRecaptchaFail(t *testing.T) {
	re, err := New("123", "456")
	if err != nil {
		t.Fail()
		t.Logf("TestRecaptcha : new : err [%s]", err)
	}
	err = re.Verify("testresponse")
	if err == nil {
		t.Fail()
		t.Logf("TestRecaptcha : new : err [%s]", err)
	}
}
