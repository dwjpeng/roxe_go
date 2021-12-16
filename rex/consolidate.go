package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewConsolidate(owner roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("consolidate"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Consolidate{
			Owner: owner,
		}),
	}
}

type Consolidate struct {
	Owner roxe.AccountName
}
