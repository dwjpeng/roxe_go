package roxe_test

import (
	"context"
	"encoding/json"
	"fmt"

	roxe "github.com/dwjpeng/roxe_go"
)

func ExampleAPI_GetInfo() {
	api := roxe.New(getAPIURL())

	info, err := api.GetInfo(context.Background())
	if err != nil {
		panic(fmt.Errorf("get info: %w", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %w", err))
	}

	fmt.Println(string(bytes))
}
