package test

import (
	"coinbase-adv/client"
	"fmt"
	"os"
	"testing"
)

func Test_GetPrice(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	currency := "BTC-USD"

	cln := client.NewClient(&creds)

	bPr, err := cln.GetPrice(currency, "BUY")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	sPr, err := cln.GetPrice(currency, "SELL")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	fmt.Printf("%s,  b = %f,  s = %f", currency, *bPr, *sPr)

}

func Test_GetQuote(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	currency := "BTC-USD"

	cln := client.NewClient(&creds)

	q, err := cln.GetQuote(currency)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	fmt.Printf("%s,  b = %f,  s = %f", currency, q.Buy, q.Sell)

}
