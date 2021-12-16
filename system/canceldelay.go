package system

import "github.com/dwjpeng/roxe_go"

// NewCancelDelay creates an action from the `roxe.system` contract
// called `canceldelay`.
//
// `canceldelay` allows you to cancel a deferred transaction,
// previously sent to the chain with a `delay_sec` larger than 0.  You
// need to sign with cancelingAuth, to cancel a transaction signed
// with that same authority.
func NewCancelDelay(cancelingAuth roxe.PermissionLevel, transactionID roxe.Checksum256) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("canceldelay"),
		Authorization: []roxe.PermissionLevel{
			cancelingAuth,
		},
		ActionData: roxe.NewActionData(CancelDelay{
			CancelingAuth: cancelingAuth,
			TransactionID: transactionID,
		}),
	}

	return a
}

// CancelDelay represents the native `canceldelay` action, through the
// system contract.
type CancelDelay struct {
	CancelingAuth roxe.PermissionLevel `json:"canceling_auth"`
	TransactionID roxe.Checksum256     `json:"trx_id"`
}
