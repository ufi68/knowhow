package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// go mod init example.com/m
// go mod tidy
// go get "github.com/t-tomalak/logrus-easy-formatter"

func main() {

	// Receive all Arguments from command line and assign to variables
	args := os.Args
	// loop over all arguments and check if a flag is present
	for i := 0; i < len(args); i++ {
		if args[i] == "-f" {
			fmt.Println("Flag -f is present")
			// check if next argument is present
		}
	}

	log.SetOutput(os.Stdout)
	log.Println("Starting up")

	// log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"code": "I00001",
			"foo":  "foo",
			"bar":  "bar",
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
	// Calls panic() after logging - not reached here
	log.Panic("I'm bailing.")

}
