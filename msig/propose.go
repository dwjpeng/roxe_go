package msig

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewPropose returns a `propose` action that lives on the
// `roxe.msig` contract.
func NewPropose(proposer roxe.AccountName, proposalName roxe.Name, requested []roxe.PermissionLevel, transaction *roxe.Transaction) *roxe.Action {
	return &roxe.Action{
		Account: roxe.AccountName("roxe.msig"),
		Name:    roxe.ActionName("propose"),
		Authorization: []roxe.PermissionLevel{
			{Actor: proposer, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Propose{proposer, proposalName, requested, transaction}),
	}
}

type Propose struct {
	Proposer     roxe.AccountName       `json:"proposer"`
	ProposalName roxe.Name              `json:"proposal_name"`
	Requested    []roxe.PermissionLevel `json:"requested"`
	Transaction  *roxe.Transaction      `json:"trx"`
}
