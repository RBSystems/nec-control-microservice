package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"

	"github.com/byuoitav/nec-control-microservice/helpers"
	"github.com/labstack/echo"
)

////////////////////////////////////////
//Power Controls
////////////////////////////////////////

//PowerOn helps with turining on the projector
func PowerOn(context echo.Context) error {
	address := context.Param("address") //This address will make a fine addition
	err := helpers.PowerOn(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Power{"on"}) //Return JSON for on

}

//PowerStandby helps with turining on the projector
func PowerStandby(context echo.Context) error {
	address := context.Param("address") //Nab that there address
	err := helpers.PowerStandby(address)
	//Do the error checking
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Power{"standby"}) //Return JSON for standby

}

//PowerStatus reports the running status of the projector, on or standby
func PowerStatus(context echo.Context) error {
	address := context.Param("address")

	status, err := helpers.GetPowerStatus(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

////////////////////////////////////////
//Input Controls
////////////////////////////////////////

//SetInputPort sets the input port on the projector
func SetInputPort(context echo.Context) error {
	port := context.Param("port") //Get the port and store it to use later.
	address := context.Param("address")

	log.L.Infof("Switching input for %s to %s...", address, port) //Tell us what you plan on doing

	err := helpers.SetInput(address, port) //Use SetInput to change to the selected port
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}

	log.L.Infof("Done!") //Finished Changing inputs
	return context.JSON(http.StatusOK, status.Input{port})

}

//InputStatus helps us get which input the projector is on
func InputStatus(context echo.Context) error {
	address := context.Param("address")

	status, err := helpers.GetInputStatus(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

//DisplayBlank turns on the Onscreen mute, don't know quite what that means
func DisplayBlank(context echo.Context) error {
	address := context.Param("address")
	log.L.Infof("Blanking Display on %s...", address)

	err := helpers.SetBlank(address, true)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Blanked{true})

}

//DisplayUnBlank turns off the mysterious Onscreen mute, again, don't know quite what that means
func DisplayUnBlank(context echo.Context) error {
	log.L.Infof("Blanking Display...")

	address := context.Param("address")

	err := helpers.SetBlank(address, false)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Blanked{false})
}

//BlankedStatus lets us see de way
func BlankedStatus(context echo.Context) error {
	address := context.Param("address")

	status, err := helpers.GetBlankStatus(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

////////////////////////////////////////
//Volume Controls
////////////////////////////////////////

//SetVolume sends a command to set the projector volume
func SetVolume(context echo.Context) error {
	address := context.Param("address")
	volumeLevel := context.Param("level")

	//change the volume from a string to an int
	level, err := strconv.Atoi(volumeLevel)
	if err != nil {
		return context.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid volume level %s. Must be in range 0-100. %s", volumeLevel, err.Error()))
	}

	err = helpers.SetVolume(address, level)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Volume{level})

}

//VolumeLevel gets us how noisy things are getting
func VolumeLevel(context echo.Context) error {

	address := context.Param("address")

	level, err := helpers.GetVolumeLevel(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, level)
}

//Mute makes the projector be quiet
func Mute(context echo.Context) error {
	address := context.Param("address")

	err := helpers.SetMute(address, true)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Mute{true})
}

//UnMute makes the projector noisy again
func UnMute(context echo.Context) error {
	address := context.Param("address")

	err := helpers.SetMute(address, false)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Mute{false})
}

//MuteStatus returns the Mute status, stating if mute is on or off
func MuteStatus(context echo.Context) error {
	address := context.Param("address")

	status, err := helpers.GetMuteStatus(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}
