// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Accept a VPC peering connection request. To accept a request, the VPC peering
// connection must be in the pending-acceptance state, and you must be the owner of
// the peer VPC. Use DescribeVpcPeeringConnections () to view your outstanding VPC
// peering connection requests. For an inter-Region VPC peering connection request,
// you must accept the VPC peering connection in the Region of the accepter VPC.
func (c *Client) AcceptVpcPeeringConnection(ctx context.Context, params *AcceptVpcPeeringConnectionInput, optFns ...func(*Options)) (*AcceptVpcPeeringConnectionOutput, error) {
	stack := middleware.NewStack("AcceptVpcPeeringConnection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpAcceptVpcPeeringConnectionMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptVpcPeeringConnection(options.Region), middleware.Before)
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
			OperationName: "AcceptVpcPeeringConnection",
			Err:           err,
		}
	}
	out := result.(*AcceptVpcPeeringConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptVpcPeeringConnectionInput struct {

	// The ID of the VPC peering connection. You must specify this parameter in the
	// request.
	VpcPeeringConnectionId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type AcceptVpcPeeringConnectionOutput struct {

	// Information about the VPC peering connection.
	VpcPeeringConnection *types.VpcPeeringConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpAcceptVpcPeeringConnectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpAcceptVpcPeeringConnection{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpAcceptVpcPeeringConnection{}, middleware.After)
}

func newServiceMetadataMiddleware_opAcceptVpcPeeringConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AcceptVpcPeeringConnection",
	}
}
