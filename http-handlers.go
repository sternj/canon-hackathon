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
	//if r.Method != http.MethodPost {
	//	w.WriteHeader(400)
	//	return
	//}
	//decoder := json.NewDecoder(r.Body)
	//var params newSession
	//_ = decoder.Decode(params)
	//
}

func switchConn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var params changeParams
	_ = decoder.Decode(&params)
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
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("DECODE ERROR")
		return
	}

	resp, err2 := http.Get(params.Uri)
	fmt.Println(params.Uri)
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
	conns.addConnection(&rtpConnection{ videoChan: videoCh, sess: session})
	w.WriteHeader(202)
	w.Write(session.Bytes())


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
	Uri string
}

type changeParams struct {
	newPrimary int
}

type newSession struct {
	ip string
	videoPort int
}