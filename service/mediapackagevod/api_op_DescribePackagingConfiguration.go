// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a description of a MediaPackage VOD PackagingConfiguration resource.
func (c *Client) DescribePackagingConfiguration(ctx context.Context, params *DescribePackagingConfigurationInput, optFns ...func(*Options)) (*DescribePackagingConfigurationOutput, error) {
	if params == nil {
		params = &DescribePackagingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePackagingConfiguration", params, optFns, c.addOperationDescribePackagingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePackagingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePackagingConfigurationInput struct {

	// The ID of a MediaPackage VOD PackagingConfiguration resource.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DescribePackagingConfigurationOutput struct {

	// The ARN of the PackagingConfiguration.
	Arn *string

	// A CMAF packaging configuration.
	CmafPackage *types.CmafPackage

	// The time the PackagingConfiguration was created.
	CreatedAt *string

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage

	// The ID of the PackagingConfiguration.
	Id *string

	// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
	MssPackage *types.MssPackage

	// The ID of a PackagingGroup.
	PackagingGroupId *string

	// A collection of tags associated with a resource
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePackagingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackagingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackagingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePackagingConfiguration"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpDescribePackagingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackagingConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opDescribePackagingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePackagingConfiguration",
	}
}
