package packethandler

import (
	"github.com/bnch/bancho/packets"
)

type packetCollection struct {
	content []packets.Packet
}

func (p *packetCollection) Push(val ...packets.Packet) {
	p.content = append(p.content, val...)
}

func (p *packetCollection) Pop() (val *packets.Packet) {
	if len(p.content) == 0 {
		return nil
	}
	val = &p.content[0]
	p.content = p.content[1:]
	return
}
