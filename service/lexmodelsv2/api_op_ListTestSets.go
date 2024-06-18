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

// The list of the test sets
func (c *Client) ListTestSets(ctx context.Context, params *ListTestSetsInput, optFns ...func(*Options)) (*ListTestSetsOutput, error) {
	if params == nil {
		params = &ListTestSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTestSets", params, optFns, c.addOperationListTestSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTestSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestSetsInput struct {

	// The maximum number of test sets to return in each page. If there are fewer
	// results than the max page size, only the actual number of results are returned.
	MaxResults *int32

	// If the response from the ListTestSets operation contains more results than
	// specified in the maxResults parameter, a token is returned in the response. Use
	// that token in the nextToken parameter to return the next page of results.
	NextToken *string

	// The sort order for the list of test sets.
	SortBy *types.TestSetSortBy

	noSmithyDocumentSerde
}

type ListTestSetsOutput struct {

	// A token that indicates whether there are more results to return in a response
	// to the ListTestSets operation. If the nextToken field is present, you send the
	// contents as the nextToken parameter of a ListTestSets operation request to get
	// the next page of results.
	NextToken *string

	// The selected test sets in a list of test sets.
	TestSets []types.TestSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTestSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTestSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTestSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTestSets"); err != nil {
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
	if err = addOpListTestSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTestSets(options.Region), middleware.Before); err != nil {
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

// ListTestSetsPaginatorOptions is the paginator options for ListTestSets
type ListTestSetsPaginatorOptions struct {
	// The maximum number of test sets to return in each page. If there are fewer
	// results than the max page size, only the actual number of results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTestSetsPaginator is a paginator for ListTestSets
type ListTestSetsPaginator struct {
	options   ListTestSetsPaginatorOptions
	client    ListTestSetsAPIClient
	params    *ListTestSetsInput
	nextToken *string
	firstPage bool
}

// NewListTestSetsPaginator returns a new ListTestSetsPaginator
func NewListTestSetsPaginator(client ListTestSetsAPIClient, params *ListTestSetsInput, optFns ...func(*ListTestSetsPaginatorOptions)) *ListTestSetsPaginator {
	if params == nil {
		params = &ListTestSetsInput{}
	}

	options := ListTestSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTestSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTestSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTestSets page.
func (p *ListTestSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTestSetsOutput, error) {
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
	result, err := p.client.ListTestSets(ctx, &params, optFns...)
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

// ListTestSetsAPIClient is a client that implements the ListTestSets operation.
type ListTestSetsAPIClient interface {
	ListTestSets(context.Context, *ListTestSetsInput, ...func(*Options)) (*ListTestSetsOutput, error)
}

var _ ListTestSetsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTestSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTestSets",
	}
}
