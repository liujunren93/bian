package spot

import (
	"github.com/shopspring/decimal"
)

type FlexibleTradeRequest struct {
	ProductID     string          `json:"productId"`     // required
	Amount        decimal.Decimal `json:"amount"`        //required
	AutoSubscribe bool            `json:"autoSubscribe"` // default is true
	SourceAccount string          `json:"sourceAccount"` // 	SPOT,FUND,ALL, 默认 SPOT
}

type FlexibleTradeResponse struct {
	RedeemID int  `json:"redeemId"`
	Success  bool `json:"success"`
}

type FlexibleRedeemRequest struct {
	ProductId     string          `json:"productId"` // required
	RedeemAll     bool            `json:"redeemAll"` // 默认 false
	Amount        decimal.Decimal `json:"amount"`
	SourceAccount string          `json:"sourceAccount"` // 	SPOT,FUND,ALL, 默认 SPOT
}

type RedeemResponse struct {
	RedeemID int  `json:"redeemId"`
	Success  bool `json:"success"`
}

type LockedTradeRequest struct {
	ProjectId     string          `json:"projectId"`
	Amount        decimal.Decimal `json:"amount"`        //required
	AutoSubscribe bool            `json:"autoSubscribe"` // default is true
	SourceAccount string          `json:"sourceAccount"` // 	SPOT,FUND,ALL, 默认 SPOT
}

type LockedTradeResponse struct {
	PurchaseID int    `json:"purchaseId"`
	PositionID string `json:"positionId"`
	Success    bool   `json:"success"`
}

type LockedRedeemRequest struct {
	PositionID string `json:"positionId"`
}

type FlexibleUserDataRequest struct {
	Asset     string `json:"asset"`
	ProductID string `json:"productId"`
	Current   int    `json:"current"` //  当前查询页。 开始值 1，默认:1
	Size      int    `json:"size"`    //默认：10，最大：100
}

type FlexibleUserDataResponse struct {
	Rows  []FlexibleUserDataRows `json:"rows"`
	Total int                    `json:"total"`
}
type FlexibleUserDataTierAnnualPercentageRate struct {
	Zero5BTC  float64 `json:"0-5BTC"`
	Five10BTC float64 `json:"5-10BTC"`
}
type FlexibleUserDataRows struct {
	TotalAmount                    string                                   `json:"totalAmount"`
	TierAnnualPercentageRate       FlexibleUserDataTierAnnualPercentageRate `json:"tierAnnualPercentageRate"`
	LatestAnnualPercentageRate     string                                   `json:"latestAnnualPercentageRate"`
	YesterdayAirdropPercentageRate string                                   `json:"yesterdayAirdropPercentageRate"`
	Asset                          string                                   `json:"asset"`
	AirDropAsset                   string                                   `json:"airDropAsset"`
	CanRedeem                      bool                                     `json:"canRedeem"`
	CollateralAmount               string                                   `json:"collateralAmount"`
	ProductID                      string                                   `json:"productId"`
	YesterdayRealTimeRewards       string                                   `json:"yesterdayRealTimeRewards"`
	CumulativeBonusRewards         string                                   `json:"cumulativeBonusRewards"`
	CumulativeRealTimeRewards      string                                   `json:"cumulativeRealTimeRewards"`
	CumulativeTotalRewards         string                                   `json:"cumulativeTotalRewards"`
	AutoSubscribe                  bool                                     `json:"autoSubscribe"`
}

type LockedUserDataRequest struct {
	Asset     string `json:"asset"`
	ProductID string `json:"productId"`
	ProjectID string `json:"projectId"`
	Current   int    `json:"current"` //  当前查询页。 开始值 1，默认:1
	Size      int    `json:"size"`    //默认：10，最大：100
}

type LockedUserDataResponse struct {
	Rows  []LockedUserDataRows `json:"rows"`
	Total int                  `json:"total"`
}
type LockedUserDataReStakeInfo struct {
	ReStakeRate           string `json:"reStakeRate"`
	ReStakeAmount         string `json:"reStakeAmount"`
	ReStakeDuration       string `json:"reStakeDuration"`
	ReStakeApr            string `json:"reStakeApr"`
	EstRewards            string `json:"estRewards"`
	ReStakeRewardsEndDate string `json:"reStakeRewardsEndDate"`
	ReStakeDeliverDate    string `json:"reStakeDeliverDate"`
}
type LockedUserDataRows struct {
	PositionID            string                    `json:"positionId"`
	ProjectID             string                    `json:"projectId"`
	Asset                 string                    `json:"asset"`
	Amount                string                    `json:"amount"`
	PurchaseTime          string                    `json:"purchaseTime"`
	Duration              string                    `json:"duration"`
	AccrualDays           string                    `json:"accrualDays"`
	RewardAsset           string                    `json:"rewardAsset"`
	Apy                   string                    `json:"APY"`
	RewardAmt             string                    `json:"rewardAmt"`
	ExtraRewardAsset      string                    `json:"extraRewardAsset"`
	ExtraRewardAPR        string                    `json:"extraRewardAPR"`
	EstExtraRewardAmt     string                    `json:"estExtraRewardAmt"`
	NextPay               string                    `json:"nextPay"`
	NextPayDate           string                    `json:"nextPayDate"`
	PayPeriod             string                    `json:"payPeriod"`
	RedeemAmountEarly     string                    `json:"redeemAmountEarly"`
	RewardsEndDate        string                    `json:"rewardsEndDate"`
	DeliverDate           string                    `json:"deliverDate"`
	RedeemPeriod          string                    `json:"redeemPeriod"`
	RedeemingAmt          string                    `json:"redeemingAmt"`
	PartialAmtDeliverDate string                    `json:"partialAmtDeliverDate"`
	CanRedeemEarly        bool                      `json:"canRedeemEarly"`
	CanFastRedemption     bool                      `json:"canFastRedemption"`
	AutoSubscribe         bool                      `json:"autoSubscribe"`
	Type                  string                    `json:"type"`
	Status                string                    `json:"status"`
	CanReStake            bool                      `json:"canReStake"`
	ReStakeInfo           LockedUserDataReStakeInfo `json:"reStakeInfo"`
}
