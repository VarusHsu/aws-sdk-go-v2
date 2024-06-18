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

// Creates a provisioning template.
//
// Requires permission to access the [CreateProvisioningTemplate] action.
//
// [CreateProvisioningTemplate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) CreateProvisioningTemplate(ctx context.Context, params *CreateProvisioningTemplateInput, optFns ...func(*Options)) (*CreateProvisioningTemplateOutput, error) {
	if params == nil {
		params = &CreateProvisioningTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProvisioningTemplate", params, optFns, c.addOperationCreateProvisioningTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProvisioningTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProvisioningTemplateInput struct {

	// The role ARN for the role associated with the provisioning template. This IoT
	// role grants permission to provision a device.
	//
	// This member is required.
	ProvisioningRoleArn *string

	// The JSON formatted contents of the provisioning template.
	//
	// This member is required.
	TemplateBody *string

	// The name of the provisioning template.
	//
	// This member is required.
	TemplateName *string

	// The description of the provisioning template.
	Description *string

	// True to enable the provisioning template, otherwise false.
	Enabled *bool

	// Creates a pre-provisioning hook template. Only supports template of type
	// FLEET_PROVISIONING . For more information about provisioning template types, see [type]
	// .
	//
	// [type]: https://docs.aws.amazon.com/iot/latest/apireference/API_CreateProvisioningTemplate.html#iot-CreateProvisioningTemplate-request-type
	PreProvisioningHook *types.ProvisioningHook

	// Metadata which can be used to manage the provisioning template.
	//
	// For URI Request parameters use format: ...key1=value1&key2=value2...
	//
	// For the CLI command-line parameter use format: &&tags
	// "key1=value1&key2=value2..."
	//
	// For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags []types.Tag

	// The type you define in a provisioning template. You can create a template with
	// only one type. You can't change the template type after its creation. The
	// default value is FLEET_PROVISIONING . For more information about provisioning
	// template, see: [Provisioning template].
	//
	// [Provisioning template]: https://docs.aws.amazon.com/iot/latest/developerguide/provision-template.html
	Type types.TemplateType

	noSmithyDocumentSerde
}

type CreateProvisioningTemplateOutput struct {

	// The default version of the provisioning template.
	DefaultVersionId *int32

	// The ARN that identifies the provisioning template.
	TemplateArn *string

	// The name of the provisioning template.
	TemplateName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProvisioningTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProvisioningTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProvisioningTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateProvisioningTemplate"); err != nil {
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
	if err = addOpCreateProvisioningTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProvisioningTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProvisioningTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateProvisioningTemplate",
	}
}
