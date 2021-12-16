package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewRentNet(
	from roxe.AccountName,
	receiver roxe.AccountName,
	loanPayment roxe.Asset,
	loanFund roxe.Asset,
) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("rentnet"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(RentNet{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentNet struct {
	From        roxe.AccountName
	Receiver    roxe.AccountName
	LoanPayment roxe.Asset
	LoanFund    roxe.Asset
}
