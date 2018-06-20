package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	se "github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/log"
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
	return context.JSON(http.StatusOK, se.PowerStatus{"on"}) //Return JSON for on
}

//PowerStandby helps with turining on the projector
func PowerStandby(context echo.Context) error {
	address := context.Param("address") //Nab that there address
	err := helpers.PowerStandby(address)
	//Do the error checking
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, se.PowerStatus{"standby"}) //Return JSON for standby
}

//PowerStatus reports the running status of the projector, on or standby
func PowerStatus(context echo.Context) error {
	//address := context.Param("address")

	// status, err := helpers.GetPowerStatus(address)
	// if err != nil {
	// 	return context.JSON(http.StatusInternalServerError, err.Error())
	// }

	// return context.JSON(http.StatusOK, status)
	return nil
}

////////////////////////////////////////
//Input Controls
////////////////////////////////////////

//SetInput sets the input port on the projector
func SetInput(context echo.Context) error {
	port := context.Param("port") //Get the port and store it to use later.
	address := context.Param("address")

	err := helpers.SetInputPort(address, port) //Use SetInput to change to the selected port
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	return context.JSON(http.StatusOK, se.Input{port})
}

//DisplayBlank turns on the Onscreen mute, don't know quite what that means
func DisplayBlank(context echo.Context) error {
	address := context.Param("address")
	log.L.Infof("Blanking Display...")

	err := helpers.SetBlank(address, true)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.BlankedStatus{true})
}

//DisplayUnBlank turns off the mysterious Onscreen mute, again, don't know quite what that means
func DisplayUnBlank(context echo.Context) error {
	log.L.Infof("Blanking Display...")

	address := context.Param("address")

	err := helpers.SetBlank(address, false)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.BlankedStatus{false})
}

////////////////////////////////////////
//Volume Controls
////////////////////////////////////////

//SetVolume sends a command to set the projector volume
func SetVolume(context echo.Context) error {
	address := context.Param("address")
	volumeLevel := context.Param("level")

	//Make the volume level an int instead of a string.
	volumeAsInt, err := strconv.Atoi(volumeLevel)
	if err != nil {
		return context.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid volume level %s. Must be in range 0-100. %s", volumeLevel, err.Error()))
	}
	err = helpers.SetVolume(address, volumeAsInt)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, se.Volume{volumeAsInt})
}

//Mute makes the projector be quiet
func Mute(context echo.Context) error {
	address := context.Param("address")

	err := helpers.SetMute(address, true)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.MuteStatus{true})
}

//UnMute makes the projector noisy again
func UnMute(context echo.Context) error {
	address := context.Param("address")

	err := helpers.SetMute(address, false)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.MuteStatus{false})
}
