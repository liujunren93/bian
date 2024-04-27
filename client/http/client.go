package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/liujunren93/bian/client"
	"github.com/liujunren93/bian/entitys"
)

type Config struct {
	BaseURL   string
	ApiKey    string
	SecretKey string
}
type Client struct {
	cfg    Config
	client *http.Client
}

func NewClient(conf Config) *Client {
	return &Client{cfg: conf, client: http.DefaultClient}
}

func (c *Client) doParams(path string, queryPramas, bodyPramas any) (urlQuery string, red io.Reader, err error) {

	if queryPramas == nil && bodyPramas == nil {
		queryPramas = entitys.RequestParam{}
	}

	queryStr, queryBody, err := client.Sign(queryPramas, bodyPramas, c.cfg.SecretKey)
	if err != nil {
		return
	}
	redbuf := &bytes.Buffer{}

	if queryBody != nil {
		data, err := json.Marshal(queryBody)
		if err != nil {
			return "", nil, err
		}
		redbuf.Write(data)
		red = redbuf
	}
	if queryPramas != nil {
		urlQuery += path + "?" + queryStr
	}

	return
}

func (c *Client) Post(path string, header http.Header, queryPramas, bodyPramas any) (*http.Response, error) {
	urlQuery, body, err := c.doParams(path, queryPramas, bodyPramas)

	req, err := http.NewRequest(http.MethodPost, c.cfg.BaseURL+urlQuery, body)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}
	req.Header.Set("X-MBX-APIKEY", c.cfg.ApiKey)
	return c.client.Do(req)
}

func (c *Client) Get(path string, header http.Header, queryPramas, bodyPramas any) (*http.Response, error) {
	urlQuery, body, err := c.doParams(path, queryPramas, bodyPramas)

	req, err := http.NewRequest(http.MethodGet, c.cfg.BaseURL+"/"+urlQuery, body)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}

	req.Header.Set("X-MBX-APIKEY", c.cfg.ApiKey)
	return c.client.Do(req)
}
