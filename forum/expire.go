package forum

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewExpire is an action to expire a proposal ahead of its natural death.
func NewExpire(proposer roxe.AccountName, proposalName roxe.Name) *roxe.Action {
	a := &roxe.Action{
		Account: ForumAN,
		Name:    ActN("expire"),
		Authorization: []roxe.PermissionLevel{
			{Actor: proposer, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Expire{
			ProposalName: proposalName,
		}),
	}
	return a
}

// Expire represents the `roxe.forum::propose` action.
type Expire struct {
	ProposalName roxe.Name `json:"proposal_name"`
}
