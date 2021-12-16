package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewSellRAM will sell at current market price a given number of
// bytes of RAM.
func NewSellRAM(account roxe.AccountName, bytes uint64) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("sellram"),
		Authorization: []roxe.PermissionLevel{
			{Actor: account, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(SellRAM{
			Account: account,
			Bytes:   bytes,
		}),
	}
	return a
}

// SellRAM represents the `roxe.system::sellram` action.
type SellRAM struct {
	Account roxe.AccountName `json:"account"`
	Bytes   uint64           `json:"bytes"`
}
