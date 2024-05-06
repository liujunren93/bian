package websocket

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/liujunren93/share_utils/client/websocket"
)

type Config struct {
	BaseURL   string
	ApiKey    string
	SecretKey string
}
type Client struct {
	cfg Config
}

type Msg struct {
	MsgType int
	Msg     []byte
	err     error
}

func NewClient(cfg Config) *Client {
	var client = Client{cfg: cfg}

	return &client
}

func (c *Client) Subscribe(ctx context.Context, path string, header http.Header, streams []string, collback func(*websocket.Msg, error)) error {

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
	cli, err := websocket.NewClient(u, websocket.WithHeader(header), websocket.WithPingInterval(time.Second*30))
	if err == nil {
		return err
	}
	cli.WriteMessage(websocket.TextMessage, buf)
	cli.ReadMessage(ctx, collback)
	return nil
}
