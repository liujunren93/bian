package spot

import "github.com/liujunren93/bian/utils"

func (s *Spot) OrderTrade(req OrderTradeRequest) (rsp *OrderBaseResponse, err error) {
	api := "/api/v3/order"
	res, err := s.apiClient().Post(api, nil, req, nil)
	if err != nil {
		return
	}
	rsp = new(OrderBaseResponse)

	err = utils.ParseResult(res, rsp)
	return
}
