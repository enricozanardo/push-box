package main

import (
	"github.com/goinggo/tracelog"
	"os"
	"net/http"
	"github.com/onezerobinary/push-box/handler"
	"github.com/onezerobinary/push-box/notification"
	"github.com/spf13/viper"
)

const (
	DEFAULT_PORT = "8801"
)

func main() {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	//development environment
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		tracelog.Errorf(err, "main", "main", "Error reading config file")
	}

	tracelog.Warning("main", "main", "Using config file")

	// Start the PushServer
	go notification.StartPushService()

	// Start the API for mobile calls
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	// API calls
	http.HandleFunc("/token", handler.TokenHandler)
	http.HandleFunc("/device", handler.DeviceHandler)
	http.HandleFunc("/position", handler.PositionHandler)
	http.HandleFunc("/status", handler.StatusHandler)
	http.HandleFunc("/mobile", handler.MobileHandler)

	tracelog.Trace("main", "main", "Starting app on port 8801")

	http.ListenAndServe(":"+port, nil)
}
