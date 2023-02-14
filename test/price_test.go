package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/QuantFu-Inc/coinbase-adv/client"
)

func Test_GetPrice(t *testing.T) {
	ctx := context.Background()
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	currency := "BTC-USD"

	cln := client.NewClient(&creds)

	bPr, err := cln.GetPrice(ctx, currency, "BUY")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	sPr, err := cln.GetPrice(ctx, currency, "SELL")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	fmt.Printf("%s,  b = %f,  s = %f", currency, *bPr, *sPr)

}

func Test_GetQuote(t *testing.T) {
	ctx := context.Background()
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	currency := "BTC-USD"

	cln := client.NewClient(&creds)

	q, err := cln.GetQuote(ctx, currency)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	fmt.Printf("%s,  b = %f,  s = %f", currency, q.Buy, q.Sell)
}

func Test_GetExchangeRates(t *testing.T) {
	ctx := context.Background()
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	currency := "USD"

	cln := client.NewClient(&creds)

	exd, err := cln.GetExchangeRate(ctx, currency)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	for s, r := range exd.GetRates() {
		fmt.Printf("%s :  %s , %s\n", currency, s, r)
	}
}
