package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewDelegateBW returns a `delegatebw` action that lives on the
// `roxe.system` contract.
func NewDelegateBW(from, receiver roxe.AccountName, stakeCPU, stakeNet roxe.Asset, transfer bool) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("delegatebw"),
		Authorization: []roxe.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(DelegateBW{
			From:     from,
			Receiver: receiver,
			StakeNet: stakeNet,
			StakeCPU: stakeCPU,
			Transfer: roxe.Bool(transfer),
		}),
	}
}

// DelegateBW represents the `roxe.system::delegatebw` action.
type DelegateBW struct {
	From     roxe.AccountName `json:"from"`
	Receiver roxe.AccountName `json:"receiver"`
	StakeNet roxe.Asset       `json:"stake_net_quantity"`
	StakeCPU roxe.Asset       `json:"stake_cpu_quantity"`
	Transfer roxe.Bool        `json:"transfer"`
}
