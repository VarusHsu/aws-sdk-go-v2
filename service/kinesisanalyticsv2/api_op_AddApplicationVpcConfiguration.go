// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a Virtual Private Cloud (VPC) configuration to the application.
// Applications can use VPCs to store and access resources securely.
//
// Note the following about VPC configurations for Managed Service for Apache
// Flink applications:
//
//   - VPC configurations are not supported for SQL applications.
//
//   - When a VPC is added to a Managed Service for Apache Flink application, the
//     application can no longer be accessed from the Internet directly. To enable
//     Internet access to the application, add an Internet gateway to your VPC.
func (c *Client) AddApplicationVpcConfiguration(ctx context.Context, params *AddApplicationVpcConfigurationInput, optFns ...func(*Options)) (*AddApplicationVpcConfigurationOutput, error) {
	if params == nil {
		params = &AddApplicationVpcConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddApplicationVpcConfiguration", params, optFns, c.addOperationAddApplicationVpcConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddApplicationVpcConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddApplicationVpcConfigurationInput struct {

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// Description of the VPC to add to the application.
	//
	// This member is required.
	VpcConfiguration *types.VpcConfiguration

	// A value you use to implement strong concurrency for application updates. You
	// must provide the ApplicationVersionID or the ConditionalToken . You get the
	// application's current ConditionalToken using DescribeApplication. For better concurrency support,
	// use the ConditionalToken parameter instead of CurrentApplicationVersionId .
	ConditionalToken *string

	// The version of the application to which you want to add the VPC configuration.
	// You must provide the CurrentApplicationVersionId or the ConditionalToken . You
	// can use the DescribeApplicationoperation to get the current application version. If the version
	// specified is not the current version, the ConcurrentModificationException is
	// returned. For better concurrency support, use the ConditionalToken parameter
	// instead of CurrentApplicationVersionId .
	CurrentApplicationVersionId *int64

	noSmithyDocumentSerde
}

type AddApplicationVpcConfigurationOutput struct {

	// The ARN of the application.
	ApplicationARN *string

	// Provides the current application version. Managed Service for Apache Flink
	// updates the ApplicationVersionId each time you update the application.
	ApplicationVersionId *int64

	// Operation ID for tracking AddApplicationVpcConfiguration request
	OperationId *string

	// The parameters of the new VPC configuration.
	VpcConfigurationDescription *types.VpcConfigurationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddApplicationVpcConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddApplicationVpcConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddApplicationVpcConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddApplicationVpcConfiguration"); err != nil {
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
	if err = addOpAddApplicationVpcConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddApplicationVpcConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddApplicationVpcConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddApplicationVpcConfiguration",
	}
}
