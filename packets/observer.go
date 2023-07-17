package packets

import (
	"time"
)

type Observer interface {
	OnWritePublishPacket(start, end time.Time, p *PublishPacket)
	OnReadPublishPacket(start, end time.Time, p *PublishPacket)

	OnPacketStage(startNs, endNs int64, stage string, p *PublishPacket)
}

func SetObserver(h Observer) {
	InternalOb.ob = h
}

var (
	InternalOb = &eventObserver{}
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

func (rph *eventObserver) OnPacketStage(startNs, endNs int64, stage string, cp ControlPacket) {
	switch m := cp.(type) {
	case *PublishPacket:
		if rph.ob != nil {
			rph.ob.OnPacketStage(startNs, endNs, stage, m)
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
