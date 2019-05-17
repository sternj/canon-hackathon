package main

import (
	"fmt"
	"net"
	"sync"
)

type primarySender struct {
	mut       sync.RWMutex
	videoSock *net.UDPConn
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
		videoSock: nil,
		primary:   nil,
	}
}

func (p *primarySender) initPrimarySend(destIP string, videoSockAddr int) {
	p.mut.Lock()
	defer p.mut.Unlock()
	videoAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", destIP, videoSockAddr))
	videoSock, _ := net.DialUDP("udp", nil, videoAddr)
	p.videoSock = videoSock

}

func (p *primarySender) isPrimaryNil() bool {
	p.mut.Lock()
	defer p.mut.Unlock()
	return p.primary == nil
}
