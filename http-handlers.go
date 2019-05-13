package main

import (
	"encoding/json"
	"fmt"
	"github.com/pixelbender/go-sdp/sdp"
	"io/ioutil"
	"net/http"
	"strconv"
)

func acceptNewConn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}
}

func switchConn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var params changeParams
	_ = decoder.Decode(params)
	newPrimary := params.newPrimary
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
	err := decoder.Decode(params)
	if err != nil {
		return
	}

	resp, err2 := http.Get(params.uri)
	defer resp.Body.Close()
	if err2 != nil {
		w.WriteHeader(500)
		//send back a 500
	}
	if resp.StatusCode != 200 {
		w.WriteHeader(resp.StatusCode)
		//send resp.statusCode
	}

	body, decodeErr := ioutil.ReadAll(resp.Body)
	if decodeErr != nil {
		w.WriteHeader(404)
		//probably send 404
	}

	session, _ := sdp.ParseString(string(body))
	audioCh, videoCh := channelsAndListenersFromSDP(session)
	conns.addConnection(&rtpConnection{audioChan: audioCh, videoChan: videoCh, sess: session})
	w.WriteHeader(202)
	fmt.Fprintf(w, session.String())
}

func initSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		return
	}
	audioSendPort, _ := strconv.Atoi(r.URL.Query().Get("audio_port"))
	videoSendPort, _ := strconv.Atoi(r.URL.Query().Get("video_port"))
	sessionIndex, _ := strconv.Atoi(r.URL.Query().Get("session_index"))
	sessionIp := r.URL.Query().Get("session_ip")
	session := conns.getElem(sessionIndex)
	fmt.Fprintf(w, readSessionWithSpecPorts(session.sess, videoSendPort, audioSendPort).String())
	session.initSends(sessionIp, audioSendPort, videoSendPort)
}
type newConn struct {
	uri string
}

type changeParams struct {
	newPrimary int
}

type newSession struct {
	ip string
	videoPort int
	audioPort int
}