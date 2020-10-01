// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your NAT gateways.
func (c *Client) DescribeNatGateways(ctx context.Context, params *DescribeNatGatewaysInput, optFns ...func(*Options)) (*DescribeNatGatewaysOutput, error) {
	stack := middleware.NewStack("DescribeNatGateways", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeNatGatewaysMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNatGateways(options.Region), middleware.Before)
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
			OperationName: "DescribeNatGateways",
			Err:           err,
		}
	}
	out := result.(*DescribeNatGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNatGatewaysInput struct {

	// The token for the next page of results.
	NextToken *string

	// One or more filters.
	//
	//     * nat-gateway-id - The ID of the NAT gateway.
	//
	//     *
	// state - The state of the NAT gateway (pending | failed | available | deleting |
	// deleted).
	//
	//     * subnet-id - The ID of the subnet in which the NAT gateway
	// resides.
	//
	//     * tag: - The key/value combination of a tag assigned to the
	// resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	//     * tag-key - The key of a tag assigned to the resource. Use this
	// filter to find all resources assigned a tag with a specific key, regardless of
	// the tag value.
	//
	//     * vpc-id - The ID of the VPC in which the NAT gateway
	// resides.
	Filter []*types.Filter

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// One or more NAT gateway IDs.
	NatGatewayIds []*string
}

type DescribeNatGatewaysOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the NAT gateways.
	NatGateways []*types.NatGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeNatGatewaysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeNatGateways{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeNatGateways{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNatGateways(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeNatGateways",
	}
}
