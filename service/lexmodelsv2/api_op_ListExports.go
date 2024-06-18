// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the exports for a bot, bot locale, or custom vocabulary. Exports are kept
// in the list for 7 days.
func (c *Client) ListExports(ctx context.Context, params *ListExportsInput, optFns ...func(*Options)) (*ListExportsOutput, error) {
	if params == nil {
		params = &ListExportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExports", params, optFns, c.addOperationListExportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExportsInput struct {

	// The unique identifier that Amazon Lex assigned to the bot.
	BotId *string

	// The version of the bot to list exports for.
	BotVersion *string

	// Provides the specification of a filter used to limit the exports in the
	// response to only those that match the filter specification. You can only specify
	// one filter and one string to filter on.
	Filters []types.ExportFilter

	// Specifies the resources that should be exported. If you don't specify a
	// resource type in the filters parameter, both bot locales and custom
	// vocabularies are exported.
	LocaleId *string

	// The maximum number of exports to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListExports operation contains more results that
	// specified in the maxResults parameter, a token is returned in the response.
	//
	// Use the returned token in the nextToken parameter of a ListExports request to
	// return the next page of results. For a complete set of results, call the
	// ListExports operation until the nextToken returned in the response is null.
	NextToken *string

	// Determines the field that the list of exports is sorted by. You can sort by the
	// LastUpdatedDateTime field in ascending or descending order.
	SortBy *types.ExportSortBy

	noSmithyDocumentSerde
}

type ListExportsOutput struct {

	// The unique identifier assigned to the bot by Amazon Lex.
	BotId *string

	// The version of the bot that was exported.
	BotVersion *string

	// Summary information for the exports that meet the filter criteria specified in
	// the request. The length of the list is specified in the maxResults parameter.
	// If there are more exports available, the nextToken field contains a token to
	// get the next page of results.
	ExportSummaries []types.ExportSummary

	// The locale specified in the request.
	LocaleId *string

	// A token that indicates whether there are more results to return in a response
	// to the ListExports operation. If the nextToken field is present, you send the
	// contents as the nextToken parameter of a ListExports operation request to get
	// the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListExports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListExports{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListExports"); err != nil {
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
	if err = addOpListExportsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExports(options.Region), middleware.Before); err != nil {
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

// ListExportsPaginatorOptions is the paginator options for ListExports
type ListExportsPaginatorOptions struct {
	// The maximum number of exports to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExportsPaginator is a paginator for ListExports
type ListExportsPaginator struct {
	options   ListExportsPaginatorOptions
	client    ListExportsAPIClient
	params    *ListExportsInput
	nextToken *string
	firstPage bool
}

// NewListExportsPaginator returns a new ListExportsPaginator
func NewListExportsPaginator(client ListExportsAPIClient, params *ListExportsInput, optFns ...func(*ListExportsPaginatorOptions)) *ListExportsPaginator {
	if params == nil {
		params = &ListExportsInput{}
	}

	options := ListExportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExportsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExports page.
func (p *ListExportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExportsOutput, error) {
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
	result, err := p.client.ListExports(ctx, &params, optFns...)
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

// ListExportsAPIClient is a client that implements the ListExports operation.
type ListExportsAPIClient interface {
	ListExports(context.Context, *ListExportsInput, ...func(*Options)) (*ListExportsOutput, error)
}

var _ ListExportsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListExports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListExports",
	}
}
