// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The CreateWorkerBlock operation allows you to prevent a Worker from working on
// your HITs. For example, you can block a Worker who is producing poor quality
// work. You can block up to 100,000 Workers.
func (c *Client) CreateWorkerBlock(ctx context.Context, params *CreateWorkerBlockInput, optFns ...func(*Options)) (*CreateWorkerBlockOutput, error) {
	stack := middleware.NewStack("CreateWorkerBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateWorkerBlockMiddlewares(stack)
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
	addOpCreateWorkerBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkerBlock(options.Region), middleware.Before)
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
			OperationName: "CreateWorkerBlock",
			Err:           err,
		}
	}
	out := result.(*CreateWorkerBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkerBlockInput struct {

	// A message explaining the reason for blocking the Worker. This parameter enables
	// you to keep track of your Workers. The Worker does not see this message.
	//
	// This member is required.
	Reason *string

	// The ID of the Worker to block.
	//
	// This member is required.
	WorkerId *string
}

type CreateWorkerBlockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateWorkerBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkerBlock{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkerBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateWorkerBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "CreateWorkerBlock",
	}
}
