package rex

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewUnstakeToREX(
	owner roxe.AccountName,
	receiver roxe.AccountName,
	fromNet roxe.Asset,
	fromCPU roxe.Asset,
) *roxe.Action {
	return &roxe.Action{
		Account: REXAN,
		Name:    ActN("unstaketorex"),
		Authorization: []roxe.PermissionLevel{
			{Actor: owner, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(UnstakeToREX{
			Owner:    owner,
			Receiver: receiver,
			FromNet:  fromNet,
			FromCPU:  fromCPU,
		}),
	}
}

type UnstakeToREX struct {
	Owner    roxe.AccountName
	Receiver roxe.AccountName
	FromNet  roxe.Asset
	FromCPU  roxe.Asset
}
