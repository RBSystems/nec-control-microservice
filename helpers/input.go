package helpers

import (
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
)

//SetInput sets the input on the projector
func SetInput(address, port string) error {
	//Map for all the different input hex values
	inputs := map[string]byte{
		"vga1":      0x01,
		"vga2":      0x02,
		"Video":     0x06,
		"Component": 0x10, // AKA YPbPr
		"hdmi1":     0x1A,
		"hdmi2":     0x1B,
		"hdbaset1":  0x20, // AKA HDBaseT
	}

	baseInputCommand := commands["ChangeInput"] //The base command to change input, the input value at position 6 will need to change based on input

	//Switch statment to handle all the input cases
	switch port {
	case "vga1":
		baseInputCommand[6] = inputs["VGA1"]
	case "vga2":
		baseInputCommand[6] = inputs["VGA2"]
	case "Video":
		baseInputCommand[6] = inputs["Video"]
	case "Component":
		baseInputCommand[6] = inputs["Component"]
	case "hdmi1":
		baseInputCommand[6] = inputs["hdmi1"]
	case "hdmi2":
		baseInputCommand[6] = inputs["hdmi2"]
	case "hdbaset1":
		baseInputCommand[6] = inputs["hdbaset1"]
	default:
		break
	}

	command := baseInputCommand
	SendCommand(command, address) //Send the ring for Frodo to deliver

	log.L.Infof("Projector Says: %v", SendCommand) //Print da response!

	return nil
}

//GetInputStatus tells us which input the projector is currently on
func GetInputStatus(address string) (status.Input, error) {

	log.L.Infof("Getting input status from %s ", address)
	command := commands["InputStatus"] //Hex command to get the Input status

	response, err := SendCommand(command, address) //Execute the command, DEW IT
	log.L.Infof("Projector Says: %v", response)    //Print da response!
	if err != nil {
		return status.Input{}, err
	}

	//Have to declare first before using in if statements...so I'll leave it blank
	statuss := status.Input{""}

	responseInputs := map[string]byte{
		"vga1":      0x01,
		"vga2":      0x02,
		"Video":     0x06,
		"Component": 0x10, // AKA YPbPr
		"hdmi":      0x21,
		"hdbaset":   0x27, // AKA HDBaseT
	}

	//Really need to figure out what this byte 33 means... I REMEMBERED!!!! Byte 33 is the hex command for the 21h
	if response[7] == byte(1) && response[8] == responseInputs["hdmi"] {
		statuss = status.Input{
			Input: "hdmi1",
		}
	} else if response[7] == byte(2) && response[8] == responseInputs["hdmi"] {
		statuss = status.Input{
			Input: "hdmi2",
		}
	} else if response[7] == byte(1) && response[8] == responseInputs["hdbaset"] {
		statuss = status.Input{
			Input: "hdbaset1",
		}
	}
	return statuss, nil
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
func GetBlankStatus(address string) (status.Blanked, error) {

	log.L.Infof("Getting blanked status of %s...", address) //Print that the device is powering on

	command := commands["MuteStatus"] //Hex command to get the blanked status (MuteStatus also handles this case)

	response, err := SendCommand(command, address) //Execute the command, DEW IT
	log.L.Debugf("Projector Says: %v", response)   //Print da response!
	if err != nil {

		return status.Blanked{}, err
	}

	var status status.Blanked

	//According to the documentation the 5th byte handles the picture mute
	if response[5] == byte(1) {
		status.Blanked = true
	} else {
		status.Blanked = false
	}
	return status, nil
}
