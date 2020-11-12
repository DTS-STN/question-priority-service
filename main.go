package main

import (
	_ "github.com/DTS-STN/question-priority-service/docs"
	"github.com/DTS-STN/question-priority-service/server"
	"os"
)

// @title Question Prioritization Service
// @version 1.0
// @description This is a service to return questions that when answered will return elegible benefits.

// @host https://qps-dev.dev.dts-stn.com
// @BasePath /

func main() {
	//Start the service
	server.Main(os.Args)
}
