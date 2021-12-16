package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewRegProxy returns a `regproxy` action that lives on the
// `roxe.system` contract.
func NewRegProxy(proxy roxe.AccountName, isProxy bool) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("regproxy"),
		Authorization: []roxe.PermissionLevel{
			{Actor: proxy, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(RegProxy{
			Proxy:   proxy,
			IsProxy: isProxy,
		}),
	}
}

// RegProxy represents the `roxe.system::regproxy` action
type RegProxy struct {
	Proxy   roxe.AccountName `json:"proxy"`
	IsProxy bool             `json:"isproxy"`
}
