package helpers

import (
	"github.com/byuoitav/common/log"
)

//PowerOn sends the command for the power to turn on
func PowerOn(address string) error {
	log.L.Infof("Setting power of %v to on...", address) //Print that the device is powering on
	command := commands["PowerOn"]                       //Hex command to turn on the projector

	SendCommand(command, address) //Execute the command, DEW IT
	return nil
}

//PowerStandby sends the command for the power to go to sleep
func PowerStandby(address string) error {
	log.L.Infof("Setting power of %v to standby...", address) //Print that the device is powering on
	command := commands["PowerOn"]                            //Hex command to turn off the projector

	SendCommand(command, address) //Execute the command, DEW IT
	return nil
}

//GetPowerStatus will give the power status of the projector
func GetPowerStatus(address string) error {
	log.L.Infof("Getting power status of %s...", address) //Print the device status

	command := commands["PowerStatus"] //Hex command to get the power status

	SendCommand(command, address) //Execute the command, DEW IT

	return nil
}
