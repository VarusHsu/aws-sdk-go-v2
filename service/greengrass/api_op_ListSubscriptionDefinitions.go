// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of subscription definitions.
func (c *Client) ListSubscriptionDefinitions(ctx context.Context, params *ListSubscriptionDefinitionsInput, optFns ...func(*Options)) (*ListSubscriptionDefinitionsOutput, error) {
	stack := middleware.NewStack("ListSubscriptionDefinitions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListSubscriptionDefinitionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscriptionDefinitions(options.Region), middleware.Before)
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
			OperationName: "ListSubscriptionDefinitions",
			Err:           err,
		}
	}
	out := result.(*ListSubscriptionDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSubscriptionDefinitionsInput struct {

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// The maximum number of results to be returned per request.
	MaxResults *string
}

type ListSubscriptionDefinitionsOutput struct {

	// Information about a definition.
	Definitions []*types.DefinitionInformation

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListSubscriptionDefinitionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListSubscriptionDefinitions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListSubscriptionDefinitions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSubscriptionDefinitions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListSubscriptionDefinitions",
	}
}
