package soumuradio

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

// Client is API Client structure
type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client

	Logger *log.Logger
}

// NewClient is constructer
// 必須の情報が与えられているか、期待するものかをチェックする
func NewClient(urlStr, logger *log.Logger) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", urlStr)
	}

	// 必須で無いので、設定されていなければデフォルト値を設定
	if logger == nil {
		logger = log.New(os.Stderr, "[Log]", log.LstdFlags)
	}

	return &Client{
		BaseURL:    parsedURL,
		HTTPClient: http.DefaultClient,
		Logger:     logger,
	}, nil
}
