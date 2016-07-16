package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func main() {
	svc := cloudwatchlogs.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))

	params := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName: aws.String("/var/log/messages"),
		//EndTime:       aws.Int64(1),
		//FilterPattern: aws.String("FilterPattern"),
		Interleaved: aws.Bool(true),
		Limit:       aws.Int64(50),
		//LogStreamNames: []*string{
		//  aws.String("LogStreamName"), // Required
		// More values...
		//},
		//NextToken: aws.String("NextToken"),
		//StartTime: aws.Int64(1),
	}
	resp, err := svc.FilterLogEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
