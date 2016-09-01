####[Google re-captcha](https://www.google.com/recaptcha/intro) in [Go](http://golang.org).

[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/jimmy-go/recaptcha.svg?branch=master)](https://travis-ci.org/jimmy-go/recaptcha)
[![Go Report Card](https://goreportcard.com/badge/github.com/jimmy-go/recaptcha)](https://goreportcard.com/report/github.com/jimmy-go/recaptcha)
[![GoDoc](http://godoc.org/github.com/jimmy-go/recaptcha?status.png)](http://godoc.org/github.com/jimmy-go/recaptcha)
[![Coverage Status](https://coveralls.io/repos/github/jimmy-go/recaptcha/badge.svg?branch=master)](https://coveralls.io/github/jimmy-go/recaptcha?branch=master)

#####Install:
```
go get gopkg.in/jimmy-go/recaptcha.v0
```

*this package contains a breaking change*

```
    Before: func Verify(r *http.Request) bool
    Now: func Verify(response string) error
```

#####Usage:
```
    # declare a new validator
    re, err := recaptcha.New("YourKey", "YourSecret")
    check errors...


    captchaResponse := r.FormValue("g-recaptcha-response")
    err := re.Verify(captchaResponse)
    check errors...
```

The tests can't verify recaptcha solution so for real
world usage before using this package run the example:
```
# Configure your recaptcha to allow localhost.

# Run the example.
go run example/main.go -recaptcha-key="YourKey" -recaptcha-secret="YourSecret"
2016/09/01 13:43:56 Starting server on http://localhost:8100. Check example by opening this url in browser.
```

#####License:

The MIT License (MIT)

Copyright (c) 2016 github.com/haisum, Haisum (haisumbhatti@gmail.com)

changes by: github.com/jimmy-go.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
