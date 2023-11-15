// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more subnets to the set of configured subnets for the specified
// load balancer. The load balancer evenly distributes requests across all
// registered subnets. For more information, see Add or Remove Subnets for Your
// Load Balancer in a VPC (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-manage-subnets.html)
// in the Classic Load Balancers Guide.
func (c *Client) AttachLoadBalancerToSubnets(ctx context.Context, params *AttachLoadBalancerToSubnetsInput, optFns ...func(*Options)) (*AttachLoadBalancerToSubnetsOutput, error) {
	if params == nil {
		params = &AttachLoadBalancerToSubnetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachLoadBalancerToSubnets", params, optFns, c.addOperationAttachLoadBalancerToSubnetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachLoadBalancerToSubnetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for AttachLoaBalancerToSubnets.
type AttachLoadBalancerToSubnetsInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The IDs of the subnets to add. You can add only one subnet per Availability
	// Zone.
	//
	// This member is required.
	Subnets []string

	noSmithyDocumentSerde
}

// Contains the output of AttachLoadBalancerToSubnets.
type AttachLoadBalancerToSubnetsOutput struct {

	// The IDs of the subnets attached to the load balancer.
	Subnets []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachLoadBalancerToSubnetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAttachLoadBalancerToSubnets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachLoadBalancerToSubnets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AttachLoadBalancerToSubnets"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAttachLoadBalancerToSubnetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachLoadBalancerToSubnets(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAttachLoadBalancerToSubnets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AttachLoadBalancerToSubnets",
	}
}
