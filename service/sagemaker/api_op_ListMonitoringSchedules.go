// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns list of all monitoring schedules.
func (c *Client) ListMonitoringSchedules(ctx context.Context, params *ListMonitoringSchedulesInput, optFns ...func(*Options)) (*ListMonitoringSchedulesOutput, error) {
	stack := middleware.NewStack("ListMonitoringSchedules", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListMonitoringSchedulesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMonitoringSchedules(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListMonitoringSchedules",
			Err:           err,
		}
	}
	out := result.(*ListMonitoringSchedulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMonitoringSchedulesInput struct {

	// The maximum number of jobs to return in the response. The default value is 10.
	MaxResults *int32

	// A filter that returns only monitoring schedules modified after a specified time.
	LastModifiedTimeAfter *time.Time

	// Name of a specific endpoint to fetch schedules for.
	EndpointName *string

	// Whether to sort results by Status, CreationTime, ScheduledTime field. The
	// default is CreationTime.
	SortBy types.MonitoringScheduleSortKey

	// Whether to sort the results in Ascending or Descending order. The default is
	// Descending.
	SortOrder types.SortOrder

	// The token returned if the response is truncated. To retrieve the next set of job
	// executions, use it in the next request.
	NextToken *string

	// A filter that returns only monitoring schedules created after a specified time.
	CreationTimeAfter *time.Time

	// Filter for monitoring schedules whose name contains a specified string.
	NameContains *string

	// A filter that returns only monitoring schedules created before a specified time.
	CreationTimeBefore *time.Time

	// A filter that returns only monitoring schedules modified before a specified
	// time.
	StatusEquals types.ScheduleStatus

	// A filter that returns only monitoring schedules modified before a specified
	// time.
	LastModifiedTimeBefore *time.Time
}

type ListMonitoringSchedulesOutput struct {

	// A JSON array in which each element is a summary for a monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleSummaries []*types.MonitoringScheduleSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of jobs, use it in the subsequent reques
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListMonitoringSchedulesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListMonitoringSchedules{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMonitoringSchedules{}, middleware.After)
}

func newServiceMetadataMiddleware_opListMonitoringSchedules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListMonitoringSchedules",
	}
}
