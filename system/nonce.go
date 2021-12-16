package system

import "github.com/dwjpeng/roxe_go"

// NewNonce returns a `nonce` action that lives on the
// `roxe.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `roxe-bios` boot process by the
// `roxe.system` contract.
func NewNonce(nonce string) *roxe.Action {
	a := &roxe.Action{
		Account:       AN("roxe"),
		Name:          ActN("nonce"),
		Authorization: []roxe.PermissionLevel{
			//{Actor: AN("roxe"), Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}
