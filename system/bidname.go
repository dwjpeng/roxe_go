package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewBidname(bidder, newname roxe.AccountName, bid roxe.Asset) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("bidname"),
		Authorization: []roxe.PermissionLevel{
			{Actor: bidder, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(Bidname{
			Bidder:  bidder,
			Newname: newname,
			Bid:     bid,
		}),
	}
	return a
}

// Bidname represents the `roxe.system_contract::bidname` action.
type Bidname struct {
	Bidder  roxe.AccountName `json:"bidder"`
	Newname roxe.AccountName `json:"newname"`
	Bid     roxe.Asset       `json:"bid"` // specified in ROC
}
