// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves details for the specified phone number order, such as order creation
// timestamp, phone numbers in E.164 format, product type, and order status.
func (c *Client) GetPhoneNumberOrder(ctx context.Context, params *GetPhoneNumberOrderInput, optFns ...func(*Options)) (*GetPhoneNumberOrderOutput, error) {
	stack := middleware.NewStack("GetPhoneNumberOrder", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetPhoneNumberOrderMiddlewares(stack)
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
	addOpGetPhoneNumberOrderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPhoneNumberOrder(options.Region), middleware.Before)
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
			OperationName: "GetPhoneNumberOrder",
			Err:           err,
		}
	}
	out := result.(*GetPhoneNumberOrderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPhoneNumberOrderInput struct {

	// The ID for the phone number order.
	//
	// This member is required.
	PhoneNumberOrderId *string
}

type GetPhoneNumberOrderOutput struct {

	// The phone number order details.
	PhoneNumberOrder *types.PhoneNumberOrder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetPhoneNumberOrderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetPhoneNumberOrder{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPhoneNumberOrder{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPhoneNumberOrder(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetPhoneNumberOrder",
	}
}
