package log

import (
	"log"
	"os"
)

var (
	// Info for infomation output
	Info *log.Logger
	// Error for Error message output
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}
