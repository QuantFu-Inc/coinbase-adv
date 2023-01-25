package client

import (
	"coinbase-adv/model"
	"net/url"
)

// ListAccounts -- list user's accounts
func (c *Client) ListAccounts(limit *int32, cursor *string) (*model.ListAccountsResponse, error) {
	var (
		u, _        = url.Parse(CoinbaseAdv_Endpoint + "/brokerage/accounts")
		response    model.ListAccountsResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	if limit != nil {
		c.addInt32Param(queryParams, "limit", *limit)
	}
	if cursor != nil {
		c.addStringParam(queryParams, "cursor", *cursor)
	}

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return &response, err
	}
	return &response, err
}
