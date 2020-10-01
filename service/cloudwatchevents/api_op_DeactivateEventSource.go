// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// You can use this operation to temporarily stop receiving events from the
// specified partner event source. The matching event bus is not deleted. When you
// deactivate a partner event source, the source goes into PENDING state. If it
// remains in PENDING state for more than two weeks, it is deleted. To activate a
// deactivated partner event source, use ActivateEventSource ().
func (c *Client) DeactivateEventSource(ctx context.Context, params *DeactivateEventSourceInput, optFns ...func(*Options)) (*DeactivateEventSourceOutput, error) {
	stack := middleware.NewStack("DeactivateEventSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeactivateEventSourceMiddlewares(stack)
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
	addOpDeactivateEventSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeactivateEventSource(options.Region), middleware.Before)
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
			OperationName: "DeactivateEventSource",
			Err:           err,
		}
	}
	out := result.(*DeactivateEventSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeactivateEventSourceInput struct {

	// The name of the partner event source to deactivate.
	//
	// This member is required.
	Name *string
}

type DeactivateEventSourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeactivateEventSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeactivateEventSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeactivateEventSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeactivateEventSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DeactivateEventSource",
	}
}
