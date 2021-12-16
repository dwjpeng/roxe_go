package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewFundCPULoan(from roxe.AccountName, loanNumber uint64, payment roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("fundcpuloan"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(FundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundCPULoan struct {
	From       roxe.AccountName
	LoanNumber uint64
	Payment    roxe.Asset
}
