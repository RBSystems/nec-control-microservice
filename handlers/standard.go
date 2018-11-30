package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
	"github.com/byuoitav/common/v2/auth"

	"github.com/byuoitav/nec-control-microservice/helpers"
	"github.com/labstack/echo"
)

////////////////////////////////////////
//Power Controls
////////////////////////////////////////

//PowerOn helps with turining on the projector
func PowerOn(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")                  //Get the address of the display
	log.L.Infof("Setting power of %v to on...", address) //Print that the device is powering on

	err := helpers.PowerOn(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Power{"on"}) //Return JSON for on

}

//PowerStandby helps with turining on the projector
func PowerStandby(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")                       //Nab that there address
	log.L.Infof("Setting power of %v to standby...", address) //Print that the device is powering off

	err := helpers.PowerStandby(address)
	//Do the error checking
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status.Power{"standby"}) //Return JSON for standby

}

//PowerStatus reports the running status of the projector, on or standby
func PowerStatus(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "read-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")
	log.L.Infof("Getting power status of %s...", address) //Print the device status

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
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

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
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "read-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")
	log.L.Infof("Getting input status from %s ", address)

	status, err := helpers.GetInputStatus(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

//DisplayBlank turns on the Onscreen mute, don't know quite what that means
func DisplayBlank(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")
	log.L.Infof("Blanking Display on %s...", address)

	err := helpers.SetBlank(address, true)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.L.Infof("Done!") //The way is open
	return context.JSON(http.StatusOK, status.Blanked{true})

}

//DisplayUnBlank turns off the mysterious Onscreen mute, again, don't know quite what that means
func DisplayUnBlank(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")
	log.L.Infof("Unblanking Display on %s...", address)

	err := helpers.SetBlank(address, false)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	log.L.Infof("Done!") //The path has been closed
	return context.JSON(http.StatusOK, status.Blanked{false})
}

//BlankedStatus lets us see de way
func BlankedStatus(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "read-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")
	log.L.Infof("Getting blanked status of %s...", address) //Print the blank status

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
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

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
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "read-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")

	level, err := helpers.GetVolumeLevel(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, level)
}

//Mute makes the projector be quiet
func Mute(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")
	log.L.Infof("Muting %s...", address)

	err := helpers.SetMute(address, true)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	log.L.Infof("Done!") //The path has been closed

	return context.JSON(http.StatusOK, status.Mute{true})
}

//UnMute makes the projector noisy again
func UnMute(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "write-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")
	log.L.Infof("Unmuting %s...", address)

	err := helpers.SetMute(address, false)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	log.L.Infof("Done!") //The way is open

	return context.JSON(http.StatusOK, status.Mute{false})
}

//MuteStatus returns the Mute status, stating if mute is on or off
func MuteStatus(context echo.Context) error {
	if ok, err := auth.CheckAuthForLocalEndpoints(context, "read-state"); !ok {
		if err != nil {
			log.L.Warnf("Problem getting auth: %v", err.Error())
		}
		return context.String(http.StatusUnauthorized, "unauthorized")
	}

	address := context.Param("address")
	log.L.Infof("Getting mute status of %s...", address) //Print the mute status of the display

	status, err := helpers.GetMuteStatus(address)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

// hasActiveInput will return if the display is displaying or not
func getActiveInput(context echo.Context) error {
	address := context.Param("address")

	active, err := helpers.ActiveInput(address)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, active)
}
