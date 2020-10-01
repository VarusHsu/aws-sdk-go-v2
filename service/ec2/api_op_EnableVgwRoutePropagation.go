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

// Enables a virtual private gateway (VGW) to propagate routes to the specified
// route table of a VPC.
func (c *Client) EnableVgwRoutePropagation(ctx context.Context, params *EnableVgwRoutePropagationInput, optFns ...func(*Options)) (*EnableVgwRoutePropagationOutput, error) {
	stack := middleware.NewStack("EnableVgwRoutePropagation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpEnableVgwRoutePropagationMiddlewares(stack)
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
	addOpEnableVgwRoutePropagationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableVgwRoutePropagation(options.Region), middleware.Before)
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
			OperationName: "EnableVgwRoutePropagation",
			Err:           err,
		}
	}
	out := result.(*EnableVgwRoutePropagationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for EnableVgwRoutePropagation.
type EnableVgwRoutePropagationInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the virtual private gateway that is attached to a VPC. The virtual
	// private gateway must be attached to the same VPC that the routing tables are
	// associated with.
	//
	// This member is required.
	GatewayId *string

	// The ID of the route table. The routing table must be associated with the same
	// VPC that the virtual private gateway is attached to.
	//
	// This member is required.
	RouteTableId *string
}

type EnableVgwRoutePropagationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpEnableVgwRoutePropagationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpEnableVgwRoutePropagation{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpEnableVgwRoutePropagation{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableVgwRoutePropagation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableVgwRoutePropagation",
	}
}
