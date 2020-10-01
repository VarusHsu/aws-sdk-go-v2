// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a job execution.
func (c *Client) DescribeJobExecution(ctx context.Context, params *DescribeJobExecutionInput, optFns ...func(*Options)) (*DescribeJobExecutionOutput, error) {
	stack := middleware.NewStack("DescribeJobExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeJobExecutionMiddlewares(stack)
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
	addOpDescribeJobExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobExecution(options.Region), middleware.Before)
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
			OperationName: "DescribeJobExecution",
			Err:           err,
		}
	}
	out := result.(*DescribeJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJobExecutionInput struct {

	// The unique identifier you assigned to this job when it was created.
	//
	// This member is required.
	JobId *string

	// The name of the thing on which the job execution is running.
	//
	// This member is required.
	ThingName *string

	// A string (consisting of the digits "0" through "9" which is used to specify a
	// particular job execution on a particular device.
	ExecutionNumber *int64
}

type DescribeJobExecutionOutput struct {

	// Information about the job execution.
	Execution *types.JobExecution

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeJobExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobExecution{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeJobExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeJobExecution",
	}
}
