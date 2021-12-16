package msig

import roxe "github.com/dwjpeng/roxe_go"

type ProposalRow struct {
	ProposalName       roxe.Name              `json:"proposal_name"`
	RequestedApprovals []roxe.PermissionLevel `json:"requested_approvals"`
	ProvidedApprovals  []roxe.PermissionLevel `json:"provided_approvals"`
	PackedTransaction  roxe.HexBytes          `json:"packed_transaction"`
}
