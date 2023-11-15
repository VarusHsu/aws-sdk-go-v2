// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the specified provisioned product.
func (c *Client) DescribeProvisionedProduct(ctx context.Context, params *DescribeProvisionedProductInput, optFns ...func(*Options)) (*DescribeProvisionedProductOutput, error) {
	if params == nil {
		params = &DescribeProvisionedProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProvisionedProduct", params, optFns, c.addOperationDescribeProvisionedProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProvisionedProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeProvisionedProductAPI input structure. AcceptLanguage - [Optional] The
// language code for localization. Id - [Optional] The provisioned product
// identifier. Name - [Optional] Another provisioned product identifier. Customers
// must provide either Id or Name.
type DescribeProvisionedProductInput struct {

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	// The provisioned product identifier. You must provide the name or ID, but not
	// both. If you do not provide a name or ID, or you provide both name and ID, an
	// InvalidParametersException will occur.
	Id *string

	// The name of the provisioned product. You must provide the name or ID, but not
	// both. If you do not provide a name or ID, or you provide both name and ID, an
	// InvalidParametersException will occur.
	Name *string

	noSmithyDocumentSerde
}

type DescribeProvisionedProductOutput struct {

	// Any CloudWatch dashboards that were created when provisioning the product.
	CloudWatchDashboards []types.CloudWatchDashboard

	// Information about the provisioned product.
	ProvisionedProductDetail *types.ProvisionedProductDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeProvisionedProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProvisionedProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProvisionedProduct{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeProvisionedProduct"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProvisionedProduct(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeProvisionedProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeProvisionedProduct",
	}
}
