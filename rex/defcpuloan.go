package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewDefundCPULoan(from roxe.AccountName, loanNumber uint64, amount roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("defcpuloan"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(DefundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundCPULoan struct {
	From       roxe.AccountName
	LoanNumber uint64
	Amount     roxe.Asset
}
