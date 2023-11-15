// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Given an identity (an email address or a domain), enables or disables whether
// Amazon SES forwards bounce and complaint notifications as email. Feedback
// forwarding can only be disabled when Amazon Simple Notification Service (Amazon
// SNS) topics are specified for both bounces and complaints. Feedback forwarding
// does not apply to delivery notifications. Delivery notifications are only
// available through Amazon SNS. You can execute this operation no more than once
// per second. For more information about using notifications with Amazon SES, see
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications.html)
// .
func (c *Client) SetIdentityFeedbackForwardingEnabled(ctx context.Context, params *SetIdentityFeedbackForwardingEnabledInput, optFns ...func(*Options)) (*SetIdentityFeedbackForwardingEnabledOutput, error) {
	if params == nil {
		params = &SetIdentityFeedbackForwardingEnabledInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetIdentityFeedbackForwardingEnabled", params, optFns, c.addOperationSetIdentityFeedbackForwardingEnabledMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetIdentityFeedbackForwardingEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to enable or disable whether Amazon SES forwards you
// bounce and complaint notifications through email. For information about email
// feedback forwarding, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications-email.html)
// .
type SetIdentityFeedbackForwardingEnabledInput struct {

	// Sets whether Amazon SES forwards bounce and complaint notifications as email.
	// true specifies that Amazon SES forwards bounce and complaint notifications as
	// email, in addition to any Amazon SNS topic publishing otherwise specified. false
	// specifies that Amazon SES publishes bounce and complaint notifications only
	// through Amazon SNS. This value can only be set to false when Amazon SNS topics
	// are set for both Bounce and Complaint notification types.
	//
	// This member is required.
	ForwardingEnabled bool

	// The identity for which to set bounce and complaint notification forwarding.
	// Examples: user@example.com , example.com .
	//
	// This member is required.
	Identity *string

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type SetIdentityFeedbackForwardingEnabledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetIdentityFeedbackForwardingEnabledMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetIdentityFeedbackForwardingEnabled{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetIdentityFeedbackForwardingEnabled{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetIdentityFeedbackForwardingEnabled"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSetIdentityFeedbackForwardingEnabledValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityFeedbackForwardingEnabled(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSetIdentityFeedbackForwardingEnabled(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetIdentityFeedbackForwardingEnabled",
	}
}
