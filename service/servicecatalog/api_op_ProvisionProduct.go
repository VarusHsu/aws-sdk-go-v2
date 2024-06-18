// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Provisions the specified product.
//
// A provisioned product is a resourced instance of a product. For example,
// provisioning a product that's based on an CloudFormation template launches an
// CloudFormation stack and its underlying resources. You can check the status of
// this request using DescribeRecord.
//
// If the request contains a tag key with an empty list of values, there's a tag
// conflict for that key. Don't include conflicted keys as tags, or this will cause
// the error "Parameter validation failed: Missing required parameter in
// Tags[N]:Value".
//
// When provisioning a product that's been added to a portfolio, you must grant
// your user, group, or role access to the portfolio. For more information, see [Granting users access]in
// the Service Catalog User Guide.
//
// [Granting users access]: https://docs.aws.amazon.com/servicecatalog/latest/adminguide/catalogs_portfolios_users.html
func (c *Client) ProvisionProduct(ctx context.Context, params *ProvisionProductInput, optFns ...func(*Options)) (*ProvisionProductOutput, error) {
	if params == nil {
		params = &ProvisionProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ProvisionProduct", params, optFns, c.addOperationProvisionProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ProvisionProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ProvisionProductInput struct {

	// An idempotency token that uniquely identifies the provisioning request.
	//
	// This member is required.
	ProvisionToken *string

	// A user-friendly name for the provisioned product. This value must be unique for
	// the Amazon Web Services account and cannot be updated after the product is
	// provisioned.
	//
	// This member is required.
	ProvisionedProductName *string

	// The language code.
	//
	//   - jp - Japanese
	//
	//   - zh - Chinese
	AcceptLanguage *string

	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related
	// events.
	NotificationArns []string

	// The path identifier of the product. This value is optional if the product has a
	// default path, and required if the product has more than one path. To list the
	// paths for a product, use ListLaunchPaths. You must provide the name or ID, but not both.
	PathId *string

	// The name of the path. You must provide the name or ID, but not both.
	PathName *string

	// The product identifier. You must provide the name or ID, but not both.
	ProductId *string

	// The name of the product. You must provide the name or ID, but not both.
	ProductName *string

	// The identifier of the provisioning artifact. You must provide the name or ID,
	// but not both.
	ProvisioningArtifactId *string

	// The name of the provisioning artifact. You must provide the name or ID, but not
	// both.
	ProvisioningArtifactName *string

	// Parameters specified by the administrator that are required for provisioning
	// the product.
	ProvisioningParameters []types.ProvisioningParameter

	// An object that contains information about the provisioning preferences for a
	// stack set.
	ProvisioningPreferences *types.ProvisioningPreferences

	// One or more tags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ProvisionProductOutput struct {

	// Information about the result of provisioning the product.
	RecordDetail *types.RecordDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationProvisionProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpProvisionProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpProvisionProduct{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ProvisionProduct"); err != nil {
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
	if err = addIdempotencyToken_opProvisionProductMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpProvisionProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opProvisionProduct(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpProvisionProduct struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpProvisionProduct) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpProvisionProduct) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ProvisionProductInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ProvisionProductInput ")
	}

	if input.ProvisionToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ProvisionToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opProvisionProductMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpProvisionProduct{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opProvisionProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ProvisionProduct",
	}
}
