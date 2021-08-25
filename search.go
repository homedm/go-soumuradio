package soumuradio

import (
	"context"
	"time"
)

// GetRadioLicenseListSearchByDate get radio list from DF to DT.
func (c *Client) GetRadioLicenseListSearchByDate(ctx context.Context, opt ListOpts, DF, DT time.Time) (LicenseList, error) {
	opt.ST = License

	opt.DF = DF // start date
	opt.DT = DT // end date

	return c.GetRadioLicenseList(ctx, opt)
}

// GetRadioLicenseListSearchByCallsign get radio list searched by callsign
func (c *Client) GetRadioLicenseListSearchByCallsign(ctx context.Context, opt ListOpts, str string) (LicenseList, error) {
	opt.ST = License

	opt.MA = str // search target callsign

	return c.GetRadioLicenseList(ctx, opt)
}

// GetRadioLicenseListSearchByDate get radio list from DF to DT.
func (c *Client) GetRadioRegistrationListSearchByDate(ctx context.Context, opt ListOpts, DF, DT time.Time) (LicenseList, error) {
	opt.ST = Registration

	opt.DF = DF // start date
	opt.DT = DT // end date

	return c.GetRadioLicenseList(ctx, opt)
}

// GetRadioLicenseListSearchByCallsign get radio list searched by callsign
func (c *Client) GetRadioRegistrationListSearchByCallsign(ctx context.Context, opt ListOpts, str string) (LicenseList, error) {
	opt.ST = Registration

	opt.MA = str // search target callsign

	return c.GetRadioLicenseList(ctx, opt)
}
