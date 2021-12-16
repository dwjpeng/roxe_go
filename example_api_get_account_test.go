package roxe_test

import (
	"context"
	"encoding/json"
	"fmt"

	roxe "github.com/dwjpeng/roxe_go"
)

func ExampleAPI_GetAccount() {
	api := roxe.New(getAPIURL())

	account := roxe.AccountName("roxe.rex")
	info, err := api.GetAccount(context.Background(), account)
	if err != nil {
		if err == roxe.ErrNotFound {
			fmt.Printf("unknown account: %s", account)
			return
		}

		panic(fmt.Errorf("get account: %w", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %w", err))
	}

	fmt.Println(string(bytes))
}
