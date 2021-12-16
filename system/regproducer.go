package system

import (
	roxe "github.com/dwjpeng/roxe_go"
	"github.com/dwjpeng/roxe_go/ecc"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `roxe.system` contract.
func NewRegProducer(producer roxe.AccountName, producerKey ecc.PublicKey, url string, location uint16) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("regproducer"),
		Authorization: []roxe.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(RegProducer{
			Producer:    producer,
			ProducerKey: producerKey,
			URL:         url,
			Location:    location,
		}),
	}
}

// RegProducer represents the `roxe.system::regproducer` action
type RegProducer struct {
	Producer    roxe.AccountName `json:"producer"`
	ProducerKey ecc.PublicKey    `json:"producer_key"`
	URL         string           `json:"url"`
	Location    uint16           `json:"location"` // what,s the meaning of that anyway ?
}
