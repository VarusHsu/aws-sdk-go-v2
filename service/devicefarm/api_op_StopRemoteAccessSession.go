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

// Ends a specified remote access session.
func (c *Client) StopRemoteAccessSession(ctx context.Context, params *StopRemoteAccessSessionInput, optFns ...func(*Options)) (*StopRemoteAccessSessionOutput, error) {
	stack := middleware.NewStack("StopRemoteAccessSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopRemoteAccessSessionMiddlewares(stack)
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
	addOpStopRemoteAccessSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopRemoteAccessSession(options.Region), middleware.Before)
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
			OperationName: "StopRemoteAccessSession",
			Err:           err,
		}
	}
	out := result.(*StopRemoteAccessSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to stop the remote access session.
type StopRemoteAccessSessionInput struct {

	// The Amazon Resource Name (ARN) of the remote access session to stop.
	//
	// This member is required.
	Arn *string
}

// Represents the response from the server that describes the remote access session
// when AWS Device Farm stops the session.
type StopRemoteAccessSessionOutput struct {

	// A container that represents the metadata from the service about the remote
	// access session you are stopping.
	RemoteAccessSession *types.RemoteAccessSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopRemoteAccessSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopRemoteAccessSession{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopRemoteAccessSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopRemoteAccessSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "StopRemoteAccessSession",
	}
}
