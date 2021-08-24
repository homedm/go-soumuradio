package soumuradio

import "time"

type LicenseList struct {
	LicenseInfo    []LicenseInfo
	TotalCount     int
	LastUpdateDate time.Time
}

// LicenseInfo is the structure of radio license information.
type LicenseInfo struct {
	No                      int
	Name                    string
	TdfkCd                  string
	Note                    string
	Address                 string
	ValidTerms              time.Time
	RadioEquipmentLocation  string
	RadioSpec               RadioSpec
	RadioStationPurpose     string
	PermittedOperatingHours string
	LicenseDate             time.Time
	BroadcastMatter         string
	CommMatter              string
	Office                  string
	CommPartner             string
	StartingLimit           string
	BroadcastDistrict       string
	WorkPersonName          string
	IdentificationSignals   string
	RadioStationNumber      string
	RadioStationCategory    string
	MovementArea            string
	LicenseNumber           string
}

// RadioSpec is the structure consisting of radio wave type, frequency, and
// antenna power.
type RadioSpec struct {
	RadioFormat string
	Freq        float64
	Power       float64
}
