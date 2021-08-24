// This example print radio station information which searched by str.
// Example:
//   go run main.go [str]
//   go run main.go jj1
//   go run main.go JJ1HGP

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tomato3713/soumuradio"
	"github.com/tomato3713/soumuradio/radiotype"
)

func main() {
	str := os.Args[1]

	cli, err := soumuradio.NewClient(http.DefaultClient)
	if err != nil {
		os.Exit(1)
	}

	opts := soumuradio.NewLicenseListOptions(radiotype.Amateur, true)
	result, err := cli.GetRadioWithCallsign(nil, opts, str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range result {
		fmt.Printf("ListInfo:\n")
		fmt.Printf(" No: %+v\n", v.ListInfo.No)
		fmt.Printf(" Name: %+v\n", v.ListInfo.Name)
		fmt.Printf(" TdfkCd: %+v\n", v.ListInfo.TdfkCd)
		fmt.Printf(" Radio Station Purpose: %+v\n", v.ListInfo.RadioStationPurpose)
		fmt.Printf(" LicenseDate: %+v\n", v.ListInfo.LicenseDate)
		// 		fmt.Printf(" ItCd: %+v\n", v.ListInfo.ItCd)
		// 		fmt.Printf(" Registration Date: %+v\n", v.ListInfo.RegistrationDate)
		// 		fmt.Printf(" Valid Terms: %+v\n", v.ListInfo.ValidTerms)
		fmt.Printf("DetailInfo:\n")
		fmt.Printf("  Name: %+v\n", v.DetailInfo.Name)
		fmt.Printf("  Address: %+v\n", v.DetailInfo.Address)
		fmt.Printf("  Office: %+v\n", v.DetailInfo.Office)
		fmt.Printf("  RadioStationCategory: %+v\n", v.DetailInfo.RadioStationCategory)
		fmt.Printf("  LicenseNumber: %+v\n", v.DetailInfo.LicenseNumber)
		fmt.Printf("  LicenseDate: %+v\n", v.DetailInfo.LicenseDate)
		fmt.Printf("  ValidTerms: %+v\n", v.DetailInfo.ValidTerms)
		fmt.Printf("  RadioStationPurpose: %+v\n", v.DetailInfo.RadioStationPurpose)
		fmt.Printf("  RadioStationNumber: %+v\n", v.DetailInfo.RadioStationNumber)
		fmt.Printf("  PermittedOperatingHours: %+v\n", v.DetailInfo.PermittedOperatingHours)
		fmt.Printf("  StartingLimit: %+v\n", v.DetailInfo.StartingLimit)
		fmt.Printf("  BroadcastMatter: %+v\n", v.DetailInfo.BroadcastMatter)
		fmt.Printf("  BroadcastDistrict: %+v\n", v.DetailInfo.BroadcastDistrict)
		fmt.Printf("  CommMatter: %+v\n", v.DetailInfo.CommMatter)
		fmt.Printf("  CommPartner: %+v\n", v.DetailInfo.CommPartner)
		fmt.Printf("  IdentificationSignals: %+v\n", v.DetailInfo.IdentificationSignals)
		fmt.Printf("  RadioEquipmentLocation: %+v\n", v.DetailInfo.RadioEquipmentLocation)
		fmt.Printf("  MovementArea: %+v\n", v.DetailInfo.MovementArea)
		fmt.Printf("  RadioSpec1: %+v\n", v.DetailInfo.RadioSpec1)
		// 		fmt.Printf("  RadioSpec2: %+v\n", v.DetailInfo.RadioSpec2)
		fmt.Printf("  WorkPersonName: %+v\n", v.DetailInfo.WorkPersonName)
		fmt.Printf("  Note: %+v\n", v.DetailInfo.Note)
		// 		fmt.Printf("  RadioEquipmentStandard: %+v\n", v.DetailInfo.RadioEquipmentStandard)
		// 		fmt.Printf("  RegistrationDate: %+v\n", v.DetailInfo.RegistrationDate)
		// 		fmt.Printf("  RegistrationNumber: %+v\n", v.DetailInfo.RegistrationNumber)
	}
}
