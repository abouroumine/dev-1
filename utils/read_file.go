package utils

import (
	"dev-1/constants"
	"os"
)

func ReadFile(bytes chan []byte, err chan error) {
	b, er := os.ReadFile(constants.FILENAME)
	err <- er
	bytes <- b
}
