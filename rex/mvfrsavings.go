package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewMoveFromSavings(owner roxe.AccountName, rex roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("mvfrsavings"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(MoveFromSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveFromSavings struct {
	Owner roxe.AccountName
	REX   roxe.Asset
}
