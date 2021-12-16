package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewDeposit(owner roxe.AccountName, amount roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("deposit"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Deposit{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Deposit struct {
	Owner  roxe.AccountName
	Amount roxe.Asset
}
