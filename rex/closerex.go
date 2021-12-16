package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewCloseREX(owner roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("closerex"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(CloseREX{
			Ownwer: owner,
		}),
	}
}

type CloseREX struct {
	Ownwer roxe.AccountName
}
