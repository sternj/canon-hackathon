package main

import (
	"fmt"
	"github.com/pixelbender/go-sdp/sdp"
	"gopkg.in/eapache/channels.v1"
	"net"
)

type rtpConnection struct {
	sess      *sdp.Session
	videoChan *channels.RingChannel
}

func (r *rtpConnection) initSends(ip string, videoPort int) {
	//audioAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", IP, audioPort))
	videoAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, videoPort))
	//audioConn, _ := net.DialUDP("udp", nil, audioAddr)
	videoSock, _ := net.DialUDP("udp", nil, videoAddr)

	go func() { //Video thread
		for {
			videoPacket := (<-r.videoChan.Out()).([]byte)

			primary.mut.Lock()
			//fmt.Println( primary.videoSock != nil)
			//fmt.Println(primary.primary == r)
			if primary.primary == r && primary.videoSock != nil {
				//fmt.Println(primary.videoSock == nil)
				//fmt.Println("SENDING TO PRIMARY")
				pkt := make([]byte, len(videoPacket))
				copy(pkt, videoPacket)
				n, err := primary.videoSock.Write(pkt)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(n)
			}
			primary.mut.Unlock()
			videoSock.Write(videoPacket)
		}

	}()

	//go func() { //timed send
	//	packet := (<-delayQueue.Out()).(*delayItem)
	//	enqTime := time.Unix(packet.timestamp,0)
	//	for time.Since(enqTime).Seconds() > 10 {
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
}
