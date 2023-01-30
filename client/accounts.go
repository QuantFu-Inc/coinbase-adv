package client

import (
	"fmt"
	"net/url"
	"quantfu.com/coinbase-adv/model"
)

type ListAccountsParams struct {
	Limit  *int32
	Cursor *string
}

// ListAccounts -- list user's accounts
func (c *Client) ListAccounts(p *ListAccountsParams) (*model.ListAccountsResponse, error) {
	var (
		u, _        = url.Parse(CoinbaseAdv_V3Endpoint + "/brokerage/accounts")
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

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err
}

type GetAccountResponse struct {
	Account model.Account
}

// GetAccount -- get account details
func (c *Client) GetAccount(uuid string) (*model.Account, error) {
	var (
		u, _        = url.Parse(CoinbaseAdv_V3Endpoint + fmt.Sprintf("/brokerage/accounts/%s", uuid))
		response    GetAccountResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response.Account, err
}
