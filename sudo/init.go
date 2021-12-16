package sudo

import roxe "github.com/dwjpeng/roxe_go"

func init() {
	roxe.RegisterAction(AN("roxe.wrap"), ActN("exec"), Exec{})
}

var AN = roxe.AN
var ActN = roxe.ActN
