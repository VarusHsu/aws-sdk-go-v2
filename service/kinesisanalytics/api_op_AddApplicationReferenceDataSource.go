// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data Analytics API V2 Documentation.
//
// Adds a reference data source to an existing application.
//
// Amazon Kinesis Analytics reads reference data (that is, an Amazon S3 object)
// and creates an in-application table within your application. In the request, you
// provide the source (S3 bucket name and object key name), name of the
// in-application table to create, and the necessary mapping information that
// describes how data in Amazon S3 object maps to columns in the resulting
// in-application table.
//
// For conceptual information, see [Configuring Application Input]. For the limits on data sources you can add to
// your application, see [Limits].
//
// This operation requires permissions to perform the
// kinesisanalytics:AddApplicationOutput action.
//
// [Limits]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/limits.html
// [Configuring Application Input]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html
func (c *Client) AddApplicationReferenceDataSource(ctx context.Context, params *AddApplicationReferenceDataSourceInput, optFns ...func(*Options)) (*AddApplicationReferenceDataSourceOutput, error) {
	if params == nil {
		params = &AddApplicationReferenceDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationReferenceDataSource", params, optFns, c.addOperationAddApplicationReferenceDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationReferenceDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationReferenceDataSourceInput struct {

	// Name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// Version of the application for which you are adding the reference data source.
	// You can use the [DescribeApplication]operation to get the current application version. If the
	// version specified is not the current version, the
	// ConcurrentModificationException is returned.
	//
	// [DescribeApplication]: https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The reference data source can be an object in your Amazon S3 bucket. Amazon
	// Kinesis Analytics reads the object and copies the data into the in-application
	// table that is created. You provide an S3 bucket, object key name, and the
	// resulting in-application table that is created. You must also provide an IAM
	// role with the necessary permissions that Amazon Kinesis Analytics can assume to
	// read the object from your S3 bucket on your behalf.
	//
	// This member is required.
	ReferenceDataSource *types.ReferenceDataSource

	noSmithyDocumentSerde
}

type AddApplicationReferenceDataSourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddApplicationReferenceDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationReferenceDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationReferenceDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddApplicationReferenceDataSource"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpAddApplicationReferenceDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationReferenceDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationReferenceDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddApplicationReferenceDataSource",
	}
}
