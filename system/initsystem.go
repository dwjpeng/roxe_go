package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewInitSystem returns a `init` action that lives on the
// `roxe.system` contract.
func NewInitSystem(version roxe.Varuint32, core roxe.Symbol) *roxe.Action {
	return &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("init"),
		Authorization: []roxe.PermissionLevel{
			{
				Actor:      AN("roxe"),
				Permission: roxe.PermissionName("active"),
			},
		},
		ActionData: roxe.NewActionData(Init{
			Version: version,
			Core:    core,
		}),
	}
}

// Init represents the `roxe.system::init` action
type Init struct {
	Version roxe.Varuint32 `json:"version"`
	Core    roxe.Symbol    `json:"core"`
}
