// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the delivery channel. Before you can delete the delivery channel, you
// must stop the configuration recorder by using the StopConfigurationRecorder ()
// action.
func (c *Client) DeleteDeliveryChannel(ctx context.Context, params *DeleteDeliveryChannelInput, optFns ...func(*Options)) (*DeleteDeliveryChannelOutput, error) {
	stack := middleware.NewStack("DeleteDeliveryChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteDeliveryChannelMiddlewares(stack)
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
	addOpDeleteDeliveryChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDeliveryChannel(options.Region), middleware.Before)
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
			OperationName: "DeleteDeliveryChannel",
			Err:           err,
		}
	}
	out := result.(*DeleteDeliveryChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DeleteDeliveryChannel () action. The action accepts the
// following data, in JSON format.
type DeleteDeliveryChannelInput struct {

	// The name of the delivery channel to delete.
	//
	// This member is required.
	DeliveryChannelName *string
}

type DeleteDeliveryChannelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteDeliveryChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDeliveryChannel{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDeliveryChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDeliveryChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteDeliveryChannel",
	}
}
