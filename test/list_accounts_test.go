package test

import (
	"fmt"
	"github.com/QuantFu-Inc/coinbase-adv/client"
	"os"
	"testing"
)

func Test_ListAccounts(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	rsp, err := cln.ListAccounts(nil)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	// should be at least 1
	if len(rsp.Accounts) < 1 {
		t.FailNow()
	}

	for _, a := range rsp.Accounts {
		if a.AvailableBalance.GetValue() > 0 {
			fmt.Printf("[%s] [%s] %f\n", a.GetUuid(), a.GetCurrency(), a.AvailableBalance.GetValue())
		}
	}

	for rsp.GetHasNext() {

		p := client.ListAccountsParams{
			Cursor: rsp.Cursor,
		}

		rsp, err = cln.ListAccounts(&p)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		for _, a := range rsp.Accounts {
			if a.AvailableBalance.GetValue() > 0 {
				fmt.Printf("[%s] [%s] %f\n", a.GetUuid(), a.GetCurrency(), a.AvailableBalance.GetValue())
			}
		}
	}
}
