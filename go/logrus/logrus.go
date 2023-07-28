package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// go mod init example.com/m
// go mod tidy
// go get "github.com/t-tomalak/logrus-easy-formatter"

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Starting up")

	// log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Info("Something happened")

	log.SetFormatter(&log.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.SetLevel(log.TraceLevel)

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")

}
