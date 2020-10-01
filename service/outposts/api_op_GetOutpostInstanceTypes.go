// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the instance types for the specified Outpost.
func (c *Client) GetOutpostInstanceTypes(ctx context.Context, params *GetOutpostInstanceTypesInput, optFns ...func(*Options)) (*GetOutpostInstanceTypesOutput, error) {
	stack := middleware.NewStack("GetOutpostInstanceTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetOutpostInstanceTypesMiddlewares(stack)
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
	addOpGetOutpostInstanceTypesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOutpostInstanceTypes(options.Region), middleware.Before)
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
			OperationName: "GetOutpostInstanceTypes",
			Err:           err,
		}
	}
	out := result.(*GetOutpostInstanceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOutpostInstanceTypesInput struct {

	// The pagination token.
	NextToken *string

	// The ID of the Outpost.
	//
	// This member is required.
	OutpostId *string

	// The maximum page size.
	MaxResults *int32
}

type GetOutpostInstanceTypesOutput struct {

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string

	// The ID of the Outpost.
	OutpostId *string

	// Information about the instance types.
	InstanceTypes []*types.InstanceTypeItem

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetOutpostInstanceTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetOutpostInstanceTypes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetOutpostInstanceTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetOutpostInstanceTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "outposts",
		OperationName: "GetOutpostInstanceTypes",
	}
}
