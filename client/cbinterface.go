package client

import (
	"github.com/QuantFu-Inc/coinbase-adv/model"
	"net/http"
)

type CoinbaseClient interface {

	// coinbase adv api funcs
	ListAccounts(p *ListAccountsParams) (*model.ListAccountsResponse, error)
	GetAccount(uuid string) (*model.Account, error)
	ListFills(p *ListFillsParams) (*model.ListFillsResponse, error)
	GetOrder(id string) (*model.GetOrderResponse, error)
	CancelOrders(ids []string) (*model.CancelOrderResponse, error)
	CreateOrder(p *model.CreateOrderRequest) (*model.CreateOrderResponse, error)
	ListOrders(p *ListOrdersParams) (*model.ListOrdersResponse, error)
	GetPrice(currency string, side string) (*float64, error)
	GetQuote(currency string) (*Quote, error)
	GetExchangeRate(currency string) (*model.GetExchangeRateResponseData, error)
	GetProduct(productId string) (*model.GetProductResponse, error)

	// utility funcs
	CheckAuthentication(req *http.Request, body []byte)
	HttpClient() *http.Client
	IsTokenValid(timestamp int64) bool
	SetRateLimit(ms int64)
}
