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

// Adds or updates a sending authorization policy for the specified identity (an
// email address or a domain). This API is for the identity owner only. If you have
// not verified the identity, this API will return an error. Sending authorization
// is a feature that enables an identity owner to authorize other senders to use
// its identities. For information about using sending authorization, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
// You can execute this operation no more than once per second.
func (c *Client) PutIdentityPolicy(ctx context.Context, params *PutIdentityPolicyInput, optFns ...func(*Options)) (*PutIdentityPolicyOutput, error) {
	stack := middleware.NewStack("PutIdentityPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPutIdentityPolicyMiddlewares(stack)
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
	addOpPutIdentityPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutIdentityPolicy(options.Region), middleware.Before)
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
			OperationName: "PutIdentityPolicy",
			Err:           err,
		}
	}
	out := result.(*PutIdentityPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to add or update a sending authorization policy for an
// identity. Sending authorization is an Amazon SES feature that enables you to
// authorize other senders to use your identities. For information, see the Amazon
// SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
type PutIdentityPolicyInput struct {

	// The text of the policy in JSON format. The policy cannot exceed 4 KB. For
	// information about the syntax of sending authorization policies, see the Amazon
	// SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).
	//
	// This member is required.
	Policy *string

	// The name of the policy. The policy name cannot exceed 64 characters and can only
	// include alphanumeric characters, dashes, and underscores.
	//
	// This member is required.
	PolicyName *string

	// The identity that the policy will apply to. You can specify an identity by using
	// its name or by using its Amazon Resource Name (ARN). Examples: user@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com. To
	// successfully call this API, you must own the identity.
	//
	// This member is required.
	Identity *string
}

// An empty element returned on a successful request.
type PutIdentityPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPutIdentityPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPutIdentityPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPutIdentityPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutIdentityPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutIdentityPolicy",
	}
}
