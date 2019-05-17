package main

import (
	"fmt"
	"net/http"
)

var conns *connectionList
var primary *primarySender

/*
TODO: Remove audio support
TODO: Account for short UDP writes

*/

func main() {
	conns = newConnectionList()
	primary = newPrimarySender()

	http.HandleFunc("/add_connection", initConnStream)
	http.HandleFunc("/index.sdp", initSession)
	http.HandleFunc("/switch_connection", switchConn)
	http.HandleFunc("/new_connection", acceptNewConn)
	http.HandleFunc("/halt_connection", haltConnStream)
	http.HandleFunc("/primary.sdp", getPrimarySDP)
	fmt.Println("LISTENING ON 4567")
	_ = http.ListenAndServe(":4567", nil)
}

//func main() {
//	resp, _ := http.Get("http://129.64.183.140:8080/ccapi/ver100/shooting/liveview/rtpsessiondesc")
//		//fmt.Println(err)
//	//resp, _ := http.Get("http://129.64.157.53:8080/camera.sdp")
//	defer resp.Body.Close()
//	body, err2 := ioutil.ReadAll(resp.Body)
//	if err2 != nil {
//		return
//	}
//	bodyText := string(body)
//	fmt.Println(bodyText)
//	sess, _ := sdp.ParseString(bodyText)
//	fmt.Println(sess.Media[0].Format[0].Name)
//	m := sdp_mutex{ mut: sync.RWMutex{}}
//	m.writeSession(*sess)
//	audioCh, videoCh := channelsAndListenersFromSDP(m.readSession())
//	var jsonstr = []byte(`{"ipaddress": "129.64.124.165", "action": "start"}`)
//	fmt.Println(string(jsonstr))
//	response,_ := http.Post("http://129.64.183.140:8080/ccapi/ver100/shooting/liveview/rtp", "application/json", bytes.NewBuffer(jsonstr))
//	x,_  := ioutil.ReadAll(response.Body)
//	fmt.Println(string(x))
//	http.HandleFunc("/index.sdp", func(w http.ResponseWriter, r *http.Request) {
//		audioSendPort, _ := strconv.Atoi(r.URL.Query().Get("audio_port"))
//		videoSendPort, _ := strconv.Atoi(r.URL.Query().Get("video_port"))
//		fmt.Println(audioSendPort)
//		WriteFromChannelToDestSock(audioSendPort, "localhost", audioCh, "audio")
//		WriteFromChannelToDestSock(videoSendPort, "localhost", videoCh, "video")
//		//AlternatingSockWrite(audioSendPort, videoSendPort, "localhost", audioCh, videoCh)
//		numWritten, err := fmt.Fprint(w, m.readSessionWithSpecPorts(videoSendPort, audioSendPort))
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println(numWritten)
//	})
//	_ = http.ListenAndServe(":4567", nil)
//}
