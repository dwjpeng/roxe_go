package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewBuyRAM(payer, receiver roxe.AccountName, roxeQuantity uint64) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("buyram"),
		Authorization: []roxe.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(BuyRAM{
			Payer:    payer,
			Receiver: receiver,
			Quantity: roxe.NewROXEAsset(int64(roxeQuantity)),
		}),
	}
	return a
}

// BuyRAM represents the `roxe.system::buyram` action.
type BuyRAM struct {
	Payer    roxe.AccountName `json:"payer"`
	Receiver roxe.AccountName `json:"receiver"`
	Quantity roxe.Asset       `json:"quant"` // specified in ROXE
}
