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

// NumOpts is Options of Number API
type NumOpts struct {
	// Start count
	ST string
	// Type of radio station
	OW string
	// only ST=2, Planning of radio station equipment
	OW2 string
	// Comprehensive communication station
	IT string
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
	params.Add("ST", opt.ST)
	params.Add("OW", opt.OW)
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
