package main

import (
	"github.com/pixelbender/go-sdp/sdp"
	"sync"
)
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	resp, _ := http.Get("http://129.64.183.140:8080/ccapi/ver100/shooting/liveview/rtpsessiondesc")
//		//fmt.Println(err)
//	defer resp.Body.Close()
//	body, err2 := ioutil.ReadAll(resp.Body)
//	if err2 != nil {
//		return
//	}
//	body_text := string(body)
//	fmt.Println(body_text)
//	sess, _ := sdp.ParseString(body_text)
//	fmt.Println(sess.Media[0].Format[0].Name)
//	m := sdp_mutex{ mut: sync.RWMutex{}}
//	m.write_session(*sess)
//	fmt.Println(sess.Version)
//	fmt.Println(m.session.Version)
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, m.session.String())
//	})
//	for m  := range(sess.Media) {
//		fmt.Println(sess.Media[m].Port)
//	}
//	//http.ListenAndServe(":8080", nil)
//	audioconn, _ := net.ResolveUDPAddr("udp", ":12010")
//	videoconn, _ := net.ResolveUDPAddr("udp", ":12000")
//	go func() {
//		defer wg.Done()
//		video_conn, _ := net.ListenUDP("udp", videoconn)
//		video_buff := make([]byte, 4096)
//		for {
//
//			n,_,_  := video_conn.ReadFromUDP(video_buff)
//			fmt.Printf("%d READ FROM VIDEO\n", n)
//		}
//	}()
//	go func() {
//		defer wg.Done()
//		audio_conn, _ := net.ListenUDP("udp", audioconn)
//		audio_buff := make([]byte, 1500)
//		for {
//			n,_,_ := audio_conn.ReadFromUDP(audio_buff)
//			fmt.Printf("%d READ FROM AUDIO\n", n)
//		}
//	}()
//	var jsonstr = []byte(`{"ipaddress": "129.64.146.174", "action": "start"}`)
//	fmt.Println(string(jsonstr))
//	resp, err := http.Post("http://129.64.183.140:8080/ccapi/ver100/shooting/liveview/rtp", "application/json", bytes.NewBuffer(jsonstr))
//	if err != nil {
//		panic("aaaa")
//	}
//	fmt.Println(resp.Status)
//	//var result map[string]interface{}
//	defer resp.Body.Close()
//	body, bodyerr := ioutil.ReadAll(resp.Body)
//	if bodyerr != nil {
//		panic("")
//	}
//	fmt.Println(string(body))
//	//fmt.Println(result)
//	fmt.Printf("AAAAAAAA")
//	wg.Wait()
//
//}

type sdp_mutex struct {
	session *sdp.Session
	mut sync.RWMutex
}

func (sdp *sdp_mutex ) writeSession(sess sdp.Session) {
	sdp.mut.Lock()
	defer sdp.mut.Unlock()
	sdp.session = &sess


}

func (sdp *sdp_mutex) readSession() sdp.Session {
	sdp.mut.RLock()
	defer sdp.mut.RUnlock()
	return *sdp.session
}

func (sdp_s *sdp_mutex) readSessionWithSpecPorts(videoport int, audioport int) *sdp.Session {
	sess, _ := sdp.ParseString(sdp_s.session.String())
	for i := range sess.Media {
		if sess.Media[i].Type == "audio" {
			sess.Media[i].Port = audioport
		} else if sess.Media[i].Type == "video" {
			sess.Media[i].Port = videoport
		}
	}
	return sess
}


