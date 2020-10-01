// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified policy from the specified certificate. Note: This API is
// deprecated. Please use DetachPolicy () instead.
func (c *Client) DetachPrincipalPolicy(ctx context.Context, params *DetachPrincipalPolicyInput, optFns ...func(*Options)) (*DetachPrincipalPolicyOutput, error) {
	stack := middleware.NewStack("DetachPrincipalPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDetachPrincipalPolicyMiddlewares(stack)
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
	addOpDetachPrincipalPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachPrincipalPolicy(options.Region), middleware.Before)
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
			OperationName: "DetachPrincipalPolicy",
			Err:           err,
		}
	}
	out := result.(*DetachPrincipalPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DetachPrincipalPolicy operation.
type DetachPrincipalPolicyInput struct {

	// The principal. If the principal is a certificate, specify the certificate ARN.
	// If the principal is an Amazon Cognito identity, specify the identity ID.
	//
	// This member is required.
	Principal *string

	// The name of the policy to detach.
	//
	// This member is required.
	PolicyName *string
}

type DetachPrincipalPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDetachPrincipalPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDetachPrincipalPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDetachPrincipalPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachPrincipalPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DetachPrincipalPolicy",
	}
}
