package soumuradio

import (
	"context"

	"github.com/tomato3713/soumuradio/internal"
)

// GetListAPI is Method for getting number
func (c *Client) GetListAPI(ctx context.Context, opt ListOpts) (*internal.Lists, error) {
	spath := "/musen/list"

	req, err := c.newRequest(ctx, "GET", spath, opt, nil)
	if err != nil {
		logf("try URL: %v", req.URL)
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err = checkStatusCode(res); err != nil {
		return nil, err
	}

	var lists internal.Lists
	if err := decodeBody(res, &lists, nil); err != nil {
		return nil, err
	}

	return &lists, nil
}

func (c *Client) GetRadioLicenseList(ctx context.Context, opt ListOpts) (LicenseList, error) {
	opt.ST = 1
	result, err := c.GetListAPI(ctx, opt)
	if err != nil {
		return LicenseList{}, err
	}

	ret, err := convertInternalLists2LicenseList(result)
	if err != nil {
		return LicenseList{}, err
	}

	return ret, nil
}
