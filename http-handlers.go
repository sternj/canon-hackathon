package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pixelbender/go-sdp/sdp"
	"io/ioutil"
	"net/http"
	"strconv"
)

func acceptNewConn(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodPost {
	//	w.WriteHeader(400)
	//	return
	//}
	//decoder := json.NewDecoder(r.Body)
	//var params newSession
	//_ = decoder.Decode(params)
	//
}
func getPrimarySDP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		return
	}
	sdpString := `v=0
o=- 0 0 IN IP4 129.64.183.140
s=RTP Session
c=IN IP4 0.0.0.0
t=0 0
a=control *
m=video 12000 RTP/AVP 103
a=rtpmap:103 H264/90000`
	session, _ := sdp.ParseString(sdpString)

	videoSendPort, _ := strconv.Atoi(r.URL.Query().Get("video_port"))
	sessionIp := r.URL.Query().Get("session_ip")
	w.Write(readSessionWithSpecPorts(session, videoSendPort).Bytes())
	primary.initPrimarySend(sessionIp, videoSendPort)
	fmt.Println("REQ DONE")
}

func switchConn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var params changeParams
	_ = decoder.Decode(&params)
	newPrimary := params.NewPrimary
	if prim := conns.getElem(newPrimary); prim != nil {
		primary.primary = prim
	}
}

//func getConnList(w http.ResponseWriter, r *http.Request) {
//
//}

func initConnStream(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var params newConn
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("DECODE ERROR")
		return
	}
	resp, err2 := http.Get(params.InitSessionUri)
	//fmt.Println(params.InitSessionUri)
	if err2 != nil {
		fmt.Println(err2)
		w.WriteHeader(500)
		return
		//send back a 500
	}
	if resp.StatusCode != 200 {
		fmt.Println("breaking")
		w.WriteHeader(resp.StatusCode)
		return
		//send resp.statusCode
	}

	body, decodeErr := ioutil.ReadAll(resp.Body)
	if decodeErr != nil {
		w.WriteHeader(404)
		//probably send 404
		return
	}

	session, _ := sdp.ParseString(string(body))

	removeAudioFromSession(session)
	videoCh := channelsAndListenersFromSDP(session)
	c := &rtpConnection{videoChan: videoCh, sess: session}
	delay := &rtpConnection{videoChan: delayQueue, sess: session}
	conns.addConnection(c)
	conns.addConnection(delay)
	if primary.isPrimaryNil() {
		initDelay()
		fmt.Println("SETTING PRIMARY")
		primary.resetPrimary(c)
	} else {
		fmt.Println("NOT SETTING PRIM")
	}
	if params.StartStreamURI != "" {
		var jsonstr = []byte(fmt.Sprintf(`{"ipaddress": "%s", "action": "start"}`, GetOutboundIP().String()))
		fmt.Println(string(jsonstr))
		_, _ = http.Post(params.StartStreamURI, "application/json", bytes.NewBuffer(jsonstr))

	}
	w.WriteHeader(202)
	w.Write(session.Bytes())
}

func haltConnStream(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var params haltStream
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("DECODE ERROR")
		return
	}
	var jsonstr = []byte(fmt.Sprintf(`{"ipaddress": "%s", "action": "stop"}`, GetOutboundIP().String()))
	fmt.Println(string(jsonstr))
	_, _ = http.Post(params.HaltStreamURI, "application/json", bytes.NewBuffer(jsonstr))
}

func initSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		return
	}
	videoSendPort, _ := strconv.Atoi(r.URL.Query().Get("video_port"))
	sessionIndex, _ := strconv.Atoi(r.URL.Query().Get("session_index"))
	sessionIp := r.URL.Query().Get("session_ip")
	session := conns.getElem(sessionIndex)
	w.Write(readSessionWithSpecPorts(session.sess, videoSendPort).Bytes())
	session.initSends(sessionIp, videoSendPort)
}

type newConn struct {
	InitSessionUri string
	StartStreamURI string
}

type haltStream struct {
	HaltStreamURI string
}

type changeParams struct {
	NewPrimary int
}

type newSession struct {
	IP        string
	VideoPort int
}
