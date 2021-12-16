package system

import (
	"github.com/dwjpeng/roxe_go"
	"github.com/dwjpeng/roxe_go/ecc"
)

// NewNewAccount returns a `newaccount` action that lives on the
// `roxe.system` contract.
func NewNewAccount(creator, newAccount roxe.AccountName, publicKey ecc.PublicKey) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("newaccount"),
		Authorization: []roxe.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: roxe.Authority{
				Threshold: 1,
				Keys: []roxe.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []roxe.PermissionLevelWeight{},
			},
			Active: roxe.Authority{
				Threshold: 1,
				Keys: []roxe.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []roxe.PermissionLevelWeight{},
			},
		}),
	}
}

// NewDelegatedNewAccount returns a `newaccount` action that lives on the
// `roxe.system` contract. It is filled with an authority structure that
// delegates full control of the new account to an already existing account.
func NewDelegatedNewAccount(creator, newAccount roxe.AccountName, delegatedTo roxe.AccountName) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("newaccount"),
		Authorization: []roxe.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: roxe.Authority{
				Threshold: 1,
				Keys:      []roxe.KeyWeight{},
				Accounts: []roxe.PermissionLevelWeight{
					roxe.PermissionLevelWeight{
						Permission: roxe.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
			Active: roxe.Authority{
				Threshold: 1,
				Keys:      []roxe.KeyWeight{},
				Accounts: []roxe.PermissionLevelWeight{
					roxe.PermissionLevelWeight{
						Permission: roxe.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
		}),
	}
}

// NewCustomNewAccount returns a `newaccount` action that lives on the
// `roxe.system` contract. You can specify your own `owner` and
// `active` permissions.
func NewCustomNewAccount(creator, newAccount roxe.AccountName, owner, active roxe.Authority) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("newaccount"),
		Authorization: []roxe.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner:   owner,
			Active:  active,
		}),
	}
}

// NewAccount represents a `newaccount` action on the `roxe.system`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewAccount struct {
	Creator roxe.AccountName `json:"creator"`
	Name    roxe.AccountName `json:"name"`
	Owner   roxe.Authority   `json:"owner"`
	Active  roxe.Authority   `json:"active"`
}
