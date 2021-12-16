package forum

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewUnVote is an action representing the action to undoing a current vote
func NewUnVote(voter roxe.AccountName, proposalName roxe.Name) *roxe.Action {
	a := &roxe.Action{
		Account: ForumAN,
		Name:    ActN("unvote"),
		Authorization: []roxe.PermissionLevel{
			{Actor: voter, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(UnVote{
			Voter:        voter,
			ProposalName: proposalName,
		}),
	}
	return a
}

// UnVote represents the `roxe.forum::unvote` action.
type UnVote struct {
	Voter        roxe.AccountName `json:"voter"`
	ProposalName roxe.Name        `json:"proposal_name"`
}
