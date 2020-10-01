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

// Replaces the set of policies associated with the specified port on which the EC2
// instance is listening with a new set of policies. At this time, only the
// back-end server authentication policy type can be applied to the instance ports;
// this policy type is composed of multiple public key policies. Each time you use
// SetLoadBalancerPoliciesForBackendServer to enable the policies, use the
// PolicyNames parameter to list the policies that you want to enable. You can use
// DescribeLoadBalancers () or DescribeLoadBalancerPolicies () to verify that the
// policy is associated with the EC2 instance.  <p>For more information about
// enabling back-end instance authentication, see <a
// href="https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-create-https-ssl-load-balancer.html#configure_backendauth_clt">Configure
// Back-end Instance Authentication</a> in the <i>Classic Load Balancers Guide</i>.
// For more information about Proxy Protocol, see <a
// href="https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-proxy-protocol.html">Configure
// Proxy Protocol Support</a> in the <i>Classic Load Balancers Guide</i>.</p>
func (c *Client) SetLoadBalancerPoliciesForBackendServer(ctx context.Context, params *SetLoadBalancerPoliciesForBackendServerInput, optFns ...func(*Options)) (*SetLoadBalancerPoliciesForBackendServerOutput, error) {
	stack := middleware.NewStack("SetLoadBalancerPoliciesForBackendServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetLoadBalancerPoliciesForBackendServerMiddlewares(stack)
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
	addOpSetLoadBalancerPoliciesForBackendServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetLoadBalancerPoliciesForBackendServer(options.Region), middleware.Before)
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
			OperationName: "SetLoadBalancerPoliciesForBackendServer",
			Err:           err,
		}
	}
	out := result.(*SetLoadBalancerPoliciesForBackendServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetLoadBalancerPoliciesForBackendServer.
type SetLoadBalancerPoliciesForBackendServerInput struct {

	// The port number associated with the EC2 instance.
	//
	// This member is required.
	InstancePort *int32

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The names of the policies. If the list is empty, then all current polices are
	// removed from the EC2 instance.
	//
	// This member is required.
	PolicyNames []*string
}

// Contains the output of SetLoadBalancerPoliciesForBackendServer.
type SetLoadBalancerPoliciesForBackendServerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetLoadBalancerPoliciesForBackendServerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetLoadBalancerPoliciesForBackendServer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetLoadBalancerPoliciesForBackendServer{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetLoadBalancerPoliciesForBackendServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "SetLoadBalancerPoliciesForBackendServer",
	}
}
