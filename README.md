go-trafikverket
===============

A simple CLI tool and Go library for interacting with Trafikverket's Förarprov
API.

Usage
-----

### CLI

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ sh
go get -u github.com/mandrean/go-trafikverket
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Then just run `go-trafikverket` in your shell.

#### Commands

| **Commands**                           | **Alias(es)**         | **Description**         |
|----------------------------------------|-----------------------|-------------------------|
| go-trafikverket list                   | l                     |                         |
| **List Subcommands**                   |                       |                         |
| go-trafikverket list licenceCategories | licenseCategories, lc | List licence categories |
| go-trafikverket list locations         | l                     | List exam locations     |
| go-trafikverket list occasions         | o                     | List exam occasions     |

### Library

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ sh
go get -u github.com/mandrean/go-trafikverket
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Then import the package in your code:

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ go
package main

import (
    log
    "github.com/mandrean/go-trafikverket/pkg"
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
