// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Launch Configuration Templates, filtered by Launch Configuration
// Template IDs
func (c *Client) DescribeLaunchConfigurationTemplates(ctx context.Context, params *DescribeLaunchConfigurationTemplatesInput, optFns ...func(*Options)) (*DescribeLaunchConfigurationTemplatesOutput, error) {
	if params == nil {
		params = &DescribeLaunchConfigurationTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLaunchConfigurationTemplates", params, optFns, c.addOperationDescribeLaunchConfigurationTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLaunchConfigurationTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLaunchConfigurationTemplatesInput struct {

	// Request to filter Launch Configuration Templates list by Launch Configuration
	// Template ID.
	LaunchConfigurationTemplateIDs []string

	// Maximum results to be returned in DescribeLaunchConfigurationTemplates.
	MaxResults int32

	// Next pagination token returned from DescribeLaunchConfigurationTemplates.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeLaunchConfigurationTemplatesOutput struct {

	// List of items returned by DescribeLaunchConfigurationTemplates.
	Items []types.LaunchConfigurationTemplate

	// Next pagination token returned from DescribeLaunchConfigurationTemplates.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLaunchConfigurationTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeLaunchConfigurationTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeLaunchConfigurationTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLaunchConfigurationTemplates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLaunchConfigurationTemplates(options.Region), middleware.Before); err != nil {
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

// DescribeLaunchConfigurationTemplatesAPIClient is a client that implements the
// DescribeLaunchConfigurationTemplates operation.
type DescribeLaunchConfigurationTemplatesAPIClient interface {
	DescribeLaunchConfigurationTemplates(context.Context, *DescribeLaunchConfigurationTemplatesInput, ...func(*Options)) (*DescribeLaunchConfigurationTemplatesOutput, error)
}

var _ DescribeLaunchConfigurationTemplatesAPIClient = (*Client)(nil)

// DescribeLaunchConfigurationTemplatesPaginatorOptions is the paginator options
// for DescribeLaunchConfigurationTemplates
type DescribeLaunchConfigurationTemplatesPaginatorOptions struct {
	// Maximum results to be returned in DescribeLaunchConfigurationTemplates.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLaunchConfigurationTemplatesPaginator is a paginator for
// DescribeLaunchConfigurationTemplates
type DescribeLaunchConfigurationTemplatesPaginator struct {
	options   DescribeLaunchConfigurationTemplatesPaginatorOptions
	client    DescribeLaunchConfigurationTemplatesAPIClient
	params    *DescribeLaunchConfigurationTemplatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeLaunchConfigurationTemplatesPaginator returns a new
// DescribeLaunchConfigurationTemplatesPaginator
func NewDescribeLaunchConfigurationTemplatesPaginator(client DescribeLaunchConfigurationTemplatesAPIClient, params *DescribeLaunchConfigurationTemplatesInput, optFns ...func(*DescribeLaunchConfigurationTemplatesPaginatorOptions)) *DescribeLaunchConfigurationTemplatesPaginator {
	if params == nil {
		params = &DescribeLaunchConfigurationTemplatesInput{}
	}

	options := DescribeLaunchConfigurationTemplatesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLaunchConfigurationTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLaunchConfigurationTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLaunchConfigurationTemplates page.
func (p *DescribeLaunchConfigurationTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLaunchConfigurationTemplatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeLaunchConfigurationTemplates(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeLaunchConfigurationTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLaunchConfigurationTemplates",
	}
}
