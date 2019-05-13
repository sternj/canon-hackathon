package main

import (
	"fmt"
	"github.com/pixelbender/go-sdp/sdp"
	"gopkg.in/eapache/channels.v1"
	"net"
)

type rtpConnection struct {
	sess *sdp.Session
	audioChan *channels.RingChannel
	videoChan *channels.RingChannel
	videoSock *net.UDPConn
	audioSock *net.UDPConn
}

func (r *rtpConnection) initSends(ip string, audioPort int, videoPort int) {
	audioAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, audioPort))
	videoAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, videoPort))
	audioConn, _ := net.DialUDP("udp", nil, audioAddr)
	videoConn, _:= net.DialUDP("udp", nil, videoAddr)
	go func() { //Audio thread
		for {
			audioPacket := (<-r.audioChan.Out()).([]byte)
			primary.mut.Lock()
			if primary.primary == nil {
				primary.primary = r
			}
			if primary.primary == r {
				pkt := make([]byte, len(audioPacket))
				copy(pkt, audioPacket)
				_,_ = primary.primary.videoSock.Write(pkt)
			}
			primary.mut.Unlock()
			_,_ = audioConn.Write(audioPacket)
		}
	}()

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

