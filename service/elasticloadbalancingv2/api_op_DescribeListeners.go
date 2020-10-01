// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified listeners or the listeners for the specified Application
// Load Balancer or Network Load Balancer. You must specify either a load balancer
// or one or more listeners. For an HTTPS or TLS listener, the output includes the
// default certificate for the listener. To describe the certificate list for the
// listener, use DescribeListenerCertificates ().
func (c *Client) DescribeListeners(ctx context.Context, params *DescribeListenersInput, optFns ...func(*Options)) (*DescribeListenersOutput, error) {
	stack := middleware.NewStack("DescribeListeners", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeListenersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeListeners(options.Region), middleware.Before)
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
			OperationName: "DescribeListeners",
			Err:           err,
		}
	}
	out := result.(*DescribeListenersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeListenersInput struct {

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The Amazon Resource Names (ARN) of the listeners.
	ListenerArns []*string

	// The maximum number of results to return with this call.
	PageSize *int32

	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string
}

type DescribeListenersOutput struct {

	// Information about the listeners.
	Listeners []*types.Listener

	// If there are additional results, this is the marker for the next set of results.
	// Otherwise, this is null.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeListenersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeListeners{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeListeners{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeListeners(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeListeners",
	}
}
