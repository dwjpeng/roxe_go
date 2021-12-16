package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewMoveToSavings(owner roxe.AccountName, rex roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("mvtosavings"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(MoveToSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveToSavings struct {
	Owner roxe.AccountName
	REX   roxe.Asset
}
