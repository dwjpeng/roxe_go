package forum

import roxe "github.com/dwjpeng/roxe_go"

func init() {
	roxe.RegisterAction(ForumAN, ActN("clnproposal"), CleanProposal{})
	roxe.RegisterAction(ForumAN, ActN("expire"), Expire{})
	roxe.RegisterAction(ForumAN, ActN("post"), Post{})
	roxe.RegisterAction(ForumAN, ActN("propose"), Propose{})
	roxe.RegisterAction(ForumAN, ActN("status"), Status{})
	roxe.RegisterAction(ForumAN, ActN("unpost"), UnPost{})
	roxe.RegisterAction(ForumAN, ActN("unvote"), UnVote{})
	roxe.RegisterAction(ForumAN, ActN("vote"), Vote{})
}

var AN = roxe.AN
var PN = roxe.PN
var ActN = roxe.ActN

var ForumAN = AN("roxe.forum")
