package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewRentCPU(
	from roxe.AccountName,
	receiver roxe.AccountName,
	loanPayment roxe.Asset,
	loanFund roxe.Asset,
) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("rentcpu"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(RentCPU{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentCPU struct {
	From        roxe.AccountName
	Receiver    roxe.AccountName
	LoanPayment roxe.Asset
	LoanFund    roxe.Asset
}
