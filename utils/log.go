package utils

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/pborman/uuid"
	"path/filepath"
)

const defaultLogPath = "/run/runc.log"

var (
	logger = NewLogger()
)

func GetLogger() *logrus.Logger {
	return logger
}

func NewLogger() *logrus.Logger {
	var (
		w io.Writer = ioutil.Discard
	)

	fname := filepath.Join("/tmp", getRandomString() + ".log")

	f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_SYNC, 0x644)
	if err == nil {
		//if open file ok, set writer to file, else to /dev/null
		w = f
	}

	return &logrus.Logger{
		Out:       w,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}
}

func getRandomString() string {
	return uuid.NewRandom().String()
}