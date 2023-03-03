package packets

import (
	"time"
)

type Hook interface {
	OnReadPublishPacket(startTime, endTime time.Time, p *PublishPacket)
}

func SetHook(h Hook) {
	internalReadPacketHook.hook = h
}

var (
	internalReadPacketHook = &eventObserver{}
)

type eventObserver struct {
	hook Hook
}

func (rph *eventObserver) OnReadNewPacket(startTime, endTime time.Time, cp ControlPacket) {
	switch m := cp.(type) {
	case *PublishPacket:
		rph.hook.OnReadPublishPacket(startTime, endTime, m)

	case *PingrespPacket:
	case *SubackPacket:
	case *UnsubackPacket:
	case *PubackPacket:
	case *PubrecPacket:
	case *PubrelPacket:
	case *PubcompPacket:
		//ignore
	}
}
