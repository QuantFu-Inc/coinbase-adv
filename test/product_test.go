package test

import (
	"QuantFu-Inc/coinbase-adv/client"
	"fmt"
	"os"
	"testing"
)

func Test_GetProduct(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	product := "BTC-USD"

	cln := client.NewClient(&creds)

	rsp, err := cln.GetProduct(product)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	fmt.Printf("%s,  price = %f, vol = %f,  status = %s", product, *rsp.Price, *rsp.Volume24h, *rsp.Status)

}
