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

// Rejects a transit gateway peering attachment request.
func (c *Client) RejectTransitGatewayPeeringAttachment(ctx context.Context, params *RejectTransitGatewayPeeringAttachmentInput, optFns ...func(*Options)) (*RejectTransitGatewayPeeringAttachmentOutput, error) {
	stack := middleware.NewStack("RejectTransitGatewayPeeringAttachment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpRejectTransitGatewayPeeringAttachmentMiddlewares(stack)
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
	addOpRejectTransitGatewayPeeringAttachmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRejectTransitGatewayPeeringAttachment(options.Region), middleware.Before)
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
			OperationName: "RejectTransitGatewayPeeringAttachment",
			Err:           err,
		}
	}
	out := result.(*RejectTransitGatewayPeeringAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RejectTransitGatewayPeeringAttachmentInput struct {

	// The ID of the transit gateway peering attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type RejectTransitGatewayPeeringAttachmentOutput struct {

	// The transit gateway peering attachment.
	TransitGatewayPeeringAttachment *types.TransitGatewayPeeringAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpRejectTransitGatewayPeeringAttachmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpRejectTransitGatewayPeeringAttachment{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpRejectTransitGatewayPeeringAttachment{}, middleware.After)
}

func newServiceMetadataMiddleware_opRejectTransitGatewayPeeringAttachment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RejectTransitGatewayPeeringAttachment",
	}
}
