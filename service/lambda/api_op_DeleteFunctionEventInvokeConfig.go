// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the configuration for asynchronous invocation for a function, version,
// or alias. To configure options for asynchronous invocation, use
// PutFunctionEventInvokeConfig ().
func (c *Client) DeleteFunctionEventInvokeConfig(ctx context.Context, params *DeleteFunctionEventInvokeConfigInput, optFns ...func(*Options)) (*DeleteFunctionEventInvokeConfigOutput, error) {
	stack := middleware.NewStack("DeleteFunctionEventInvokeConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteFunctionEventInvokeConfigMiddlewares(stack)
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
	addOpDeleteFunctionEventInvokeConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFunctionEventInvokeConfig(options.Region), middleware.Before)
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
			OperationName: "DeleteFunctionEventInvokeConfig",
			Err:           err,
		}
	}
	out := result.(*DeleteFunctionEventInvokeConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFunctionEventInvokeConfigInput struct {

	// A version number or alias name.
	Qualifier *string

	// The name of the Lambda function, version, or alias. Name formats
	//
	//     * Function
	// name - my-function (name-only), my-function:v1 (with alias).
	//
	//     * Function ARN
	// - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//     * Partial ARN
	// - 123456789012:function:my-function.
	//
	// You can append a version number or alias
	// to any of the formats. The length constraint applies only to the full ARN. If
	// you specify only the function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string
}

type DeleteFunctionEventInvokeConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteFunctionEventInvokeConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFunctionEventInvokeConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFunctionEventInvokeConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteFunctionEventInvokeConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "DeleteFunctionEventInvokeConfig",
	}
}
