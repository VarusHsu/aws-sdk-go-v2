// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds Sources to flow
func (c *Client) AddFlowSources(ctx context.Context, params *AddFlowSourcesInput, optFns ...func(*Options)) (*AddFlowSourcesOutput, error) {
	stack := middleware.NewStack("AddFlowSources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAddFlowSourcesMiddlewares(stack)
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
	addOpAddFlowSourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddFlowSources(options.Region), middleware.Before)
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
			OperationName: "AddFlowSources",
			Err:           err,
		}
	}
	out := result.(*AddFlowSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to add sources to the flow.
type AddFlowSourcesInput struct {

	// A list of sources that you want to add.
	//
	// This member is required.
	Sources []*types.SetSourceRequest

	// The flow that you want to mutate.
	//
	// This member is required.
	FlowArn *string
}

type AddFlowSourcesOutput struct {

	// The ARN of the flow that these sources were added to.
	FlowArn *string

	// The details of the newly added sources.
	Sources []*types.Source

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAddFlowSourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAddFlowSources{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAddFlowSources{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddFlowSources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "AddFlowSources",
	}
}
