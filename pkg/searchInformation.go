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
	SearchInformationRequest struct {
		BookingSession BookingSession `yaml:"bookingSession"`
	}

	SearchInformationResponse struct {
		Data struct {
			CanBookLicence bool   `yaml:"canBookLicence"`
			LicenceID      uint64 `yaml:"licenceId"`
			Licences       []struct {
				ID          uint64 `yaml:"id"`
				Name        string `yaml:"name"`
				Description string `yaml:"description"`
				Category    string `yaml:"category"`
				Icon        string `yaml:"icon"`
			} `yaml:"licences"`
			LicenceCategories []LicenceCategory `yaml:"licenceCategories"`
			LocationID        uint64            `yaml:"locationId"`
			Locations         []Location        `yaml:"locations"`
			TimeIntervalID    uint64            `yaml:"timeIntervalId"`
			TimeIntervals     []struct {
				ID        uint64    `yaml:"id"`
				StartDate time.Time `yaml:"startDate"`
				Name      string    `yaml:"name"`
			} `yaml:"timeIntervals"`
			ShowLanguage bool   `yaml:"showLanguage"`
			LanguageID   uint64 `yaml:"languageId"`
			Languages    []struct {
				ID          uint64 `yaml:"id"`
				Name        string `yaml:"name"`
				LocationIDs []int  `yaml:"locationIds"`
			} `yaml:"languages"`
			ShowVehicleType bool   `yaml:"showVehicleType"`
			VehicleTypeID   uint64 `yaml:"vehicleTypeId"`
			VehicleTypes    []struct {
				ID   int    `yaml:"id"`
				Name string `yaml:"name"`
			} `yaml:"vehicleTypes"`
			ShowTachographType bool   `yaml:"showTachographType"`
			TachographTypeID   uint64 `yaml:"tachographTypeId"`
			TachographTypes    []struct {
				ID   int    `yaml:"id"`
				Name string `yaml:"name"`
			} `yaml:"tachographTypes"`
			ShowOccasionChoices bool   `yaml:"showOccasionChoices"`
			OccasionChoiceID    uint64 `yaml:"occasionChoiceId"`
			OccasionChoices     []struct {
				ID   uint64 `yaml:"id"`
				Name string `yaml:"name"`
			} `yaml:"occasionChoices"`
			ShowExaminationType bool   `yaml:"showExaminationType"`
			ExaminationTypeID   uint64 `yaml:"examinationTypeId"`
			ExaminationTypes    []struct {
				ID   uint64 `yaml:"id"`
				Name string `yaml:"name"`
			} `yaml:"examinationTypes"`
		} `yaml:"data"`
		Status int    `yaml:"status"`
		URL    string `yaml:"url"`
	}

	Location struct {
		ID      uint64 `yaml:"id"`
		Name    string `yaml:"name"`
		Address struct {
			StreetAddress1 string `yaml:"streetAddress1"`
			StreetAddress2 string `yaml:"streetAddress2"`
			ZipCode        string `yaml:"zipCode"`
			City           string `yaml:"city"`
			CareOf         string `yaml:"careOf"`
		} `yaml:"address"`
		Coordinates struct {
			Latitude  float64 `yaml:"latitude"`
			Longitude float64 `yaml:"longitude"`
		} `yaml:"coordinates"`
	}
)

// SearchInformation searches and returns different types of available information
// associated with the provided social security number, like available licence categories, exam locations etc.
func (tc *TrafikverketClient) SearchInformation(body *SearchInformationRequest) (*SearchInformationResponse, *http.Response, error) {
	// create request
	req, err := NewRequest("POST", "/search-information", &body)
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
	var resp SearchInformationResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, res, err
	}

	b, _ := json.Marshal(resp)
	log.Debugln(string(b))

	return &resp, res, nil
}

// Locations returns the available examination locations for the provided social security number
func (tc *TrafikverketClient) Locations(body *SearchInformationRequest) (*[]Location, *http.Response, error) {
	resp, res, err := tc.SearchInformation(body)
	if err != nil {
		return nil, res, err
	}

	return &resp.Data.Locations, res, nil
}
