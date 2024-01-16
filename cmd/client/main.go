package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"taverok/hostver/internal/client/app"
	"taverok/hostver/internal/client/node"
)

var pushUrl string

func main() {
	parseFlags()
	appService := app.Service{}
	nodeService := node.Service{AppService: appService}

	currentNode, err := nodeService.CurrentNode()
	if err != nil {
		log.Panic(err)
	}

	jsonBody, err := currentNode.AsJson()
	if err != nil {
		log.Panic(err)
	}

	if pushUrl == "" {
		fmt.Println(jsonBody)
		return
	}

	send(err, jsonBody)
}

func send(err error, json string) {
	res, err := http.Post(pushUrl, "application/json", strings.NewReader(json))
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		log.Panicf("Push failed with status code %d", res.StatusCode)
	}
}

func parseFlags() {
	flag.StringVar(&pushUrl, "u", "", "Push URL")
	flag.Parse()
}
