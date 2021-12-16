package forum

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// CleanProposal is an action to flush proposal and allow RAM used by it.
func NewCleanProposal(cleaner roxe.AccountName, proposalName roxe.Name, maxCount uint64) *roxe.Action {
	a := &roxe.Action{
		Account: ForumAN,
		Name:    ActN("clnproposal"),
		Authorization: []roxe.PermissionLevel{
			{Actor: cleaner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(CleanProposal{
			ProposalName: proposalName,
			MaxCount:     maxCount,
		}),
	}
	return a
}

// CleanProposal represents the `roxe.forum::clnproposal` action.
type CleanProposal struct {
	ProposalName roxe.Name `json:"proposal_name"`
	MaxCount     uint64    `json:"max_count"`
}
