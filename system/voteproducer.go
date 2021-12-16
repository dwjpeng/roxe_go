package system

import "github.com/dwjpeng/roxe_go"

// NewNonce returns a `nonce` action that lives on the
// `roxe.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `roxe-bios` boot process by the
// `roxe.system` contract.
func NewVoteProducer(voter roxe.AccountName, proxy roxe.AccountName, producers ...roxe.AccountName) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("voteproducer"),
		Authorization: []roxe.PermissionLevel{
			{Actor: voter, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(
			VoteProducer{
				Voter:     voter,
				Proxy:     proxy,
				Producers: producers,
			},
		),
	}
	return a
}

// VoteProducer represents the `roxe.system::voteproducer` action
type VoteProducer struct {
	Voter     roxe.AccountName   `json:"voter"`
	Proxy     roxe.AccountName   `json:"proxy"`
	Producers []roxe.AccountName `json:"producers"`
}
