// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the current status and future status of all offerings purchased by an AWS
// account. The response indicates how many offerings are currently available and
// the offerings that will be available in the next period. The API returns a
// NotEligible error if the user is not permitted to invoke the operation. If you
// must be able to invoke this operation, contact aws-devicefarm-support@amazon.com
// (mailto:aws-devicefarm-support@amazon.com).
func (c *Client) GetOfferingStatus(ctx context.Context, params *GetOfferingStatusInput, optFns ...func(*Options)) (*GetOfferingStatusOutput, error) {
	stack := middleware.NewStack("GetOfferingStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetOfferingStatusMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOfferingStatus(options.Region), middleware.Before)
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
			OperationName: "GetOfferingStatus",
			Err:           err,
		}
	}
	out := result.(*GetOfferingStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to retrieve the offering status for the specified
// customer or account.
type GetOfferingStatusInput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Returns the status result for a device offering.
type GetOfferingStatusOutput struct {

	// When specified, gets the offering status for the current period.
	Current map[string]*types.OfferingStatus

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// When specified, gets the offering status for the next period.
	NextPeriod map[string]*types.OfferingStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetOfferingStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetOfferingStatus{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOfferingStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetOfferingStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "GetOfferingStatus",
	}
}
