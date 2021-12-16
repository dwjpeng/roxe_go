package system

import "github.com/dwjpeng/roxe_go"

// NewLinkAuth creates an action from the `roxe.system` contract
// called `linkauth`.
//
// `linkauth` allows you to attach certain permission to the given
// `code::actionName`. With this set on-chain, you can use the
// `requiredPermission` to sign transactions for `code::actionName`
// and not rely on your `active` (which might be more sensitive as it
// can sign anything) for the given operation.
func NewLinkAuth(account, code roxe.AccountName, actionName roxe.ActionName, requiredPermission roxe.PermissionName) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("linkauth"),
		Authorization: []roxe.PermissionLevel{
			{
				Actor:      account,
				Permission: roxe.PermissionName("active"),
			},
		},
		ActionData: roxe.NewActionData(LinkAuth{
			Account:     account,
			Code:        code,
			Type:        actionName,
			Requirement: requiredPermission,
		}),
	}

	return a
}

// LinkAuth represents the native `linkauth` action, through the
// system contract.
type LinkAuth struct {
	Account     roxe.AccountName    `json:"account"`
	Code        roxe.AccountName    `json:"code"`
	Type        roxe.ActionName     `json:"type"`
	Requirement roxe.PermissionName `json:"requirement"`
}
