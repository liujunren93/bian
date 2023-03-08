package websocket

import (
	"fmt"
	"net/http"

	"github.com/liujunren93/bian/client"

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

type wsClient struct {
	conf  client.Config
	wsmap map[string]*websocket.Conn
}

func NewClient(conf client.Config) *wsClient {
	var cli = wsClient{conf: conf}

	return &cli
}

func (cli *wsClient) newClient(path string, hash string, header http.Header) (*websocket.Conn, error) {
	// cli.wsmu.RLock()

	if cli.wsmap == nil {
		cli.wsmap = make(map[string]*websocket.Conn)
	}
	if c, ok := cli.wsmap[hash]; ok {
		return c, nil
	}
	if len(header) == 0 {
		header = http.Header{}
	}
	header.Add("X-MBX-APIKEY", cli.conf.ApiKey)
	ws, _, err := websocket.DefaultDialer.Dial(cli.conf.BaseApi+path, header)
	if err != nil {
		return nil, err
	}
	cli.wsmap[hash] = ws
	return ws, nil
}

func (cli *wsClient) SendSign(path string, header http.Header, req Params) error {

	cc, err := cli.newClient(path, req.Hash(), header)
	if err != nil {
		return err
	}

	client.Sign(req.Params.(client.Signer), cli.conf.ApiSecret)
	return cc.WriteJSON(req)
}

func (cli *wsClient) Pongs() {

}

func (cli *wsClient) Send(path string, header http.Header, req Params) error {
	if len(path) == 0 {
		path = req.Method
	}
	cc, err := cli.newClient(path, path, header)
	if err != nil {
		return err
	}

	return cc.WriteJSON(&req)
}

func (cli *wsClient) Receiver(path string, f func([]byte)) error {
	cc, err := cli.newClient(path, path, nil)
	if err != nil {
		return err
	}
	go func() {
		for {
			_, data, err := cc.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err) {
					return
				}
			}

			f(data)
		}
	}()
	return nil

}
