package soumuradio

import (
	"context"
	"strconv"
	"time"

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

	var ret LicenseList

	var licenses []LicenseInfo
	for _, e := range result.Musen {
		var l LicenseInfo
		l.No, err = strconv.Atoi(e.ListInfo.No)
		if err != nil {
			return LicenseList{}, err
		}
		l.Name = e.ListInfo.Name
		l.RadioStationPurpose = e.ListInfo.RadioStationPurpose
		l.TdfkCd = e.ListInfo.TdfkCd
		l.LicenseDate, err = time.Parse("2006-01-02", e.ListInfo.LicenseDate)

		l.Note = e.DetailInfo.Note
		l.Address = e.DetailInfo.Address
		l.ValidTerms, err = time.Parse("2006-01-02まで", e.DetailInfo.ValidTerms)
		if err != nil {
			return LicenseList{}, err
		}
		l.RadioEquipmentLocation = e.DetailInfo.RadioEquipmentLocation
		if e.DetailInfo.RadioSpec1 != "" {
			l.RadioSpec, err = parseRadioSpec(e.DetailInfo.RadioSpec1)
			if err != nil {
				return LicenseList{}, err
			}
		}
		l.PermittedOperatingHours = e.DetailInfo.PermittedOperatingHours
		l.BroadcastMatter = e.DetailInfo.BroadcastMatter
		l.CommMatter = e.DetailInfo.CommMatter
		l.Office = e.DetailInfo.Office
		l.CommPartner = e.DetailInfo.CommPartner
		l.StartingLimit = e.DetailInfo.StartingLimit
		l.BroadcastDistrict = e.DetailInfo.BroadcastDistrict
		l.WorkPersonName = e.DetailInfo.WorkPersonName
		l.IdentificationSignals = e.DetailInfo.IdentificationSignals
		l.RadioStationNumber = e.DetailInfo.RadioStationNumber
		l.RadioStationCategory = e.DetailInfo.RadioStationCategory
		l.MovementArea = e.DetailInfo.MovementArea
		l.LicenseNumber = e.DetailInfo.LicenseNumber

		licenses = append(licenses, l)
	}
	ret.LicenseInfo = licenses

	ret.LastUpdateDate, err = time.Parse("2006-01-02", result.MusenInformation.LastUpdateDate)
	if err != nil {
		return LicenseList{}, err
	}
	ret.TotalCount, err = strconv.Atoi(result.MusenInformation.TotalCount)
	if err != nil {
		return LicenseList{}, err
	}

	return ret, nil
}
