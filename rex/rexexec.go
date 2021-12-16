package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewREXExec(user roxe.AccountName, max uint16) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("rexexec"),
		Authorization: []roxe.PermissionLevel{
			{Actor: user, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(REXExec{
			User: user,
			Max:  max,
		}),
	}
}

type REXExec struct {
	User roxe.AccountName
	Max  uint16
}
