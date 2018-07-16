package helpers

import (
	"fmt"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
)

//SetVolume makes the projector louder or quieter
func SetVolume(address string, volumeLevel int) error {
	log.L.Infof("Setting volume of %s to %v", address, volumeLevel)

	//We will need to map these values differently for a scale of 0 to 100 at another time.
	if volumeLevel > 100 || volumeLevel < 0 {
		err := fmt.Errorf("Invalid volume level %v: must be in range 0-31", volumeLevel)
		log.L.Errorf(err.Error())

		return err
	}

	//MAKE THINGS A FLOAT AND ALL YOUR PROBLEMS WILL FLOAT AWAY!
	var level float64

	//Check the volume to see it from 1-100
	log.L.Infof("Volume Level is: %v", volumeLevel)

	//Normalization the get the volume on a 1-100 scale -have to cast volume level to float64 for rounding purposes
	level = (float64(volumeLevel) / 100) * 31

	//make the volume int to a byte
	volume := byte(level)

	//Check to see what volume the projector is being set to.
	log.L.Infof("Volume Level: %v", volume)

	//Hex command to change the projector volume, this isn't even his final form! (Just a temporary holder for original value)
	tempArray := commands["Volume"]
	//Now there are two of them!? This is getting out of hand!
	volumeCommand := make([]byte, len(tempArray))
	//Copy the original array in to the new one as to not change the original
	copy(volumeCommand, tempArray)

	//Set the 8th byte to change the volume based on documentation
	volumeCommand[8] = volume

	//calculate the checksum
	checkSum := getChecksum(volumeCommand)

	//Set the 10th byte to be the checksum as per documentation requirements
	volumeCommand[10] = checkSum

	//Deliver the package
	response, err := SendCommand(volumeCommand, address) //Execute the command, DEW IT
	log.L.Debugf("Projector Says: %v", response)         //Print da response!
	if err != nil {
		return err
	}

	//Thats it, party is over.
	return nil
}

//GetVolumeLevel does just that...or does it?!?
func GetVolumeLevel(address string) (status.Volume, error) {

	log.L.Infof("Getting voulme status of %s...", address) //Print that the device is powering on

	tempArray := commands["VolumeLevel"]
	//Now there are two of them!? This is getting out of hand!
	command := make([]byte, len(tempArray))
	//Copy the original array in to the new one as to not change the original
	copy(command, tempArray)

	//calculate the checksum
	checkSum := getChecksum(command)

	//Set the final value of the array to the checksum as documentation requires
	command[8] = checkSum

	log.L.Debugf("Command Sent to Projector: %v \n", command)

	response, err := SendCommand(command, address)   //Execute the command, DEW IT
	log.L.Debugf("Projector response: %v", response) //Print da response!
	if err != nil {
		return status.Volume{}, err

	}

	//The 12th value of the response is the current volume level, translate that to an int
	volumeLevel := float64(response[12])

	//Renormalize to make it on a scale from 1-100
	levelFloat := (volumeLevel / 31) * 100

	//Change the level to an int because thats what the status.Volume return type is
	level := int(levelFloat)

	//Return the statusevaluator with the current volume level
	return status.Volume{Volume: level}, nil

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
func GetMuteStatus(address string) (status.Mute, error) {

	log.L.Infof("Getting mute status of %s...", address) //Print that the device is powering on

	command := commands["MuteStatus"] //Hex command to get the Mute status

	response, err := SendCommand(command, address)   //Execute the command, DEW IT
	log.L.Debugf("Projector Response: %v", response) //Print da response!
	if err != nil {
		return status.Mute{}, err

	}

	//According to the documentation the 6th byte handles the picture mute
	if response[6] == byte(1) {
		return status.Mute{Muted: true}, nil
	} else {
		return status.Mute{Muted: false}, nil
	}
}
