package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tomato3713/soumuradio"
)

func main() {
	soumuradio.DebugEnable()
	cli, err := soumuradio.NewClient(http.DefaultClient)
	if err != nil {
		os.Exit(1)
	}

	opts := soumuradio.NewLicenseListOptions(soumuradio.Amateur, false)
	result, err := cli.GetRadioList(nil, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range result {
		fmt.Printf("Station Name: %s\n\tLicense Date: %s\n", v.ListInfo.Name, v.ListInfo.LicenseDate)
	}
}
