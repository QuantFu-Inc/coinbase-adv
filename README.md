# go-coinbase-adv

### Go implementation for Coinbase Advanced Trading API 

  - See /openapi for v3 specification of the models. This is used for model generation only.
  - See /client for actual API calls. 
  - See /test for example calls to APIs

  - Supports authentication via API key or via OAuth token.


```
// ListAccounts Example (/test/list_accounts_test.go)

//devToken := os.Getenv("CB-ACTOKEN")
//creds := client.Credentials{AccessToken: devToken}

creds := client.Credentials{
  ApiKey:      os.Getenv("CB-APIKEY"),
  ApiSKey:     os.Getenv("CB-SKEY"),
  AccessToken: os.Getenv("CB-ACTOKEN"),
}

cln := client.NewClient(&creds)

limit := int32(1)
p := client.ListAccountsParams{
  Limit: &limit,
}

rsp, err := cln.ListAccounts(&p)
if err != nil {
  fmt.Println(err)
  return
}
```

  
