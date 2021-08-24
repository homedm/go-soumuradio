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
	Note                    string    // Detail only
	Address                 string    // Detail only
	ValidTerms              time.Time // Detail only
	RadioEquipmentLocation  string    // Detail only
	RadioSpec               RadioSpec // Detail only
	RadioStationPurpose     string
	PermittedOperatingHours string // Detail only
	LicenseDate             time.Time
	BroadcastMatter         string // Detail only
	CommMatter              string
	Office                  string // Detail only
	CommPartner             string // Detail only
	StartingLimit           string // Detail only
	BroadcastDistrict       string // Detail only
	WorkPersonName          string // Detail only
	IdentificationSignals   string // Detail only
	RadioStationNumber      string // Detail only
	RadioStationCategory    string // Detail only
	MovementArea            string // Detail only
	LicenseNumber           string // Detail only
}

// RadioSpec is the structure consisting of radio wave type, frequency, and
// antenna power.
type RadioSpec struct {
	RadioFormat string
	Freq        float64
	Power       float64
}
