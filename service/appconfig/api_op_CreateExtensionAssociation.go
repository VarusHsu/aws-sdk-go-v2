// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// When you create an extension or configure an Amazon Web Services authored
// extension, you associate the extension with an AppConfig application,
// environment, or configuration profile. For example, you can choose to run the
// AppConfig deployment events to Amazon SNS Amazon Web Services authored extension
// and receive notifications on an Amazon SNS topic anytime a configuration
// deployment is started for a specific application. Defining which extension to
// associate with an AppConfig resource is called an extension association. An
// extension association is a specified relationship between an extension and an
// AppConfig resource, such as an application or a configuration profile. For more
// information about extensions and associations, see Working with AppConfig
// extensions (https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html)
// in the AppConfig User Guide.
func (c *Client) CreateExtensionAssociation(ctx context.Context, params *CreateExtensionAssociationInput, optFns ...func(*Options)) (*CreateExtensionAssociationOutput, error) {
	if params == nil {
		params = &CreateExtensionAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateExtensionAssociation", params, optFns, c.addOperationCreateExtensionAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateExtensionAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExtensionAssociationInput struct {

	// The name, the ID, or the Amazon Resource Name (ARN) of the extension.
	//
	// This member is required.
	ExtensionIdentifier *string

	// The ARN of an application, configuration profile, or environment.
	//
	// This member is required.
	ResourceIdentifier *string

	// The version number of the extension. If not specified, AppConfig uses the
	// maximum version of the extension.
	ExtensionVersionNumber *int32

	// The parameter names and values defined in the extensions. Extension parameters
	// marked Required must be entered for this field.
	Parameters map[string]string

	// Adds one or more tags for the specified extension association. Tags are
	// metadata that help you categorize resources in different ways, for example, by
	// purpose, owner, or environment. Each tag consists of a key and an optional
	// value, both of which you define.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateExtensionAssociationOutput struct {

	// The system-generated Amazon Resource Name (ARN) for the extension.
	Arn *string

	// The ARN of the extension defined in the association.
	ExtensionArn *string

	// The version number for the extension defined in the association.
	ExtensionVersionNumber int32

	// The system-generated ID for the association.
	Id *string

	// The parameter names and values defined in the association.
	Parameters map[string]string

	// The ARNs of applications, configuration profiles, or environments defined in
	// the association.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateExtensionAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateExtensionAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateExtensionAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateExtensionAssociation"); err != nil {
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
	if err = addOpCreateExtensionAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExtensionAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateExtensionAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateExtensionAssociation",
	}
}
