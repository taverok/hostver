package main

import (
	"encoding/json"
	"fmt"
	"taverok/hostver/internal/client/app"
	"taverok/hostver/internal/client/node"
)

func main() {
	appService := app.Service{}
	nodeService := node.Service{AppService: appService}

	currentNode, err := nodeService.CurrentNode()
	if err != nil {
		panic(err)
	}

	indent, err := json.MarshalIndent(currentNode, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(indent))
}
