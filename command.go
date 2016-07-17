package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type command struct {
	region       string
	logGroupName string
	start        int64
	limit        int64
	tail         bool
	interleaved  bool
	help         bool
}

// ParseCommand parses the command line and creates a new command to run.
func parseCommand() *command {
	startParam := "1 minute ago"

	command := &command{interleaved: true, limit: 50, tail: false}
	flag.StringVar(&command.region, "region", "", "AWS region to request logs from")
	flag.StringVar(&command.logGroupName, "group", "", "Log group name to read from")
	flag.StringVar(&startParam, "start", "1 minute ago", "The ISO 8601 time that log events should start from")
	flag.Int64Var(&command.limit, "limit", 50, "Number of messages to request")
	flag.BoolVar(&command.tail, "tail", false, "Read log messages continuously")
	flag.BoolVar(&command.interleaved, "interleaved", true, "Interleave log messages between sources")
	flag.Parse()

	if command.help {
		usage()
	}

	if command.region == "" || command.logGroupName == "" {
		usage()
	}

	if startParam == "1 minute ago" {
		command.start = time.Now().Add(1 * time.Minute).Unix()
	}

	return command
}

func usage() {
	fmt.Println("cwl - A command line tool for reviewing Amazon CloudWatch Logs")
	fmt.Println()
	fmt.Println("Parameters:")
	flag.PrintDefaults()
	os.Exit(-1)
}
