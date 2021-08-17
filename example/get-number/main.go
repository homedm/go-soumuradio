package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tomato3713/go-soumuradio/soumu"
)

func main() {
	cli, err := soumu.NewClient(http.DefaultClient)
	if err != nil {
		os.Exit(1)
	}

	opts := soumu.NewNumOptions(soumu.License, soumu.Amateur)
	result, err := cli.GetTotalCount(nil, opts)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Amateur Radio TotalCount is %d\n", result)
}
