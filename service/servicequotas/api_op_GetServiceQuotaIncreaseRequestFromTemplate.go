// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the specified quota increase request in your quota
// request template.
func (c *Client) GetServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, params *GetServiceQuotaIncreaseRequestFromTemplateInput, optFns ...func(*Options)) (*GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	if params == nil {
		params = &GetServiceQuotaIncreaseRequestFromTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceQuotaIncreaseRequestFromTemplate", params, optFns, c.addOperationGetServiceQuotaIncreaseRequestFromTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceQuotaIncreaseRequestFromTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceQuotaIncreaseRequestFromTemplateInput struct {

	// Specifies the Amazon Web Services Region for which you made the request.
	//
	// This member is required.
	AwsRegion *string

	// Specifies the quota identifier. To find the quota code for a specific quota,
	// use the ListServiceQuotas operation, and look for the QuotaCode response in the
	// output for the quota you want.
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServices operation.
	//
	// This member is required.
	ServiceCode *string

	noSmithyDocumentSerde
}

type GetServiceQuotaIncreaseRequestFromTemplateOutput struct {

	// Information about the quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceQuotaIncreaseRequestFromTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetServiceQuotaIncreaseRequestFromTemplate"); err != nil {
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
	if err = addOpGetServiceQuotaIncreaseRequestFromTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceQuotaIncreaseRequestFromTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceQuotaIncreaseRequestFromTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetServiceQuotaIncreaseRequestFromTemplate",
	}
}
