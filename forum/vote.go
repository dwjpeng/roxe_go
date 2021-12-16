package forum

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewVote is an action representing a simple vote to be broadcast
// through the chain network.
func NewVote(voter roxe.AccountName, proposalName roxe.Name, voteValue uint8, voteJSON string) *roxe.Action {
	a := &roxe.Action{
		Account: ForumAN,
		Name:    ActN("vote"),
		Authorization: []roxe.PermissionLevel{
			{Actor: voter, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Vote{
			Voter:        voter,
			ProposalName: proposalName,
			Vote:         voteValue,
			VoteJSON:     voteJSON,
		}),
	}
	return a
}

// Vote represents the `roxe.forum::vote` action.
type Vote struct {
	Voter        roxe.AccountName `json:"voter"`
	ProposalName roxe.Name        `json:"proposal_name"`
	Vote         uint8            `json:"vote"`
	VoteJSON     string           `json:"vote_json"`
}
