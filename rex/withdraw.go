package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewWithdraw(owner roxe.AccountName, amount roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("withdraw"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Withdraw{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Withdraw struct {
	Owner  roxe.AccountName
	Amount roxe.Asset
}
