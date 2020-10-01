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

// Detaches a virtual private gateway from a VPC. You do this if you're planning to
// turn off the VPC and not use it anymore. You can confirm a virtual private
// gateway has been completely detached from a VPC by describing the virtual
// private gateway (any attachments to the virtual private gateway are also
// described). You must wait for the attachment's state to switch to detached
// before you can delete the VPC or attach a different VPC to the virtual private
// gateway.
func (c *Client) DetachVpnGateway(ctx context.Context, params *DetachVpnGatewayInput, optFns ...func(*Options)) (*DetachVpnGatewayOutput, error) {
	stack := middleware.NewStack("DetachVpnGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDetachVpnGatewayMiddlewares(stack)
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
	addOpDetachVpnGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachVpnGateway(options.Region), middleware.Before)
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
			OperationName: "DetachVpnGateway",
			Err:           err,
		}
	}
	out := result.(*DetachVpnGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DetachVpnGateway.
type DetachVpnGatewayInput struct {

	// The ID of the virtual private gateway.
	//
	// This member is required.
	VpnGatewayId *string

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DetachVpnGatewayOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDetachVpnGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDetachVpnGateway{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDetachVpnGateway{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachVpnGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DetachVpnGateway",
	}
}
