// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates a destination. This operation is used only to create
// destinations for cross-account subscriptions. A destination encapsulates a
// physical resource (such as an Amazon Kinesis stream) and enables you to
// subscribe to a real-time stream of log events for a different account, ingested
// using PutLogEvents
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html).
// Through an access policy, a destination controls what is written to it. By
// default, PutDestination does not set any access policy with the destination,
// which means a cross-account user cannot call PutSubscriptionFilter
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutSubscriptionFilter.html)
// against this destination. To enable this, the destination owner must call
// PutDestinationPolicy
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestinationPolicy.html)
// after PutDestination.
func (c *Client) PutDestination(ctx context.Context, params *PutDestinationInput, optFns ...func(*Options)) (*PutDestinationOutput, error) {
	stack := middleware.NewStack("PutDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutDestinationMiddlewares(stack)
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
	addOpPutDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutDestination(options.Region), middleware.Before)
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
			OperationName: "PutDestination",
			Err:           err,
		}
	}
	out := result.(*PutDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutDestinationInput struct {

	// The ARN of an Amazon Kinesis stream to which to deliver matching log events.
	//
	// This member is required.
	TargetArn *string

	// A name for the destination.
	//
	// This member is required.
	DestinationName *string

	// The ARN of an IAM role that grants CloudWatch Logs permissions to call the
	// Amazon Kinesis PutRecord operation on the destination stream.
	//
	// This member is required.
	RoleArn *string
}

type PutDestinationOutput struct {

	// The destination.
	Destination *types.Destination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutDestination{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutDestination",
	}
}
