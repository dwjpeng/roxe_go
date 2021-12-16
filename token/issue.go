package token

import roxe "github.com/dwjpeng/roxe_go"

func NewIssue(to roxe.AccountName, quantity roxe.Asset, memo string) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe.token"),
		Name:    ActN("issue"),
		Authorization: []roxe.PermissionLevel{
			{Actor: to, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(Issue{
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Issue represents the `issue` struct on the `roxe.token` contract.
type Issue struct {
	To       roxe.AccountName `json:"to"`
	Quantity roxe.Asset       `json:"quantity"`
	Memo     string           `json:"memo"`
}
