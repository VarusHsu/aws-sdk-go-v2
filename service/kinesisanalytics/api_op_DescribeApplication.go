// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation (). Returns information about a specific Amazon
// Kinesis Analytics application. If you want to retrieve a list of all
// applications in your account, use the ListApplications
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_ListApplications.html)
// operation. This operation requires permissions to perform the
// kinesisanalytics:DescribeApplication action. You can use DescribeApplication to
// get the current application versionId, which you need to call other operations
// such as Update.
func (c *Client) DescribeApplication(ctx context.Context, params *DescribeApplicationInput, optFns ...func(*Options)) (*DescribeApplicationOutput, error) {
	stack := middleware.NewStack("DescribeApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeApplicationMiddlewares(stack)
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
	addOpDescribeApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplication(options.Region), middleware.Before)
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
			OperationName: "DescribeApplication",
			Err:           err,
		}
	}
	out := result.(*DescribeApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeApplicationInput struct {

	// Name of the application.
	//
	// This member is required.
	ApplicationName *string
}

//
type DescribeApplicationOutput struct {

	// Provides a description of the application, such as the application Amazon
	// Resource Name (ARN), status, latest version, and input and output configuration
	// details.
	//
	// This member is required.
	ApplicationDetail *types.ApplicationDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApplication{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DescribeApplication",
	}
}
