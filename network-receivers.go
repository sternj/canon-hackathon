package main



import (
	"fmt"
	"github.com/pixelbender/go-sdp/sdp"
	"gopkg.in/eapache/channels.v1"
	"net"
	"os"
)

func readRtpFromUdpSock(listenPort int, sendChannel* channels.RingChannel, r_type string) {
	connection, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open socket")
		return
	}
	go func() {
		listenConnection, _ := net.ListenUDP("udp", connection)
		buff := make([]byte, 4096)
		for {
			n,_,err := listenConnection.ReadFromUDP(buff)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR READING FROM UDP")
			}
			readBytes := buff[:n]
			toEnqueue := make([]byte, len(readBytes))
			copy(toEnqueue, readBytes)
			sendChannel.In() <- toEnqueue
			//if r_type == "audio" {
			//	fmt.Println(n)
			//}
		}
	}()
}

func WriteFromChannelToDestSock(port int, ip string, recvChannel* channels.RingChannel, r_type string) {
	addr, err1 := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, port))
	if err1 != nil {
		fmt.Printf("ERROR\n")
	}
	go func() {
		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR ON %s", r_type)
		}
		defer conn.Close()
		for {
			packet := <-recvChannel.Out()
			a := packet.([]byte)
			conn.Write(a)

		}

	}()
}

func AlternatingSockWrite(audioPort int, videoPort int, ip string, audioRecvChannel* channels.RingChannel, videoRecvChannel* channels.RingChannel) {
	audioAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, audioPort))
	videoAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, videoPort))
	go func() {
		audioConn, _ := net.DialUDP("udp", nil, audioAddr)
		videoConn, _ := net.DialUDP("udp", nil, videoAddr)
		defer audioConn.Close()
		defer videoConn.Close()
		for {
			audioPacket := (<-audioRecvChannel.Out()).([]byte)
			videoPacket := (<-videoRecvChannel.Out()).([]byte)
			audioConn.Write(audioPacket)
			videoConn.Write(videoPacket)
		}
	}()
}

func  channelsAndListenersFromSDP(session *sdp.Session) *channels.RingChannel {

	videoChannel := channels.NewRingChannel(1000)

	readRtpFromUdpSock(getVideoPort(session), videoChannel, "video")
	return videoChannel
}
