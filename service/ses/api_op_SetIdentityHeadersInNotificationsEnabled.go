// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Given an identity (an email address or a domain), sets whether Amazon SES
// includes the original email headers in the Amazon Simple Notification Service
// (Amazon SNS) notifications of a specified type. You can execute this operation
// no more than once per second. For more information about using notifications
// with Amazon SES, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
func (c *Client) SetIdentityHeadersInNotificationsEnabled(ctx context.Context, params *SetIdentityHeadersInNotificationsEnabledInput, optFns ...func(*Options)) (*SetIdentityHeadersInNotificationsEnabledOutput, error) {
	stack := middleware.NewStack("SetIdentityHeadersInNotificationsEnabled", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetIdentityHeadersInNotificationsEnabledMiddlewares(stack)
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
	addOpSetIdentityHeadersInNotificationsEnabledValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityHeadersInNotificationsEnabled(options.Region), middleware.Before)
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
			OperationName: "SetIdentityHeadersInNotificationsEnabled",
			Err:           err,
		}
	}
	out := result.(*SetIdentityHeadersInNotificationsEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to set whether Amazon SES includes the original email
// headers in the Amazon SNS notifications of a specified type. For information
// about notifications, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications-via-sns.html).
type SetIdentityHeadersInNotificationsEnabledInput struct {

	// Sets whether Amazon SES includes the original email headers in Amazon SNS
	// notifications of the specified notification type. A value of true specifies that
	// Amazon SES will include headers in notifications, and a value of false specifies
	// that Amazon SES will not include headers in notifications. This value can only
	// be set when NotificationType is already set to use a particular Amazon SNS
	// topic.
	//
	// This member is required.
	Enabled *bool

	// The notification type for which to enable or disable headers in notifications.
	//
	// This member is required.
	NotificationType types.NotificationType

	// The identity for which to enable or disable headers in notifications. Examples:
	// user@example.com, example.com.
	//
	// This member is required.
	Identity *string
}

// An empty element returned on a successful request.
type SetIdentityHeadersInNotificationsEnabledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetIdentityHeadersInNotificationsEnabledMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetIdentityHeadersInNotificationsEnabled{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetIdentityHeadersInNotificationsEnabled{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetIdentityHeadersInNotificationsEnabled(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SetIdentityHeadersInNotificationsEnabled",
	}
}
