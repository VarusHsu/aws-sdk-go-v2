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

// Describes one or more of your subnets. For more information, see Your VPC and
// Subnets (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in
// the Amazon Virtual Private Cloud User Guide.
func (c *Client) DescribeSubnets(ctx context.Context, params *DescribeSubnetsInput, optFns ...func(*Options)) (*DescribeSubnetsOutput, error) {
	stack := middleware.NewStack("DescribeSubnets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeSubnetsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSubnets(options.Region), middleware.Before)
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
			OperationName: "DescribeSubnets",
			Err:           err,
		}
	}
	out := result.(*DescribeSubnetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSubnetsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more filters.
	//
	//     * availability-zone - The Availability Zone for the
	// subnet. You can also use availabilityZone as the filter name.
	//
	//     *
	// availability-zone-id - The ID of the Availability Zone for the subnet. You can
	// also use availabilityZoneId as the filter name.
	//
	//     *
	// available-ip-address-count - The number of IPv4 addresses in the subnet that are
	// available.
	//
	//     * cidr-block - The IPv4 CIDR block of the subnet. The CIDR block
	// you specify must exactly match the subnet's CIDR block for information to be
	// returned for the subnet. You can also use cidr or cidrBlock as the filter
	// names.
	//
	//     * default-for-az - Indicates whether this is the default subnet for
	// the Availability Zone. You can also use defaultForAz as the filter name.
	//
	//     *
	// ipv6-cidr-block-association.ipv6-cidr-block - An IPv6 CIDR block associated with
	// the subnet.
	//
	//     * ipv6-cidr-block-association.association-id - An association
	// ID for an IPv6 CIDR block associated with the subnet.
	//
	//     *
	// ipv6-cidr-block-association.state - The state of an IPv6 CIDR block associated
	// with the subnet.
	//
	//     * owner-id - The ID of the AWS account that owns the
	// subnet.
	//
	//     * state - The state of the subnet (pending | available).
	//
	//     *
	// subnet-arn - The Amazon Resource Name (ARN) of the subnet.
	//
	//     * subnet-id -
	// The ID of the subnet.
	//
	//     * tag: - The key/value combination of a tag assigned
	// to the resource. Use the tag key in the filter name and the tag value as the
	// filter value. For example, to find all resources that have a tag with the key
	// Owner and the value TeamA, specify tag:Owner for the filter name and TeamA for
	// the filter value.
	//
	//     * tag-key - The key of a tag assigned to the resource.
	// Use this filter to find all resources assigned a tag with a specific key,
	// regardless of the tag value.
	//
	//     * vpc-id - The ID of the VPC for the subnet.
	Filters []*types.Filter

	// One or more subnet IDs. Default: Describes all your subnets.
	SubnetIds []*string
}

type DescribeSubnetsOutput struct {

	// Information about one or more subnets.
	Subnets []*types.Subnet

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeSubnetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeSubnets{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSubnets{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSubnets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSubnets",
	}
}
