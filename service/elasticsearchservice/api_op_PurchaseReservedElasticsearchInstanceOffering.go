// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows you to purchase reserved Elasticsearch instances.
func (c *Client) PurchaseReservedElasticsearchInstanceOffering(ctx context.Context, params *PurchaseReservedElasticsearchInstanceOfferingInput, optFns ...func(*Options)) (*PurchaseReservedElasticsearchInstanceOfferingOutput, error) {
	stack := middleware.NewStack("PurchaseReservedElasticsearchInstanceOffering", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpPurchaseReservedElasticsearchInstanceOfferingMiddlewares(stack)
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
	addOpPurchaseReservedElasticsearchInstanceOfferingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseReservedElasticsearchInstanceOffering(options.Region), middleware.Before)
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
			OperationName: "PurchaseReservedElasticsearchInstanceOffering",
			Err:           err,
		}
	}
	out := result.(*PurchaseReservedElasticsearchInstanceOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for parameters to PurchaseReservedElasticsearchInstanceOffering
type PurchaseReservedElasticsearchInstanceOfferingInput struct {

	// The number of Elasticsearch instances to reserve.
	InstanceCount *int32

	// The ID of the reserved Elasticsearch instance offering to purchase.
	//
	// This member is required.
	ReservedElasticsearchInstanceOfferingId *string

	// A customer-specified identifier to track this reservation.
	//
	// This member is required.
	ReservationName *string
}

// Represents the output of a PurchaseReservedElasticsearchInstanceOffering
// operation.
type PurchaseReservedElasticsearchInstanceOfferingOutput struct {

	// Details of the reserved Elasticsearch instance which was purchased.
	ReservedElasticsearchInstanceId *string

	// The customer-specified identifier used to track this reservation.
	ReservationName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpPurchaseReservedElasticsearchInstanceOfferingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpPurchaseReservedElasticsearchInstanceOffering{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpPurchaseReservedElasticsearchInstanceOffering{}, middleware.After)
}

func newServiceMetadataMiddleware_opPurchaseReservedElasticsearchInstanceOffering(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "PurchaseReservedElasticsearchInstanceOffering",
	}
}
