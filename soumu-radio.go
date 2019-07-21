package soumuradio

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
)

// Client is API Client structure
type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client

	Logger *log.Logger
}

// OF is indicates API Response to JSON format
const OF = 2

var (
	// ErrStatusCode is HTTP Response Status code error
	ErrStatusCode = errors.New("HTTP Response Status error")
)

// NewClient is constructer
// 必須の情報が与えられているか、期待するものかをチェックする
func NewClient(urlStr string, logger *log.Logger) (*Client, error) {
	if urlStr == "" {
		urlStr = "https://www.tele.soumu.go.jp"
	}

	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
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

var version = "0.1"
var usrAgent = fmt.Sprintf("SoumuRadioGoClient/%s (%s)", version, runtime.Version())

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, spath)

	fmt.Println(u.String())

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("User-Agent", usrAgent)

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}, f *os.File) error {
	defer resp.Body.Close()

	if f != nil {
		resp.Body = ioutil.NopCloser(io.TeeReader(resp.Body, f))
		defer f.Close()
	}

	decorder := json.NewDecoder(resp.Body)
	return decorder.Decode(out)
}
