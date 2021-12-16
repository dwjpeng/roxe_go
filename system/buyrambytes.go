package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewBuyRAMBytes will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewBuyRAMBytes(payer, receiver roxe.AccountName, bytes uint32) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("buyrambytes"),
		Authorization: []roxe.PermissionLevel{
			{Actor: payer, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(BuyRAMBytes{
			Payer:    payer,
			Receiver: receiver,
			Bytes:    bytes,
		}),
	}
	return a
}

// BuyRAMBytes represents the `roxe.system::buyrambytes` action.
type BuyRAMBytes struct {
	Payer    roxe.AccountName `json:"payer"`
	Receiver roxe.AccountName `json:"receiver"`
	Bytes    uint32           `json:"bytes"`
}
