package msig

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewApprove returns a `approve` action that lives on the
// `roxe.msig` contract.
func NewApprove(proposer roxe.AccountName, proposalName roxe.Name, level roxe.PermissionLevel) *roxe.Action {
	return &roxe.Action{
		Account:       roxe.AccountName("roxe.msig"),
		Name:          roxe.ActionName("approve"),
		Authorization: []roxe.PermissionLevel{level},
		ActionData:    roxe.NewActionData(Approve{proposer, proposalName, level}),
	}
}

type Approve struct {
	Proposer     roxe.AccountName     `json:"proposer"`
	ProposalName roxe.Name            `json:"proposal_name"`
	Level        roxe.PermissionLevel `json:"level"`
}
