// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation configures Amazon Route 53 to automatically renew the specified
// domain before the domain registration expires. The cost of renewing your domain
// registration is billed to your Amazon Web Services account.
//
// The period during which you can renew a domain name varies by TLD. For a list
// of TLDs and their renewal policies, see [Domains That You Can Register with Amazon Route 53]in the Amazon Route 53 Developer Guide.
// Route 53 requires that you renew before the end of the renewal period so we can
// complete processing before the deadline.
//
// [Domains That You Can Register with Amazon Route 53]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html
func (c *Client) EnableDomainAutoRenew(ctx context.Context, params *EnableDomainAutoRenewInput, optFns ...func(*Options)) (*EnableDomainAutoRenewOutput, error) {
	if params == nil {
		params = &EnableDomainAutoRenewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableDomainAutoRenew", params, optFns, c.addOperationEnableDomainAutoRenewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableDomainAutoRenewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableDomainAutoRenewInput struct {

	// The name of the domain that you want to enable automatic renewal for.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

type EnableDomainAutoRenewOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableDomainAutoRenewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableDomainAutoRenew{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableDomainAutoRenew{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EnableDomainAutoRenew"); err != nil {
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
	if err = addOpEnableDomainAutoRenewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableDomainAutoRenew(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableDomainAutoRenew(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EnableDomainAutoRenew",
	}
}
