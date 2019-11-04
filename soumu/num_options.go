package soumu

import (
	"net/url"
	"strconv"
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

// NumOpts is Options of Number API
type NumOpts struct {
	// start count
	SC int
	// DC is data count
	DC int
	// ST is station target
	ST TargetInfo
	// OW is Type of radio station
	OW RadioStationType
	// OW2 is only ST=2, Planning of radio station equipment
	// TODO: make structure
	OW2 string
	// Comprehensive communication station
	IT ComprehensiveCommunicationStation
	// only ST=1, Prefecture / city
	// Reference: http://www.soumu.go.jp/main_content/000632830.pdf
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
	// SK is Sorting Key
	SK SortingKey
}

// NewNumOptions returns an initialized NumOpts structure.
func NewNumOptions(st TargetInfo, ow RadioStationType) NumOpts {
	opts := NumOpts{}
	opts.ST = st
	opts.OW = ow
	opts.SC = 1
	opts.DC = 3
	return opts
}

func (opt NumOpts) encodeOption() url.Values {
	params := url.Values{}
	params.Add("ST", strconv.Itoa(int(opt.ST)))
	if opt.OW != "" {
		params.Add("OW", string(opt.OW))
	}
	if opt.IT != "" {
		params.Add("IT", string(opt.IT))
	}
	if opt.MK != "" {
		params.Add("MK", opt.MK)
	}
	if opt.FF != "" {
		params.Add("FF", opt.FF)
	}
	if opt.TF != "" {
		params.Add("TF", opt.TF)
	}
	if opt.HZ != "" {
		params.Add("HZ", opt.HZ)
	}
	if opt.NA != "" {
		params.Add("NA", opt.NA)
	}
	if !opt.DT.IsZero() {
		params.Add("DT", opt.DT.Format("20060102"))
	}
	if !opt.DF.IsZero() {
		params.Add("DF", opt.DF.Format("20060102")) // only ST == 1 options
	}
	if opt.ST == 1 {
		if opt.HCV != "" {
			params.Add("HCV", opt.HCV)
		}
		if opt.TSNJK != "" {
			params.Add("TSNJK", opt.TSNJK)
		}
		if opt.KHS != "" {
			params.Add("KHS", opt.KHS)
		}
		if opt.MA != "" {
			params.Add("MA", opt.MA)
		}
		if opt.AS != "" {
			params.Add("AS", opt.AS)
		}
	}
	if opt.ST == 2 {
		// only ST=2, Planning of radio station equipment
		if opt.OW2 != "" {
			params.Add("OW2", opt.OW2)
		}
	}
	return params
}
