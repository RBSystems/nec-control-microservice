package helpers

import (
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/status"
)

//PowerOn sends the command for the power to turn on
func PowerOn(address string) error {
	log.L.Infof("Setting power of %v to on...", address) //Print that the device is powering on
	command := commands["PowerOn"]                       //Hex command to turn on the projector
	response, err := SendCommand(command, address)       //Execute the command, DEW IT
	log.L.Debugf("%v", response)
	if err != nil {
		log.L.Info("Nope Didn't work! - %v", err.Error())
		return err
	}

	return nil
}

//PowerStandby sends the command for the power to go to sleep
func PowerStandby(address string) error {
	log.L.Infof("Setting power of %v to standby...", address) //Print that the device is powering on
	command := commands["Standby"]                            //Hex command to turn off the projector
	response, err := SendCommand(command, address)            //Execute the command, DEW IT
	log.L.Debugf("%v", response)
	if err != nil {
		log.L.Info("Nope Didn't work! - %v", err.Error())
		return err
	}

	return nil
}

//GetPowerStatus will give the power status of the projector
func GetPowerStatus(address string) (status.Power, error) {

	log.L.Infof("Getting power status of %s...", address) //Print the device status
	command := commands["PowerStatus"]                    //Hex command to get the power status

	response, err := SendCommand(command, address)   //Execute the command, DEW IT
	log.L.Debugf("Projector Response: %v", response) //Print da response!
	if err != nil {
		return status.Power{}, err
	}

	var status status.Power

	if response[7] == byte(0) {
		status.Power = "standby"
	} else if response[7] == byte(1) {
		status.Power = "on"
	}

	return status, nil
}
