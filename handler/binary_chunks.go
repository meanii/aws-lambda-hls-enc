package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/meanii/aws-lambda-hls-enc/config"
	"github.com/meanii/aws-lambda-hls-enc/helper"
)

// BinaryChunks is the handler for binary chunks
// It will sign the URL and redirect to the signed URL
func BinaryChunksRedirection(cfUrl url.URL, w http.ResponseWriter, r *http.Request) {
	signedUrlSchema := helper.URLSigner{
		KeyID:         config.ConfigEnv.KeyID,
		PrivKeyBase64: config.ConfigEnv.PrivateKeyBase64,
		Hour:          config.ConfigEnv.ExpireTime,
	}

	signedChunkedUrl, err := signedUrlSchema.Signer(cfUrl.String())
	if err != nil {
		fmt.Printf("Error signing URL: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error signing URL: %s\n", err)))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Location", signedChunkedUrl)
	http.Redirect(w, r, signedChunkedUrl, http.StatusFound)
}
