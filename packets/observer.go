package packets

import (
	"time"
)

type Observer interface {
	OnWritePublishPacket(start, end time.Time, p *PublishPacket)
	OnReadPublishPacket(start, end time.Time, p *PublishPacket)
}

func SetObserver(h Observer) {
	internalOb.ob = h
}

var (
	internalOb = &eventObserver{}
)

type eventObserver struct {
	ob Observer
}

func (rph *eventObserver) OnWritePacket(start, end time.Time, cp ControlPacket) {
	switch m := cp.(type) {
	case *PublishPacket:
		if rph.ob != nil {
			rph.ob.OnWritePublishPacket(start, end, m)
		}

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

func (rph *eventObserver) OnReadPacket(start, end time.Time, cp ControlPacket) {
	switch m := cp.(type) {
	case *PublishPacket:
		if rph.ob != nil {
			rph.ob.OnReadPublishPacket(start, end, m)
		}

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
