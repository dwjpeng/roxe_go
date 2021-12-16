package token

import "github.com/dwjpeng/roxe_go"

func init() {
	roxe.RegisterAction(AN("roxe.token"), ActN("transfer"), Transfer{})
	roxe.RegisterAction(AN("roxe.token"), ActN("issue"), Issue{})
	roxe.RegisterAction(AN("roxe.token"), ActN("create"), Create{})
}
