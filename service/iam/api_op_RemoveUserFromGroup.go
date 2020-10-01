// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified user from the specified group.
func (c *Client) RemoveUserFromGroup(ctx context.Context, params *RemoveUserFromGroupInput, optFns ...func(*Options)) (*RemoveUserFromGroupOutput, error) {
	stack := middleware.NewStack("RemoveUserFromGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRemoveUserFromGroupMiddlewares(stack)
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
	addOpRemoveUserFromGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveUserFromGroup(options.Region), middleware.Before)
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
			OperationName: "RemoveUserFromGroup",
			Err:           err,
		}
	}
	out := result.(*RemoveUserFromGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveUserFromGroupInput struct {

	// The name of the user to remove. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string

	// The name of the group to update. This parameter allows (through its regex
	// pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting of
	// upper and lowercase alphanumeric characters with no spaces. You can also include
	// any of the following characters: _+=,.@-
	//
	// This member is required.
	GroupName *string
}

type RemoveUserFromGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRemoveUserFromGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRemoveUserFromGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveUserFromGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveUserFromGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "RemoveUserFromGroup",
	}
}
