// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use the SearchDataSources operation to search for data sources that belong to
// an account.
func (c *Client) SearchDataSources(ctx context.Context, params *SearchDataSourcesInput, optFns ...func(*Options)) (*SearchDataSourcesOutput, error) {
	if params == nil {
		params = &SearchDataSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchDataSources", params, optFns, c.addOperationSearchDataSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchDataSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchDataSourcesInput struct {

	// The Amazon Web Services account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The filters to apply to the search.
	//
	// This member is required.
	Filters []types.DataSourceSearchFilter

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchDataSourcesOutput struct {

	// A DataSourceSummaries object that returns a summary of a data source.
	DataSourceSummaries []types.DataSourceSummary

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchDataSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchDataSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchDataSources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchDataSources"); err != nil {
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
	if err = addOpSearchDataSourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchDataSources(options.Region), middleware.Before); err != nil {
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

// SearchDataSourcesPaginatorOptions is the paginator options for SearchDataSources
type SearchDataSourcesPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchDataSourcesPaginator is a paginator for SearchDataSources
type SearchDataSourcesPaginator struct {
	options   SearchDataSourcesPaginatorOptions
	client    SearchDataSourcesAPIClient
	params    *SearchDataSourcesInput
	nextToken *string
	firstPage bool
}

// NewSearchDataSourcesPaginator returns a new SearchDataSourcesPaginator
func NewSearchDataSourcesPaginator(client SearchDataSourcesAPIClient, params *SearchDataSourcesInput, optFns ...func(*SearchDataSourcesPaginatorOptions)) *SearchDataSourcesPaginator {
	if params == nil {
		params = &SearchDataSourcesInput{}
	}

	options := SearchDataSourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchDataSourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchDataSourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchDataSources page.
func (p *SearchDataSourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchDataSourcesOutput, error) {
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
	result, err := p.client.SearchDataSources(ctx, &params, optFns...)
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

// SearchDataSourcesAPIClient is a client that implements the SearchDataSources
// operation.
type SearchDataSourcesAPIClient interface {
	SearchDataSources(context.Context, *SearchDataSourcesInput, ...func(*Options)) (*SearchDataSourcesOutput, error)
}

var _ SearchDataSourcesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opSearchDataSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchDataSources",
	}
}
