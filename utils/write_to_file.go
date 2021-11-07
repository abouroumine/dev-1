package utils

import (
	"dev-1/constants"
	"os"
)

func WriteToFile(done chan bool, bytes chan []byte, err chan error) {
	err <- os.WriteFile(constants.FILENAME, <-bytes, 0666)
	<-done
}
