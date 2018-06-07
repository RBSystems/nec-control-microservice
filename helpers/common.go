package helpers

import (
	"net"
	"time"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
)

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
	getConnection(address)
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

	byteArray = append(byteArray, newbyteArray...)
	return byteArray, nil
}
