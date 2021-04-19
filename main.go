package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

var Log Logger = Logger{}

func main() {
	// Create configuration object
	configuration := Configuration{}
	// Initialize logging
	logfolder := "/var/log/mac-api"
	_, err := os.Stat(logfolder)
	if os.IsNotExist(err) || os.IsPermission(err) {
		logfolder = "log"
	}
	Log.Initialize(strings.Join([]string{logfolder, "/log.txt"}, ""))

	config := ""
	config1 := "/etc/mac-api/config.conf"
	config2 := "config/config.conf"
	// Error checking for config files
	_, err1 := os.Stat(config1)
	_, err2 := os.Stat(config2)
	if err1 == nil {
		config = config1
	} else if err2 == nil {
		config = config2
	} else if os.IsNotExist(err1) && os.IsNotExist(err2) {
		Log.Logger.Info().Msg("No configuration file found. Using Commandline parameter.")
	} else if !os.IsNotExist(err1) && os.IsPermission(err1) {
		Log.Logger.Warn().Str("path", config1).Msg("Unable to use configuration file. No permission to access the configuration file.")
	} else if !os.IsNotExist(err2) && os.IsPermission(err2) {
		Log.Logger.Warn().Str("path", config2).Msg("Unable to use configuration file. No permission to access the configuration file.")
	} else if err1 != nil {
		Log.Logger.Warn().Str("error", err1.Error()).Msg("Error while accessing the configuration file.")
	} else if err2 != nil {
		Log.Logger.Warn().Str("error", err2.Error()).Msg("Error while accessing the configuration file.")
	}
	// Try to parse the configuration file if exists
	if config != "" {
		body, err := ioutil.ReadFile(config)
		if err != nil {
			Log.Logger.Warn().Str("error", err.Error()).Msg("Error while reading the configuration file.")
		}
		err = yaml.Unmarshal([]byte(body), &configuration)
		if err != nil {
			Log.Logger.Warn().Str("error", err.Error()).Msg("Error while parsing the configuration file.")
		}
		// Set default configuration parameter
	} else {
		configuration.Address = ":8080"
		configuration.TimeInterval = 86400
		configuration.Logging.Debug = false
	}
	// Commandline flags
	flag.StringVar(&configuration.Address, "port", configuration.Address, "Address for the API to run on (default: ':8080')")
	flag.IntVar(&configuration.TimeInterval, "timeinterval", configuration.TimeInterval, "Time interval when the data should be refreshed in seconds (default: 86400)")
	flag.BoolVar(&configuration.Logging.Debug, "debug", configuration.Logging.Debug, "Option to run the API in debugging mode")
	flag.Parse()
	// Check if debug log should be enabeled
	if configuration.Logging.Debug {
		Log.EnableDebug(true)
	}

	// Create app worker
	a := App{}
	a.Initialize("http://standards-oui.ieee.org/oui/oui.csv", configuration.Address, configuration.TimeInterval)
	//
	a.Run()
}
