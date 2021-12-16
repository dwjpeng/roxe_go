package system

import "github.com/dwjpeng/roxe_go"

// NewUnlinkAuth creates an action from the `roxe.system` contract
// called `unlinkauth`.
//
// `unlinkauth` detaches a previously set permission from a
// `code::actionName`. See `linkauth`.
func NewUnlinkAuth(account, code roxe.AccountName, actionName roxe.ActionName) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("unlinkauth"),
		Authorization: []roxe.PermissionLevel{
			{
				Actor:      account,
				Permission: roxe.PermissionName("active"),
			},
		},
		ActionData: roxe.NewActionData(UnlinkAuth{
			Account: account,
			Code:    code,
			Type:    actionName,
		}),
	}

	return a
}

// UnlinkAuth represents the native `unlinkauth` action, through the
// system contract.
type UnlinkAuth struct {
	Account roxe.AccountName `json:"account"`
	Code    roxe.AccountName `json:"code"`
	Type    roxe.ActionName  `json:"type"`
}
