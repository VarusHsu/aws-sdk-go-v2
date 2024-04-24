// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration of a running DataSync task execution. Currently, the
// only Option that you can modify with UpdateTaskExecution is BytesPerSecond (https://docs.aws.amazon.com/datasync/latest/userguide/API_Options.html#DataSync-Type-Options-BytesPerSecond)
// , which throttles bandwidth for a running or queued task execution.
func (c *Client) UpdateTaskExecution(ctx context.Context, params *UpdateTaskExecutionInput, optFns ...func(*Options)) (*UpdateTaskExecutionOutput, error) {
	if params == nil {
		params = &UpdateTaskExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTaskExecution", params, optFns, c.addOperationUpdateTaskExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTaskExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTaskExecutionInput struct {

	// Indicates how your transfer task is configured. These options include how
	// DataSync handles files, objects, and their associated metadata during your
	// transfer. You also can specify how to verify data integrity, set bandwidth
	// limits for your task, among other options. Each option has a default value.
	// Unless you need to, you don't have to configure any option before calling
	// StartTaskExecution (https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html)
	// . You also can override your task options for each task execution. For example,
	// you might want to adjust the LogLevel for an individual execution.
	//
	// This member is required.
	Options *types.Options

	// Specifies the Amazon Resource Name (ARN) of the task execution that you're
	// updating.
	//
	// This member is required.
	TaskExecutionArn *string

	noSmithyDocumentSerde
}

type UpdateTaskExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTaskExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTaskExecution"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateTaskExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTaskExecution(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateTaskExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTaskExecution",
	}
}
