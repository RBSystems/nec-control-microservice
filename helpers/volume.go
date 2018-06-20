package helpers

import (
	"fmt"

	"github.com/byuoitav/common/log"
)

//SetVolume makes the projector louder or quieter
func SetVolume(address string, volumeLevel int) error {
	log.L.Infof("Setting volume of %s to %v", address, volumeLevel)

	if volumeLevel > 100 || volumeLevel < 0 {
		err := fmt.Errorf("Invalid volume level %v: must be in range 0-100", volumeLevel)
		log.L.Errorf(err.Error())

		return err
	}

	volume := byte(volumeLevel) //make the volume int to a byte

	volumeCommand := commands["Volume"] //Hex command to change the projector volume

	volumeCommand[8] = volume //Set the 8th byte to change the volume based on documentation

	//Deliver the package
	SendCommand(volumeCommand, address)

	return nil
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
