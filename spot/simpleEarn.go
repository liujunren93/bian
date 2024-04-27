package spot

import (
	"github.com/liujunren93/bian/utils"
)

// FlexibleList 活期列表
// func (s *Spot) FlexibleList(pagisSize, offset int) {
// 	api := "/sapi/v1/simple-earn/flexible/list"
// 	res, err := s.apiClient().Get(api, nil, map[string]any{}, nil)
// 	data, err := io.ReadAll(res.Body)
// 	fmt.Println(string(data), err)
// }

// FlexibleUserData 获取活期产品持仓
func (s *Spot) FlexibleUserData(req FlexibleUserDataRequest) (response *FlexibleUserDataResponse, err error) {
	response = new(FlexibleUserDataResponse)
	api := "/sapi/v1/simple-earn/flexible/position"
	res, err := s.apiClient().Get(api, nil, req, nil)
	if err != nil {
		return
	}
	err = utils.ParseResult(res, &response)
	return
}

// FlexibleList 活期订购
func (s *Spot) FlexibleTrade(req FlexibleTradeRequest) (response *FlexibleTradeResponse, err error) {
	response = new(FlexibleTradeResponse)
	api := "/sapi/v1/simple-earn/flexible/subscribe"
	res, err := s.apiClient().Post(api, nil, req, nil)
	if err != nil {
		return
	}
	err = utils.ParseResult(res, &response)
	return
}

// FlexibleRedeem 活期赎回
func (s *Spot) FlexibleRedeem(req FlexibleRedeemRequest) (response *RedeemResponse, err error) {
	response = new(RedeemResponse)
	api := "/sapi/v1/simple-earn/flexible/redeem"
	res, err := s.apiClient().Post(api, nil, req, nil)
	if err != nil {
		return
	}
	err = utils.ParseResult(res, &response)
	return
}

// #########定期##########
// LockedUserData 用户定期数据
func (s *Spot) LockedUserData(req LockedUserDataRequest) (response *LockedUserDataResponse, err error) {
	response = new(LockedUserDataResponse)
	api := "/sapi/v1/simple-earn/locked/position"
	res, err := s.apiClient().Get(api, nil, req, nil)
	if err != nil {
		return
	}
	err = utils.ParseResult(res, &response)
	return
}

// LockedTrade 购买定期

// LockedTrade 定期购买
func (s *Spot) LockedTrade(req LockedTradeRequest) (response *LockedTradeResponse, err error) {
	response = new(LockedTradeResponse)
	api := "/sapi/v1/simple-earn/locked/subscribe"
	res, err := s.apiClient().Post(api, nil, req, nil)
	if err != nil {
		return
	}
	err = utils.ParseResult(res, &response)
	return
}

// LockedRedeem 定期赎回
func (s *Spot) LockedRedeem(req LockedRedeemRequest) (response *RedeemResponse, err error) {
	response = new(RedeemResponse)
	api := "/sapi/v1/simple-earn/locked/redeem"
	res, err := s.apiClient().Post(api, nil, req, nil)
	if err != nil {
		return
	}
	err = utils.ParseResult(res, &response)
	return
}
