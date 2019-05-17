package main

import (
	"fmt"
	"github.com/pixelbender/go-sdp/sdp"
	"net/http"
)

var cons *connectionList
var primary *primarySender

/*
TODO: Remove audio support
TODO: Account for short UDP writes

*/

func main() {
	cons = newConnectionList()
	primary = newPrimarySender()

	http.HandleFunc("/add_connection", initConnStream)
	http.HandleFunc("/index.sdp", initSession)
	http.HandleFunc("/switch_connection", switchConn)
	http.HandleFunc("/new_connection", acceptNewConn)
	http.HandleFunc("/halt_connection", haltConnStream)
	http.HandleFunc("/primary.sdp", getPrimarySDP)
	http.HandleFunc("/test.sdp", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("REQ")
		sdpString := `v=0
o=- 0 0 IN IP4 129.64.183.140
s=RTP Session
c=IN IP4 0.0.0.0
t=0 0
a=control *
m=video 12000 RTP/AVP 103
a=rtpmap:103 H264/90000`
		session, _ := sdp.ParseString(sdpString)
		w.Write(session.Bytes())
	})
	fmt.Println("LISTENING ON 4567")
	_ = http.ListenAndServe(":4567", nil)
}
