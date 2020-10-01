// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a reference data source configuration from the specified SQL-based
// Amazon Kinesis Data Analytics application's configuration. If the application is
// running, Kinesis Data Analytics immediately removes the in-application table
// that you created using the AddApplicationReferenceDataSource () operation.
func (c *Client) DeleteApplicationReferenceDataSource(ctx context.Context, params *DeleteApplicationReferenceDataSourceInput, optFns ...func(*Options)) (*DeleteApplicationReferenceDataSourceOutput, error) {
	stack := middleware.NewStack("DeleteApplicationReferenceDataSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteApplicationReferenceDataSourceMiddlewares(stack)
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
	addOpDeleteApplicationReferenceDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationReferenceDataSource(options.Region), middleware.Before)
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
			OperationName: "DeleteApplicationReferenceDataSource",
			Err:           err,
		}
	}
	out := result.(*DeleteApplicationReferenceDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationReferenceDataSourceInput struct {

	// The ID of the reference data source. When you add a reference data source to
	// your application using the AddApplicationReferenceDataSource (), Kinesis Data
	// Analytics assigns an ID. You can use the DescribeApplication () operation to get
	// the reference ID.
	//
	// This member is required.
	ReferenceId *string

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// The current application version. You can use the DescribeApplication ()
	// operation to get the current application version. If the version specified is
	// not the current version, the ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64
}

type DeleteApplicationReferenceDataSourceOutput struct {

	// The updated version ID of the application.
	ApplicationVersionId *int64

	// The application Amazon Resource Name (ARN).
	ApplicationARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteApplicationReferenceDataSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationReferenceDataSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationReferenceDataSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteApplicationReferenceDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationReferenceDataSource",
	}
}
