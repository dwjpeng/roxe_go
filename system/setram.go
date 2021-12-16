package system

import (
	roxe "github.com/dwjpeng/roxe_go"
)

func NewSetRAM(maxRAMSize uint64) *roxe.Action {
	a := &roxe.Action{
		Account: AN("roxe"),
		Name:    ActN("setram"),
		Authorization: []roxe.PermissionLevel{
			{
				Actor:      AN("roxe"),
				Permission: roxe.PermissionName("active"),
			},
		},
		ActionData: roxe.NewActionData(SetRAM{
			MaxRAMSize: roxe.Uint64(maxRAMSize),
		}),
	}
	return a
}

// SetRAM represents the hard-coded `setram` action.
type SetRAM struct {
	MaxRAMSize roxe.Uint64 `json:"max_ram_size"`
}
