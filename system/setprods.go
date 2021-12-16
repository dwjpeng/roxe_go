package system

import (
	roxe "github.com/dwjpeng/roxe_go"
	"github.com/dwjpeng/roxe_go/ecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `roxe.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `roxe-bios` boot process by the
// `roxe.system` contract.
func NewSetProds(producers []ProducerKey) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("setprods"),
		Authorization: []roxe.PermissionLevel{
			{Actor: AN("roxe"), Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `roxe.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    roxe.AccountName `json:"producer_name"`
	BlockSigningKey ecc.PublicKey    `json:"block_signing_key"`
}
