// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the domain's endpoint options, specifically whether all requests to the
// domain must arrive over HTTPS. For more information, see [Configuring Domain Endpoint Options]in the Amazon
// CloudSearch Developer Guide.
//
// [Configuring Domain Endpoint Options]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-domain-endpoint-options.html
func (c *Client) DescribeDomainEndpointOptions(ctx context.Context, params *DescribeDomainEndpointOptionsInput, optFns ...func(*Options)) (*DescribeDomainEndpointOptionsOutput, error) {
	if params == nil {
		params = &DescribeDomainEndpointOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomainEndpointOptions", params, optFns, c.addOperationDescribeDomainEndpointOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainEndpointOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeDomainEndpointOptions operation. Specify the name of the domain
// you want to describe. To show the active configuration and exclude any pending
// changes, set the Deployed option to true .
type DescribeDomainEndpointOptionsInput struct {

	// A string that represents the name of a domain.
	//
	// This member is required.
	DomainName *string

	// Whether to retrieve the latest configuration (which might be in a Processing
	// state) or the current, active configuration. Defaults to false .
	Deployed *bool

	noSmithyDocumentSerde
}

// The result of a DescribeDomainEndpointOptions request. Contains the status and
// configuration of a search domain's endpoint options.
type DescribeDomainEndpointOptionsOutput struct {

	// The status and configuration of a search domain's endpoint options.
	DomainEndpointOptions *types.DomainEndpointOptionsStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDomainEndpointOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDomainEndpointOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDomainEndpointOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDomainEndpointOptions"); err != nil {
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
	if err = addOpDescribeDomainEndpointOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomainEndpointOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomainEndpointOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDomainEndpointOptions",
	}
}
