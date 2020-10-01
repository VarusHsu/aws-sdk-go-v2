// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deprecated. Use the DeleteIdentity operation to delete email addresses and
// domains.
func (c *Client) DeleteVerifiedEmailAddress(ctx context.Context, params *DeleteVerifiedEmailAddressInput, optFns ...func(*Options)) (*DeleteVerifiedEmailAddressOutput, error) {
	stack := middleware.NewStack("DeleteVerifiedEmailAddress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteVerifiedEmailAddressMiddlewares(stack)
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
	addOpDeleteVerifiedEmailAddressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVerifiedEmailAddress(options.Region), middleware.Before)
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
			OperationName: "DeleteVerifiedEmailAddress",
			Err:           err,
		}
	}
	out := result.(*DeleteVerifiedEmailAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete an email address from the list of email addresses
// you have attempted to verify under your AWS account.
type DeleteVerifiedEmailAddressInput struct {

	// An email address to be removed from the list of verified addresses.
	//
	// This member is required.
	EmailAddress *string
}

type DeleteVerifiedEmailAddressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteVerifiedEmailAddressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteVerifiedEmailAddress{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteVerifiedEmailAddress{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVerifiedEmailAddress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteVerifiedEmailAddress",
	}
}
