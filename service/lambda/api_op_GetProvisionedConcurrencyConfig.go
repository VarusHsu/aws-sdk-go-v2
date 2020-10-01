// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the provisioned concurrency configuration for a function's alias or
// version.
func (c *Client) GetProvisionedConcurrencyConfig(ctx context.Context, params *GetProvisionedConcurrencyConfigInput, optFns ...func(*Options)) (*GetProvisionedConcurrencyConfigOutput, error) {
	stack := middleware.NewStack("GetProvisionedConcurrencyConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetProvisionedConcurrencyConfigMiddlewares(stack)
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
	addOpGetProvisionedConcurrencyConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetProvisionedConcurrencyConfig(options.Region), middleware.Before)
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
			OperationName: "GetProvisionedConcurrencyConfig",
			Err:           err,
		}
	}
	out := result.(*GetProvisionedConcurrencyConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProvisionedConcurrencyConfigInput struct {

	// The version number or alias name.
	//
	// This member is required.
	Qualifier *string

	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// my-function.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//     * Partial ARN -
	// 123456789012:function:my-function.
	//
	// The length constraint applies only to the
	// full ARN. If you specify only the function name, it is limited to 64 characters
	// in length.
	//
	// This member is required.
	FunctionName *string
}

type GetProvisionedConcurrencyConfigOutput struct {

	// The amount of provisioned concurrency requested.
	RequestedProvisionedConcurrentExecutions *int32

	// The amount of provisioned concurrency available.
	AvailableProvisionedConcurrentExecutions *int32

	// For failed allocations, the reason that provisioned concurrency could not be
	// allocated.
	StatusReason *string

	// The date and time that a user last updated the configuration, in ISO 8601 format
	// (https://www.iso.org/iso-8601-date-and-time-format.html).
	LastModified *string

	// The status of the allocation process.
	Status types.ProvisionedConcurrencyStatusEnum

	// The amount of provisioned concurrency allocated.
	AllocatedProvisionedConcurrentExecutions *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetProvisionedConcurrencyConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetProvisionedConcurrencyConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProvisionedConcurrencyConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetProvisionedConcurrencyConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetProvisionedConcurrencyConfig",
	}
}
