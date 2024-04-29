package spot

import (
	"context"
	"encoding/json"

	"github.com/liujunren93/bian/entitys"
)

func (s *Spot) SubscribeKline(ctx context.Context, streams []string, callback func(data entitys.Kine, err error)) {
	client := s.getWsClient()
	client.Subscribe("/stream", nil, streams...)
	client.Response(ctx, func(data []byte, err error) {
		if err != nil {
			callback(entitys.Kine{}, err)
		} else {
			var k entitys.Kine
			err := json.Unmarshal(data, &k)
			callback(k, err)
		}
	})

}
