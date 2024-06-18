// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an Proton component. A component is an infrastructure extension for a
// service instance.
//
// For more information about components, see [Proton components] in the Proton User Guide.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func (c *Client) CreateComponent(ctx context.Context, params *CreateComponentInput, optFns ...func(*Options)) (*CreateComponentOutput, error) {
	if params == nil {
		params = &CreateComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComponent", params, optFns, c.addOperationCreateComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateComponentInput struct {

	// A path to a manifest file that lists the Infrastructure as Code (IaC) file,
	// template language, and rendering engine for infrastructure that a custom
	// component provisions.
	//
	// This value conforms to the media type: application/yaml
	//
	// This member is required.
	Manifest *string

	// The customer-provided name of the component.
	//
	// This member is required.
	Name *string

	// A path to the Infrastructure as Code (IaC) file describing infrastructure that
	// a custom component provisions.
	//
	// Components support a single IaC file, even if you use Terraform as your
	// template language.
	//
	// This value conforms to the media type: application/yaml
	//
	// This member is required.
	TemplateFile *string

	// The client token for the created component.
	ClientToken *string

	// An optional customer-provided description of the component.
	Description *string

	// The name of the Proton environment that you want to associate this component
	// with. You must specify this when you don't specify serviceInstanceName and
	// serviceName .
	EnvironmentName *string

	// The name of the service instance that you want to attach this component to. If
	// you don't specify this, the component isn't attached to any service instance.
	// Specify both serviceInstanceName and serviceName or neither of them.
	ServiceInstanceName *string

	// The name of the service that serviceInstanceName is associated with. If you
	// don't specify this, the component isn't attached to any service instance.
	// Specify both serviceInstanceName and serviceName or neither of them.
	ServiceName *string

	// The service spec that you want the component to use to access service inputs.
	// Set this only when you attach the component to a service instance.
	//
	// This value conforms to the media type: application/yaml
	ServiceSpec *string

	// An optional list of metadata items that you can associate with the Proton
	// component. A tag is a key-value pair.
	//
	// For more information, see [Proton resources and tagging] in the Proton User Guide.
	//
	// [Proton resources and tagging]: https://docs.aws.amazon.com/proton/latest/userguide/resources.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateComponentOutput struct {

	// The detailed data of the created component.
	//
	// This member is required.
	Component *types.Component

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateComponent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateComponent"); err != nil {
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
	if err = addIdempotencyToken_opCreateComponentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComponent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateComponent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateComponent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateComponentInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateComponentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateComponent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateComponent",
	}
}
