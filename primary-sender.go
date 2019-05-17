package main

import (
	"encoding/binary"
	"fmt"
	"gopkg.in/eapache/channels.v1"
	"net"
	"sync"
)

type primarySender struct {
	mut       sync.RWMutex
	//videoSock *net.UDPConn
	primaryChannel *channels.RingChannel
	primary   *rtpConnection
}

func (p *primarySender) resetPrimary(prim *rtpConnection) {
	p.mut.Lock()
	defer p.mut.Unlock()
	p.primary = prim
}

func newPrimarySender() *primarySender {
	return &primarySender{
		mut:       sync.RWMutex{},
		primaryChannel: channels.NewRingChannel(100000),
		primary:   nil,
	}
}

func (p *primarySender) initPrimarySend(destIP string, videoSockAddr int) {
	p.mut.Lock()
	defer p.mut.Unlock()
	var seq uint16 = 0

	videoAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", destIP, videoSockAddr))
	videoSock, _ := net.DialUDP("udp", nil, videoAddr)
	go func() {
		fmt.Println("INITING PRIMARY SEND")
		for {
			b := make([]byte, 2)
			binary.BigEndian.PutUint16(b, seq)
			seq ++
			pkt := (<-p.primaryChannel.Out()).([]byte)
			pkt[2] = b[0]
			pkt[3] = b[1]
			videoSock.Write(pkt)
		}
	}()
}

func (p *primarySender) isPrimaryNil() bool {
	p.mut.Lock()
	defer p.mut.Unlock()
	return p.primary == nil
}
