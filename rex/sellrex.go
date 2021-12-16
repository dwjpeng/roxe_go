package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewSellREX(from roxe.AccountName, rex roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("sellrex"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(SellREX{
			From: from,
			REX:  rex,
		}),
	}
}

type SellREX struct {
	From roxe.AccountName
	REX  roxe.Asset
}
