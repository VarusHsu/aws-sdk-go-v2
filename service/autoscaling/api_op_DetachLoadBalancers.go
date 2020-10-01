// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches one or more Classic Load Balancers from the specified Auto Scaling
// group. This operation detaches only Classic Load Balancers. If you have
// Application Load Balancers or Network Load Balancers, use the
// DetachLoadBalancerTargetGroups () API instead. When you detach a load balancer,
// it enters the Removing state while deregistering the instances in the group.
// When all instances are deregistered, then you can no longer describe the load
// balancer using the DescribeLoadBalancers () API call. The instances remain
// running.
func (c *Client) DetachLoadBalancers(ctx context.Context, params *DetachLoadBalancersInput, optFns ...func(*Options)) (*DetachLoadBalancersOutput, error) {
	stack := middleware.NewStack("DetachLoadBalancers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDetachLoadBalancersMiddlewares(stack)
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
	addOpDetachLoadBalancersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachLoadBalancers(options.Region), middleware.Before)
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
			OperationName: "DetachLoadBalancers",
			Err:           err,
		}
	}
	out := result.(*DetachLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachLoadBalancersInput struct {

	// The names of the load balancers. You can specify up to 10 load balancers.
	//
	// This member is required.
	LoadBalancerNames []*string

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string
}

type DetachLoadBalancersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDetachLoadBalancersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDetachLoadBalancers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDetachLoadBalancers{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachLoadBalancers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DetachLoadBalancers",
	}
}
