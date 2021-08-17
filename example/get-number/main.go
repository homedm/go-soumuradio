package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tomato3713/soumuradio"
)

func main() {
	cli, err := soumuradio.NewClient(http.DefaultClient)
	if err != nil {
		os.Exit(1)
	}

	opts := soumuradio.NewNumOptions(soumuradio.License, soumuradio.Amateur)
	result, err := cli.GetTotalCount(nil, opts)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Amateur Radio TotalCount is %d\n", result)
}
