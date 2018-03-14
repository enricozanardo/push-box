package main

import (
	"github.com/goinggo/tracelog"
	"os"
	"net/http"
	"github.com/onezerobinary/push-box/handler"
)

const (
	DEFAULT_PORT = "8801"
)

func main() {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/token", handler.TokenHandler)

	tracelog.Trace("main", "main", "Starting app on port 8801")

	http.ListenAndServe(":"+port, nil)
}
