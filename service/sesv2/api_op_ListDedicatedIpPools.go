// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all of the dedicated IP pools that exist in your AWS account in the current
// Region.
func (c *Client) ListDedicatedIpPools(ctx context.Context, params *ListDedicatedIpPoolsInput, optFns ...func(*Options)) (*ListDedicatedIpPoolsOutput, error) {
	stack := middleware.NewStack("ListDedicatedIpPools", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListDedicatedIpPoolsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDedicatedIpPools(options.Region), middleware.Before)
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
			OperationName: "ListDedicatedIpPools",
			Err:           err,
		}
	}
	out := result.(*ListDedicatedIpPoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain a list of dedicated IP pools.
type ListDedicatedIpPoolsInput struct {

	// A token returned from a previous call to ListDedicatedIpPools to indicate the
	// position in the list of dedicated IP pools.
	NextToken *string

	// The number of results to show in a single call to ListDedicatedIpPools. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int32
}

// A list of dedicated IP pools.
type ListDedicatedIpPoolsOutput struct {

	// A token that indicates that there are additional IP pools to list. To view
	// additional IP pools, issue another request to ListDedicatedIpPools, passing this
	// token in the NextToken parameter.
	NextToken *string

	// A list of all of the dedicated IP pools that are associated with your AWS
	// account in the current Region.
	DedicatedIpPools []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListDedicatedIpPoolsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListDedicatedIpPools{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListDedicatedIpPools{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDedicatedIpPools(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListDedicatedIpPools",
	}
}
