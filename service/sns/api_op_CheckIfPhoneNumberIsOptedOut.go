// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accepts a phone number and indicates whether the phone holder has opted out of
// receiving SMS messages from your account. You cannot send SMS messages to a
// number that is opted out. To resume sending messages, you can opt in the number
// by using the OptInPhoneNumber action.
func (c *Client) CheckIfPhoneNumberIsOptedOut(ctx context.Context, params *CheckIfPhoneNumberIsOptedOutInput, optFns ...func(*Options)) (*CheckIfPhoneNumberIsOptedOutOutput, error) {
	stack := middleware.NewStack("CheckIfPhoneNumberIsOptedOut", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCheckIfPhoneNumberIsOptedOutMiddlewares(stack)
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
	addOpCheckIfPhoneNumberIsOptedOutValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCheckIfPhoneNumberIsOptedOut(options.Region), middleware.Before)
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
			OperationName: "CheckIfPhoneNumberIsOptedOut",
			Err:           err,
		}
	}
	out := result.(*CheckIfPhoneNumberIsOptedOutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CheckIfPhoneNumberIsOptedOut action.
type CheckIfPhoneNumberIsOptedOutInput struct {

	// The phone number for which you want to check the opt out status.
	//
	// This member is required.
	PhoneNumber *string
}

// The response from the CheckIfPhoneNumberIsOptedOut action.
type CheckIfPhoneNumberIsOptedOutOutput struct {

	// Indicates whether the phone number is opted out:
	//
	//     * true – The phone number
	// is opted out, meaning you cannot publish SMS messages to it.
	//
	//     * false – The
	// phone number is opted in, meaning you can publish SMS messages to it.
	IsOptedOut *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCheckIfPhoneNumberIsOptedOutMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCheckIfPhoneNumberIsOptedOut{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCheckIfPhoneNumberIsOptedOut{}, middleware.After)
}

func newServiceMetadataMiddleware_opCheckIfPhoneNumberIsOptedOut(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "CheckIfPhoneNumberIsOptedOut",
	}
}
