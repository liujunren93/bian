package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Config struct {
	BaseURL   string
	ApiKey    string
	SecretKey string
}
type Client struct {
	cfg     Config
	conn    *websocket.Conn
	header  http.Header
	path    string
	streams []string
}

func NewClient(cfg Config) *Client {
	var client = Client{cfg: cfg}

	return &client
}

func (c *Client) Subscribe(path string, header http.Header, streams ...string) error {
	c.path = path
	c.header = header
	c.streams = streams
	data := map[string]interface{}{
		"method": "SUBSCRIBE",
		"id":     time.Now().Unix(),
		"params": streams,
	}
	buf, _ := json.Marshal(data)
	var u = c.cfg.BaseURL + path + "?streams="
	for _, stream := range streams {
		u += stream + "/"
	}
	u = strings.TrimRight(u, "/")

	conn, _, err := websocket.DefaultDialer.Dial(u, header)
	if err != nil {
		return err
	}
	conn.WriteMessage(websocket.TextMessage, buf)
	c.conn = conn
	return nil
}

func (c *Client) Unsubscribe(params ...string) error {
	panic("")
}

func (c *Client) Response(ctx context.Context, f func(data []byte, err error)) {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				err := c.conn.WriteMessage(websocket.CloseMessage, nil)
				if err != nil {
					f(nil, err)
				}
				return
			case <-ticker.C:
				err := c.conn.WriteMessage(websocket.PingMessage, nil)
				if err != nil {
					f(nil, err)
				}
			default:
				msgType, data, err := c.conn.ReadMessage()
				fmt.Println(msgType, string(data), err)
				if err != nil {
					f(nil, err)
					return
				}
				switch msgType {
				case websocket.TextMessage:
					f(data, err)
				case websocket.CloseMessage:
					f(data, err)
					return
				case websocket.PingMessage:
					err := c.conn.WriteMessage(websocket.PongMessage, nil)
					if err != nil {
						f(nil, err)
					}
				}

			}

		}
	}()

}
