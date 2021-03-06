package p2p

import (
	"github.com/dwjpeng/roxe_go"
)

type Envelope struct {
	Sender   *Peer
	Receiver *Peer
	Packet   *roxe.Packet `json:"envelope"`
}

func NewEnvelope(sender *Peer, receiver *Peer, packet *roxe.Packet) *Envelope {
	return &Envelope{
		Sender:   sender,
		Receiver: receiver,
		Packet:   packet,
	}
}
