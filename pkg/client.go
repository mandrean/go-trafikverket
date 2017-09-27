// Copyright © 2017 Sebastian Mandrean <sebastian.mandrean@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package pkg

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

const (
	TRAFIKVERKET_BASE_URL = "https://fp.trafikverket.se"
	TRAFIKVERKET_BOKA_URL = TRAFIKVERKET_BASE_URL + "/Boka/"
)

type (
	TrafikverketClient struct {
		*http.Client
	}

	BookingSession struct {
		SocialSecurityNumber string `yaml:"socialSecurityNumber"`
		LicenceID            int    `yaml:"licenceId"`
		BookingModeID        int    `yaml:"bookingModeId"`
		IgnoreDebt           bool   `yaml:"ignoreDebt"`
		ExaminationTypeID    int    `yaml:"examinationTypeId"`
	}
)

// NewClient creates a new TrafikverketClient
func NewClient() *TrafikverketClient {
	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &TrafikverketClient{
		&http.Client{
			Timeout:   time.Second * 10,
			Transport: t,
		},
	}
}

// NewPostRequest creates a new *http.Request for the payload and sets the required headers
func NewRequest(method string, resource string, payload interface{}) (*http.Request, error) {
	// encode as json
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(payload)

	// join url
	u, err := url.Parse(TRAFIKVERKET_BOKA_URL)
	u.Path = path.Join(u.Path, resource)
	s := u.String()

	// create request
	req, err := http.NewRequest(strings.ToUpper(method), s, b)
	if err != nil {
		return nil, err
	}

	// set headers
	req.Header.Set("Origin", TRAFIKVERKET_BASE_URL)
	req.Header.Set("Referer", TRAFIKVERKET_BOKA_URL)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	log.Debugf("%v %v", req.Method, req.URL)
	log.Debugln(b.String())

	return req, nil
}
