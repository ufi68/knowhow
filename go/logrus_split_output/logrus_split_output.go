package main

import (
	"os"

	logrus "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func main() {
	infoLogger := &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% [%lvl%] %msg% %customField% \n",
		},
	}
	infoLogger.SetLevel(logrus.TraceLevel)

	// Send to StdErr
	errLogger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% [%lvl%] %msg% %customField% \n",
		},
	}

	infoLogger.Printf("Log message")

	infoLogger.WithFields(
		logrus.Fields{
			"code": "I00001",
			"foo":  "foo",
			"bar":  "bar",
		},
	).Info("Something happened")

	errLogger.WithFields(
		logrus.Fields{
			"code": "I00001",
			"foo":  "foo",
			"bar":  "bar",
		},
	).Error("Something happened")

	errLogger.Printf("Error message")

	errLogger.Error("Error message")

}
