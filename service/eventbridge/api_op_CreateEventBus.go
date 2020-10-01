// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new event bus within your account. This can be a custom event bus
// which you can use to receive events from your custom applications and services,
// or it can be a partner event bus which can be matched to a partner event source.
func (c *Client) CreateEventBus(ctx context.Context, params *CreateEventBusInput, optFns ...func(*Options)) (*CreateEventBusOutput, error) {
	stack := middleware.NewStack("CreateEventBus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateEventBusMiddlewares(stack)
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
	addOpCreateEventBusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventBus(options.Region), middleware.Before)
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
			OperationName: "CreateEventBus",
			Err:           err,
		}
	}
	out := result.(*CreateEventBusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventBusInput struct {

	// If you are creating a partner event bus, this specifies the partner event source
	// that the new event bus will be matched with.
	EventSourceName *string

	// The name of the new event bus. Event bus names cannot contain the / character.
	// You can't use the name default for a custom event bus, as this name is already
	// used for your account's default event bus. If this is a partner event bus, the
	// name must exactly match the name of the partner event source that this event bus
	// is matched to.
	//
	// This member is required.
	Name *string

	// Tags to associate with the event bus.
	Tags []*types.Tag
}

type CreateEventBusOutput struct {

	// The ARN of the new event bus.
	EventBusArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateEventBusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEventBus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEventBus{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateEventBus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "CreateEventBus",
	}
}
