package msig

import (
	"github.com/dwjpeng/roxe_go"
)

func init() {
	roxe.RegisterAction(AN("roxe.msig"), ActN("propose"), &Propose{})
	roxe.RegisterAction(AN("roxe.msig"), ActN("approve"), &Approve{})
	roxe.RegisterAction(AN("roxe.msig"), ActN("unapprove"), &Unapprove{})
	roxe.RegisterAction(AN("roxe.msig"), ActN("cancel"), &Cancel{})
	roxe.RegisterAction(AN("roxe.msig"), ActN("exec"), &Exec{})
}

var AN = roxe.AN
var PN = roxe.PN
var ActN = roxe.ActN
