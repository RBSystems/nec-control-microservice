package main

import (
	"net/http"

	"github.com/byuoitav/common"
	"github.com/byuoitav/nec-control-microservice/handlers"
)

func main() {

	// log.SetLevel("debug")

	port := ":8020"
	router := common.NewRouter()

	// Use the `secure` routing group to require authentication
	//secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	//Control endpoints
	router.GET("/:address/power/on", handlers.PowerOn)
	router.GET("/:address/power/standby", handlers.PowerStandby)
	router.GET("/:address/volume/set/:level", handlers.SetVolume)
	router.GET("/:address/volume/mute", handlers.Mute)
	router.GET("/:address/volume/unmute", handlers.UnMute)
	router.GET("/:address/display/blank", handlers.DisplayBlank)
	router.GET("/:address/display/unblank", handlers.DisplayUnBlank)
	router.GET("/:address/input/:port", handlers.SetInputPort)

	//status endpoints
	router.GET("/:address/volume/level", handlers.VolumeLevel)
	router.GET("/:address/volume/mute/status", handlers.MuteStatus)
	router.GET("/:address/power/status", handlers.PowerStatus)
	router.GET("/:address/display/status", handlers.BlankedStatus)
	router.GET("/:address/input/current", handlers.InputStatus)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
