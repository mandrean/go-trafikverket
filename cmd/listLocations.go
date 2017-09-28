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
)

// locationsCmd represents the locations command
var locationsCmd = &cobra.Command{
	Use:     "locations",
	Aliases: []string{"l"},
	Short:   "List exam locations",
	Long:    ``,
	Run:     locations,
}

func init() {
	listCmd.AddCommand(locationsCmd)

	locationsCmd.Flags().StringVarP(&socialSecurityNumber, "social-security-number", "S", "", "(Required) Social security number")
	locationsCmd.Flags().IntVarP(&licenceID, "licence-id", "t", 5, "(Optional) License ID/type")
	locationsCmd.Flags().IntVarP(&bookingModeID, "booking-mode-id", "B", 0, "(Optional) Booking mode ID/type")
	locationsCmd.Flags().BoolVarP(&ignoreDebt, "ignore-debt", "I", false, "(Optional) Ignore debt")
	locationsCmd.Flags().IntVarP(&examinationTypeID, "examination-type-id", "E", 0, "(Optional) Examination ID/type")
}

func locations(cmd *cobra.Command, args []string) {
	// create client
	tc := pkg.NewClient()

	// check required flag
	if socialSecurityNumber == "" {
		log.Errorln("--social-security-number/-S is required!")
		return
	}

	// create payload
	body := &pkg.SearchInformationRequest{
		BookingSession: pkg.BookingSession{
			SocialSecurityNumber: socialSecurityNumber,
			LicenceID:            licenceID,
			BookingModeID:        bookingModeID,
			IgnoreDebt:           ignoreDebt,
			ExaminationTypeID:    examinationTypeID,
		},
	}

	// fetch locations
	ls, _, err := tc.Locations(body)
	if err != nil {
		log.Errorln(err)
		return
	}

	// print results
	switch Output {
	case "wide":
		printLocationsWide(ls)
		break
	case "json":
		printJSON(ls)
		break
	case "yaml":
		printYAML(ls)
	default:
		printLocationsWide(ls)
	}
}

func printLocationsWide(ls *[]pkg.Location) {
	table := uitable.New()
	table.MaxColWidth = 50

	table.AddRow("ID", "CITY", "STREET", "COORDINATES")
	for _, l := range *ls {
		c := fmt.Sprintf("%v, %v", l.Coordinates.Latitude, l.Coordinates.Longitude)
		table.AddRow(l.ID, l.Address.City, l.Address.StreetAddress1, c)
	}
	fmt.Println(table)
}
