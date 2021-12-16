package sudo

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewExec creates an `exec` action, found in the `roxe.wrap`
// contract.
//
// Given an `roxe.Transaction`, call `roxe.MarshalBinary` on it first,
// pass the resulting bytes as `roxe.HexBytes` here.
func NewExec(executer roxe.AccountName, transaction roxe.Transaction) *roxe.Action {
	a := &roxe.Action{
		Account: roxe.AccountName("roxe.wrap"),
		Name:    roxe.ActionName("exec"),
		Authorization: []roxe.PermissionLevel{
			{Actor: executer, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Exec{
			Executer:    executer,
			Transaction: transaction,
		}),
	}
	return a
}

// Exec represents the `roxe.system::exec` action.
type Exec struct {
	Executer    roxe.AccountName `json:"executer"`
	Transaction roxe.Transaction `json:"trx"`
}
