package soumu

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
