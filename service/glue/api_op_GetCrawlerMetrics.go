// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves metrics about specified crawlers.
func (c *Client) GetCrawlerMetrics(ctx context.Context, params *GetCrawlerMetricsInput, optFns ...func(*Options)) (*GetCrawlerMetricsOutput, error) {
	if params == nil {
		params = &GetCrawlerMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCrawlerMetrics", params, optFns, c.addOperationGetCrawlerMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCrawlerMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCrawlerMetricsInput struct {

	// A list of the names of crawlers about which to retrieve metrics.
	CrawlerNameList []string

	// The maximum size of a list to return.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCrawlerMetricsOutput struct {

	// A list of metrics for the specified crawler.
	CrawlerMetricsList []types.CrawlerMetrics

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCrawlerMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCrawlerMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCrawlerMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCrawlerMetrics"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCrawlerMetrics(options.Region), middleware.Before); err != nil {
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

// GetCrawlerMetricsAPIClient is a client that implements the GetCrawlerMetrics
// operation.
type GetCrawlerMetricsAPIClient interface {
	GetCrawlerMetrics(context.Context, *GetCrawlerMetricsInput, ...func(*Options)) (*GetCrawlerMetricsOutput, error)
}

var _ GetCrawlerMetricsAPIClient = (*Client)(nil)

// GetCrawlerMetricsPaginatorOptions is the paginator options for GetCrawlerMetrics
type GetCrawlerMetricsPaginatorOptions struct {
	// The maximum size of a list to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCrawlerMetricsPaginator is a paginator for GetCrawlerMetrics
type GetCrawlerMetricsPaginator struct {
	options   GetCrawlerMetricsPaginatorOptions
	client    GetCrawlerMetricsAPIClient
	params    *GetCrawlerMetricsInput
	nextToken *string
	firstPage bool
}

// NewGetCrawlerMetricsPaginator returns a new GetCrawlerMetricsPaginator
func NewGetCrawlerMetricsPaginator(client GetCrawlerMetricsAPIClient, params *GetCrawlerMetricsInput, optFns ...func(*GetCrawlerMetricsPaginatorOptions)) *GetCrawlerMetricsPaginator {
	if params == nil {
		params = &GetCrawlerMetricsInput{}
	}

	options := GetCrawlerMetricsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCrawlerMetricsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCrawlerMetricsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCrawlerMetrics page.
func (p *GetCrawlerMetricsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCrawlerMetricsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetCrawlerMetrics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCrawlerMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCrawlerMetrics",
	}
}
