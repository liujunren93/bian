package websocket

import (
	"fmt"
	"net/http"

	"github.com/liujunren93/bian/client"

	"github.com/gorilla/websocket"
)

type wsClient struct {
	conf  client.Config
	wsmap map[string]*websocket.Conn
}

func NewClient(conf client.Config) *wsClient {
	var cli = wsClient{conf: conf}

	return &cli
}

func (cli *wsClient) newClient(path string, header http.Header) (*websocket.Conn, error) {
	// cli.wsmu.RLock()
	if cli.wsmap == nil {
		cli.wsmap = make(map[string]*websocket.Conn)
	}
	if c, ok := cli.wsmap[path]; ok {
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
	cli.wsmap[path] = ws
	return ws, nil
}

func (cli *wsClient) SendSign(path string, header http.Header, data client.Params) error {
	if header == nil {
		header = http.Header{}
	}

	header.Add("X-MBX-APIKEY", cli.conf.ApiKey)
	cc, err := cli.newClient(path, header)
	if err != nil {
		return err
	}

	client.Sign(data, cli.conf.ApiSecret)
	return cc.WriteJSON(data)
}

func (cli *wsClient) Pongs() {

}

func (cli *wsClient) Send(method, path string, header http.Header, data interface{}) error {

	cc, err := cli.newClient(path, header)
	if err != nil {
		return err
	}

	var req = map[string]interface{}{
		"id":     1,
		"method": method,
	}
	if data != nil {
		req["params"] = data
	}
	return cc.WriteJSON(&req)
}

func (cli *wsClient) Receiver(path string, f func([]byte)) error {
	cc, err := cli.newClient(path, nil)
	if err != nil {
		return err
	}
	go func() {
		for {
			_, data, err := cc.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err) {
					fmt.Println("IsUnexpectedCloseError", err)
					return
				}
			}

			f(data)
		}
	}()
	return nil

}
