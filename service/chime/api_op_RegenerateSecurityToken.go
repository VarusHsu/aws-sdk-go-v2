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

// Regenerates the security token for a bot.
func (c *Client) RegenerateSecurityToken(ctx context.Context, params *RegenerateSecurityTokenInput, optFns ...func(*Options)) (*RegenerateSecurityTokenOutput, error) {
	stack := middleware.NewStack("RegenerateSecurityToken", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRegenerateSecurityTokenMiddlewares(stack)
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
	addOpRegenerateSecurityTokenValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegenerateSecurityToken(options.Region), middleware.Before)
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
			OperationName: "RegenerateSecurityToken",
			Err:           err,
		}
	}
	out := result.(*RegenerateSecurityTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegenerateSecurityTokenInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The bot ID.
	//
	// This member is required.
	BotId *string
}

type RegenerateSecurityTokenOutput struct {

	// A resource that allows Enterprise account administrators to configure an
	// interface to receive events from Amazon Chime.
	Bot *types.Bot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRegenerateSecurityTokenMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRegenerateSecurityToken{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRegenerateSecurityToken{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegenerateSecurityToken(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "RegenerateSecurityToken",
	}
}
