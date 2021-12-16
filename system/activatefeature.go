package system

import (
	"github.com/dwjpeng/roxe_go"
)

func NewActivateFeature(featureDigest roxe.Checksum256) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("activate"),
		Authorization: []roxe.PermissionLevel{
			{Actor: AN("roxe"), Permission: PN("active")},
		},
		ActionData: roxe.NewActionData(Activate{
			FeatureDigest: featureDigest,
		}),
	}
}

// Activate represents a `activate` action on the `roxe` contract.
type Activate struct {
	FeatureDigest roxe.Checksum256 `json:"feature_digest"`
}
