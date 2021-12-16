package token

import roxe "github.com/dwjpeng/roxe_go"

func NewTransfer(from, to roxe.AccountName, quantity roxe.Asset, memo string) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe.token"),
		Name:    ActN("transfer"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `roxe.token` contract.
type Transfer struct {
	From     roxe.AccountName `json:"from"`
	To       roxe.AccountName `json:"to"`
	Quantity roxe.Asset       `json:"quantity"`
	Memo     string           `json:"memo"`
}
