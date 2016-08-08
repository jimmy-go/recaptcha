// Package main contains an example of go recapcha usage.
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
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jimmy-go/recaptcha"
)

var (
	key    = flag.String("recaptcha-key", "", "Recaptcha key.")
	secret = flag.String("recaptcha-secret", "", "Recaptcha secret.")
)

func main() {
	flag.Parse()

	re, err := recaptcha.New(*key, *secret)
	if err != nil {
		panic(err)
	}
	// render site.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, siteTemplate, *key)
	})
	// process captcha
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		captchaResponse := r.FormValue("g-recaptcha-response")
		err := re.Verify(captchaResponse)
		if err != nil {
			w.Write([]byte("invalid captcha response: " + err.Error()))
			return
		}
		w.Write([]byte("captcha response is valid"))
	})
	// start server.
	log.Printf("Starting server on http://localhost:8100. Check example by opening this url in browser.\n")
	log.Fatal(http.ListenAndServe(":8100", nil))
}

const (
	siteTemplate = `
<html>
<head>
	<script src='https://www.google.com/recaptcha/api.js'></script>
</head>
<body>
	<form action="/submit" method="post">
		<div class="g-recaptcha" data-sitekey="%s"></div>
		<input type="submit">
	</form>
</body>
</html>`
)
