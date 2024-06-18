// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the quota increase request for the specified quota from your quota
// request template.
func (c *Client) DeleteServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, params *DeleteServiceQuotaIncreaseRequestFromTemplateInput, optFns ...func(*Options)) (*DeleteServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	if params == nil {
		params = &DeleteServiceQuotaIncreaseRequestFromTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteServiceQuotaIncreaseRequestFromTemplate", params, optFns, c.addOperationDeleteServiceQuotaIncreaseRequestFromTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteServiceQuotaIncreaseRequestFromTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteServiceQuotaIncreaseRequestFromTemplateInput struct {

	// Specifies the Amazon Web Services Region for which the request was made.
	//
	// This member is required.
	AwsRegion *string

	// Specifies the quota identifier. To find the quota code for a specific quota,
	// use the ListServiceQuotasoperation, and look for the QuotaCode response in the output for the
	// quota you want.
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServicesoperation.
	//
	// This member is required.
	ServiceCode *string

	noSmithyDocumentSerde
}

type DeleteServiceQuotaIncreaseRequestFromTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteServiceQuotaIncreaseRequestFromTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteServiceQuotaIncreaseRequestFromTemplate"); err != nil {
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
	if err = addOpDeleteServiceQuotaIncreaseRequestFromTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteServiceQuotaIncreaseRequestFromTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteServiceQuotaIncreaseRequestFromTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteServiceQuotaIncreaseRequestFromTemplate",
	}
}
