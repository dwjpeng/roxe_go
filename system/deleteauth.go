package system

import "github.com/dwjpeng/roxe_go"

// NewDeleteAuth creates an action from the `roxe.system` contract
// called `deleteauth`.
//
// You cannot delete the `owner` or `active` permissions.  Also, if a
// permission is still linked through a previous `updatelink` action,
// you will need to `unlinkauth` first.
func NewDeleteAuth(account roxe.AccountName, permission roxe.PermissionName) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("deleteauth"),
		Authorization: []roxe.PermissionLevel{
			{Actor: account, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(DeleteAuth{
			Account:    account,
			Permission: permission,
		}),
	}

	return a
}

// DeleteAuth represents the native `deleteauth` action, reachable
// through the `roxe.system` contract.
type DeleteAuth struct {
	Account    roxe.AccountName    `json:"account"`
	Permission roxe.PermissionName `json:"permission"`
}
