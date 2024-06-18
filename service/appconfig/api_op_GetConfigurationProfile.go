// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about a configuration profile.
func (c *Client) GetConfigurationProfile(ctx context.Context, params *GetConfigurationProfileInput, optFns ...func(*Options)) (*GetConfigurationProfileOutput, error) {
	if params == nil {
		params = &GetConfigurationProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfigurationProfile", params, optFns, c.addOperationGetConfigurationProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfigurationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfigurationProfileInput struct {

	// The ID of the application that includes the configuration profile you want to
	// get.
	//
	// This member is required.
	ApplicationId *string

	// The ID of the configuration profile that you want to get.
	//
	// This member is required.
	ConfigurationProfileId *string

	noSmithyDocumentSerde
}

type GetConfigurationProfileOutput struct {

	// The application ID.
	ApplicationId *string

	// The configuration profile description.
	Description *string

	// The configuration profile ID.
	Id *string

	// The Amazon Resource Name of the Key Management Service key to encrypt new
	// configuration data versions in the AppConfig hosted configuration store. This
	// attribute is only used for hosted configuration types. To encrypt data managed
	// in other configuration stores, see the documentation for how to specify an KMS
	// key for that particular service.
	KmsKeyArn *string

	// The Key Management Service key identifier (key ID, key alias, or key ARN)
	// provided when the resource was created or updated.
	KmsKeyIdentifier *string

	// The URI location of the configuration.
	LocationUri *string

	// The name of the configuration profile.
	Name *string

	// The ARN of an IAM role with permission to access the configuration at the
	// specified LocationUri .
	RetrievalRoleArn *string

	// The type of configurations contained in the profile. AppConfig supports feature
	// flags and freeform configurations. We recommend you create feature flag
	// configurations to enable or disable new features and freeform configurations to
	// distribute configurations to an application. When calling this API, enter one of
	// the following values for Type :
	//
	//     AWS.AppConfig.FeatureFlags
	//
	//     AWS.Freeform
	Type *string

	// A list of methods for validating the configuration.
	Validators []types.Validator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConfigurationProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfigurationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetConfigurationProfile"); err != nil {
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
	if err = addOpGetConfigurationProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfigurationProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConfigurationProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetConfigurationProfile",
	}
}
