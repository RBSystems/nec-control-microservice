package helpers

import (
	"fmt"

	se "github.com/byuoitav/av-api/statusevaluators"
	"github.com/byuoitav/common/log"
)

//SetVolume makes the projector louder or quieter
func SetVolume(address string, volumeLevel int) error {
	log.L.Infof("Setting volume of %s to %v", address, volumeLevel)

	if volumeLevel > 32 || volumeLevel < 0 {
		err := fmt.Errorf("Invalid volume level %v: must be in range 0-100", volumeLevel)
		log.L.Errorf(err.Error())

		return err
	}

	volume := byte(volumeLevel) //make the volume int to a byte

	volumeCommand := commands["Volume"] //Hex command to change the projector volume

	volumeCommand[8] = volume //Set the 8th byte to change the volume based on documentation

	//Deliver the package
	response, err := SendCommand(volumeCommand, address)     //Execute the command, DEW IT
	log.L.Infof("The Super Cool Hex Chain is: %v", response) //Print da hex!
	if err != nil {
		return err
	}

	return nil
}

//GetVolumeLevel does just that...or does it?!?
func GetVolumeLevel(address string) (se.Volume, error) {
	log.L.Infof("Getting voulme status of %s...", address) //Print that the device is powering on

	command := commands["VolumeLevel"] //Hex command to get the volume level

	response, err := SendCommand(command, address)           //Execute the command, DEW IT
	log.L.Infof("The Super Cool Hex Chain is: %v", response) //Print da hex!
	if err != nil {
		return se.Volume{}, err
	}

	//TODO: Pick up from here and finish implementation
	// level, err := strconv.Atoi(fields[0])
	// if err != nil {
	// 	return se.Volume{}, err
	// }
	return se.Volume{Volume: 1}, nil
}

//SetMute makes things talk or be silent
func SetMute(address string, muted bool) error {

	var command []byte
	//If muted value is true, mute the projector, else unmute
	if muted {
		log.L.Infof("Muting %s...", address)
		command = commands["MuteOn"] //Hex command to mute the projector
	} else {
		log.L.Infof("Un-muting %s", address)
		command = commands["MuteOff"] //Hex command to mute the projector
	}

	SendCommand(command, address)

	return nil
}

//GetMuteStatus returns if the projector mute status
func GetMuteStatus(address string) (se.MuteStatus, error) {
	log.L.Infof("Getting mute status of %s...", address) //Print that the device is powering on

	command := commands["MuteStatus"] //Hex command to get the Mute status

	response, err := SendCommand(command, address)           //Execute the command, DEW IT
	log.L.Infof("The Super Cool Hex Chain is: %v", response) //Print da hex!
	if err != nil {
		return se.MuteStatus{}, err
	}

	//According to the documentation the 6th byte handles the picture mute
	if response[6] == byte(1) {
		return se.MuteStatus{Muted: true}, nil
	} else {
		return se.MuteStatus{Muted: false}, nil
	}
}
