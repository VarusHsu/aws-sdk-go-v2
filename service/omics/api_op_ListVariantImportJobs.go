// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of variant import jobs.
func (c *Client) ListVariantImportJobs(ctx context.Context, params *ListVariantImportJobsInput, optFns ...func(*Options)) (*ListVariantImportJobsOutput, error) {
	if params == nil {
		params = &ListVariantImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVariantImportJobs", params, optFns, c.addOperationListVariantImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVariantImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVariantImportJobsInput struct {

	// A filter to apply to the list.
	Filter *types.ListVariantImportJobsFilter

	// A list of job IDs.
	Ids []string

	// The maximum number of import jobs to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVariantImportJobsOutput struct {

	// A pagination token that's included if more results are available.
	NextToken *string

	// A list of jobs.
	VariantImportJobs []types.VariantImportJobItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVariantImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVariantImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVariantImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVariantImportJobs"); err != nil {
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
	if err = addEndpointPrefix_opListVariantImportJobsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVariantImportJobs(options.Region), middleware.Before); err != nil {
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

// ListVariantImportJobsPaginatorOptions is the paginator options for
// ListVariantImportJobs
type ListVariantImportJobsPaginatorOptions struct {
	// The maximum number of import jobs to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVariantImportJobsPaginator is a paginator for ListVariantImportJobs
type ListVariantImportJobsPaginator struct {
	options   ListVariantImportJobsPaginatorOptions
	client    ListVariantImportJobsAPIClient
	params    *ListVariantImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListVariantImportJobsPaginator returns a new ListVariantImportJobsPaginator
func NewListVariantImportJobsPaginator(client ListVariantImportJobsAPIClient, params *ListVariantImportJobsInput, optFns ...func(*ListVariantImportJobsPaginatorOptions)) *ListVariantImportJobsPaginator {
	if params == nil {
		params = &ListVariantImportJobsInput{}
	}

	options := ListVariantImportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVariantImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVariantImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVariantImportJobs page.
func (p *ListVariantImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVariantImportJobsOutput, error) {
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

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListVariantImportJobs(ctx, &params, optFns...)
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

type endpointPrefix_opListVariantImportJobsMiddleware struct {
}

func (*endpointPrefix_opListVariantImportJobsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListVariantImportJobsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListVariantImportJobsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListVariantImportJobsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListVariantImportJobsAPIClient is a client that implements the
// ListVariantImportJobs operation.
type ListVariantImportJobsAPIClient interface {
	ListVariantImportJobs(context.Context, *ListVariantImportJobsInput, ...func(*Options)) (*ListVariantImportJobsOutput, error)
}

var _ ListVariantImportJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListVariantImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVariantImportJobs",
	}
}
