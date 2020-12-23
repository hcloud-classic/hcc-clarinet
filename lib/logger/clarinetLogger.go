package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	goUser "os/user"
	"sync"
	"time"
)

// LogName : Log folder name. Also used as log prefix.
var LogName = "clarinet"
var userLogLocation string

// Logger : Pointer of logger
var Logger *log.Logger
var once sync.Once

// FpLog : File pointer of logger
var FpLog *os.File

func setUserLogFilePath() {
	curUser, _ := goUser.Current()
	userLogLocation = curUser.HomeDir + "/.hcc/clarinet/log/"
}

// CreateDirIfNotExist : Make directory if not exist
func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// Prepare : Prepare logger
func Prepare() bool {
	var err error
	returnValue := false

	// Create directory if not exist
	setUserLogFilePath()
	if _, err = os.Stat(userLogLocation); os.IsNotExist(err) {
		err = createDirIfNotExist(userLogLocation)
		if err != nil {
			log.Fatal(err)
		}
	}

	now := time.Now()

	year := fmt.Sprintf("%d", now.Year())
	month := fmt.Sprintf("%02d", now.Month())
	day := fmt.Sprintf("%02d", now.Day())

	date := year + month + day

	FpLog, err := os.OpenFile(userLogLocation+
		LogName+"_"+date+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		Logger = log.New(io.Writer(os.Stdout), LogName+"_logger: ", log.Ldate|log.Ltime)
		return false
	}

	Logger = log.New(io.MultiWriter(FpLog, os.Stdout), LogName+"_logger: ", log.Ldate|log.Ltime)

	returnValue = true

	return returnValue
}
