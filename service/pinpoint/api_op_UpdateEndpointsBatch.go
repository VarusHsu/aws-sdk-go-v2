// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new batch of endpoints for an application or updates the settings and
// attributes of a batch of existing endpoints for an application. You can also use
// this operation to define custom attributes for a batch of endpoints. If an
// update includes one or more values for a custom attribute, Amazon Pinpoint
// replaces (overwrites) any existing values with the new values.
func (c *Client) UpdateEndpointsBatch(ctx context.Context, params *UpdateEndpointsBatchInput, optFns ...func(*Options)) (*UpdateEndpointsBatchOutput, error) {
	stack := middleware.NewStack("UpdateEndpointsBatch", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateEndpointsBatchMiddlewares(stack)
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
	addOpUpdateEndpointsBatchValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpointsBatch(options.Region), middleware.Before)
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
			OperationName: "UpdateEndpointsBatch",
			Err:           err,
		}
	}
	out := result.(*UpdateEndpointsBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointsBatchInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// Specifies a batch of endpoints to create or update and the settings and
	// attributes to set or change for each endpoint.
	//
	// This member is required.
	EndpointBatchRequest *types.EndpointBatchRequest
}

type UpdateEndpointsBatchOutput struct {

	// Provides information about an API request or response.
	//
	// This member is required.
	MessageBody *types.MessageBody

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateEndpointsBatchMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEndpointsBatch{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEndpointsBatch{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateEndpointsBatch(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateEndpointsBatch",
	}
}
