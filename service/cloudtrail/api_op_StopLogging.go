// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Suspends the recording of AWS API calls and log file delivery for the specified
// trail. Under most circumstances, there is no need to use this action. You can
// update a trail without stopping it first. This action is the only way to stop
// recording. For a trail enabled in all regions, this operation must be called
// from the region in which the trail was created, or an InvalidHomeRegionException
// will occur. This operation cannot be called on the shadow trails (replicated
// trails in other regions) of a trail enabled in all regions.
func (c *Client) StopLogging(ctx context.Context, params *StopLoggingInput, optFns ...func(*Options)) (*StopLoggingOutput, error) {
	stack := middleware.NewStack("StopLogging", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopLoggingMiddlewares(stack)
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
	addOpStopLoggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopLogging(options.Region), middleware.Before)
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
			OperationName: "StopLogging",
			Err:           err,
		}
	}
	out := result.(*StopLoggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Passes the request to CloudTrail to stop logging AWS API calls for the specified
// account.
type StopLoggingInput struct {

	// Specifies the name or the CloudTrail ARN of the trail for which CloudTrail will
	// stop logging AWS API calls. The format of a trail ARN is:
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// This member is required.
	Name *string
}

// Returns the objects or data listed below if successful. Otherwise, returns an
// error.
type StopLoggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopLoggingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopLogging{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopLogging{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopLogging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "StopLogging",
	}
}
