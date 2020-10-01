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

// Cancels one or more Spot Instance requests. Canceling a Spot Instance request
// does not terminate running Spot Instances associated with the request.
func (c *Client) CancelSpotInstanceRequests(ctx context.Context, params *CancelSpotInstanceRequestsInput, optFns ...func(*Options)) (*CancelSpotInstanceRequestsOutput, error) {
	stack := middleware.NewStack("CancelSpotInstanceRequests", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCancelSpotInstanceRequestsMiddlewares(stack)
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
	addOpCancelSpotInstanceRequestsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelSpotInstanceRequests(options.Region), middleware.Before)
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
			OperationName: "CancelSpotInstanceRequests",
			Err:           err,
		}
	}
	out := result.(*CancelSpotInstanceRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CancelSpotInstanceRequests.
type CancelSpotInstanceRequestsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more Spot Instance request IDs.
	//
	// This member is required.
	SpotInstanceRequestIds []*string
}

// Contains the output of CancelSpotInstanceRequests.
type CancelSpotInstanceRequestsOutput struct {

	// One or more Spot Instance requests.
	CancelledSpotInstanceRequests []*types.CancelledSpotInstanceRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCancelSpotInstanceRequestsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCancelSpotInstanceRequests{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCancelSpotInstanceRequests{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelSpotInstanceRequests(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CancelSpotInstanceRequests",
	}
}
