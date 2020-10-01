// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimeservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes session information for a specified bot, alias, and user ID.
func (c *Client) DeleteSession(ctx context.Context, params *DeleteSessionInput, optFns ...func(*Options)) (*DeleteSessionOutput, error) {
	stack := middleware.NewStack("DeleteSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteSessionMiddlewares(stack)
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
	addOpDeleteSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSession(options.Region), middleware.Before)
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
			OperationName: "DeleteSession",
			Err:           err,
		}
	}
	out := result.(*DeleteSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSessionInput struct {

	// The identifier of the user associated with the session data.
	//
	// This member is required.
	UserId *string

	// The name of the bot that contains the session data.
	//
	// This member is required.
	BotName *string

	// The alias in use for the bot that contains the session data.
	//
	// This member is required.
	BotAlias *string
}

type DeleteSessionOutput struct {

	// The alias in use for the bot associated with the session data.
	BotAlias *string

	// The unique identifier for the session.
	SessionId *string

	// The ID of the client application user.
	UserId *string

	// The name of the bot associated with the session data.
	BotName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSession{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteSession",
	}
}
