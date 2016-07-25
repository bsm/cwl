package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func main() {
	command := parseCommand()

	config := &aws.Config{Region: aws.String(command.region)}
	if command.profile != "" {
		cred := credentials.NewSharedCredentials("", command.profile)
		_, err := cred.Get()
		if err != nil {
			fmt.Printf("Error: Could not retreive profile `%s` from credentials.\n", command.profile)
			os.Exit(-1)
		}

		config.Credentials = cred
	}

	svc := cloudwatchlogs.New(session.New(config))

	nextToken := readAndPrintLogItems(svc, command, nil)
	for nextToken != nil {
		nextToken = readAndPrintLogItems(svc, command, nextToken)
	}
}

func readAndPrintLogItems(svc *cloudwatchlogs.CloudWatchLogs, comm *command,
	nextToken *string) *string {

	params := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName:  aws.String(comm.logGroupName),
		FilterPattern: aws.String(comm.filter),
		Interleaved:   aws.Bool(comm.interleaved),
		Limit:         aws.Int64(comm.limit),
		//LogStreamNames: []*string{
		//  aws.String("LogStreamName"), // Required
		// More values...
		//},
		NextToken: nextToken,
		StartTime: aws.Int64(comm.start.UTC().Unix() * 1000),
		EndTime:   aws.Int64(comm.end.UTC().Unix() * 1000),
	}
	resp, err := svc.FilterLogEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil
	}

	printLogItems(resp.Events)

	return resp.NextToken
}

func printLogItems(events []*cloudwatchlogs.FilteredLogEvent) {
	for _, event := range events {
		fmt.Printf("%v\n", *event.Message)
	}
}
