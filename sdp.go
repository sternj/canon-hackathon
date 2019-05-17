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

