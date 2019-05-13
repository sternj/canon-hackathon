package main

import (
	"net"
	"sync"
)

type primarySender struct {
	mut sync.RWMutex
	videoSock *net.UDPConn
	audioSock *net.UDPConn
	primary *rtpConnection

}

func (p *primarySender) resetPrimary(prim *rtpConnection) {
	p.mut.Lock()
	defer p.mut.Unlock()
	p.primary = prim
}

func newPrimarySender() *primarySender {
	return &primarySender{
		mut: sync.RWMutex{},
		videoSock: nil,
		audioSock: nil,
		primary: nil,
	}
}

