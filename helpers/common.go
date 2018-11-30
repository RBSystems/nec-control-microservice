package helpers

import (
	"net"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

var commands = map[string][]byte{
	"PowerOn":        {0x02, 0x00, 0x00, 0x00, 0x00, 0x02}, // Powers on the projector
	"Standby":        {0x02, 0x01, 0x00, 0x00, 0x00, 0x03}, // Puts the projector on standby
	"ChangeInput":    {0x02, 0x03, 0x00, 0x00, 0x02, 0x01, 0x00, 0x00},
	"MuteOn":         {0x02, 0x12, 0x00, 0x00, 0x00, 0x14},
	"MuteOff":        {0x02, 0x13, 0x00, 0x00, 0x00, 0x15},
	"Volume":         {0x03, 0x10, 0x00, 0x00, 0x05, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00}, // Used for changing the volume level of the projector
	"ScreenBlankOn":  {0x02, 0x10, 0x00, 0x00, 0x00, 0x12},                               // Blanks the screen
	"ScreenBlankOff": {0x02, 0x11, 0x00, 0x00, 0x00, 0x13},                               // Unblanks the screen
	"PowerStatus":    {0x00, 0x85, 0x00, 0x00, 0x01, 0x01, 0x87},                         // Used for getting the projector power status
	"MuteStatus":     {0x00, 0x85, 0x00, 0x00, 0x01, 0x03, 0x89},                         // Used for getting the mute status
	"VolumeLevel":    {0x03, 0x04, 0x00, 0x00, 0x03, 0x05, 0x00, 0x00, 0x00},             // Used for getting the volume level
	"InputStatus":    {0x00, 0x85, 0x00, 0x00, 0x01, 0x02, 0x88},                         // Used for retreiving the current input
	"ActiveInput":    {0x00, 0xBF, 0x00, 0x00, 0x01, 0x02, 0xC2},                         // Used for getting the ActiveInput Status of the projector
}

// getConnection establishes a TCP connection with the projector
func getConnection(address string) (*net.TCPConn, *nerr.E) {
	log.L.Debugf("Getting connection for %v", address)
	radder, err := net.ResolveTCPAddr("tcp", address+":7142") //Resolve the TCP connection on Port 7142 that is required by NEC
	if err != nil {
		nerr.Translate(err).Addf("Could not get connection for %v", address)
	}

	conn, err := net.DialTCP("tcp", nil, radder)
	if err != nil {
		nerr.Translate(err).Addf("Could not get connection for %v", address)
	}
	log.L.Debugf("Done!")
	return conn, nil
}

// SendCommand sends the byte array to the desired address of projector
func SendCommand(command []byte, address string) ([]byte, *nerr.E) {
	log.L.Debugf("Sending command %x, to %v", command, address)
	conn, err := getConnection(address)
	if err != nil {
		return []byte{}, err.Addf("Could not send command")
	}

	conn.SetWriteDeadline(time.Now().Add(3 * time.Second))

	numwritten, commandError := conn.Write(command)
	if numwritten != len(command) {
		return []byte{}, err.Addf("The command written was not the same length as the given command")
	}
	if commandError != nil {
		return []byte{}, nerr.Translate(err).Addf("Could not get value")
	}

	byteArray := make([]byte, 5)
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	numRead, readError := conn.Read(byteArray)
	if numRead != len(byteArray) {
		return []byte{}, nerr.Create("Couldn't read back command response", "error")
	}
	if readError != nil {
		return []byte{}, nerr.Translate(err).Addf("Could not get byte array value")
	}

	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	newbyteArray := make([]byte, (uint8)(byteArray[4])+1)
	numRead, readError = conn.Read(newbyteArray)
	if numRead != len(newbyteArray) {
		return []byte{}, nerr.Create("Couldn't read back command response", "error")
	}
	if readError != nil {
		return []byte{}, nerr.Translate(err).Addf("Could not get new byte array value")
	}

	//byteArray is the response back from the projector
	byteArray = append(byteArray, newbyteArray...)

	//close the connection
	defer conn.Close()

	//This is the response, which is also a byte array, and the nil error
	return byteArray, nil
}

//getChecksum returns the checksum value for the end of the hex array
func getChecksum(command []byte) byte {
	var checksum byte
	for i := 0; i < len(command); i++ {
		checksum += command[i]
	}
	return checksum
}
