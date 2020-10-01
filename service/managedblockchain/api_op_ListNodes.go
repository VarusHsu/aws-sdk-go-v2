// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the nodes within a network.
func (c *Client) ListNodes(ctx context.Context, params *ListNodesInput, optFns ...func(*Options)) (*ListNodesOutput, error) {
	stack := middleware.NewStack("ListNodes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListNodesMiddlewares(stack)
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
	addOpListNodesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListNodes(options.Region), middleware.Before)
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
			OperationName: "ListNodes",
			Err:           err,
		}
	}
	out := result.(*ListNodesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNodesInput struct {

	// The unique identifier of the member who owns the nodes to list.
	//
	// This member is required.
	MemberId *string

	// An optional status specifier. If provided, only nodes currently in this status
	// are listed.
	Status types.NodeStatus

	// The unique identifier of the network for which to list nodes.
	//
	// This member is required.
	NetworkId *string

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// The maximum number of nodes to list.
	MaxResults *int32
}

type ListNodesOutput struct {

	// An array of NodeSummary objects that contain configuration properties for each
	// node.
	Nodes []*types.NodeSummary

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListNodesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListNodes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListNodes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListNodes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "ListNodes",
	}
}
