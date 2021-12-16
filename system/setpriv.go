package system

import roxe "github.com/dwjpeng/roxe_go"

// NewSetPriv returns a `setpriv` action that lives on the
// `roxe.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `roxe-bios` boot process by the
// `roxe.system` contract.
func NewSetPriv(account roxe.AccountName) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("setpriv"),
		Authorization: []roxe.PermissionLevel{
			{Actor: AN("roxe"), Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(SetPriv{
			Account: account,
			IsPriv:  roxe.Bool(true),
		}),
	}
	return a
}

// SetPriv sets privileged account status. Used in the bios boot mechanism.
type SetPriv struct {
	Account roxe.AccountName `json:"account"`
	IsPriv  roxe.Bool        `json:"is_priv"`
}
