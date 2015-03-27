package main

import (
	"flag"
	"github.com/fcheslack/crossref"
	"github.com/k0kubun/pp"
	"log"
)

func main() {
	flag.Parse()
	doi := flag.Arg(0)
	if doi == "" {
		log.Fatal("No DOI specified")
	}
	work, err := crossref.GetWork(doi)
	if err != nil {
		log.Fatal(err)
	}
	pp.Print(work)
}
