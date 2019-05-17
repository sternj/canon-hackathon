package main

import (
	"fmt"
	"gopkg.in/eapache/channels.v1"
	"time"
)

type delayItem struct {
	timestamp int64
	item      []byte
}

var delayIngest = channels.NewRingChannel(1000000)
var delayQueue = channels.NewRingChannel(1000000)

func initDelay() {
	fmt.Println("INITING DELAY QUEUE")
	go func() { //timed send
		for {
			packet := (<-delayIngest.Out()).(*delayItem)
			//fmt.Println("got packet")
			enqTime := time.Unix(packet.timestamp, 0)
			for time.Since(enqTime).Seconds() < 10 {
				time.Sleep(1 * time.Millisecond)
			}
			//fmt.Println("ENQUEUEING DELAY")
			delayQueue.In() <- packet.item
		}

	}()
}

//func initSendsDelay(ip string, videoPort int) {
//
//}
