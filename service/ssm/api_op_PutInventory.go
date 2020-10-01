// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Bulk update custom inventory items on one more instance. The request adds an
// inventory item, if it doesn't already exist, or updates an inventory item, if it
// does exist.
func (c *Client) PutInventory(ctx context.Context, params *PutInventoryInput, optFns ...func(*Options)) (*PutInventoryOutput, error) {
	stack := middleware.NewStack("PutInventory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutInventoryMiddlewares(stack)
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
	addOpPutInventoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutInventory(options.Region), middleware.Before)
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
			OperationName: "PutInventory",
			Err:           err,
		}
	}
	out := result.(*PutInventoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutInventoryInput struct {

	// The inventory items that you want to add or update on instances.
	//
	// This member is required.
	Items []*types.InventoryItem

	// An instance ID where you want to add or update inventory items.
	//
	// This member is required.
	InstanceId *string
}

type PutInventoryOutput struct {

	// Information about the request.
	Message *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutInventoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutInventory{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutInventory{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutInventory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "PutInventory",
	}
}
