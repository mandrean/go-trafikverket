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

package cmd

import (
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/mandrean/trafikverket-forarprov/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

// occasionsCmd represents the locations command
var occasionsCmd = &cobra.Command{
	Use:     "occasions",
	Aliases: []string{"o"},
	Short:   "List exam occasions",
	Long:    ``,
	Run:     occasions,
}

func init() {
	listCmd.AddCommand(occasionsCmd)

	occasionsCmd.Flags().StringVarP(&socialSecurityNumber, "social-security-number", "S", "", "(Required) Social security number")
	occasionsCmd.Flags().IntVarP(&licenceID, "licence-id", "t", 5, "(Optional) License ID/type")
	occasionsCmd.Flags().IntVarP(&bookingModeID, "booking-mode-id", "B", 0, "(Optional) Booking mode ID/type")
	occasionsCmd.Flags().BoolVarP(&ignoreDebt, "ignore-debt", "I", false, "(Optional) Ignore debt")

	occasionsCmd.Flags().StringVarP(&startDate, "start-date", "D", "", "(Optional) Start date")
	occasionsCmd.Flags().IntVarP(&locationID, "location-id", "L", 0, "(Required) Location ID")
	occasionsCmd.Flags().IntVarP(&languageID, "language-id", "l", 13, "(Optional) Language ID")
	occasionsCmd.Flags().IntVarP(&vehicleTypeID, "vehicle-type-id", "V", 1, "(Optional) Vehicle type ID")
	occasionsCmd.Flags().IntVarP(&tachographTypeID, "tachograph-type-id", "T", 1, "(Optional) Tachograph type ID")
	occasionsCmd.Flags().IntVarP(&occasionChoiceID, "occasion-choice-id", "O", 1, "(Optional) Occasion choice ID")
	occasionsCmd.Flags().IntVarP(&examinationTypeID, "examination-type-id", "E", 0, "(Optional) Examination type ID")
}

func occasions(cmd *cobra.Command, args []string) {
	// create client
	tc := pkg.NewClient()

	// check required flags
	missing := false
	if socialSecurityNumber == "" {
		log.Errorln("--social-security-number/-S is required!")
		missing = true
	}
	if locationID == 0 {
		log.Errorln("--location-id/-L is required!")
		missing = true
	}
	if missing {
		return
	}

	// create payload
	t, _ := time.Parse(time.RFC3339, startDate)
	body := &pkg.OccasionBundlesRequest{
		BookingSession: pkg.BookingSession{
			SocialSecurityNumber: socialSecurityNumber,
			LicenceID:            licenceID,
			BookingModeID:        bookingModeID,
			IgnoreDebt:           ignoreDebt,
			ExaminationTypeID:    examinationTypeID,
		},
		OccasionBundleQuery: pkg.OccasionBundleQuery{
			StartDate:         t,
			LocationID:        locationID,
			LanguageID:        languageID,
			VehicleTypeID:     vehicleTypeID,
			TachographTypeID:  tachographTypeID,
			OccasionChoiceID:  occasionChoiceID,
			ExaminationTypeID: examinationTypeID,
		},
	}

	// fetch occasions
	os, _, err := tc.Occasions(body)
	if err != nil {
		log.Errorln(err)
		return
	}

	// print results
	switch Output {
	case "wide":
		printOccasionsWide(os)
		break
	case "json":
		printJSON(os)
		break
	case "yaml":
		printYAML(os)
	default:
		printOccasionsWide(os)
	}
}

func printOccasionsWide(os *[]pkg.Occasion) {
	table := uitable.New()
	table.MaxColWidth = 50

	table.AddRow("NAME", "TYPE", "DATE", "TIME", "COST")
	for _, o := range *os {
		table.AddRow(o.LocationName, o.Name, o.Date, o.Time, o.Cost+o.CostText)
	}
	fmt.Println(table)
}
