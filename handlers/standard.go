package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/nec-control-microservice/helpers"
	"github.com/labstack/echo"
)

//Map for all the different command hex values
var commands = map[string][]byte{
	"PowerOn":        {0x02, 0x00, 0x00, 0x00, 0x00, 0x02},
	"StandBy":        {0x02, 0x01, 0x00, 0x00, 0x00, 0x03},
	"ChangeInput":    {0x02, 0x03, 0x00, 0x00, 0x02, 0x01, 0x00, 0x00},
	"MuteOn":         {0x02, 0x12, 0x00, 0x00, 0x00, 0x14},
	"MuteOff":        {0x02, 0x13, 0x00, 0x00, 0x00, 0x15},
	"Volume":         {0x03, 0x10, 0x00, 0x00, 0x05, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00},
	"ScreenBlankOn":  {0x02, 0x14, 0x00, 0x00, 0x00, 0x16},
	"ScreenBlankOff": {0x02, 0x15, 0x00, 0x00, 0x00, 0x17},
	"PowerStatus":    {0x00, 0x85, 0x00, 0x00, 0x01, 0x01, 0x87},
}

////////////////////////////////////////
//Power Controls
////////////////////////////////////////

//PowerOn helps with turining on the projector
func PowerOn(context echo.Context) error {
	log.L.Infof("Powering on %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")
	data := commands["PowerOn"] //Hex command to turn on the projector

	helpers.SendCommand(data, address)
	return nil
}

//PowerStandby helps with turining on the projector
func PowerStandby(context echo.Context) error {
	log.L.Infof("Going on standby %s...", context.Param("address")) //Print that the device is powering on
	address := context.Param("address")
	command := commands["StandBy"]        //Hex command to turn the projector on standby
	helpers.SendCommand(command, address) //Send the command
	return nil
}

//PowerStatus reports the running status of the projector, on or standby
func PowerStatus(context echo.Context) error {
	log.L.Infof("Getting power status of %s...", context.Param("address")) //Print the device status
	address := context.Param("address")
	command := commands["PowerStatus"] //Hex command to get the power status

	helpers.SendCommand(command, address) //Send the command

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

	baseInputCommand := commands["ChangeInput"] //The base command to change input, the input value at position 6 will need to change based on input

	//Map for all the different input hex values
	inputs := map[string]byte{
		"VGA1":      0x01,
		"VGA2":      0x02,
		"Video":     0x06,
		"Component": 0x10, // AKA YPbPr
		"HDMI1":     0x1A,
		"HDMI2":     0x1B,
		"LAN":       0x20, // AKA HDBaseT
	}

	//Switch statment to handle all the input cases
	switch port {
	case "VGA1":
		baseInputCommand[6] = inputs["VGA1"]
	case "VGA2":
		baseInputCommand[6] = inputs["VGA2"]
	case "Video":
		baseInputCommand[6] = inputs["Video"]
	case "Component":
		baseInputCommand[6] = inputs["Component"]
	case "HDMI1":
		baseInputCommand[6] = inputs["HDMI1"]
	case "HDMI2":
		baseInputCommand[6] = inputs["HDMI2"]
	case "LAN":
		baseInputCommand[6] = inputs["LAN"]
	default:
		break
	}

	command := baseInputCommand

	helpers.SendCommand(command, address) //Send the ring for Frodo to deliver
	return nil
}

//DisplayBlank turns on the Onscreen mute, don't know quite what that means
func DisplayBlank(context echo.Context) error {
	log.L.Infof("Blanking the screen on  %s...", context.Param("address"))
	address := context.Param("address")
	command := commands["ScreenBlankOn"]  //Hex command to turn the projector screen blank on
	helpers.SendCommand(command, address) //Send the command
	return nil
}

//DisplayUnBlank turns off the mysterious Onscreen mute, again, don't know quite what that means
func DisplayUnBlank(context echo.Context) error {
	log.L.Infof("Un-blanking screen on  %s...", context.Param("address"))
	address := context.Param("address")
	command := commands["ScreenBlankOff"] //Hex command to turn the projector screen blank off
	helpers.SendCommand(command, address) //Send the command
	return nil
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
		log.L.Errorf("Error: %v", err.Error())                           //Print out the error is being received
		return context.JSON(http.StatusInternalServerError, err.Error()) //Return that error and a server error
	}
	volume := byte(volumeAsInt) //make the volume int to a byte

	volumeCommand := commands["Volume"] //Hex command to change the projector volume

	volumeCommand[8] = volume //Set the 8th byte to change the volume based on documentation

	//Deliver the package
	helpers.SendCommand(volumeCommand, address)
	return nil
}

//Mute makes the projector be quiet
func Mute(context echo.Context) error {
	log.L.Infof("Muting %s...", context.Param("address"))
	address := context.Param("address")
	command := commands["MuteOn"]         //Hex command to mute the projector
	helpers.SendCommand(command, address) //Send the command
	return nil
}

//UnMute makes the projector noisy again
func UnMute(context echo.Context) error {
	log.L.Infof("Un-muting %s...", context.Param("address"))
	address := context.Param("address")
	command := commands["MuteOff"]        //Hex command to unmute the projector
	helpers.SendCommand(command, address) //Send the command
	return nil
}
