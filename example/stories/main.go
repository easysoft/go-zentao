package main

import (
	"log"

	"github.com/easysoft/go-zentao"
)

func main() {
	zt, err := zentao.NewBasicAuthClient(
		"admin",
		"jaege1ugh4ooYip7",
		zentao.WithBaseURL("https://zentao-easysoft.cloud.okteto.net"),
		zentao.WithDevMode(),
		zentao.WithDumpAll(),
		zentao.WithoutProxy(),
	)
	if err != nil {
		log.Fatal(err)
	}
	pds, _, err := zt.Products.List()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Products count: %v", len(pds.Products))
	pgs, _, err := zt.Programs.List("")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Programs count: %v", len(pgs.Programs))
}
