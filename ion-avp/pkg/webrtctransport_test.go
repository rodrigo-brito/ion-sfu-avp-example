package avp

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/pion/rtcp"
	"github.com/pion/transport/test"
	"github.com/pion/webrtc/v3"
	"github.com/stretchr/testify/assert"
)

func waitForBuilder(transport *WebRTCTransport, tid string) chan struct{} {
	done := make(chan struct{})
	go func() {
		for {
			transport.mu.RLock()
			if transport.builders[tid] != nil {
				transport.mu.RUnlock()
				close(done)
				return
			}
			transport.mu.RUnlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return done
}

func waitForRemoveTrack(transport *WebRTCTransport, tid string) chan struct{} {
	done := make(chan struct{})
	go func() {
		for {
			transport.mu.RLock()
			if transport.builders[tid] == nil {
				transport.mu.RUnlock()
				close(done)
				return
			}
			transport.mu.RUnlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return done
}

func signal(t *testing.T, transport *WebRTCTransport, remote *webrtc.PeerConnection) {
	offer, err := remote.CreateOffer(nil)
	gatherComplete := webrtc.GatheringCompletePromise(remote)
	assert.NoError(t, err)
	assert.NoError(t, remote.SetLocalDescription(offer))
	<-gatherComplete

	assert.NoError(t, transport.SetRemoteDescription(offer))

	answer, err := transport.CreateAnswer()
	assert.NoError(t, err)
	assert.NoError(t, transport.SetLocalDescription(answer))
	assert.NoError(t, remote.SetRemoteDescription(answer))
}

func TestNewWebRTCTransport(t *testing.T) {
	report := test.CheckRoutines(t)
	defer report()

	me := webrtc.MediaEngine{}
	me.RegisterDefaultCodecs()
	api := webrtc.NewAPI(webrtc.WithMediaEngine(me))
	remote, err := api.NewPeerConnection(webrtc.Configuration{})
	assert.NoError(t, err)
	defer remote.Close()

	track, err := remote.NewTrack(webrtc.DefaultPayloadTypeOpus, rand.Uint32(), "audio", "pion")
	assert.NoError(t, err)

	_, err = remote.AddTrack(track)
	assert.NoError(t, err)

	Init(map[string]ElementFun{"test-eid": testFunc})

	closed := make(chan struct{})
	transport := NewWebRTCTransport("id", Config{})

	transport.OnClose(func() {
		close(closed)
	})

	err = transport.Process("123", "tid", "test-eid", []byte{})
	assert.NoError(t, err)

	signal(t, transport, remote)

	assert.NoError(t, transport.Close())
	assert.NotNil(t, transport)

	sendRTPUntilDone(closed, t, []*webrtc.Track{track})
}

func TestNewWebRTCTransportWithBuilder(t *testing.T) {
	report := test.CheckRoutines(t)
	defer report()

	me := webrtc.MediaEngine{}
	me.RegisterDefaultCodecs()
	api := webrtc.NewAPI(webrtc.WithMediaEngine(me))
	remote, err := api.NewPeerConnection(webrtc.Configuration{})
	assert.NoError(t, err)

	track, err := remote.NewTrack(webrtc.DefaultPayloadTypeOpus, rand.Uint32(), "audio", "pion")
	assert.NoError(t, err)

	_, err = remote.AddTrack(track)
	assert.NoError(t, err)

	onTransportCloseFired, onTransportCloseFunc := context.WithCancel(context.Background())

	Init(map[string]ElementFun{"test-eid": testFunc})

	transport := NewWebRTCTransport("id", Config{})
	assert.NotNil(t, transport)

	transport.OnClose(func() {
		onTransportCloseFunc()
	})

	err = transport.Process("123", "tid", "test-eid", []byte{})
	assert.NoError(t, err)

	err = transport.Process("123", "tid", "test-eid", []byte{})
	assert.NoError(t, err)

	signal(t, transport, remote)

	transport.OnICECandidate(func(c *webrtc.ICECandidate) {})

	assert.NoError(t, transport.AddICECandidate(webrtc.ICECandidateInit{Candidate: "1986380506 99 udp 2 10.0.75.1 53634 typ host generation 0 network-id 2"}))
	assert.NoError(t, transport.Close())
	<-onTransportCloseFired.Done()

	assert.NoError(t, transport.pc.Close())
	assert.NoError(t, remote.Close())
}

func TestNewWebRTCTransportWithOnNegotiation(t *testing.T) {
	report := test.CheckRoutines(t)
	defer report()

	me := webrtc.MediaEngine{}
	me.RegisterDefaultCodecs()
	api := webrtc.NewAPI(webrtc.WithMediaEngine(me))
	remote, err := api.NewPeerConnection(webrtc.Configuration{})
	assert.NoError(t, err)

	Init(map[string]ElementFun{"test-eid": testFunc})

	transport := NewWebRTCTransport("id", Config{})
	assert.NotNil(t, transport)

	negotiated := make(chan struct{})
	remote.OnNegotiationNeeded(func() {
		signal(t, transport, remote)
		close(negotiated)
	})

	track, err := remote.NewTrack(webrtc.DefaultPayloadTypeOpus, rand.Uint32(), "audio", "pion")
	assert.NoError(t, err)

	sender, err := remote.AddTrack(track)
	assert.NoError(t, err)

	<-negotiated

	err = remote.RemoveTrack(sender)
	assert.NoError(t, err)

	assert.NoError(t, remote.WriteRTCP([]rtcp.Packet{&rtcp.ReceiverEstimatedMaximumBitrate{Bitrate: 10000000, SenderSSRC: track.SSRC()}}))
	assert.NoError(t, transport.Close())
	assert.NoError(t, remote.Close())
}

func TestNewWebRTCTransportWithExpectedBuilder(t *testing.T) {
	report := test.CheckRoutines(t)
	defer report()

	me := webrtc.MediaEngine{}
	me.RegisterDefaultCodecs()
	api := webrtc.NewAPI(webrtc.WithMediaEngine(me))
	remote, err := api.NewPeerConnection(webrtc.Configuration{})
	assert.NoError(t, err)

	tid := "tid"
	track, err := remote.NewTrack(webrtc.DefaultPayloadTypeVP8, rand.Uint32(), tid, "pion")
	assert.NoError(t, err)

	sender, err := remote.AddTrack(track)
	assert.NoError(t, err)

	Init(map[string]ElementFun{"test-eid": testFunc})

	transport := NewWebRTCTransport("id", Config{})
	assert.NotNil(t, transport)

	err = transport.Process("123", tid, "test-eid", []byte{})
	assert.NoError(t, err)

	offer, err := remote.CreateOffer(nil)
	assert.NoError(t, remote.SetLocalDescription(offer))
	gatherComplete := webrtc.GatheringCompletePromise(remote)
	assert.NoError(t, err)
	assert.NotNil(t, offer)
	<-gatherComplete
	assert.NoError(t, transport.SetRemoteDescription(*remote.LocalDescription()))
	answer, err := transport.CreateAnswer()
	assert.NoError(t, err)
	err = transport.SetLocalDescription(answer)
	assert.NoError(t, err)
	assert.NoError(t, remote.SetRemoteDescription(*transport.pc.LocalDescription()))

	done := waitForBuilder(transport, tid)
	sendRTPUntilDone(done, t, []*webrtc.Track{track})

	err = transport.Process("123", tid, "test-eid", []byte{})
	assert.NoError(t, err)

	assert.NoError(t, remote.RemoveTrack(sender))
	assert.NoError(t, signalPair(remote, transport.pc))

	done = waitForRemoveTrack(transport, tid)

	sendRTPUntilDone(done, t, []*webrtc.Track{track})

	assert.NoError(t, transport.Close())
	assert.NoError(t, remote.Close())
}
