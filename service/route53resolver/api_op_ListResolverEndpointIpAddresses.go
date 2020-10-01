// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the IP addresses for a specified resolver endpoint.
func (c *Client) ListResolverEndpointIpAddresses(ctx context.Context, params *ListResolverEndpointIpAddressesInput, optFns ...func(*Options)) (*ListResolverEndpointIpAddressesOutput, error) {
	stack := middleware.NewStack("ListResolverEndpointIpAddresses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListResolverEndpointIpAddressesMiddlewares(stack)
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
	addOpListResolverEndpointIpAddressesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverEndpointIpAddresses(options.Region), middleware.Before)
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
			OperationName: "ListResolverEndpointIpAddresses",
			Err:           err,
		}
	}
	out := result.(*ListResolverEndpointIpAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverEndpointIpAddressesInput struct {

	// For the first ListResolverEndpointIpAddresses request, omit this value. If the
	// specified resolver endpoint has more than MaxResults IP addresses, you can
	// submit another ListResolverEndpointIpAddresses request to get the next group of
	// IP addresses. In the next request, specify the value of NextToken from the
	// previous response.
	NextToken *string

	// The maximum number of IP addresses that you want to return in the response to a
	// ListResolverEndpointIpAddresses request. If you don't specify a value for
	// MaxResults, Resolver returns up to 100 IP addresses.
	MaxResults *int32

	// The ID of the resolver endpoint that you want to get IP addresses for.
	//
	// This member is required.
	ResolverEndpointId *string
}

type ListResolverEndpointIpAddressesOutput struct {

	// The value that you specified for MaxResults in the request.
	MaxResults *int32

	// If the specified endpoint has more than MaxResults IP addresses, you can submit
	// another ListResolverEndpointIpAddresses request to get the next group of IP
	// addresses. In the next request, specify the value of NextToken from the previous
	// response.
	NextToken *string

	// The IP addresses that DNS queries pass through on their way to your network
	// (outbound endpoint) or on the way to Resolver (inbound endpoint).
	IpAddresses []*types.IpAddressResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListResolverEndpointIpAddressesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverEndpointIpAddresses{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverEndpointIpAddresses{}, middleware.After)
}

func newServiceMetadataMiddleware_opListResolverEndpointIpAddresses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListResolverEndpointIpAddresses",
	}
}
