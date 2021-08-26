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
	result, err := cli.GetRadioLicenseListSearchByCallsign(nil, opts, str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range result.LicenseInfo {
		fmt.Printf("ListInfo:\n")
		fmt.Printf(" No: %+v\n", v.No)
		fmt.Printf(" Name: %+v\n", v.Name)
		fmt.Printf(" TdfkCd: %+v\n", v.TdfkCd)
		fmt.Printf(" Radio Station Purpose: %+v\n", v.RadioStationPurpose)
		fmt.Printf(" LicenseDate: %+v\n", v.LicenseDate)
		// 		fmt.Printf(" ItCd: %+v\n", v.ItCd)
		// 		fmt.Printf(" Registration Date: %+v\n", v.RegistrationDate)
		// 		fmt.Printf(" Valid Terms: %+v\n", v.ValidTerms)
		fmt.Printf("DetailInfo:\n")
		fmt.Printf("  Name: %+v\n", v.Name)
		fmt.Printf("  Address: %+v\n", v.Address)
		fmt.Printf("  Office: %+v\n", v.Office)
		fmt.Printf("  RadioStationCategory: %+v\n", v.RadioStationCategory)
		fmt.Printf("  LicenseNumber: %+v\n", v.LicenseNumber)
		fmt.Printf("  LicenseDate: %+v\n", v.LicenseDate)
		fmt.Printf("  ValidTerms: %+v\n", v.ValidTerms)
		fmt.Printf("  RadioStationPurpose: %+v\n", v.RadioStationPurpose)
		fmt.Printf("  RadioStationNumber: %+v\n", v.RadioStationNumber)
		fmt.Printf("  PermittedOperatingHours: %+v\n", v.PermittedOperatingHours)
		fmt.Printf("  StartingLimit: %+v\n", v.StartingLimit)
		fmt.Printf("  BroadcastMatter: %+v\n", v.BroadcastMatter)
		fmt.Printf("  BroadcastDistrict: %+v\n", v.BroadcastDistrict)
		fmt.Printf("  CommMatter: %+v\n", v.CommMatter)
		fmt.Printf("  CommPartner: %+v\n", v.CommPartner)
		fmt.Printf("  IdentificationSignals: %+v\n", v.IdentificationSignals)
		fmt.Printf("  RadioEquipmentLocation: %+v\n", v.RadioEquipmentLocation)
		fmt.Printf("  MovementArea: %+v\n", v.MovementArea)
		fmt.Printf("  RadioSpec\n")
		for _, rs := range v.RadioSpec {
			fmt.Printf("  - %q, %v, %vW\n", rs.RadioFormat, rs.Freq, rs.Power)
		}
		fmt.Printf("  WorkPersonName: %+v\n", v.WorkPersonName)
		fmt.Printf("  Note: %+v\n", v.Note)
		// 		fmt.Printf("  RadioEquipmentStandard: %+v\n", v.RadioEquipmentStandard)
		// 		fmt.Printf("  RegistrationDate: %+v\n", v.RegistrationDate)
		// 		fmt.Printf("  RegistrationNumber: %+v\n", v.RegistrationNumber)
	}
}
