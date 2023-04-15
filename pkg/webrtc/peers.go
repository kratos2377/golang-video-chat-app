package webrtc

import (
	"sync"
	"videochat/pkg/chat"

	"github.com/pion/webrtc/v3"
)

type Room struct {
	Peers
	Hub *chat.Hub
}

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrame() {

}
