package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/QuantFu-Inc/coinbase-adv/model"
)

type ListAccountsParams struct {
	Limit  *int32
	Cursor *string
}

// ListAccounts -- list user's accounts
func (c *Client) ListAccounts(ctx context.Context, p *ListAccountsParams) (*model.ListAccountsResponse, error) {
	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + "/brokerage/accounts")
		response    model.ListAccountsResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)
	if p != nil {
		if p.Limit != nil {
			c.addInt32Param(queryParams, "limit", *p.Limit)
		}
		if p.Cursor != nil {
			c.addStringParam(queryParams, "cursor", *p.Cursor)
		}
	}

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err
}

type GetAccountResponse struct {
	Account model.Account
}

// GetAccount -- get account details
func (c *Client) GetAccount(ctx context.Context, uuid string) (*model.Account, error) {
	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + fmt.Sprintf("/brokerage/accounts/%s", uuid))
		response    GetAccountResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response.Account, err
}
