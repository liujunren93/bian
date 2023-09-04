package websocket

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/liujunren93/bian/client"
	"github.com/liujunren93/bian/utils"

	"github.com/gorilla/websocket"
)

type Params struct {
	ID     interface{} `json:"id,omitempty"`
	Method string      `json:"method,omitempty"`
	Params interface{} `json:"params,omitempty"`
}

func (r Params) Hash() string {
	return fmt.Sprintf("%s_%v", r.Method, r.ID)
}

type WsClient struct {
	conf  client.Config
	wsmap map[string]*websocket.Conn
}

func NewClient(conf client.Config) *WsClient {
	var cli = WsClient{conf: conf}

	return &cli
}
func (cli *WsClient) Close(hash string) {
	fmt.Println("Close:", hash)
	fmt.Println("Close:", cli.wsmap)
	err := cli.wsmap[hash].Close()
	if err != nil {
		fmt.Println("ws close error:", err)
	}
}

func (cli *WsClient) newClient(path string, key string, header http.Header) (ws *websocket.Conn, hash string, err error) {
	// cli.wsmu.RLock()
	hash, _ = utils.Md5(key)
	if cli.wsmap == nil {
		cli.wsmap = make(map[string]*websocket.Conn)
	}
	if c, ok := cli.wsmap[hash]; ok {
		return c, hash, nil
	}
	if len(header) == 0 {
		header = http.Header{}
	}
	header.Add("X-MBX-APIKEY", cli.conf.ApiKey)
	if cli.conf.Proxy != "" {
		websocket.DefaultDialer.Proxy = func(r *http.Request) (*url.URL, error) {
			return url.Parse(cli.conf.Proxy)
		}
	}
	ws, res, err := websocket.DefaultDialer.Dial(cli.conf.BaseApi, header)
	if err != nil {
		fmt.Println(res)
		return nil, hash, err
	}
	cli.wsmap[hash] = ws
	return ws, hash, nil
}

func (cli *WsClient) SendSign(path string, header http.Header, req Params) error {

	cc, _, err := cli.newClient(path, req.Hash(), header)
	if err != nil {
		return err
	}

	client.Sign(req.Params.(client.Signer), cli.conf.ApiSecret)
	return cc.WriteJSON(req)
}

func (cli *WsClient) Pongs() {

}

func (cli *WsClient) Send(path string, header http.Header, req Params) error {
	if len(path) == 0 {
		path = req.Method
	}
	cc, _, err := cli.newClient(path, path, header)
	if err != nil {
		return err
	}
	return cc.WriteJSON(&req)
}

func (cli *WsClient) Receiver(path string, f func([]byte)) (done chan struct{}, err error) {
	cc, hash, err := cli.newClient(path, path, nil)
	if err != nil {
		return nil, err
	}
	done = make(chan struct{})
	go func(hash string) {
		for {
			_, data, err := cc.ReadMessage()
			if err != nil {
				fmt.Println("Receiver:", err)
				done <- struct{}{}
				close(done)
				cli.Close(hash)
				return
			}
			f(data)
		}
	}(hash)
	return done, nil

}
