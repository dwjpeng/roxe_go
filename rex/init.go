package rex

import roxe "github.com/dwjpeng/roxe_go"

func init() {
	roxe.RegisterAction(REXAN, ActN("buyrex"), BuyREX{})
	roxe.RegisterAction(REXAN, ActN("closerex"), CloseREX{})
	roxe.RegisterAction(REXAN, ActN("cnclrexorder"), CancelREXOrder{})
	roxe.RegisterAction(REXAN, ActN("consolidate"), Consolidate{})
	roxe.RegisterAction(REXAN, ActN("defcpuloan"), DefundCPULoan{})
	roxe.RegisterAction(REXAN, ActN("defnetloan"), DefundNetLoan{})
	roxe.RegisterAction(REXAN, ActN("deposit"), Deposit{})
	roxe.RegisterAction(REXAN, ActN("fundcpuloan"), FundCPULoan{})
	roxe.RegisterAction(REXAN, ActN("fundnetloan"), FundNetLoan{})
	roxe.RegisterAction(REXAN, ActN("mvfrsavings"), MoveFromSavings{})
	roxe.RegisterAction(REXAN, ActN("mvtosavings"), MoveToSavings{})
	roxe.RegisterAction(REXAN, ActN("rentcpu"), RentCPU{})
	roxe.RegisterAction(REXAN, ActN("rentnet"), RentNet{})
	roxe.RegisterAction(REXAN, ActN("rexexec"), REXExec{})
	roxe.RegisterAction(REXAN, ActN("sellrex"), SellREX{})
	roxe.RegisterAction(REXAN, ActN("unstaketorex"), UnstakeToREX{})
	roxe.RegisterAction(REXAN, ActN("updaterex"), UpdateREX{})
	roxe.RegisterAction(REXAN, ActN("withdraw"), Withdraw{})
}

var AN = roxe.AN
var PN = roxe.PN
var ActN = roxe.ActN

var REXAN = AN("roxe")
