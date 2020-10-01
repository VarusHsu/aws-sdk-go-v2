// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get details for a reservation.
func (c *Client) DescribeReservation(ctx context.Context, params *DescribeReservationInput, optFns ...func(*Options)) (*DescribeReservationOutput, error) {
	stack := middleware.NewStack("DescribeReservation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeReservationMiddlewares(stack)
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
	addOpDescribeReservationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservation(options.Region), middleware.Before)
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
			OperationName: "DescribeReservation",
			Err:           err,
		}
	}
	out := result.(*DescribeReservationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeReservationRequest
type DescribeReservationInput struct {

	// Unique reservation ID, e.g. '1234567'
	//
	// This member is required.
	ReservationId *string
}

// Placeholder documentation for DescribeReservationResponse
type DescribeReservationOutput struct {

	// Number of reserved resources
	Count *int32

	// One-time charge for each reserved resource, e.g. '0.0' for a NO_UPFRONT offering
	FixedPrice *float64

	// Current state of reservation, e.g. 'ACTIVE'
	State types.ReservationState

	// Units for duration, e.g. 'MONTHS'
	DurationUnits types.OfferingDurationUnits

	// User specified reservation name
	Name *string

	// AWS region, e.g. 'us-west-2'
	Region *string

	// Offering type, e.g. 'NO_UPFRONT'
	OfferingType types.OfferingType

	// Unique offering ID, e.g. '87654321'
	OfferingId *string

	// A collection of key-value pairs
	Tags map[string]*string

	// Offering description, e.g. 'HD AVC output at 10-20 Mbps, 30 fps, and standard VQ
	// in US West (Oregon)'
	OfferingDescription *string

	// Unique reservation ID, e.g. '1234567'
	ReservationId *string

	// Reservation UTC start date and time in ISO-8601 format, e.g.
	// '2018-03-01T00:00:00'
	Start *string

	// Reservation UTC end date and time in ISO-8601 format, e.g. '2019-03-01T00:00:00'
	End *string

	// Unique reservation ARN, e.g.
	// 'arn:aws:medialive:us-west-2:123456789012:reservation:1234567'
	Arn *string

	// Resource configuration details
	ResourceSpecification *types.ReservationResourceSpecification

	// Lease duration, e.g. '12'
	Duration *int32

	// Recurring usage charge for each reserved resource, e.g. '157.0'
	UsagePrice *float64

	// Currency code for usagePrice and fixedPrice in ISO-4217 format, e.g. 'USD'
	CurrencyCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeReservationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReservation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReservation{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeReservation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeReservation",
	}
}
