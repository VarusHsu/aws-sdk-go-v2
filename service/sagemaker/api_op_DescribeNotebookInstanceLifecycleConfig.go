// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns a description of a notebook instance lifecycle configuration. For
// information about notebook instance lifestyle configurations, see Step 2.1:
// (Optional) Customize a Notebook Instance
// (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html).
func (c *Client) DescribeNotebookInstanceLifecycleConfig(ctx context.Context, params *DescribeNotebookInstanceLifecycleConfigInput, optFns ...func(*Options)) (*DescribeNotebookInstanceLifecycleConfigOutput, error) {
	stack := middleware.NewStack("DescribeNotebookInstanceLifecycleConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeNotebookInstanceLifecycleConfigMiddlewares(stack)
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
	addOpDescribeNotebookInstanceLifecycleConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotebookInstanceLifecycleConfig(options.Region), middleware.Before)
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
			OperationName: "DescribeNotebookInstanceLifecycleConfig",
			Err:           err,
		}
	}
	out := result.(*DescribeNotebookInstanceLifecycleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotebookInstanceLifecycleConfigInput struct {

	// The name of the lifecycle configuration to describe.
	//
	// This member is required.
	NotebookInstanceLifecycleConfigName *string
}

type DescribeNotebookInstanceLifecycleConfigOutput struct {

	// The shell script that runs every time you start a notebook instance, including
	// when you create the notebook instance.
	OnStart []*types.NotebookInstanceLifecycleHook

	// A timestamp that tells when the lifecycle configuration was last modified.
	LastModifiedTime *time.Time

	// A timestamp that tells when the lifecycle configuration was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the lifecycle configuration.
	NotebookInstanceLifecycleConfigArn *string

	// The shell script that runs only once, when you create a notebook instance.
	OnCreate []*types.NotebookInstanceLifecycleHook

	// The name of the lifecycle configuration.
	NotebookInstanceLifecycleConfigName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeNotebookInstanceLifecycleConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNotebookInstanceLifecycleConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNotebookInstanceLifecycleConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNotebookInstanceLifecycleConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeNotebookInstanceLifecycleConfig",
	}
}
