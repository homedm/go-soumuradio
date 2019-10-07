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

type RadioStationType string

const (
	// アマチュア局
	Amateur RadioStationType = "AT"
	// 固定局 (個別免許)
	FixedStation RadioStationType = "FX"
	// 特定以外の地上基幹放送局
	NonSpecificTerrestrialBroadcastingStation RadioStationType = "BB"
	// 特定地上基幹放送試験局
	SpecifiedGroundBroadcastStation RadioStationType = "BC"
	// 特定以外の地上基幹放送試験局
	NonSpecificTerrestrialBroadcastingTestStations RadioStationType = "BD"
	// 特定地上基幹放送試験局
	SpecifiedTerrestrialBroadcastingTestStation RadioStationType = "BE"
	// 地上一般放送局
	TerrestrialGeneralBroadcastingStation RadioStationType = "BG"
	// 海岸局
	CoastStation RadioStationType = "FC"
	// 航空局
	AeronauticalStation RadioStationType = "FA"
	// 基地局 (PHS 除く)
	BaseStationExcludingPHS RadioStationType = "FB"
	// 基地局 (PHS)
	BaseStationPHS RadioStationType = "FB_PHS"
	// 携帯基地局
	PortableBaseStation RadioStationType = "FP"
	// 無線呼出局
	RadioPagingStation RadioStationType = "RP"
	// 陸上移動中継局
	LandMobileRelayStation RadioStationType = "FBR"
	// 船舶局
	ShipStation RadioStationType = "MS"
	// 特定船舶局
	SpecifiedShipStation RadioStationType = "MSS"
	// 遭難自動通報局
	DistressAutomaticReportStation RadioStationType = "DS"
	// 船上通信局
	OnBoardCommunicationStation RadioStationType = "MB"
	// 航空機局
	AircraftStation RadioStationType = "MA"
	// 陸上移動局
	LandMobileStation RadioStationType = "ML"
	// 携帯局
	PortableStation RadioStationType = "MP"
	// 移動局
	MobileStation RadioStationType = "MO"
	// 無線航行陸上局
	RadionavigationLandStation RadioStationType = "RL"
	// 無線航行移動局
	RadionavigationMobileStation RadioStationType = "RO"
	// 無線標定陸上局
	RadiolocationLandStation RadioStationType = "LR"
	// 無線標定移動局
	RadiolocationMobileStation RadioStationType = "MR"
	// 無線標識局
	RadiobeaconStation RadioStationType = "RB"
	// 地球局
	EarthStation RadioStationType = "TC"
	// 海岸地球局
	CoastEarthStation RadioStationType = "TI"
	// 航空地球局
	AeronauticalEarthStation RadioStationType = "TB"
	// 携帯基地地球局
	PortableBaseEarthStation RadioStationType = "TYP"
	// 船舶地球局
	ShipEarthStation RadioStationType = "TG"
	// 航空機地球局
	AircraftEarthStation RadioStationType = "TJ"
	// 携帯移動地球局
	PortableMobileEarthStation RadioStationType = "TUP"
	// 宇宙局
	SpaceStation RadioStationType = "ME"
	// 人工衛星局
	ArtificialSatelliteStation RadioStationType = "EKT"
	// 衛星基幹放送局
	SatelliteBasicBroadcasting RadioStationType = "EV"
	// 衛星基幹放送試験局
	SatelliteBasicBroadcastingExperimentalTestStation = "EBE"
	// 非常局
	EmergencyTrafficStation RadioStationType = "EM"
	// 実験試験局
	ExperimentalStation RadioStationType = "EX"
	// 特定実験試験局
	SpecificExperimentalStation RadioStationType = "EXT"
	// 実用化試験局 DVT
	DevelopmentTestStation RadioStationType = "DVT"
	// 簡易無線局 (パーソナル無線を除く簡易無線局)
	SimplicityRadioStationExcludingPersonal RadioStationType = "CR"
	// 簡易無線局(パーソナル無線) CR_PA
	SimplicityRadioStationPersonal RadioStationType = "CR"
	// 構内無線局
	PremisesRadioStation RadioStationType = "LO"
	// 気象援助局
	MeteorologicalAidsStation RadioStationType = "SM"
	// 標準周波数局
	StandardFrequencyStation RadioStationType = "SS"
	// 特別業務の局
	StationInTheSpecialService RadioStationType = "SP"
	// 固定局(包括免許)
	FixedStationBlanketLicense RadioStationType = "FX_H"
	// 陸上移動局(包括免許)
	LandMobileStationBlanketLicense RadioStationType = "ML_H"
	// 携帯局(包括免許)
	PortableStationBlanketLicense RadioStationType = "MP_H"
	// 地球局(包括免許)
	EarthStationBlanketLicense RadioStationType = "TC_H"
	// 携帯移動地球局(包括免許)
	PortableMobileEarthStationBlanketLicense RadioStationType = "TUP_H"
	// 基地局(包括免許) FB_H
	BaseStationExcludingBlanketLicense RadioStationType = "FB_H"
)

// 所轄総合通信局
type ComprehensiveCommunicationStation string

const (
	Hokkaido ComprehensiveCommunicationStation = "J"
	Tohoku   ComprehensiveCommunicationStation = "I"
	Kanto    ComprehensiveCommunicationStation = "A"
	Shinetu  ComprehensiveCommunicationStation = "B"
	Hokuriku ComprehensiveCommunicationStation = "D"
	Tokai    ComprehensiveCommunicationStation = "C"
	Kinki    ComprehensiveCommunicationStation = "E"
	Chugoku  ComprehensiveCommunicationStation = "F"
	Shikoku  ComprehensiveCommunicationStation = "G"
	Kyushu   ComprehensiveCommunicationStation = "H"
	Okinawa  ComprehensiveCommunicationStation = "O"
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
