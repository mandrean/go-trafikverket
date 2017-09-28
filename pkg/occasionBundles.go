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
	"time"
)

type (
	OccasionBundlesRequest struct {
		BookingSession      BookingSession      `yaml:"bookingSession"`
		OccasionBundleQuery OccasionBundleQuery `yaml:"occasionBundleQuery"`
	}

	OccasionBundleQuery struct {
		StartDate         time.Time `yaml:"startDate"`
		LocationID        int       `yaml:"locationId"`
		LanguageID        int       `yaml:"languageId"`
		VehicleTypeID     int       `yaml:"vehicleTypeId"`
		TachographTypeID  int       `yaml:"tachographTypeId"`
		OccasionChoiceID  int       `yaml:"occasionChoiceId"`
		ExaminationTypeID int       `yaml:"examinationTypeId"`
	}

	OccasionBundlesResponse struct {
		Data []struct {
			Occasions []Occasion `yaml:"occasions"`
			Cost      string     `yaml:"cost"`
		} `yaml:"data"`
		Status int    `yaml:"status"`
		URL    string `yaml:"url"`
	}

	Occasion struct {
		ExaminationID interface{} `yaml:"examinationId"`
		Duration      struct {
			Start time.Time `yaml:"start"`
			End   time.Time `yaml:"end"`
		} `yaml:"duration"`
		ExaminationTypeID int         `yaml:"examinationTypeId"`
		LocationID        int         `yaml:"locationId"`
		OccasionChoiceID  int         `yaml:"occasionChoiceId"`
		VehicleTypeID     int         `yaml:"vehicleTypeId"`
		LanguageID        int         `yaml:"languageId"`
		TachographTypeID  int         `yaml:"tachographTypeId"`
		Name              string      `yaml:"name"`
		Properties        interface{} `yaml:"properties"`
		Date              string      `yaml:"date"`
		Time              string      `yaml:"time"`
		LocationName      string      `yaml:"locationName"`
		Cost              string      `yaml:"cost"`
		CostText          string      `yaml:"costText"`
		IncreasedFee      bool        `yaml:"increasedFee"`
		IsEducatorBooking interface{} `yaml:"isEducatorBooking"`
		PlaceAddress      string      `yaml:"placeAddress"`
	}
)

// OccasionBundles returns the occasion bundles for the specified parameters
func (tc *TrafikverketClient) OccasionBundles(body *OccasionBundlesRequest) (*OccasionBundlesResponse, *http.Response, error) {
	// create request
	req, err := NewRequest("POST", "/occasion-bundles", &body)
	if err != nil {
		return nil, nil, err
	}

	// make request
	res, err := tc.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, res, errors.New(res.Status)
	}

	// decode response
	var resp OccasionBundlesResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, res, err
	}

	b, _ := json.Marshal(resp)
	log.Debugln(string(b))

	return &resp, res, nil
}

// Occasions returns the available exam occasions for the specified parameters
func (tc *TrafikverketClient) Occasions(body *OccasionBundlesRequest) (*[]Occasion, *http.Response, error) {
	resp, res, err := tc.OccasionBundles(body)
	if err != nil {
		return nil, res, err
	}

	// merge occasions
	var o []Occasion
	for _, d := range resp.Data {
		if d.Occasions != nil {
			o = append(o, d.Occasions...)
		}
	}

	return &o, res, nil
}
