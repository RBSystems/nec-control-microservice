package helpers

import (
	"github.com/byuoitav/common/log"
)

//SetInputPort sets the input on the projector
func SetInputPort(address, port string) error {
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

	log.L.Infof("Switching input for %s to %s...", address, port) //Tell us what you plan on doing

	baseInputCommand := commands["ChangeInput"] //The base command to change input, the input value at position 6 will need to change based on input

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
	SendCommand(command, address) //Send the ring for Frodo to deliver

	return nil
}

//SetBlank sets the blank status
func SetBlank(address string, blank bool) error {
	log.L.Infof("Setting blank on %s to %v", address, blank)

	var command []byte
	if blank {
		command = commands["ScreenBlankOn"] //Hex command to turn the projector screen blank on
	} else {
		command = commands["ScreenBlankOff"] //Hex command to turn the projector screen blank off
	}

	SendCommand(command, address) //Send the command

	return nil
}
