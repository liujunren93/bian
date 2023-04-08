package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/liujunren93/bian/client"
)

type HttpClient struct {
	conf client.Config
}

func NewClient(conf client.Config) *HttpClient {
	return &HttpClient{conf: conf}
}

func (c *HttpClient) do(sign bool, path, method string, header http.Header, queryData, postData client.Signer, dest interface{}) error {
	if header == nil {
		header = http.Header{}
	}
	header.Add("X-MBX-APIKEY", c.conf.ApiKey)
	path = c.conf.BaseApi + path
	if sign {
		if queryData == nil {
			queryData = client.QueryParams{}
		}
		client.Sign(queryData, c.conf.ApiSecret)
	}

	if queryData.Len() > 0 {
		path += "?" + queryData.String()
	}
	var red io.Reader
	if postData != nil && postData.Len() > 0 {
		if sign {
			client.Sign(postData, c.conf.ApiSecret)
		}

		buf, err := json.Marshal(postData)
		if err != nil {
			return err
		}
		red = bytes.NewReader(buf)
	}

	req, err := http.NewRequest(method, path, red)
	if err != nil {
		return err
	}
	req.Header = header
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(path)
	fmt.Println(string(resData))
	err = json.Unmarshal(resData, dest)
	if err != nil {
		return err
	}
	return nil
}

func (c *HttpClient) Post(path string, header http.Header, queryData, postData client.Signer, dest interface{}) error {

	return c.do(true, path, "POST", header, queryData, postData, dest)

}

func (c *HttpClient) Put(path string, header http.Header, queryData, postData client.Signer, dest interface{}) error {
	return c.do(true, path, "PUT", header, queryData, postData, dest)
}

func (c *HttpClient) PutNoSign(path string, header http.Header, queryData, postData client.Signer, dest interface{}) error {
	return c.do(false, path, "PUT", header, queryData, postData, dest)
}

func (c *HttpClient) Get(path string, header http.Header, data client.QueryParams, dest interface{}) error {
	return c.do(true, path, "GET", header, data, nil, dest)
}
