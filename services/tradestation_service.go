package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/neothedev/trading-webhook/configs"
)

var (
	config = configs.LoadConfig(os.Getenv("APP_ENV"))
	token  = os.Getenv("TSTOKEN") // Your OAuth2 access token
)

func getRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s%s", config.TradeStationBaseURL, config.TradeStationAPIRoot, path)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func GetAccounts() ([]byte, error) {
	return getRequest(config.TradeStationEndpoints.Accounts)
}

func GetBalances(accountId string) ([]byte, error) {
	return getRequest(fmt.Sprintf(config.TradeStationEndpoints.Balances, accountId))
}

func GetPositions(accountId string) ([]byte, error) {
	return getRequest(fmt.Sprintf(config.TradeStationEndpoints.Positions, accountId))
}

func GetOrders(accountId string) ([]byte, error) {
	return getRequest(fmt.Sprintf(config.TradeStationEndpoints.Orders, accountId))
}

func GetQuote(symbol string) ([]byte, error) {
	return getRequest(fmt.Sprintf(config.TradeStationEndpoints.Quote, symbol))
}

func GetBarChart(symbol string) ([]byte, error) {
	return getRequest(fmt.Sprintf(config.TradeStationEndpoints.BarChart, symbol))
}

func GetOptionChain(symbol string) ([]byte, error) {
	return getRequest(fmt.Sprintf(config.TradeStationEndpoints.OptionChain, symbol))
}

func GetOptionQuotes(symbol string) ([]byte, error) {
	return getRequest(fmt.Sprintf(config.TradeStationEndpoints.OptionQuotes, symbol))
}

func GetMarketDepth(symbol string) ([]byte, error) {
	return getRequest(fmt.Sprintf(config.TradeStationEndpoints.MarketDepth, symbol))
}
