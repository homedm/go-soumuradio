package soumuradio

import (
	"context"
	"time"
)

// GetRadioListWithDate get radio list from DF to DT.
func (c *Client) GetRadioListWithDate(ctx context.Context, opt ListOpts, DF, DT time.Time) ([]Musen, error) {
	opt.DF = DF
	opt.DT = DT

	ret, err := c.GetListAPI(ctx, opt)
	if err != nil {
		return nil, err
	}

	return ret.Musen, nil
}

// GetRadioWithCallsign get radio list searched by callsign
func (c *Client) GetRadioWithCallsign(ctx context.Context, opt ListOpts, str string) ([]Musen, error) {
	opt.MA = str

	ret, err := c.GetListAPI(ctx, opt)
	if err != nil {
		return nil, err
	}

	return ret.Musen, nil
}
