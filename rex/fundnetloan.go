package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewFundNetLoan(from roxe.AccountName, loanNumber uint64, payment roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("fundnetloan"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(FundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundNetLoan struct {
	From       roxe.AccountName
	LoanNumber uint64
	Payment    roxe.Asset
}
