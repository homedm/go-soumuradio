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

// Num is API Response for getting number
type Num struct {
	MusenInformation MusenInformation `json:"musenInformation"`
	Musen            Musen            `json:"musen"`
}

// MusenInformation is header JSON Response
type MusenInformation struct {
	TotalCount     string `json:"totalCount"`
	LastUpdateDate string `json:"lastUpdateDate"`
}

// Lists is the struct of List acquistion API
type Lists struct {
	Musen            []Musen          `json:"musen"`
	MusenInformation MusenInformation `json:"musenInformation"`
}

// Musen is the common struct
type Musen struct {
	// for Number acquisition API
	Count string `json:"count"`
	// for List acquisition API
	ListInfo   ListInfo   `json:"listInfo"`
	DetailInfo DetailInfo `json:"detailInfo"`
}

// ListInfo is the struct for List API.
type ListInfo struct {
	No   string `json:"no"`
	Name string `json:"name"`

	// only lisense information
	RadioStationPurpose string `json:"radioStationPurpose"`
	TdfkCd              string `json:"tdfkCd"`
	LicenseDate         string `json:"licenseDate"`

	// only registration information
	ValidTerms       string `json:"validTerms"`
	ItCd             string `json:"itCd"`
	RegistrationDate string `json:"registrationDate"`
}

// DetailInfo is the struct for List acquistion API
type DetailInfo struct {
	Name                   string `json:"name"`
	Note                   string `json:"note"`
	Address                string `json:"address"`
	ValidTerms             string `json:"validTerms"`
	RadioEquipmentLocation string `json:"radioEquipmentLocation"`

	// only lisense information
	RadioSpec1              string `json:"radioSpec1"`
	RadioStationPurpose     string `json:"radioStationPurpose"`
	PermittedOperatingHours string `json:"permittedOperatingHours"`
	LicenseDate             string `json:"licenseDate"`
	BroadcastMatter         string `json:"broadcastMatter"`
	CommMatter              string `json:"commMatter"`
	Office                  string `json:"office"`
	CommPartner             string `json:"commPartner"`
	StartingLimit           string `json:"startingLimit"`
	BroadcastDistrict       string `json:"broadcastDistrict"`
	WorkPersonName          string `json:"workPersonName"`
	IdentificationSignals   string `json:"identificationSignals"`
	RadioStationNumber      string `json:"radioStationNumber"`
	RadioStationCategory    string `json:"radioStationCategory"`
	MovementArea            string `json:"movementArea"`
	LicenseNumber           string `json:"licenseNumber"`

	// only registration information
	RegistrationDate       string `json:"registrationDate"`
	RadioEquipmentStandard string `json:"radioEquipmentStandard"`
	RegistrationNumber     string `json:"registrationNumber"`
	RadioSpec2             string `json:"radioSpec2"`
}

// RadioSpec is the structure consisting of radio wave type, frequency, and
// antenna power.
type RadioSpec struct {
	RadioFormat string
	Freq        float64
	Power       float64
}
