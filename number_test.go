package soumuradio

import (
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/tomato3713/soumuradio/radiotype"
)

func TestGetTotalCount(t *testing.T) {
	r, err := recorder.New("fixtures/number01")
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

	opts := NewNumOptions(License, radiotype.Amateur)
	_, err = cli.GetTotalCount(nil, opts)
	if err != nil {
		t.Fatalf("fault to GetTotalCount returned err: %v", err)
	}
}

func TestGetLastUpdateDate(t *testing.T) {
	r, err := recorder.New("fixtures/number02")
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

	opts := NewNumOptions(License, radiotype.Amateur)
	_, err = cli.GetLastUpdateDate(nil, opts)
	if err != nil {
		t.Fatalf("fault to GetLastUpdateDate returned err: %v", err)
	}
}
