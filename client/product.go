package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/QuantFu-Inc/coinbase-adv/model"
)

// GetProduct -- get product
func (c *Client) GetProduct(ctx context.Context, productId string) (*model.GetProductResponse, error) {
	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + fmt.Sprintf("/brokerage/products/%s", productId))
		response    model.GetProductResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err
}
