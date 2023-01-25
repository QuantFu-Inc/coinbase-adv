package test

import (
	"coinbase-adv/client"
	"fmt"
	"os"
	"testing"
)

func Test_ListAccounts(t *testing.T) {
	//devToken := "8e263c0bd4673380acf4b67956c1ebe5e048bbe7979e42888f1f372e03a03863"
	//
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	rsp, err := cln.ListAccounts(nil, nil)
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
			fmt.Printf("[%s] %f\n", a.GetCurrency(), a.AvailableBalance.GetValue())
		}
	}

	for rsp.GetHasNext() {

		rsp, err = cln.ListAccounts(nil, rsp.Cursor)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		for _, a := range rsp.Accounts {
			if a.AvailableBalance.GetValue() > 0 {
				fmt.Printf("[%s] %f\n", a.GetCurrency(), a.AvailableBalance.GetValue())
			}
		}
	}
}
