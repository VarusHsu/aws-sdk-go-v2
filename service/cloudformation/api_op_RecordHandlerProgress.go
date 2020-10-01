// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Reports progress of a resource handler to CloudFormation. Reserved for use by
// the CloudFormation CLI
// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
// Do not use this API in your code.
func (c *Client) RecordHandlerProgress(ctx context.Context, params *RecordHandlerProgressInput, optFns ...func(*Options)) (*RecordHandlerProgressOutput, error) {
	stack := middleware.NewStack("RecordHandlerProgress", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRecordHandlerProgressMiddlewares(stack)
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
	addOpRecordHandlerProgressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRecordHandlerProgress(options.Region), middleware.Before)
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
			OperationName: "RecordHandlerProgress",
			Err:           err,
		}
	}
	out := result.(*RecordHandlerProgressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecordHandlerProgressInput struct {

	// Reserved for use by the CloudFormation CLI
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
	//
	// This member is required.
	BearerToken *string

	// Reserved for use by the CloudFormation CLI
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
	CurrentOperationStatus types.OperationStatus

	// Reserved for use by the CloudFormation CLI
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
	ErrorCode types.HandlerErrorCode

	// Reserved for use by the CloudFormation CLI
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
	ClientRequestToken *string

	// Reserved for use by the CloudFormation CLI
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
	//
	// This member is required.
	OperationStatus types.OperationStatus

	// Reserved for use by the CloudFormation CLI
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
	ResourceModel *string

	// Reserved for use by the CloudFormation CLI
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
	StatusMessage *string
}

type RecordHandlerProgressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRecordHandlerProgressMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRecordHandlerProgress{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRecordHandlerProgress{}, middleware.After)
}

func newServiceMetadataMiddleware_opRecordHandlerProgress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "RecordHandlerProgress",
	}
}
