package soumu

import (
	"context"
	"net/url"
	"time"
)

// Num is API Response of getting number
type Num struct {
	Musen            Musen            `json:"musen"`
	MusenInformation MusenInformation `json:"musenInformation"`
}

// Musen is record information of JSON Response
type Musen struct {
	Count string `json:"count"`
}

// MusenInformation is header JSON Response
type MusenInformation struct {
	TotalCount     string `json:"totalCount"`
	LastUpdateDate string `json:"lastUpdateDate"`
}

// RadioStationType is radio station type option variable
type RadioStationType string

const (
	// Amateur is Amateur radio station
	Amateur RadioStationType = "AT"
	// FixedStation is Fixed Station
	FixedStation RadioStationType = "FX"
	// NonSpecificTerrestrialBroadcastingStation is Non specific Terrestrial Broadcasting Station
	NonSpecificTerrestrialBroadcastingStation RadioStationType = "BB"
	// SpecifiedGroundBroadcastStation is Specifed Ground Broadcasting Station
	SpecifiedGroundBroadcastStation RadioStationType = "BC"
	// NonSpecificTerrestrialBroadcastingTestStations is non specific terrestrial broadcasting test station
	NonSpecificTerrestrialBroadcastingTestStations RadioStationType = "BD"
	// SpecifiedTerrestrialBroadcastingTestStation is specified terrestrial broadcasting station
	SpecifiedTerrestrialBroadcastingTestStation RadioStationType = "BE"
	// TerrestrialGeneralBroadcastingStation is terrestrial general broadcasting station
	TerrestrialGeneralBroadcastingStation RadioStationType = "BG"
	// CoastStation is coast station.
	CoastStation RadioStationType = "FC"
	// AeronauticalStation is aeronautical station
	AeronauticalStation RadioStationType = "FA"
	// BaseStationExcludingPHS is base station excluding PHS
	BaseStationExcludingPHS RadioStationType = "FB"
	// BaseStationPHS is base station (PHS)
	BaseStationPHS RadioStationType = "FB_PHS"
	// PortableBaseStation is Portable Base Station
	PortableBaseStation RadioStationType = "FP"
	// RadioPagingStation is Radio Paging Station
	RadioPagingStation RadioStationType = "RP"
	// LandMobileRelayStation is Land Mobile Relay Station
	LandMobileRelayStation RadioStationType = "FBR"
	// ShipStation is Ship Station
	ShipStation RadioStationType = "MS"
	// SpecifiedShipStation is Specified Ship Station
	SpecifiedShipStation RadioStationType = "MSS"
	// DistressAutomaticReportStation is Distress Automatic Report Station
	DistressAutomaticReportStation RadioStationType = "DS"
	// OnBoardCommunicationStation is On Board Communication Station
	OnBoardCommunicationStation RadioStationType = "MB"
	// AircraftStation is Aircraft Station
	AircraftStation RadioStationType = "MA"
	// LandMobileStation is Land Mobile Station
	LandMobileStation RadioStationType = "ML"
	// PortableStation is Portable Station
	PortableStation RadioStationType = "MP"
	// MobileStation is Mobile Station
	MobileStation RadioStationType = "MO"
	// RadionavigationLandStation is Radio Navigation Land Station
	RadionavigationLandStation RadioStationType = "RL"
	// RadionavigationMobileStation is Radio Navigation Mobile Station
	RadionavigationMobileStation RadioStationType = "RO"
	// RadioLocationLandStation is Radio Location Land Station
	RadioLocationLandStation RadioStationType = "LR"
	// RadioLocationMobileStation is Radio Location Mobile Station
	RadioLocationMobileStation RadioStationType = "MR"
	// RadiobeaconStation is Radiobeacon Station
	RadiobeaconStation RadioStationType = "RB"
	// EarthStation is Earth Station
	EarthStation RadioStationType = "TC"
	// CoastEarthStation is Coast Earth Station
	CoastEarthStation RadioStationType = "TI"
	// AeronauticalEarthStation is Aeronautical Earth Station
	AeronauticalEarthStation RadioStationType = "TB"
	// PortableBaseEarthStation is Portable Base Earth Station
	PortableBaseEarthStation RadioStationType = "TYP"
	// ShipEarthStation is Ship Earth Station
	ShipEarthStation RadioStationType = "TG"
	// AircraftEarthStation is Aircraft Earth Station
	AircraftEarthStation RadioStationType = "TJ"
	// PortableMobileEarthStation is Portable Mobile Earth Station
	PortableMobileEarthStation RadioStationType = "TUP"
	// SpaceStation is Space Station
	SpaceStation RadioStationType = "ME"
	// ArtificialSatelliteStation is Artifical Satellite Station
	ArtificialSatelliteStation RadioStationType = "EKT"
	// SatelliteBasicBroadcasting is Satellite Basic Broadcasting
	SatelliteBasicBroadcasting RadioStationType = "EV"
	// SatelliteBasicBroadcastingExperimentalTestStation is Satellite Basic
	// Broadcasting Experimental Test Station
	SatelliteBasicBroadcastingExperimentalTestStation = "EBE"
	// EmergencyTrafficStation is Emergency Traffic Station
	EmergencyTrafficStation RadioStationType = "EM"
	// ExperimentalStation is Experimental Station
	ExperimentalStation RadioStationType = "EX"
	// SpecificExperimentalStation is Specific Experimental Station
	SpecificExperimentalStation RadioStationType = "EXT"
	// DevelopmentTestStation is Development Test Station
	DevelopmentTestStation RadioStationType = "DVT"
	// SimplicityRadioStationExcludingPersonal is Simplicity Radio Station
	// Excluding Personal Radio
	SimplicityRadioStationExcludingPersonal RadioStationType = "CR"
	// SimplicityRadioStationPersonal is Simplicity Radio Statoin Personal
	SimplicityRadioStationPersonal RadioStationType = "CR_PA"
	// PremisesRadioStation is Premises Radio Station
	PremisesRadioStation RadioStationType = "LO"
	// MeteorologicalAidsStation is Meteorological Aids Station
	MeteorologicalAidsStation RadioStationType = "SM"
	// StandardFrequencyStation is Standard Frequency Station
	StandardFrequencyStation RadioStationType = "SS"
	// StationInTheSpecialService is Station In The Special Service Station
	StationInTheSpecialService RadioStationType = "SP"
	// FixedStationBlanketLicense is Fixed Station (Blanket License)
	FixedStationBlanketLicense RadioStationType = "FX_H"
	// LandMobileStationBlanketLicense is Land Mobile Station (Blanket License)
	LandMobileStationBlanketLicense RadioStationType = "ML_H"
	// PortableStationBlanketLicense is Portable Station (Blanket License)
	PortableStationBlanketLicense RadioStationType = "MP_H"
	// EarthStationBlanketLicense is Earth Station (Blanket License)
	EarthStationBlanketLicense RadioStationType = "TC_H"
	// PortableMobileEarthStationBlanketLicense is Portable Mobile Earth Staion (Blanket License)
	PortableMobileEarthStationBlanketLicense RadioStationType = "TUP_H"
	// BaseStationExcludingBlanketLicense is Base Station Excluding (Blanket License)
	BaseStationExcludingBlanketLicense RadioStationType = "FB_H"
)

