package soumuradio

import (
	"context"
	"fmt"
	"net/url"
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
	ST string
	OW string
}

// GetNum is Method for getting number
func (c *Client) GetNum(ctx context.Context, opt NumOpts) (*Num, error) {
	spath := fmt.Sprintf("/musen/num")

	params := url.Values{}
	params.Add("ST", opt.ST)
	params.Add("OW", opt.OW)
	params.Add("OF", OF)

	req, err := c.newRequest(ctx, "GET", spath, params, nil)
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
