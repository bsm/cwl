package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func main() {
	command := parseCommand()
	svc := cloudwatchlogs.New(session.New(&aws.Config{Region: aws.String(command.region)}))

	if command.tail {
		tail(svc, command)
	} else {
		readAndPrintLogItems(svc, command, "")
	}
}

func tail(svc *cloudwatchlogs.CloudWatchLogs, comm *command) {
	var nextToken string

	for {
		fmt.Println("-----------------------")
		nextToken := readAndPrintLogItems(svc, comm, nextToken)
		if nextToken == "" {
			break
		}
		time.Sleep(5)
	}
}

func readAndPrintLogItems(svc *cloudwatchlogs.CloudWatchLogs, comm *command,
	nextToken string) string {

	var nextTokenPtr *string
	if nextToken == "" {
		nextTokenPtr = nil
	} else {
		nextTokenPtr = &nextToken
	}

	params := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName: aws.String(comm.logGroupName),
		//EndTime:       aws.Int64(1),
		//FilterPattern: aws.String("FilterPattern"),
		Interleaved: aws.Bool(comm.interleaved),
		Limit:       aws.Int64(comm.limit),
		//LogStreamNames: []*string{
		//  aws.String("LogStreamName"), // Required
		// More values...
		//},
		NextToken: nextTokenPtr,
		StartTime: aws.Int64(comm.start),
	}
	resp, err := svc.FilterLogEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return ""
	}

	printLogItems(resp.Events)

	return *resp.NextToken
}

func printLogItems(events []*cloudwatchlogs.FilteredLogEvent) {
	for i := range events {
		event := events[len(events)-i-1]
		fmt.Printf("%v\n", *event.Message)
	}
}
