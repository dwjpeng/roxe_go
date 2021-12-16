package system

import "github.com/dwjpeng/roxe_go"

// NewUpdateAuth creates an action from the `roxe.system` contract
// called `updateauth`.
//
// usingPermission needs to be `owner` if you want to modify the
// `owner` authorization, otherwise `active` will do for the rest.
func NewUpdateAuth(account roxe.AccountName, permission, parent roxe.PermissionName, authority roxe.Authority, usingPermission roxe.PermissionName) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("updateauth"),
		Authorization: []roxe.PermissionLevel{
			{
				Actor:      account,
				Permission: usingPermission,
			},
		},
		ActionData: roxe.NewActionData(UpdateAuth{
			Account:    account,
			Permission: permission,
			Parent:     parent,
			Auth:       authority,
		}),
	}

	return a
}

// UpdateAuth represents the hard-coded `updateauth` action.
//
// If you change the `active` permission, `owner` is the required parent.
//
// If you change the `owner` permission, there should be no parent.
type UpdateAuth struct {
	Account    roxe.AccountName    `json:"account"`
	Permission roxe.PermissionName `json:"permission"`
	Parent     roxe.PermissionName `json:"parent"`
	Auth       roxe.Authority      `json:"auth"`
}
