// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the Device Defender audit settings for this account.
// Settings include how audit notifications are sent and which audit checks are
// enabled or disabled.
//
// Requires permission to access the [DescribeAccountAuditConfiguration] action.
//
// [DescribeAccountAuditConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) DescribeAccountAuditConfiguration(ctx context.Context, params *DescribeAccountAuditConfigurationInput, optFns ...func(*Options)) (*DescribeAccountAuditConfigurationOutput, error) {
	if params == nil {
		params = &DescribeAccountAuditConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountAuditConfiguration", params, optFns, c.addOperationDescribeAccountAuditConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountAuditConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAuditConfigurationInput struct {
	noSmithyDocumentSerde
}

type DescribeAccountAuditConfigurationOutput struct {

	// Which audit checks are enabled and disabled for this account.
	AuditCheckConfigurations map[string]types.AuditCheckConfiguration

	// Information about the targets to which audit notifications are sent for this
	// account.
	AuditNotificationTargetConfigurations map[string]types.AuditNotificationTarget

	// The ARN of the role that grants permission to IoT to access information about
	// your devices, policies, certificates, and other items as required when
	// performing an audit.
	//
	// On the first call to UpdateAccountAuditConfiguration , this parameter is
	// required.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccountAuditConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccountAuditConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccountAuditConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAccountAuditConfiguration"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAuditConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountAuditConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAccountAuditConfiguration",
	}
}
