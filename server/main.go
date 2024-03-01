package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/grafov/m3u8"

	"github.com/meanii/aws-lambda-hls-enc/config"
	"github.com/meanii/aws-lambda-hls-enc/handler"
	"github.com/meanii/aws-lambda-hls-enc/helper"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// handle OPTIONS method
	if r.Method == http.MethodOptions {
		handler.HandlerMethodOptions(w, r)
		return
	}

	// handle HEAD method
	if r.Method == http.MethodHead {
		handler.HandlerMethodHead(w, r)
		return
	}

	requestFullUrl := r.URL.String()
	fmt.Printf("Request URL: %s\n", requestFullUrl)

	cloudfrontUrl, err := url.Parse(requestFullUrl)
	if err != nil {
		fmt.Printf("Error parsing URL: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// set the cloudfront URL host to the origin
	// example: https://lambda-function.com -> https://origin-cdn.com
	cloudfrontUrl.Host = config.ConfigEnv.Origin

	// check if has suffix .ts, *.key return the file
	if strings.HasSuffix(requestFullUrl, ".ts") || strings.HasSuffix(requestFullUrl, ".key") {
		handler.BinaryChunksRedirection(*cloudfrontUrl, w, r)
		return
	}

	// fetch m3u8 file from cloudfront
	m3u8Resp, err := helper.Fetch(cloudfrontUrl.String())
	defer m3u8Resp.Body.Close()

	// return 500 if error fetching m3u8 file
	// it may be not valid signed URL
	// master file must be signed, before calling to this lambda function
	// example: https://lambda-function.com/master.m3u8?Expires=1610000000&Signature=xxxx&Key-Pair-Id=xxxx
	// but the master.m3u8 signed with cdn-origin.com
	if err != nil {
		fmt.Printf("Error fetching m3u8: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// decode m3u8 file, this is a common library for m3u8 file parsing
	// it will return the type of the m3u8 file, and the parsed object of the m3u8 file
	p, listType, err := m3u8.DecodeFrom(m3u8Resp.Body, true)
	if err != nil {
		fmt.Printf("Error decoding m3u8: %s\n", err)
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("The provided URL might be expired Or url signed with the incorrect credentials."))
		return
	}

	// check the type of the m3u8 file, then sign the URL and return the signed URL
	// listType, m3u8.MEDIA, m3u8.MASTER are the types of the m3u8 file
	switch listType {

	case m3u8.MEDIA:
		// cast the parsed object to m3u8.MediaPlaylist
		// then call the handler.M3u8Media to return the media playlist
		mediapl := p.(*m3u8.MediaPlaylist)
		handler.M3u8Media(mediapl, w, r)
		return

	case m3u8.MASTER:
		// cast the parsed object to m3u8.MasterPlaylist
		// then call the handler.M3u8Master to return the master playlist
		masterpl := p.(*m3u8.MasterPlaylist)
		handler.M3u8Master(*cloudfrontUrl, masterpl, w, r)
		return
	}

	// if the m3u8 file is not a media or master playlist, return 500
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Error decoding m3u8: unknown list type"))
}
