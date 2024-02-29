package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/grafov/m3u8"

	"github.com/meanii/aws-lambda-hls-enc/config"
	"github.com/meanii/aws-lambda-hls-enc/helper"
)

// M3u8Media is the handler for m3u8 media playlist
// It will return the media playlist
func M3u8Media(mediapl *m3u8.MediaPlaylist, w http.ResponseWriter, _ *http.Request) {
	mediaPlaylistString := mediapl.String()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(mediaPlaylistString))
}

// M3u8Master is the handler for m3u8 master playlist
// It will return the master playlist
func M3u8Master(
	clurl url.URL,
	masterpl *m3u8.MasterPlaylist,
	w http.ResponseWriter,
	_ *http.Request,
) {
	masterPlaylistString := masterpl.String()
	for _, variant := range masterpl.Variants {

		baseName := path.Base(clurl.String())
		variantUrl := strings.Replace(clurl.String(), baseName, variant.URI, 1)

		// This is the URLSigner struct, it is used to sign the URL
		signedUrlSchema := helper.URLSigner{
			KeyID:         config.ConfigEnv.KeyID,
			PrivKeyBase64: config.ConfigEnv.PrivateKeyBase64,
			Hour:          config.ConfigEnv.ExpireTime,
		}
		signedVariantUrl, err := signedUrlSchema.Signer(variantUrl)
		if err != nil {
			fmt.Printf("Error signing URL: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error signing URL: %s\n", err)))
			return
		}

		// Replace the original variant URL with the signed variant URL
		signedVariantURL, _ := url.Parse(signedVariantUrl)
		signedChunkUrl := variant.URI + "?" + signedVariantURL.RawQuery // Add the query string
		masterPlaylistString = strings.Replace(
			masterPlaylistString,
			variant.URI,
			signedChunkUrl,
			-1,
		)

	}

	// Set the headers and write the response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/x-mpegURL")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(masterPlaylistString))
}
