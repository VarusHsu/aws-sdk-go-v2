// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends a message. Note that ConnectionToken is used for invoking this API instead
// of ParticipantToken.
func (c *Client) SendMessage(ctx context.Context, params *SendMessageInput, optFns ...func(*Options)) (*SendMessageOutput, error) {
	stack := middleware.NewStack("SendMessage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSendMessageMiddlewares(stack)
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
	addIdempotencyToken_opSendMessageMiddleware(stack, options)
	addOpSendMessageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendMessage(options.Region), middleware.Before)
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
			OperationName: "SendMessage",
			Err:           err,
		}
	}
	out := result.(*SendMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendMessageInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The type of the content. Supported types are text/plain.
	//
	// This member is required.
	ContentType *string

	// The authentication token associated with the connection.
	//
	// This member is required.
	ConnectionToken *string

	// The content of the message.
	//
	// This member is required.
	Content *string
}

type SendMessageOutput struct {

	// The ID of the message.
	Id *string

	// The time when the message was sent. It's specified in ISO 8601 format:
	// yyyy-MM-ddThh:mm:ss.SSSZ. For example, 2019-11-08T02:41:28.172Z.
	AbsoluteTime *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSendMessageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSendMessage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSendMessage{}, middleware.After)
}

type idempotencyToken_initializeOpSendMessage struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpSendMessage) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpSendMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*SendMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *SendMessageInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opSendMessageMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpSendMessage{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opSendMessage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "SendMessage",
	}
}
