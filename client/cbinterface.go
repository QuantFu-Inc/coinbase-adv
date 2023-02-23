package client

import (
	"context"
	"net/http"

	"github.com/QuantFu-Inc/coinbase-adv/model"
)

type CoinbaseClient interface {

	// coinbase adv api funcs
	ListAccounts(ctx context.Context, p *ListAccountsParams) (*model.ListAccountsResponse, error)
	GetAccount(ctx context.Context, uuid string) (*model.Account, error)
	ListFills(ctx context.Context, p *ListFillsParams) (*model.ListFillsResponse, error)
	GetOrder(ctx context.Context, id string) (*model.GetOrderResponse, error)
	CancelOrders(ctx context.Context, ids []string) (*model.CancelOrderResponse, error)
	CreateOrder(ctx context.Context, p *model.CreateOrderRequest) (*model.CreateOrderResponse, error)
	ListOrders(ctx context.Context, p *ListOrdersParams) (*model.ListOrdersResponse, error)
	GetPrice(ctx context.Context, currency string, side string) (*float64, error)
	GetQuote(ctx context.Context, currency string) (*Quote, error)
	GetExchangeRate(ctx context.Context, currency string) (*model.GetExchangeRateResponseData, error)
	GetProduct(ctx context.Context, productId string) (*model.GetProductResponse, error)

	// utility funcs
	CheckAuthentication(req *http.Request, body []byte)
	HttpClient() *http.Client
	IsTokenValid(timestamp int64) bool
	SetRateLimit(ms int64)
}
