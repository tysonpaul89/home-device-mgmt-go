package main

import "hrdm/logger"

// Main function initalizes, logger, config files, web servers, etc.
func main() {
  log := logger.NewLogger()
  
  log.Info("Hello World!")
}
