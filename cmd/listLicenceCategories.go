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

// licenceCategoriesCmd represents the licenceCategories command
var licenceCategoriesCmd = &cobra.Command{
	Use:     "licenceCategories",
	Aliases: []string{"licenseCategories", "lc"},
	Short:   "List licence categories",
	Long:    ``,
	Run:     licenceCategories,
}

func init() {
	listCmd.AddCommand(licenceCategoriesCmd)
}

func licenceCategories(cmd *cobra.Command, args []string) {
	// create client
	tc := pkg.NewClient()

	// fetch licence categories
	lcs, _, err := tc.LicenceCategories()
	if err != nil {
		log.Errorln(err)
		return
	}

	// print results
	switch Output {
	case "wide":
		printLicenceCategories(lcs)
		break
	case "json":
		printJSON(lcs)
		break
	case "yaml":
		printYAML(lcs)
	default:
		printLicenceCategories(lcs)
	}
}

func printLicenceCategories(lcs *[]pkg.LicenceCategory) {
	table := uitable.New()
	table.MaxColWidth = 50

	table.AddRow("CATEGORY", "TYPE", "DESCRIPTION")
	for _, lc := range *lcs {
		for _, l := range lc.Licences {
			table.AddRow(l.Category, l.Name, l.Description)
		}
	}
	fmt.Println(table)
}
