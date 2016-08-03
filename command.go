package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type command struct {
	profile        string
	region         string
	logGroupName   string
	filter         string
	start          time.Time
	end            time.Time
	limit          int64
	tail           bool
	refresh        int64
	help           bool
	streams        []string
	fullStreamName bool
}

// ParseCommand parses the command line and creates a new command to run.
func parseCommand() *command {
	startParam := "1 minute ago"
	endParam := "now"
	streamParam := ""

	command := &command{limit: 50, tail: false}

	flag.StringVar(&command.profile, "profile", "", "AWS credential profile to use.")
	flag.StringVar(&command.region, "region", "", "AWS region to request logs from")
	flag.StringVar(&command.logGroupName, "group", "", "Log group name to read from")
	flag.StringVar(&streamParam, "streams", "", "List of streams, comma separated.")
	flag.StringVar(&command.filter, "filter", "", "Filter pattern to apply")
	flag.StringVar(&startParam, "start", "1 minute ago", "The RFC3339 time that log events should start from")
	flag.StringVar(&endParam, "end", "now", "The RFC3339 time that log events should end")

	//flag.BoolVar(&command.tail, "tail", false, "Read log messages continuously")
	//flag.Int64Var(&command.refresh, "refresh", 5, "Refresh rate for tailing logs, in seconds.")
	flag.BoolVar(&command.fullStreamName, "fullStreamName", false, "Display full stream names.")
	flag.Parse()

	if command.help {
		usage()
	}

	if command.region == "" || command.logGroupName == "" {
		usage()
	}

	startTime, err := parseTime(startParam)
	if err != nil {
		fmt.Printf("Could not parse start time.\n")
		fmt.Println()
		usage()
	}
	command.start = startTime

	endTime, err := parseTime(endParam)
	if err != nil {
		fmt.Printf("Could not parse end time.\n")
		fmt.Println()
		usage()
	}
	command.end = endTime

	if streamParam != "" {
		command.streams = strings.Split(streamParam, ",")
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

// Parses a time string into something reasonable
func parseTime(timeString string) (time.Time, error) {
	if timeString == "now" {
		return time.Now(), nil
	}

	// handle fuzzy time strings
	re := regexp.MustCompile("([0-9]+)\\s+(second|minute|hour|day)s?\\s+ago")
	matches := re.FindStringSubmatch(timeString)
	if len(matches) == 3 {
		value, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			return time.Now(), err
		}

		var duration time.Duration
		switch matches[2] {
		case "second":
			duration = time.Second
		case "minute":
			duration = time.Minute
		case "hour":
			duration = time.Hour
		case "day":
			duration = time.Hour * 24
		default:
			return time.Now(), fmt.Errorf("Unknown time unit, %s", matches[2])
		}

		return time.Now().Add(-1 * time.Duration(value) * duration), nil
	}

	rfcTime, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return time.Now(), err
	}
	return rfcTime, nil
}
