package main

import (
	"fmt"
	"os"

	"github.com/tomato3713/go-soumuradio/soumu"
)

func main() {
	soumu.DebugEnable()
	cli, err := soumu.NewClient("")
	if err != nil {
		os.Exit(1)
	}

	opts := soumu.NewLicenseListOptions(soumu.License, soumu.Amateur, false)
	result, err := cli.GetList(nil, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v", *result)
}
