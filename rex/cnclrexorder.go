package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewCancelREXOrder(owner roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("cnclrexorder"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(CancelREXOrder{
			Owner: owner,
		}),
	}
}

type CancelREXOrder struct {
	Owner roxe.AccountName
}
