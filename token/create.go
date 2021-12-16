package token

import roxe "github.com/dwjpeng/roxe_go"

func NewCreate(issuer roxe.AccountName, maxSupply roxe.Asset) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe.token"),
		Name:    ActN("create"),
		Authorization: []roxe.PermissionLevel{
			{Actor: AN("roxe.token"), Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(Create{
			Issuer:        issuer,
			MaximumSupply: maxSupply,
		}),
	}
}

// Create represents the `create` struct on the `roxe.token` contract.
type Create struct {
	Issuer        roxe.AccountName `json:"issuer"`
	MaximumSupply roxe.Asset       `json:"maximum_supply"`
}
