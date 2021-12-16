package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewUnregProducer returns a `unregprod` action that lives on the
// `roxe.system` contract.
func NewUnregProducer(producer roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("unregprod"),
		Authorization: []roxe.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(UnregProducer{
			Producer: producer,
		}),
	}
}

// UnregProducer represents the `roxe.system::unregprod` action
type UnregProducer struct {
	Producer roxe.AccountName `json:"producer"`
}
