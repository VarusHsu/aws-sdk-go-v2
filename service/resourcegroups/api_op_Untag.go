// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes tags from a specified resource group.
func (c *Client) Untag(ctx context.Context, params *UntagInput, optFns ...func(*Options)) (*UntagOutput, error) {
	stack := middleware.NewStack("Untag", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUntagMiddlewares(stack)
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
	addOpUntagValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUntag(options.Region), middleware.Before)
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
			OperationName: "Untag",
			Err:           err,
		}
	}
	out := result.(*UntagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagInput struct {

	// The keys of the tags to be removed.
	//
	// This member is required.
	Keys []*string

	// The ARN of the resource group from which to remove tags. The command removed
	// both the specified keys and any values associated with those keys.
	//
	// This member is required.
	Arn *string
}

type UntagOutput struct {

	// The keys of the tags that were removed.
	Keys []*string

	// The ARN of the resource group from which tags have been removed.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUntagMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUntag{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUntag{}, middleware.After)
}

func newServiceMetadataMiddleware_opUntag(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "Untag",
	}
}
