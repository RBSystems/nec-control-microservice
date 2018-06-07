package handlers

import (
	"encoding/hex"
	"net/http"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/nec-control-microservice/helpers"
	"github.com/labstack/echo"
)

////////////////////////////////////////
//Power Controls
////////////////////////////////////////

//PowerOn helps with turining on the projector
func PowerOn(context echo.Context) error {
	log.L.Infof("Powering on %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")
	data, err := hex.DecodeString("020000000002") //Hex command to turn on the projector
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	helpers.SendCommand(data, address)
	return nil
}

//PowerStandby helps with turining on the projector
func PowerStandby(context echo.Context) error {
	log.L.Infof("Going on standby %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")
	data, err := hex.DecodeString("020100000003") //Hex command to turn the projector on standby
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	helpers.SendCommand(data, address)
	return nil
}

//PowerStatus reports the running status of the projector, on or standby
func PowerStatus(context echo.Context) error {
	log.L.Infof("Getting power status of %s...", context.Param("address")) //Print the device status

	address := context.Param("address")
	data, err := hex.DecodeString("00850000010187") //Hex command to request status
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	helpers.SendCommand(data, address)

	return nil
}

////////////////////////////////////////
//Input Controls
////////////////////////////////////////

//SetInputPort sets the input port on the projector
func SetInputPort(context echo.Context) error {

	return nil
}

////////////////////////////////////////
//Volume Controls
////////////////////////////////////////

//SetVolume sends a command to set the projector volume
func SetVolume(context echo.Context) error {
	address := context.Param("address")
	//volumeLevel := context.Param("level")

	data, err := hex.DecodeString("03100000050500")
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	helpers.SendCommand(data, address)
	return nil
}
