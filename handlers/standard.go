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
	port := context.Param("port") //Get the port and store it to use later.
	address := context.Param("address")
	log.L.Infof("Switching input for %s to %s...", address, port)

	baseInputCommand := "020300000201" //The base command to change input, the input value needs to be appended on the end

	//Map for all the different input hex values
	inputs := map[string]string{
		"VGA1":      "01",
		"VGA2":      "02",
		"Video":     "06",
		"Component": "10", // AKA YPbPr
		"HDMI1":     "1A",
		"HDMI2":     "1B",
		"LAN":       "20", // AKA HDBaseT
	}

	var data string //Just a place holder to make it stop screaming at me

	//Switch statment to handle all the input cases
	switch port {
	case "VGA1":
		data = baseInputCommand + inputs["VGA1"]
	case "VGA2":
		data = baseInputCommand + inputs["VGA2"]
	case "Video":
		data = baseInputCommand + inputs["Video"]
	case "Component":
		data = baseInputCommand + inputs["Component"]
	case "HDMI1":
		data = baseInputCommand + inputs["HDMI1"]
	case "HDMI2":
		data = baseInputCommand + inputs["HDMI2"]
	case "LAN":
		data = baseInputCommand + inputs["LAN"]
	default:
		data = baseInputCommand
	}

	command, err := hex.DecodeString(data) //Hex command to request status

	//Obligatory error checking...
	if err != nil {
		log.L.Errorf("Error: %v", err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}

	helpers.SendCommand(command, address) //Send the ring for Frodo to deliver
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
