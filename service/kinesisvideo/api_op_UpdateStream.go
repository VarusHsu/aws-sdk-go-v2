// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates stream metadata, such as the device name and media type. You must
// provide the stream name or the Amazon Resource Name (ARN) of the stream. To make
// sure that you have the latest version of the stream before updating it, you can
// specify the stream version. Kinesis Video Streams assigns a version to each
// stream. When you update a stream, Kinesis Video Streams assigns a new version
// number. To get the latest stream version, use the DescribeStream API.
// UpdateStream is an asynchronous operation, and takes time to complete.
func (c *Client) UpdateStream(ctx context.Context, params *UpdateStreamInput, optFns ...func(*Options)) (*UpdateStreamOutput, error) {
	stack := middleware.NewStack("UpdateStream", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateStreamMiddlewares(stack)
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
	addOpUpdateStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStream(options.Region), middleware.Before)
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
			OperationName: "UpdateStream",
			Err:           err,
		}
	}
	out := result.(*UpdateStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStreamInput struct {

	// The ARN of the stream whose metadata you want to update.
	StreamARN *string

	// The version of the stream whose metadata you want to update.
	//
	// This member is required.
	CurrentVersion *string

	// The stream's media type. Use MediaType to specify the type of content that the
	// stream contains to the consumers of the stream. For more information about media
	// types, see Media Types
	// (http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose
	// to specify the MediaType, see Naming Requirements
	// (https://tools.ietf.org/html/rfc6838#section-4.2). To play video on the console,
	// you must specify the correct video type. For example, if the video in the stream
	// is H.264, specify video/h264 as the MediaType.
	MediaType *string

	// The name of the device that is writing to the stream. In the current
	// implementation, Kinesis Video Streams does not use this name.
	DeviceName *string

	// The name of the stream whose metadata you want to update. The stream name is an
	// identifier for the stream, and must be unique for each account and region.
	StreamName *string
}

type UpdateStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateStreamMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateStream{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateStream{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "UpdateStream",
	}
}
