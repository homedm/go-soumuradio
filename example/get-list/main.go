package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tomato3713/soumuradio"
	"github.com/tomato3713/soumuradio/radiotype"
)

func main() {
	soumuradio.DebugEnable()
	cli, err := soumuradio.NewClient(http.DefaultClient)
	if err != nil {
		os.Exit(1)
	}

	opts := soumuradio.NewLicenseListOptions(radiotype.Amateur, false)
	result, err := cli.GetListAPI(nil, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v", *result)
}
