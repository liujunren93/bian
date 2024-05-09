package spot

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/liujunren93/bian/entitys"
	"github.com/liujunren93/share_utils/client/websocket"
)

func (s *Spot) SubscribeKline(ctx context.Context, streams []string, callback func(data *entitys.Kine, err error)) error {
	client := s.getWsClient()
	err := client.Subscribe(ctx, "/stream", nil, streams, func(m *websocket.Msg, err error) {
		if err != nil {
			callback(nil, err)
			return
		}
		if m.Msg == nil {
			callback(nil, err)
			return
		}
		var k entitys.KlineStreamResponse
		err = json.Unmarshal(m.Msg, &k)
		if k.Stream != "" {
			callback(&k.Data.Kline, err)
		}

	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
