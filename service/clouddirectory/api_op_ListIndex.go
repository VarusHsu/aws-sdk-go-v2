// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists objects attached to the specified index.
func (c *Client) ListIndex(ctx context.Context, params *ListIndexInput, optFns ...func(*Options)) (*ListIndexOutput, error) {
	stack := middleware.NewStack("ListIndex", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListIndexMiddlewares(stack)
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
	addOpListIndexValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIndex(options.Region), middleware.Before)
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
			OperationName: "ListIndex",
			Err:           err,
		}
	}
	out := result.(*ListIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIndexInput struct {

	// The pagination token.
	NextToken *string

	// Specifies the ranges of indexed values that you want to query.
	RangesOnIndexedValues []*types.ObjectAttributeRange

	// The consistency level to execute the request at.
	ConsistencyLevel types.ConsistencyLevel

	// The maximum number of objects in a single page to retrieve from the index during
	// a request. For more information, see Amazon Cloud Directory Limits
	// (http://docs.aws.amazon.com/clouddirectory/latest/developerguide/limits.html).
	MaxResults *int32

	// The ARN of the directory that the index exists in.
	//
	// This member is required.
	DirectoryArn *string

	// The reference to the index to list.
	//
	// This member is required.
	IndexReference *types.ObjectReference
}

type ListIndexOutput struct {

	// The pagination token.
	NextToken *string

	// The objects and indexed values attached to the index.
	IndexAttachments []*types.IndexAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListIndexMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListIndex{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListIndex{}, middleware.After)
}

func newServiceMetadataMiddleware_opListIndex(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListIndex",
	}
}
