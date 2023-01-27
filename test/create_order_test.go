package test

import (
	"coinbase-adv/client"
	"coinbase-adv/model"
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_CreateMarketOrder(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	cloid := fmt.Sprintf("unit-test-cmo-%d", time.Now().UnixMilli())
	prod := "DOGE-USD"
	side := string(model.BUY)
	quoteSz := "5"

	p := &model.CreateOrderRequest{
		ClientOrderId: &cloid,
		ProductId:     &prod,
		Side:          &side,
		OrderConfiguration: &model.CreateOrderRequestOrderConfiguration{
			MarketMarketIoc: &model.CreateOrderRequestOrderConfigurationMarketMarketIoc{
				QuoteSize: &quoteSz,
			},
		},
	}

	rsp, err := cln.CreateOrder(p)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	if *rsp.Success {

		// list fills.
		pf := &client.ListFillsParams{
			OrderId:   *rsp.OrderId,
			ProductId: *rsp.SuccessResponse.ProductId,
		}

		frsp, err := cln.ListFills(pf)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		for _, f := range frsp.Fills {

			fmt.Printf("[%s] [%s] %f  @ %f \n", f.GetProductId(), f.GetSide(), f.GetSize(), f.GetPrice())

		}

	}
}

func Test_CreateLimitOrder(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	cloid := fmt.Sprintf("unit-test-cmo-%d", time.Now().UnixMilli())
	prod := "ETH-USD"
	side := string(model.BUY)
	baseSz := "0.01"
	lmt := "1500"

	p := &model.CreateOrderRequest{
		ClientOrderId: &cloid,
		ProductId:     &prod,
		Side:          &side,
		OrderConfiguration: &model.CreateOrderRequestOrderConfiguration{
			LimitLimitGtc: &model.CreateOrderRequestOrderConfigurationLimitLimitGtc{
				LimitPrice: &lmt,
				BaseSize:   &baseSz,
			},
		},
	}

	rsp, err := cln.CreateOrder(p)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	if *rsp.Success {

		p := client.ListOrdersParams{
			OrderStatus: []string{string(model.OPEN)},
		}

		rsp, err := cln.ListOrders(&p)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		for _, o := range rsp.Orders {
			fmt.Printf("[%s] [%s] [%s] %f  @ %f \n", o.GetStatus(), o.GetProductId(), o.GetSide(), o.GetFilledSize(), o.GetAverageFilledPrice())
		}

	}

}
