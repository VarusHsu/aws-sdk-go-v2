// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of Call Analytics categories, including all rules that make up
// each category. To get detailed information about a specific Call Analytics
// category, use the operation.
func (c *Client) ListCallAnalyticsCategories(ctx context.Context, params *ListCallAnalyticsCategoriesInput, optFns ...func(*Options)) (*ListCallAnalyticsCategoriesOutput, error) {
	if params == nil {
		params = &ListCallAnalyticsCategoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCallAnalyticsCategories", params, optFns, c.addOperationListCallAnalyticsCategoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCallAnalyticsCategoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCallAnalyticsCategoriesInput struct {

	// The maximum number of Call Analytics categories to return in each page of
	// results. If there are fewer results than the value that you specify, only the
	// actual results are returned. If you don't specify a value, a default of 5 is
	// used.
	MaxResults *int32

	// If your ListCallAnalyticsCategories request returns more results than can be
	// displayed, NextToken is displayed in the response with an associated string. To
	// get the next page of results, copy this string and repeat your request,
	// including NextToken with the value of the copied string. Repeat as needed to
	// view all your results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCallAnalyticsCategoriesOutput struct {

	// Provides detailed information about your Call Analytics categories, including
	// all the rules associated with each category.
	Categories []types.CategoryProperties

	// If NextToken is present in your response, it indicates that not all results are
	// displayed. To view the next set of results, copy the string associated with the
	// NextToken parameter in your results output, then run your request again
	// including NextToken with the value of the copied string. Repeat as needed to
	// view all your results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCallAnalyticsCategoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCallAnalyticsCategories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCallAnalyticsCategories{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCallAnalyticsCategories"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCallAnalyticsCategories(options.Region), middleware.Before); err != nil {
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

// ListCallAnalyticsCategoriesAPIClient is a client that implements the
// ListCallAnalyticsCategories operation.
type ListCallAnalyticsCategoriesAPIClient interface {
	ListCallAnalyticsCategories(context.Context, *ListCallAnalyticsCategoriesInput, ...func(*Options)) (*ListCallAnalyticsCategoriesOutput, error)
}

var _ ListCallAnalyticsCategoriesAPIClient = (*Client)(nil)

// ListCallAnalyticsCategoriesPaginatorOptions is the paginator options for
// ListCallAnalyticsCategories
type ListCallAnalyticsCategoriesPaginatorOptions struct {
	// The maximum number of Call Analytics categories to return in each page of
	// results. If there are fewer results than the value that you specify, only the
	// actual results are returned. If you don't specify a value, a default of 5 is
	// used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCallAnalyticsCategoriesPaginator is a paginator for
// ListCallAnalyticsCategories
type ListCallAnalyticsCategoriesPaginator struct {
	options   ListCallAnalyticsCategoriesPaginatorOptions
	client    ListCallAnalyticsCategoriesAPIClient
	params    *ListCallAnalyticsCategoriesInput
	nextToken *string
	firstPage bool
}

// NewListCallAnalyticsCategoriesPaginator returns a new
// ListCallAnalyticsCategoriesPaginator
func NewListCallAnalyticsCategoriesPaginator(client ListCallAnalyticsCategoriesAPIClient, params *ListCallAnalyticsCategoriesInput, optFns ...func(*ListCallAnalyticsCategoriesPaginatorOptions)) *ListCallAnalyticsCategoriesPaginator {
	if params == nil {
		params = &ListCallAnalyticsCategoriesInput{}
	}

	options := ListCallAnalyticsCategoriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCallAnalyticsCategoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCallAnalyticsCategoriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCallAnalyticsCategories page.
func (p *ListCallAnalyticsCategoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCallAnalyticsCategoriesOutput, error) {
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

	result, err := p.client.ListCallAnalyticsCategories(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCallAnalyticsCategories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCallAnalyticsCategories",
	}
}
