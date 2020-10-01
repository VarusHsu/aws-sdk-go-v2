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

// Given a list of verified identities (email addresses and/or domains), returns a
// structure describing identity notification attributes. This operation is
// throttled at one request per second and can only get notification attributes for
// up to 100 identities at a time. For more information about using notifications
// with Amazon SES, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
func (c *Client) GetIdentityNotificationAttributes(ctx context.Context, params *GetIdentityNotificationAttributesInput, optFns ...func(*Options)) (*GetIdentityNotificationAttributesOutput, error) {
	stack := middleware.NewStack("GetIdentityNotificationAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetIdentityNotificationAttributesMiddlewares(stack)
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
	addOpGetIdentityNotificationAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityNotificationAttributes(options.Region), middleware.Before)
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
			OperationName: "GetIdentityNotificationAttributes",
			Err:           err,
		}
	}
	out := result.(*GetIdentityNotificationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the notification attributes for a list of
// identities you verified with Amazon SES. For information about Amazon SES
// notifications, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
type GetIdentityNotificationAttributesInput struct {

	// A list of one or more identities. You can specify an identity by using its name
	// or by using its Amazon Resource Name (ARN). Examples: user@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// This member is required.
	Identities []*string
}

// Represents the notification attributes for a list of identities.
type GetIdentityNotificationAttributesOutput struct {

	// A map of Identity to IdentityNotificationAttributes.
	//
	// This member is required.
	NotificationAttributes map[string]*types.IdentityNotificationAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetIdentityNotificationAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetIdentityNotificationAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetIdentityNotificationAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetIdentityNotificationAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetIdentityNotificationAttributes",
	}
}
