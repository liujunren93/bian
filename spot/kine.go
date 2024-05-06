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
		if m.Msg != nil {
			callback(nil, m.Err)
			return
		}
		var k entitys.Kine
		err = json.Unmarshal(m.Msg, &k)
		callback(&k, err)

	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
