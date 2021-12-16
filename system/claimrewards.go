package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewClaimRewards will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewClaimRewards(owner roxe.AccountName) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("claimrewards"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(ClaimRewards{
			Owner: owner,
		}),
	}
	return a
}

// ClaimRewards represents the `roxe.system::claimrewards` action.
type ClaimRewards struct {
	Owner roxe.AccountName `json:"owner"`
}
