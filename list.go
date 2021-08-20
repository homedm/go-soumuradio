package soumuradio

import (
	"context"
)

// GetListAPI is Method for getting number
func (c *Client) GetListAPI(ctx context.Context, opt ListOpts) (*Lists, error) {
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

	var lists Lists
	if err := decodeBody(res, &lists, nil); err != nil {
		return nil, err
	}

	return &lists, nil
}
