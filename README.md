trafikverket-forarprov
======================

A simple CLI tool and Go library for interacting with Trafikverket's Förarprov
API.

Usage
-----

### CLI

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ sh
go get -u github.com/mandrean/trafikverket-forarprov
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Then just run `trafikverket-forarprov` in your shell.

#### Commands

| **Commands**                                  | **Alias(es)**         | **Description**         |
|-----------------------------------------------|-----------------------|-------------------------|
| trafikverket-forarprov list                   | l                     |                         |
| **List Subcommands**                          |                       |                         |
| trafikverket-forarprov list licenceCategories | licenseCategories, lc | List licence categories |
| trafikverket-forarprov list locations         | l                     | List exam locations     |
| trafikverket-forarprov list occasions         | o                     | List exam occasions     |

### Library

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ sh
go get -u github.com/mandrean/trafikverket-forarprov
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Then import the package in your code:

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ go
package main

import (
    log
    "github.com/mandrean/trafikverket-forarprov/pkg"
)

func main() {
    // create client
    tc := pkg.NewClient()

    // fetch licence categories
    lcs, _, err := tc.LicenceCategories()
    if err != nil {
        log.Errorln(err)
        return
    }

    // print to console
    for _, lc := range *lcs {
        log.Println(lc.Name)
    }
}
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
