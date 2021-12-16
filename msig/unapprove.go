package msig

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewUnapprove returns a `unapprove` action that lives on the
// `roxe.msig` contract.
func NewUnapprove(proposer roxe.AccountName, proposalName roxe.Name, level roxe.PermissionLevel) *roxe.Action {
	return &roxe.Action{
		Account:       roxe.AccountName("roxe.msig"),
		Name:          roxe.ActionName("unapprove"),
		Authorization: []roxe.PermissionLevel{level},
		ActionData:    roxe.NewActionData(Unapprove{proposer, proposalName, level}),
	}
}

type Unapprove struct {
	Proposer     roxe.AccountName     `json:"proposer"`
	ProposalName roxe.Name            `json:"proposal_name"`
	Level        roxe.PermissionLevel `json:"level"`
}
