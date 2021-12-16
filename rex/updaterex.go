package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewUpdateREX(owner roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("updaterex"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(UpdateREX{
			Owner: owner,
		}),
	}
}

type UpdateREX struct {
	Owner roxe.AccountName
}
