package main

import (
	"kalle83/feedback-roulette/cmd"

	log "github.com/sirupsen/logrus"
)

func initLogging() {
	// log.Trace("Something very low level.")
	// log.Debug("Useful debugging information.")
	// log.Info("Something noteworthy happened!")
	// log.Warn("You should probably take a look at this.")
	// log.Error("Something failed but I'm not quitting.")
	// // Calls os.Exit(1) after logging
	// log.Fatal("Bye.")
	// // Calls panic() after logging
	// log.Panic("I'm bailing.")

	log.SetLevel(log.DebugLevel)
}

func main() {

	initLogging()
	cmd.Execute()
}
