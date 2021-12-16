package system

import (
	"github.com/dwjpeng/roxe_go"
)

func init() {
	roxe.RegisterAction(AN("roxe"), ActN("setcode"), SetCode{})
	roxe.RegisterAction(AN("roxe"), ActN("setabi"), SetABI{})
	roxe.RegisterAction(AN("roxe"), ActN("newaccount"), NewAccount{})
	roxe.RegisterAction(AN("roxe"), ActN("delegatebw"), DelegateBW{})
	roxe.RegisterAction(AN("roxe"), ActN("undelegatebw"), UndelegateBW{})
	roxe.RegisterAction(AN("roxe"), ActN("refund"), Refund{})
	roxe.RegisterAction(AN("roxe"), ActN("regproducer"), RegProducer{})
	roxe.RegisterAction(AN("roxe"), ActN("unregprod"), UnregProducer{})
	roxe.RegisterAction(AN("roxe"), ActN("regproxy"), RegProxy{})
	roxe.RegisterAction(AN("roxe"), ActN("voteproducer"), VoteProducer{})
	roxe.RegisterAction(AN("roxe"), ActN("claimrewards"), ClaimRewards{})
	roxe.RegisterAction(AN("roxe"), ActN("buyram"), BuyRAM{})
	roxe.RegisterAction(AN("roxe"), ActN("buyrambytes"), BuyRAMBytes{})
	roxe.RegisterAction(AN("roxe"), ActN("linkauth"), LinkAuth{})
	roxe.RegisterAction(AN("roxe"), ActN("unlinkauth"), UnlinkAuth{})
	roxe.RegisterAction(AN("roxe"), ActN("deleteauth"), DeleteAuth{})
	roxe.RegisterAction(AN("roxe"), ActN("rmvproducer"), RemoveProducer{})
	roxe.RegisterAction(AN("roxe"), ActN("setprods"), SetProds{})
	roxe.RegisterAction(AN("roxe"), ActN("setpriv"), SetPriv{})
	roxe.RegisterAction(AN("roxe"), ActN("canceldelay"), CancelDelay{})
	roxe.RegisterAction(AN("roxe"), ActN("bidname"), Bidname{})
	// roxe.RegisterAction(AN("roxe"), ActN("nonce"), &Nonce{})
	roxe.RegisterAction(AN("roxe"), ActN("sellram"), SellRAM{})
	roxe.RegisterAction(AN("roxe"), ActN("updateauth"), UpdateAuth{})
	roxe.RegisterAction(AN("roxe"), ActN("setramrate"), SetRAMRate{})
	roxe.RegisterAction(AN("roxe"), ActN("setalimits"), Setalimits{})
}

var AN = roxe.AN
var PN = roxe.PN
var ActN = roxe.ActN
