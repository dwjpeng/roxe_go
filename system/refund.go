package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewRefund returns a `refund` action that lives on the
// `roxe.system` contract.
func NewRefund(owner roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("refund"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(Refund{
			Owner: owner,
		}),
	}
}

// Refund represents the `roxe.system::refund` action
type Refund struct {
	Owner roxe.AccountName `json:"owner"`
}
