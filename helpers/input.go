package helpers

import (
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
	"github.com/byuoitav/common/status"
)

//SetInput sets the input on the projector
func SetInput(address, port string) error {
	//Map for all the different input hex values
	inputs := map[string]byte{
		"computer":  0x01,
		"video":     0x06, // Yellow cable for RYW
		"component": 0x10,
		"hdmi1":     0xA1,
		"hdmi2":     0xA2,
		"hdbaset1":  0xBF, // AKA HDBaseT
	}

	//Hex command to change the projector volume, it is going in a temporary Array as to not change its original value
	tempArray := commands["ChangeInput"]

	//Make an empty byte array the size of the original command
	inputCommand := make([]byte, len(tempArray))

	//Copy the original array in to the new one as to not change the original
	copy(inputCommand, tempArray)

	//Switch statment to handle all the input cases
	switch port {
	case "computer":
		inputCommand[6] = inputs["computer"]
	case "video":
		inputCommand[6] = inputs["video"]
	case "component":
		inputCommand[6] = inputs["component"]
	case "hdmi1":
		inputCommand[6] = inputs["hdmi1"]
	case "hdmi2":
		inputCommand[6] = inputs["hdmi2"]
	case "hdbaset1":
		inputCommand[6] = inputs["hdbaset1"]
	default:
		break
	}

	checkSum := getChecksum(inputCommand) //Calculate the checksum

	inputCommand[7] = checkSum //Put the checksum value at the end of the change input command

	log.L.Debugf("Command being sent is: %v", inputCommand) //Print out the command that will be sent in case I want to verify

	response, err := SendCommand(inputCommand, address) //Send the ring for Frodo to deliver
	log.L.Debugf("Projector Says: %v", response)        //Print da response!
	if err != nil {
		return err
	}

	return nil
}

//GetInputStatus tells us which input the projector is currently on
func GetInputStatus(address string) (status.Input, error) {

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

// ActiveInput will tell if a projector has an active input or not
func ActiveInput(address string) (bool, *nerr.E) {
	log.L.Debugf("Checking for active input for %s", address)

	// operationStatus is a map of the commands according to the documentation from NEC
	// operationStatus := map[string]byte{
	// 	"Standby":          0x00,
	// 	"StandbyError":     0x06, // If there is an error with the projector this will be the satus we want to see
	// 	"StandbyPowerSave": 0x0F,
	// 	"NetworkStandby":   0x10, // This is the standby that the projectors should be using
	// 	"PowerOn":          0x04, // I believe this means the projector is on...
	// }

	//TODO: Implement the operationStatus to start getting information back. See pg. 83 of the documentation.

	command := commands["AcitveInput"]

	response, err := SendCommand(command, address)
	if err != nil {
		return false, err.Add("Could not get an active signal")
	}
	log.L.Debugf("Projector response: %v", response)

	return false, nil
}
