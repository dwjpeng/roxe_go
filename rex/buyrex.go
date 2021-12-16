package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewBuyREX(from roxe.AccountName, amount roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("buyrex"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(BuyREX{
			From:   from,
			Amount: amount,
		}),
	}
}

type BuyREX struct {
	From   roxe.AccountName
	Amount roxe.Asset
}
