package main

import (
	"net/http"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/common"
	"github.com/byuoitav/nec-control-microservice/handlers"
	"github.com/labstack/echo"
)

func main() {

	// log.SetLevel("debug")

	port := ":8020"
	router := common.NewRouter()

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	//Control endpoints
	secure.GET("/:address/power/on", handlers.PowerOn)
	secure.GET("/:address/power/standby", handlers.PowerStandby)
	secure.GET("/:address/volume/set/:level", handlers.SetVolume)
	secure.GET("/:address/volume/mute", handlers.Mute)
	secure.GET("/:address/volume/unmute", handlers.UnMute)
	secure.GET("/:address/display/blank", handlers.DisplayBlank)
	secure.GET("/:address/display/unblank", handlers.DisplayUnBlank)
	secure.GET("/:address/input/:port", handlers.SetInputPort)

	//status endpoints
	secure.GET("/:address/volume/level", handlers.VolumeLevel)
	secure.GET("/:address/volume/mute/status", handlers.MuteStatus)
	secure.GET("/:address/power/status", handlers.PowerStatus)
	secure.GET("/:address/display/status", handlers.BlankedStatus)
	secure.GET("/:address/input/current", handlers.InputStatus)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
