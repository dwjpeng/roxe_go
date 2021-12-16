package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewUndelegateBW returns a `undelegatebw` action that lives on the
// `roxe.system` contract.
func NewUndelegateBW(from, receiver roxe.AccountName, unstakeCPU, unstakeNet roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("undelegatebw"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(UndelegateBW{
			From:       from,
			Receiver:   receiver,
			UnstakeNet: unstakeNet,
			UnstakeCPU: unstakeCPU,
		}),
	}
}

// UndelegateBW represents the `roxe.system::undelegatebw` action.
type UndelegateBW struct {
	From       roxe.AccountName `json:"from"`
	Receiver   roxe.AccountName `json:"receiver"`
	UnstakeNet roxe.Asset       `json:"unstake_net_quantity"`
	UnstakeCPU roxe.Asset       `json:"unstake_cpu_quantity"`
}
