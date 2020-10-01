// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the connections associated with your account.
func (c *Client) ListConnections(ctx context.Context, params *ListConnectionsInput, optFns ...func(*Options)) (*ListConnectionsOutput, error) {
	stack := middleware.NewStack("ListConnections", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpListConnectionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListConnections(options.Region), middleware.Before)
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
			OperationName: "ListConnections",
			Err:           err,
		}
	}
	out := result.(*ListConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectionsInput struct {

	// Filters the list of connections to those associated with a specified host.
	HostArnFilter *string

	// Filters the list of connections to those associated with a specified provider,
	// such as Bitbucket.
	ProviderTypeFilter types.ProviderType

	// The token that was returned from the previous ListConnections call, which can be
	// used to return the next set of connections in the list.
	NextToken *string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
}

type ListConnectionsOutput struct {

	// A token that can be used in the next ListConnections call. To view all items in
	// the list, continue to call this operation with each subsequent token until no
	// more nextToken values are returned.
	NextToken *string

	// A list of connections and the details for each connection, such as status,
	// owner, and provider type.
	Connections []*types.Connection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpListConnectionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpListConnections{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpListConnections{}, middleware.After)
}

func newServiceMetadataMiddleware_opListConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-connections",
		OperationName: "ListConnections",
	}
}
