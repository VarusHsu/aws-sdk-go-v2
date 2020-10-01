// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an export job for an application.
func (c *Client) CreateExportJob(ctx context.Context, params *CreateExportJobInput, optFns ...func(*Options)) (*CreateExportJobOutput, error) {
	stack := middleware.NewStack("CreateExportJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateExportJobMiddlewares(stack)
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
	addOpCreateExportJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExportJob(options.Region), middleware.Before)
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
			OperationName: "CreateExportJob",
			Err:           err,
		}
	}
	out := result.(*CreateExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExportJobInput struct {

	// Specifies the settings for a job that exports endpoint definitions to an Amazon
	// Simple Storage Service (Amazon S3) bucket.
	//
	// This member is required.
	ExportJobRequest *types.ExportJobRequest

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string
}

type CreateExportJobOutput struct {

	// Provides information about the status and settings of a job that exports
	// endpoint definitions to a file. The file can be added directly to an Amazon
	// Simple Storage Service (Amazon S3) bucket by using the Amazon Pinpoint API or
	// downloaded directly to a computer by using the Amazon Pinpoint console.
	//
	// This member is required.
	ExportJobResponse *types.ExportJobResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateExportJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateExportJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateExportJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateExportJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "CreateExportJob",
	}
}
