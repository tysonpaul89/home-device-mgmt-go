package main

import "hrdm/pkg/log"

// Main function initalizes, logger, config files, web servers, etc.
func main() {
	log.Info("This is an Info")
	log.Warn("This is a Warn")
	log.Debug("This is a Debug")
	log.Error("This is a Error")
}
