package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/meanii/aws-lambda-hls-enc/config"
	"github.com/meanii/aws-lambda-hls-enc/server"
)

func main() {
	config.InitConfig()
	http.HandleFunc("/", server.HandleRequest)

	// Start the Lambda proxy
	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
