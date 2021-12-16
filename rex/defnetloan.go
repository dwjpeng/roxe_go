package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewDefundNetLoan(from roxe.AccountName, loanNumber uint64, amount roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("defnetloan"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(DefundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundNetLoan struct {
	From       roxe.AccountName
	LoanNumber uint64
	Amount     roxe.Asset
}
