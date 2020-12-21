package binance

import (
	"context"
	"encoding/json"
)

// AllCoinsInformationService get coin info
type AllCoinsInformationService struct {
	c *Client
}

// Do send request
func (s *AllCoinsInformationService) Do(ctx context.Context, opts ...RequestOption) (res []*CoinInfo, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/config/getall",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*CoinInfo{}, err
	}
	res = make([]*CoinInfo, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CoinInfo{}, err
	}
	return res, nil
}

// CoinInfo define coin info
type CoinInfo struct {
	Coin              string         `json:"coin"`
	DepositAllEnable  bool           `json:"depositAllEnable"`
	Name              string         `json:"name"`
	NetworkList       []NetworkEntry `json:"networkList"`
	WithdrawAllEnable bool           `json:"withdrawAllEnable"`
}

type NetworkEntry struct {
	Coin                    string `json:"coin"`
	DepositEnable           bool   `json:"depositEnable"`
	Name                    string `json:"name"`
	Network                 string `json:"network"`
	SpecialTips             string `json:"specialTips"`
	WithdrawEnable          bool   `json:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawMin             string `json:"withdrawMin"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
}
