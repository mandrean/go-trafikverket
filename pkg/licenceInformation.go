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
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type (
	LicenceInformationResponse struct {
		Data struct {
			EnableSocialSecurityNumber bool              `yaml:"enableSocialSecurityNumber"`
			SocialSecurityNumber       string            `yaml:"socialSecurityNumber"`
			LicenceID                  int               `yaml:"licenceId"`
			LicenceCategories          []LicenceCategory `yaml:"licenceCategories"`
		} `yaml:"data"`
		Status int    `yaml:"status"`
		URL    string `yaml:"url"`
	}

	LicenceCategory struct {
		Name     string `yaml:"name"`
		Licences []struct {
			ID          int    `yaml:"id"`
			Name        string `yaml:"name"`
			Description string `yaml:"description"`
			Category    string `yaml:"category"`
			Icon        string `yaml:"icon"`
		} `yaml:"licences"`
	}
)

// LicenceInformation returns information about different driver's licence types available for booking exams for
func (tc *TrafikverketClient) LicenceInformation() (*LicenceInformationResponse, *http.Response, error) {
	// create request
	req, err := NewRequest("POST", "/licence-information", "{}")
	if err != nil {
		return nil, nil, err
	}

	// make request
	res, err := tc.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, res, errors.New(res.Status)
	}

	// decode response
	var resp LicenceInformationResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, res, err
	}

	b, _ := json.Marshal(resp)
	log.Debugln(string(b))

	return &resp, res, nil
}

// LicenceCategories returns the different driver's licence categories available for booking exams for
func (tc *TrafikverketClient) LicenceCategories() (*[]LicenceCategory, *http.Response, error) {
	resp, res, err := tc.LicenceInformation()
	if err != nil {
		return nil, res, err
	}

	return &resp.Data.LicenceCategories, res, nil
}
