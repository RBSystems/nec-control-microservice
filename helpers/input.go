package helpers

import (
	"github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/log"
)

//SetInput sets the input on the projector
func SetInput(address, port string) error {
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
	case "hdmi1":
		baseInputCommand[6] = inputs["HDMI1"]
	case "hdmi2":
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

//GetInputStatus tells us which input the projector is currently on
func GetInputStatus(address string) (statusevaluators.Input, error) {
	log.L.Infof("Getting input status from %s ", address)
	command := commands["InputStatus"] //Hex command to get the Input status

	response, err := SendCommand(command, address) //Execute the command, DEW IT
	log.L.Debugf("Projector Says: %v", response)   //Print da response!
	if err != nil {
		return statusevaluators.Input{}, err
	}

	//Have to declare first before using in if statements...so I'll leave it blank
	status := statusevaluators.Input{""}

	if response[7] == byte(1) && response[8] == byte(33) {
		status = statusevaluators.Input{
			Input: "hdmi1",
		}
	} else if response[7] == byte(2) && response[8] == byte(33) {
		status = statusevaluators.Input{
			Input: "hdmi2",
		}
	}
	return status, nil

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

//GetBlankStatus tells us if the picture is blanked or not.
func GetBlankStatus(address string) (statusevaluators.BlankedStatus, error) {
	log.L.Infof("Getting blanked status of %s...", address) //Print that the device is powering on

	command := commands["MuteStatus"] //Hex command to get the blanked status (MuteStatus also handles this case)

	response, err := SendCommand(command, address) //Execute the command, DEW IT
	log.L.Debugf("Projector Says: %v", response)   //Print da response!
	if err != nil {
		return statusevaluators.BlankedStatus{}, err
	}

	var status statusevaluators.BlankedStatus

	//According to the documentation the 5th byte handles the picture mute
	if response[5] == byte(1) {
		status.Blanked = true
	} else {
		status.Blanked = false
	}
	return status, nil
}
