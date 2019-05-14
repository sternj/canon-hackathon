package main

import "github.com/pixelbender/go-sdp/sdp"



func getVideoPort(session *sdp.Session) int {
	return getMedia("video", session)
}

func getMedia(media string, session *sdp.Session) int {
	for i := range session.Media {
		if session.Media[i].Type == media {
			return session.Media[i].Port
		}
	}
	return -1
}

func readSessionWithSpecPorts(s *sdp.Session, videoport int) *sdp.Session {
	sess, _ := sdp.ParseString(s.String())
	for i := range sess.Media {
		if sess.Media[i].Type == "video" {
			sess.Media[i].Port = videoport
		}
	}
	return sess
}

func removeAudioFromSession(s *sdp.Session) {
	media := make([]*sdp.Media, 0)
	for _,m := range s.Media {
		if m.Type != "audio" {
			media = append(media, m)
		}
	}
	s.Media = media
}