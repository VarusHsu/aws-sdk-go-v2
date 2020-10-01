// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Unassigns one or more IPv6 addresses from a network interface.
func (c *Client) UnassignIpv6Addresses(ctx context.Context, params *UnassignIpv6AddressesInput, optFns ...func(*Options)) (*UnassignIpv6AddressesOutput, error) {
	stack := middleware.NewStack("UnassignIpv6Addresses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpUnassignIpv6AddressesMiddlewares(stack)
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
	addOpUnassignIpv6AddressesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUnassignIpv6Addresses(options.Region), middleware.Before)
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
			OperationName: "UnassignIpv6Addresses",
			Err:           err,
		}
	}
	out := result.(*UnassignIpv6AddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnassignIpv6AddressesInput struct {

	// The ID of the network interface.
	//
	// This member is required.
	NetworkInterfaceId *string

	// The IPv6 addresses to unassign from the network interface.
	//
	// This member is required.
	Ipv6Addresses []*string
}

type UnassignIpv6AddressesOutput struct {

	// The ID of the network interface.
	NetworkInterfaceId *string

	// The IPv6 addresses that have been unassigned from the network interface.
	UnassignedIpv6Addresses []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpUnassignIpv6AddressesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpUnassignIpv6Addresses{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpUnassignIpv6Addresses{}, middleware.After)
}

func newServiceMetadataMiddleware_opUnassignIpv6Addresses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "UnassignIpv6Addresses",
	}
}
