package soumuradio

import "time"

// RadioSpec is the structure consisting of radio wave type, frequency, and
// antenna power.
type RadioSpec struct {
	RadioFormat []string
	Freq        string
	Power       float64
}

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
	RadioStationPurpose     string
	LicenseDate             time.Time
	CommMatter              string
	Note                    string      // Detail only
	Address                 string      // Detail only
	ValidTerms              time.Time   // Detail only
	RadioEquipmentLocation  string      // Detail only
	RadioSpec               []RadioSpec // Detail only
	PermittedOperatingHours string      // Detail only
	BroadcastMatter         string      // Detail only
	Office                  string      // Detail only
	CommPartner             string      // Detail only
	StartingLimit           string      // Detail only
	BroadcastDistrict       string      // Detail only
	WorkPersonName          string      // Detail only
	IdentificationSignals   string      // Detail only
	RadioStationNumber      string      // Detail only
	RadioStationCategory    string      // Detail only
	MovementArea            string      // Detail only
	LicenseNumber           string      // Detail only
}

type RegistrationList struct {
	RegistrationInfo []RegistrationInfo
	TotalCount       int
	LastUpdateDate   time.Time
}

type RegistrationInfo struct {
	No                     string
	Name                   string
	ValidTerms             string
	ItCd                   string
	RegistrationDate       string
	Note                   string // Detail only
	Address                string // Detail only
	RadioEquipmentLocation string // Detail only
	RadioEquipmentStandard string // Detail only
	RegistrationNumber     string // Detail only
	RadioSpec              string // Detail only
}
