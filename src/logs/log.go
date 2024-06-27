package logs

import (
	"log"
)

// Warning create a warning log
func Warning(err error, message string) {

	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print(err)
	log.SetPrefix("WARNING: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print(message)
}
