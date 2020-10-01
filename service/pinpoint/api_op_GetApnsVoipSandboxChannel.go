// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the status and settings of the APNs VoIP sandbox
// channel for an application.
func (c *Client) GetApnsVoipSandboxChannel(ctx context.Context, params *GetApnsVoipSandboxChannelInput, optFns ...func(*Options)) (*GetApnsVoipSandboxChannelOutput, error) {
	stack := middleware.NewStack("GetApnsVoipSandboxChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetApnsVoipSandboxChannelMiddlewares(stack)
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
	addOpGetApnsVoipSandboxChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetApnsVoipSandboxChannel(options.Region), middleware.Before)
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
			OperationName: "GetApnsVoipSandboxChannel",
			Err:           err,
		}
	}
	out := result.(*GetApnsVoipSandboxChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApnsVoipSandboxChannelInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string
}

type GetApnsVoipSandboxChannelOutput struct {

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP sandbox channel for an application.
	//
	// This member is required.
	APNSVoipSandboxChannelResponse *types.APNSVoipSandboxChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetApnsVoipSandboxChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetApnsVoipSandboxChannel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApnsVoipSandboxChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetApnsVoipSandboxChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetApnsVoipSandboxChannel",
	}
}
