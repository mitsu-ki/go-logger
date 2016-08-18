package main

import (
	log "github.com/mitsu-ki/go-logger"
)

func main() {
	log.Info("The default log level is:", "INFO")
	log.Debug("That means that this message is ignored")

	log.SetLogLevel(log.DEBUG)
	log.Debug("Changing the log level is easy")

	log.Fatal("And a 'Fatal' log call will exit the application")
}
