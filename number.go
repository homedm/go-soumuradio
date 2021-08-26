package soumuradio

import (
	"context"
	"strconv"
	"time"

	"github.com/tomato3713/soumuradio/internal"
)

// getNumberAPI is Method for Number API
func (c *Client) getNumberAPI(ctx context.Context, opt NumOpts) (*internal.Num, error) {
	spath := "/musen/num"

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

	var num internal.Num
	if err := decodeBody(res, &num, nil); err != nil {
		return nil, err
	}

	return &num, nil
}

// GetTotalCount is Method for getting number of radio
func (c *Client) GetTotalCount(ctx context.Context, opt NumOpts) (int, error) {
	ret, err := c.getNumberAPI(ctx, opt)
	if err != nil {
		return 0, err
	}

	count, err := strconv.Atoi(ret.MusenInformation.TotalCount)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *Client) GetLastUpdateDate(ctx context.Context, opt NumOpts) (time.Time, error) {
	ret, err := c.getNumberAPI(ctx, opt)
	if err != nil {
		return time.Time{}, err
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, ret.MusenInformation.LastUpdateDate)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
