package msig

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewExec returns a `exec` action that lives on the
// `roxe.msig` contract.
func NewExec(proposer roxe.AccountName, proposalName roxe.Name, executer roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: roxe.AccountName("roxe.msig"),
		Name:    roxe.ActionName("exec"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []roxe.PermissionLevel{
			{Actor: executer, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Exec{proposer, proposalName, executer}),
	}
}

type Exec struct {
	Proposer     roxe.AccountName `json:"proposer"`
	ProposalName roxe.Name        `json:"proposal_name"`
	Executer     roxe.AccountName `json:"executer"`
}
