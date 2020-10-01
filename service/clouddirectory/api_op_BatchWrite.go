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

// Performs all the write operations in a batch. Either all the operations succeed
// or none.
func (c *Client) BatchWrite(ctx context.Context, params *BatchWriteInput, optFns ...func(*Options)) (*BatchWriteOutput, error) {
	stack := middleware.NewStack("BatchWrite", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchWriteMiddlewares(stack)
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
	addOpBatchWriteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchWrite(options.Region), middleware.Before)
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
			OperationName: "BatchWrite",
			Err:           err,
		}
	}
	out := result.(*BatchWriteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchWriteInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory (). For
	// more information, see arns ().
	//
	// This member is required.
	DirectoryArn *string

	// A list of operations that are part of the batch.
	//
	// This member is required.
	Operations []*types.BatchWriteOperation
}

type BatchWriteOutput struct {

	// A list of all the responses for each batch write.
	Responses []*types.BatchWriteOperationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchWriteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchWrite{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchWrite{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchWrite(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "BatchWrite",
	}
}
