package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("This is A log file")
	log.Warn("Another log file")
	log.Fatal("Just failed...")
}
