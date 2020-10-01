// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the bots associated with the administrator's Amazon Chime Enterprise
// account ID.
func (c *Client) ListBots(ctx context.Context, params *ListBotsInput, optFns ...func(*Options)) (*ListBotsOutput, error) {
	stack := middleware.NewStack("ListBots", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListBotsMiddlewares(stack)
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
	addOpListBotsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBots(options.Region), middleware.Before)
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
			OperationName: "ListBots",
			Err:           err,
		}
	}
	out := result.(*ListBotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBotsInput struct {

	// The maximum number of results to return in a single call. The default is 10.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string
}

type ListBotsOutput struct {

	// List of bots and bot details.
	Bots []*types.Bot

	// The token to use to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListBotsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListBots{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListBots{}, middleware.After)
}

func newServiceMetadataMiddleware_opListBots(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListBots",
	}
}
