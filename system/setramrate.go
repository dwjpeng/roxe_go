package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewSetRAMRate(bytesPerBlock uint16) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("setram"),
		Authorization: []roxe.PermissionLevel{
			{
				Actor:      AN("roxe"),
				Permission: roxe.PermissionName("active"),
			},
		},
		ActionData: roxe.NewActionData(SetRAMRate{
			BytesPerBlock: bytesPerBlock,
		}),
	}
	return a
}

// SetRAMRate represents the system contract's `setramrate` action.
type SetRAMRate struct {
	BytesPerBlock uint16 `json:"bytes_per_block"`
}
