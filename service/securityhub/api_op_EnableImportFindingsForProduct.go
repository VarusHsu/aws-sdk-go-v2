// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the integration of a partner product with Security Hub. Integrated
// products send findings to Security Hub.
//
// When you enable a product integration, a permissions policy that grants
// permission for the product to send findings to Security Hub is applied.
func (c *Client) EnableImportFindingsForProduct(ctx context.Context, params *EnableImportFindingsForProductInput, optFns ...func(*Options)) (*EnableImportFindingsForProductOutput, error) {
	if params == nil {
		params = &EnableImportFindingsForProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableImportFindingsForProduct", params, optFns, c.addOperationEnableImportFindingsForProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableImportFindingsForProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableImportFindingsForProductInput struct {

	// The ARN of the product to enable the integration for.
	//
	// This member is required.
	ProductArn *string

	noSmithyDocumentSerde
}

type EnableImportFindingsForProductOutput struct {

	// The ARN of your subscription to the product to enable integrations for.
	ProductSubscriptionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableImportFindingsForProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEnableImportFindingsForProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableImportFindingsForProduct{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EnableImportFindingsForProduct"); err != nil {
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
	if err = addOpEnableImportFindingsForProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableImportFindingsForProduct(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableImportFindingsForProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EnableImportFindingsForProduct",
	}
}
