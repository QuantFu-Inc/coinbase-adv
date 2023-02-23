package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/QuantFu-Inc/coinbase-adv/client"
)

func Test_GetProduct(t *testing.T) {
	ctx := context.Background()
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	product := "BTC-USD"

	cln := client.NewClient(&creds)

	rsp, err := cln.GetProduct(ctx, product)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	fmt.Printf("%s,  price = %f, vol = %f,  status = %s", product, *rsp.Price, *rsp.Volume24h, *rsp.Status)

}
