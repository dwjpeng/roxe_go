package forum

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// Status is an action to set a status update for a given account on the forum contract.
func NewStatus(account roxe.AccountName, content string) *roxe.Action {
	a := &roxe.Action{
		Account: ForumAN,
		Name:    ActN("status"),
		Authorization: []roxe.PermissionLevel{
			{Actor: account, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Status{
			Account: account,
			Content: content,
		}),
	}
	return a
}

// Status represents the `roxe.forum::status` action.
type Status struct {
	Account roxe.AccountName `json:"account_name"`
	Content string           `json:"content"`
}
