package msig

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewCancel returns a `cancel` action that lives on the
// `roxe.msig` contract.
func NewCancel(proposer roxe.AccountName, proposalName roxe.Name, canceler roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: roxe.AccountName("roxe.msig"),
		Name:    roxe.ActionName("cancel"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []roxe.PermissionLevel{
			{Actor: canceler, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Cancel{proposer, proposalName, canceler}),
	}
}

type Cancel struct {
	Proposer     roxe.AccountName `json:"proposer"`
	ProposalName roxe.Name        `json:"proposal_name"`
	Canceler     roxe.AccountName `json:"canceler"`
}
