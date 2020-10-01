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

// Removes a VPC configuration from a Kinesis Data Analytics application.
func (c *Client) DeleteApplicationVpcConfiguration(ctx context.Context, params *DeleteApplicationVpcConfigurationInput, optFns ...func(*Options)) (*DeleteApplicationVpcConfigurationOutput, error) {
	stack := middleware.NewStack("DeleteApplicationVpcConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteApplicationVpcConfigurationMiddlewares(stack)
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
	addOpDeleteApplicationVpcConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationVpcConfiguration(options.Region), middleware.Before)
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
			OperationName: "DeleteApplicationVpcConfiguration",
			Err:           err,
		}
	}
	out := result.(*DeleteApplicationVpcConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationVpcConfigurationInput struct {

	// The current application version ID. You can retrieve the application version ID
	// using DescribeApplication ().
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The ID of the VPC configuration to delete.
	//
	// This member is required.
	VpcConfigurationId *string

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string
}

type DeleteApplicationVpcConfigurationOutput struct {

	// The ARN of the Kinesis Data Analytics application.
	ApplicationARN *string

	// The updated version ID of the application.
	ApplicationVersionId *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteApplicationVpcConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationVpcConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationVpcConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteApplicationVpcConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationVpcConfiguration",
	}
}
