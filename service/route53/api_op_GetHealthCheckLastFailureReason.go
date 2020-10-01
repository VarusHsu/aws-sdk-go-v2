// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the reason that a specified health check failed most recently.
func (c *Client) GetHealthCheckLastFailureReason(ctx context.Context, params *GetHealthCheckLastFailureReasonInput, optFns ...func(*Options)) (*GetHealthCheckLastFailureReasonOutput, error) {
	stack := middleware.NewStack("GetHealthCheckLastFailureReason", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetHealthCheckLastFailureReasonMiddlewares(stack)
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
	addOpGetHealthCheckLastFailureReasonValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetHealthCheckLastFailureReason(options.Region), middleware.Before)
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
			OperationName: "GetHealthCheckLastFailureReason",
			Err:           err,
		}
	}
	out := result.(*GetHealthCheckLastFailureReasonOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for the reason that a health check failed most recently.
type GetHealthCheckLastFailureReasonInput struct {

	// The ID for the health check for which you want the last failure reason. When you
	// created the health check, CreateHealthCheck returned the ID in the response, in
	// the HealthCheckId element. If you want to get the last failure reason for a
	// calculated health check, you must use the Amazon Route 53 console or the
	// CloudWatch console. You can't use GetHealthCheckLastFailureReason for a
	// calculated health check.
	//
	// This member is required.
	HealthCheckId *string
}

// A complex type that contains the response to a GetHealthCheckLastFailureReason
// request.
type GetHealthCheckLastFailureReasonOutput struct {

	// A list that contains one Observation element for each Amazon Route 53 health
	// checker that is reporting a last failure reason.
	//
	// This member is required.
	HealthCheckObservations []*types.HealthCheckObservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetHealthCheckLastFailureReasonMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetHealthCheckLastFailureReason{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetHealthCheckLastFailureReason{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetHealthCheckLastFailureReason(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHealthCheckLastFailureReason",
	}
}
