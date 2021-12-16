package forum

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewPropose is an action to submit a proposal for vote.
func NewPropose(proposer roxe.AccountName, proposalName roxe.Name, title string, proposalJSON string, expiresAt roxe.JSONTime) *roxe.Action {
	a := &roxe.Action{
		Account: ForumAN,
		Name:    ActN("propose"),
		Authorization: []roxe.PermissionLevel{
			{Actor: proposer, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Propose{
			Proposer:     proposer,
			ProposalName: proposalName,
			Title:        title,
			ProposalJSON: proposalJSON,
			ExpiresAt:    expiresAt,
		}),
	}
	return a
}

// Propose represents the `roxe.forum::propose` action.
type Propose struct {
	Proposer     roxe.AccountName `json:"proposer"`
	ProposalName roxe.Name        `json:"proposal_name"`
	Title        string           `json:"title"`
	ProposalJSON string           `json:"proposal_json"`
	ExpiresAt    roxe.JSONTime    `json:"expires_at"`
}
