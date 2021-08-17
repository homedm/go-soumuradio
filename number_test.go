package soumuradio

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestGetTotalCount(t *testing.T) {
	r, err := recorder.New("fixtures/soumuradio")
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

	opts := NewNumOptions(License, Amateur)
	_, err = cli.GetTotalCount(context.Background(), opts)
	if err != nil {
		t.Fatalf("fault to GetTotalCount returned err: %v", err)
	}
}

func TestGetLastUpdateDate(t *testing.T) {
	r, err := recorder.New("fixtures/soumuradio")
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

	opts := NewNumOptions(License, Amateur)
	_, err = cli.GetLastUpdateDate(context.Background(), opts)
	if err != nil {
		t.Fatalf("fault to GetLastUpdateDate returned err: %v", err)
	}
}
