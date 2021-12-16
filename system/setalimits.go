package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewSetalimits sets the account limits. Requires signature from `roxe@active` account.
func NewSetalimits(account roxe.AccountName, ramBytes, netWeight, cpuWeight int64) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("setalimit"),
		Authorization: []roxe.PermissionLevel{
			{Actor: roxe.AccountName("roxe"), Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(Setalimits{
			Account:   account,
			RAMBytes:  ramBytes,
			NetWeight: netWeight,
			CPUWeight: cpuWeight,
		}),
	}
	return a
}

// Setalimits represents the `roxe.system::setalimit` action.
type Setalimits struct {
	Account   roxe.AccountName `json:"account"`
	RAMBytes  int64            `json:"ram_bytes"`
	NetWeight int64            `json:"net_weight"`
	CPUWeight int64            `json:"cpu_weight"`
}
