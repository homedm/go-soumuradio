package soumuradio

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/tomato3713/soumuradio/radiotype"
)

func TestGetRadioLicenseList(t *testing.T) {
	r, err := recorder.New("fixtures/list1")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Stop()

	cfg := &http.Client{
		Transport: r,
	}

	cli, err := NewClient(cfg)
	if err != nil {
		t.Fatalf("fault to make new client: %v", err)
	}

	opts := NewLicenseListOptions(radiotype.Amateur, true)
	_, err = cli.GetRadioLicenseList(nil, opts)
	if err != nil {
		t.Fatalf("fault to GetRadioLicenseList returned err: %v", err)
	}
}
