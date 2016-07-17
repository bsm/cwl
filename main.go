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
		readAndPrintLogItems(svc, command, nil)
	}
}

func tail(svc *cloudwatchlogs.CloudWatchLogs, comm *command) {
	var nextToken *string

	for {
		nextToken, nextTime := readAndPrintLogItems(svc, comm, nextToken)
		if nextToken == nil {
			comm.start = time.Unix(nextTime/1000, 0)
		}
		time.Sleep(time.Duration(comm.refresh) * time.Second)
	}
}

func readAndPrintLogItems(svc *cloudwatchlogs.CloudWatchLogs, comm *command,
	nextToken *string) (*string, int64) {

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
		NextToken: nextToken,
		StartTime: aws.Int64(comm.start.UTC().Unix() * 1000),
	}
	resp, err := svc.FilterLogEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil, 0
	}

	printLogItems(resp.Events)

	return resp.NextToken, *resp.Events[len(resp.Events)-1].Timestamp
}

func printLogItems(events []*cloudwatchlogs.FilteredLogEvent) {
	for _, event := range events {
		fmt.Printf("%v\n", *event.Message)
	}
}
