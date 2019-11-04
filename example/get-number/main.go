package main

import (
	"fmt"
	"os"

	"github.com/tomato3713/go-soumuradio/soumu"
)

func main() {
	cli, err := soumu.NewClient("")
	if err != nil {
		os.Exit(1)
	}

	opts := soumu.NewNumOptions(soumu.License, soumu.Amateur)
	result, err := cli.GetNum(nil, opts)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("%+v", *result)
}