// ComprehensiveCommunicationStation is Comprehensive Communication Station
type ComprehensiveCommunicationStation string

const (
	// Hokkaido is Hokkaido Comprehensive Communication Station
	Hokkaido ComprehensiveCommunicationStation = "J"
	// Tohoku is Tohoku Comprehensive Communication Station
	Tohoku ComprehensiveCommunicationStation = "I"
	// Kanto is Kanto Comprehensive Communication Station
	Kanto ComprehensiveCommunicationStation = "A"
	// Shinetu is Shinetu Comprehensive Communication Station
	Shinetu ComprehensiveCommunicationStation = "B"
	// Hokuriku is Hokuriku Comprehensive Communication Station
	Hokuriku ComprehensiveCommunicationStation = "D"
	// Tokai is Tokai Comprehensive Communication Station
	Tokai ComprehensiveCommunicationStation = "C"
	// Kinki is Kinki Comprehensive Communication Station
	Kinki ComprehensiveCommunicationStation = "E"
	// Chugoku is Chugoku Comprehensive Communication Station
	Chugoku ComprehensiveCommunicationStation = "F"
	// Shikoku is Shikoku Comprehensive Communication Station
	Shikoku ComprehensiveCommunicationStation = "G"
	// Kyushu is Kyushu Comprehensive Communication Station
	Kyushu ComprehensiveCommunicationStation = "H"
	// Okinawa is Okinawa Comprehensive Communication Station
	Okinawa ComprehensiveCommunicationStation = "O"
)

// NumOpts is Options of Number API
type NumOpts struct {
	// Start count
	ST int
	// Type of radio station
	OW RadioStationType
	// only ST=2, Planning of radio station equipment
	OW2 string
	// Comprehensive communication station
	IT ComprehensiveCommunicationStation
	// only ST=1, Prefecture / city
	HCV string
	// aim of radio
	MK string
	// only ST=1, Communication matters
	TSNJK string
	// only ST=1, Types of basic broadcasting
	KHS string
	// frequency (start)
	// kHz: Integer part 10 digits, decimal part 6 digits or less
	// MHz: Integer part 7 digits, decimal part 9 digits or less
	// GHz: Integer part 4 digits, decimal part 12 digits or less
	FF string
	// frequency (final)
	TF string
	// Unit, kHz, MHz, GHz
	HZ string
	// name
	// max length is 100 characters
	NA string
	// only ST=1, call sign
	MA string
	// only ST=1, Ship aircraft name
	AS string
	// licensed date (start)
	DF time.Time
	// licensed date (final)
	DT time.Time
}

func (opt NumOpts) encodeOption() url.Values {
	params := url.Values{}
	params.Add("ST", string(opt.ST))
	params.Add("OW", string(opt.OW))
	params.Add("IT", string(opt.IT))
	params.Add("MK", opt.MK)
	params.Add("FF", opt.FF)
	params.Add("TF", opt.TF)
	params.Add("HZ", opt.HZ)
	params.Add("NA", opt.NA)
	params.Add("DT", opt.DT.Format("20060102"))
	params.Add("DF", opt.DF.Format("20060102")) // only ST == 1 options
	if opt.ST == 1 {
		params.Add("HCV", opt.HCV)
		params.Add("TSNJK", opt.TSNJK)
		params.Add("KHS", opt.KHS)
		params.Add("MA", opt.MA)
		params.Add("AS", opt.AS)
	}
	if opt.ST == 2 {
		// only ST=2, Planning of radio station equipment
		params.Add("OW2", opt.OW2)
	}
	return params
}

// GetNum is Method for getting number
func (c *Client) GetNum(ctx context.Context, opt NumOpts) (*Num, error) {
	spath := "/musen/num"

	req, err := c.newRequest(ctx, "GET", spath, opt, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// check status code
	if res.StatusCode < 200 || 300 <= res.StatusCode {
		return nil, ErrStatusCode
	}

	var num Num
	if err := decodeBody(res, &num, nil); err != nil {
		return nil, err
	}

	return &num, nil
}
