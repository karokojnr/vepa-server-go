package util

import (
	"fmt"
)
func GetPort() string {
	var port = GoDotEnvVariable("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3500"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
