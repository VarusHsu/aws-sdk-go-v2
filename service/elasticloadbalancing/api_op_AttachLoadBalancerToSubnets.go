// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more subnets to the set of configured subnets for the specified load
// balancer. The load balancer evenly distributes requests across all registered
// subnets. For more information, see Add or Remove Subnets for Your Load Balancer
// in a VPC
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-manage-subnets.html)
// in the Classic Load Balancers Guide.
func (c *Client) AttachLoadBalancerToSubnets(ctx context.Context, params *AttachLoadBalancerToSubnetsInput, optFns ...func(*Options)) (*AttachLoadBalancerToSubnetsOutput, error) {
	stack := middleware.NewStack("AttachLoadBalancerToSubnets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpAttachLoadBalancerToSubnetsMiddlewares(stack)
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
	addOpAttachLoadBalancerToSubnetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachLoadBalancerToSubnets(options.Region), middleware.Before)
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
			OperationName: "AttachLoadBalancerToSubnets",
			Err:           err,
		}
	}
	out := result.(*AttachLoadBalancerToSubnetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for AttachLoaBalancerToSubnets.
type AttachLoadBalancerToSubnetsInput struct {

	// The IDs of the subnets to add. You can add only one subnet per Availability
	// Zone.
	//
	// This member is required.
	Subnets []*string

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string
}

// Contains the output of AttachLoadBalancerToSubnets.
type AttachLoadBalancerToSubnetsOutput struct {

	// The IDs of the subnets attached to the load balancer.
	Subnets []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpAttachLoadBalancerToSubnetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpAttachLoadBalancerToSubnets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachLoadBalancerToSubnets{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachLoadBalancerToSubnets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "AttachLoadBalancerToSubnets",
	}
}
