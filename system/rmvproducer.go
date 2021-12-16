package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewRemoveProducer returns a `rmvproducer` action that lives on the
// `roxe.system` contract.  This is to be called by the consortium of
// BPs, to oust a BP from its place.  If you want to unregister
// yourself as a BP, use `unregprod`.
func NewRemoveProducer(producer roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("rmvproducer"),
		Authorization: []roxe.PermissionLevel{
			{Actor: AN("roxe"), Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(RemoveProducer{
			Producer: producer,
		}),
	}
}

// RemoveProducer represents the `roxe.system::rmvproducer` action
type RemoveProducer struct {
	Producer roxe.AccountName `json:"producer"`
}
