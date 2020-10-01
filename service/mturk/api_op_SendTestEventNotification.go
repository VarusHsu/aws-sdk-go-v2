// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The SendTestEventNotification operation causes Amazon Mechanical Turk to send a
// notification message as if a HIT event occurred, according to the provided
// notification specification. This allows you to test notifications without
// setting up notifications for a real HIT type and trying to trigger them using
// the website. When you call this operation, the service attempts to send the test
// notification immediately.
func (c *Client) SendTestEventNotification(ctx context.Context, params *SendTestEventNotificationInput, optFns ...func(*Options)) (*SendTestEventNotificationOutput, error) {
	stack := middleware.NewStack("SendTestEventNotification", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSendTestEventNotificationMiddlewares(stack)
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
	addOpSendTestEventNotificationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendTestEventNotification(options.Region), middleware.Before)
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
			OperationName: "SendTestEventNotification",
			Err:           err,
		}
	}
	out := result.(*SendTestEventNotificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendTestEventNotificationInput struct {

	// The notification specification to test. This value is identical to the value you
	// would provide to the UpdateNotificationSettings operation when you establish the
	// notification specification for a HIT type.
	//
	// This member is required.
	Notification *types.NotificationSpecification

	// The event to simulate to test the notification specification. This event is
	// included in the test message even if the notification specification does not
	// include the event type. The notification specification does not filter out the
	// test event.
	//
	// This member is required.
	TestEventType types.EventType
}

type SendTestEventNotificationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSendTestEventNotificationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSendTestEventNotification{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendTestEventNotification{}, middleware.After)
}

func newServiceMetadataMiddleware_opSendTestEventNotification(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "SendTestEventNotification",
	}
}
