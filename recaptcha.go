// Package recaptcha contains recaptcha validator for Go.
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
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var (
	errInvalidKey    = errors.New("recaptcha: invalid key length")
	errInvalidSecret = errors.New("recaptcha: invalid secret length")
)

const (
	// verifyURL URL to google validation service.
	verifyURL = "https://www.google.com/recaptcha/api/siteverify"
)

// Recaptcha struct for Google recaptcha validation.
// Process every request one by one for now. Later support will add concurrent
// validation.
type Recaptcha struct {
	Key    string
	Secret string

	client *http.Client
	sync.RWMutex
}

// New returns a new Recaptcha validator.
func New(key, secret string) (*Recaptcha, error) {
	if len(key) < 1 {
		return nil, errInvalidKey
	}
	if len(secret) < 1 {
		return nil, errInvalidSecret
	}
	re := &Recaptcha{
		Key:    key,
		Secret: secret,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
	return re, nil
}

// Response is google response.
type Response struct {
	ErrorCodes []string `json:"error-codes"`
}

// Verify validates recaptcha with google service.
func (r *Recaptcha) Verify(response string) error {
	r.RLock()
	defer r.RUnlock()

	v := url.Values{}
	v.Add("secret", r.Secret)
	v.Add("response", response)
	resp, err := r.client.PostForm(verifyURL, v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var res *Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	err = joinerrs(res.ErrorCodes)
	return err
}

// joinerrs joins all errors in one message or return nil.
func joinerrs(s []string) error {
	if len(s) < 1 {
		return nil
	}
	msg := strings.Join(s, ", ")
	return errors.New("recaptcha: response errors: " + msg)
}
