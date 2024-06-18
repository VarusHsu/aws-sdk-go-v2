// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the specified configuration.
func (c *Client) DescribeConfiguration(ctx context.Context, params *DescribeConfigurationInput, optFns ...func(*Options)) (*DescribeConfigurationOutput, error) {
	if params == nil {
		params = &DescribeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfiguration", params, optFns, c.addOperationDescribeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationInput struct {

	// The unique ID that Amazon MQ generates for the configuration.
	//
	// This member is required.
	ConfigurationId *string

	noSmithyDocumentSerde
}

type DescribeConfigurationOutput struct {

	// Required. The ARN of the configuration.
	Arn *string

	// Optional. The authentication strategy associated with the configuration. The
	// default is SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// Required. The date and time of the configuration revision.
	Created *time.Time

	// Required. The description of the configuration.
	Description *string

	// Required. The type of broker engine. Currently, Amazon MQ supports ACTIVEMQ and
	// RABBITMQ.
	EngineType types.EngineType

	// Required. The broker engine's version. For a list of supported engine versions,
	// see, [Supported engines].
	//
	// [Supported engines]: https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker-engine.html
	EngineVersion *string

	// Required. The unique ID that Amazon MQ generates for the configuration.
	Id *string

	// Required. The latest revision of the configuration.
	LatestRevision *types.ConfigurationRevision

	// Required. The name of the configuration. This value can contain only
	// alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~).
	// This value must be 1-150 characters long.
	Name *string

	// The list of all tags associated with this configuration.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeConfiguration"); err != nil {
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
	if err = addOpDescribeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeConfiguration",
	}
}
