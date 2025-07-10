package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

type CloudWatchLogger struct {
	client      *cloudwatchlogs.Client
	logGroup    string
	logStream   string
	sequenceTok *string
}

func NewCloudWatchLogger(logGroup, logStream string) (*CloudWatchLogger, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := cloudwatchlogs.NewFromConfig(cfg)

	return &CloudWatchLogger{
		client:    client,
		logGroup:  logGroup,
		logStream: logStream,
	}, nil
}

func (l *CloudWatchLogger) SendLog(message string) error {
	timestamp := time.Now().UnixMilli()
	input := &cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(l.logGroup),
		LogStreamName: aws.String(l.logStream),
		LogEvents: []types.InputLogEvent{
			{
				Message:   aws.String(message),
				Timestamp: aws.Int64(timestamp),
			},
		},
		SequenceToken: l.sequenceTok,
	}

	output, err := l.client.PutLogEvents(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to send log: %w", err)
	}

	l.sequenceTok = output.NextSequenceToken
	return nil
}
