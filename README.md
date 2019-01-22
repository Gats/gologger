# gologger
A simple logging package for golang projects

# Installation
Package is go gettable (Not tested yet) using go get -u github.com/TrueAfrican/gologger

# Usage

First initate the package using a log file as below

logfile := "/var/log/example_log.log"
	l, err := logger.New(logfile)
	if err != nil {
		panic(err)
	}

Write to the log as below
There are 3 types of logs provided Info, Errror and Applogs
Info - Just a log for information purposes, prefixed using INFO
Error - Error log, prfixed by ERROR
Applogs - Used for logging http request and response logs

# Examples

l.AppLogs("This is a response message", "RESPONSE")

l.AppLogs("This is a response message", "REQUEST")

l.INFO("This is an information message")

l.ERROR("This is an error message")
