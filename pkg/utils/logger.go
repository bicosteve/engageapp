package utils

import (
	"fmt"
	"log"
)

func Log(level, prefix, msg string, params ...interface{}) {
	message := fmt.Sprintf(msg, params...)
	log.Printf("%s %s %s", level, prefix, message)
}
