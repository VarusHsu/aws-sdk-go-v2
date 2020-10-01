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

// Removes the specified Availability Zones from the set of Availability Zones for
// the specified load balancer in EC2-Classic or a default VPC. For load balancers
// in a non-default VPC, use DetachLoadBalancerFromSubnets (). There must be at
// least one Availability Zone registered with a load balancer at all times. After
// an Availability Zone is removed, all instances registered with the load balancer
// that are in the removed Availability Zone go into the OutOfService state. Then,
// the load balancer attempts to equally balance the traffic among its remaining
// Availability Zones. For more information, see Add or Remove Availability Zones
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-az.html)
// in the Classic Load Balancers Guide.
func (c *Client) DisableAvailabilityZonesForLoadBalancer(ctx context.Context, params *DisableAvailabilityZonesForLoadBalancerInput, optFns ...func(*Options)) (*DisableAvailabilityZonesForLoadBalancerOutput, error) {
	stack := middleware.NewStack("DisableAvailabilityZonesForLoadBalancer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDisableAvailabilityZonesForLoadBalancerMiddlewares(stack)
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
	addOpDisableAvailabilityZonesForLoadBalancerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableAvailabilityZonesForLoadBalancer(options.Region), middleware.Before)
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
			OperationName: "DisableAvailabilityZonesForLoadBalancer",
			Err:           err,
		}
	}
	out := result.(*DisableAvailabilityZonesForLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DisableAvailabilityZonesForLoadBalancer.
type DisableAvailabilityZonesForLoadBalancerInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The Availability Zones.
	//
	// This member is required.
	AvailabilityZones []*string
}

// Contains the output for DisableAvailabilityZonesForLoadBalancer.
type DisableAvailabilityZonesForLoadBalancerOutput struct {

	// The remaining Availability Zones for the load balancer.
	AvailabilityZones []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDisableAvailabilityZonesForLoadBalancerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDisableAvailabilityZonesForLoadBalancer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableAvailabilityZonesForLoadBalancer{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableAvailabilityZonesForLoadBalancer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DisableAvailabilityZonesForLoadBalancer",
	}
}
