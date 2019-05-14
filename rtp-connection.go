package main

import (
	"fmt"
	"github.com/pixelbender/go-sdp/sdp"
	"gopkg.in/eapache/channels.v1"
	"net"
)

type rtpConnection struct {
	sess *sdp.Session
	videoChan *channels.RingChannel
	videoSock *net.UDPConn
}

func (r *rtpConnection) initSends(ip string,  videoPort int) {
	//audioAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, audioPort))
	videoAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, videoPort))
	//audioConn, _ := net.DialUDP("udp", nil, audioAddr)
	videoConn, _:= net.DialUDP("udp", nil, videoAddr)

	go func() { //Video thread
		for {
			videoPacket := (<-r.videoChan.Out()).([]byte)
			primary.mut.Lock()
			if primary.primary == r {
				pkt := make([]byte, len(videoPacket))
				copy(pkt, videoPacket)
				primary.primary.videoSock.Write(pkt)
			}
			primary.mut.Unlock()
			videoConn.Write(videoPacket)
		}

	}()
}

