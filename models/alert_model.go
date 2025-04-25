package models

type AlertPayload struct {
	Alert   string `json:"alert"`
	Ticker  string `json:"ticker"`
	NeoCloud struct {
		Lead float64 `json:"lead"`
		Lag  float64 `json:"lag"`
	} `json:"neo_cloud"`
	Ohlcv struct {
		Open   float64 `json:"open"`
		High   float64 `json:"high"`
		Low    float64 `json:"low"`
		Close  float64 `json:"close"`
		Volume float64 `json:"volume"`
	} `json:"ohlcv"`
	BarTime int64 `json:"bartime"`
}
