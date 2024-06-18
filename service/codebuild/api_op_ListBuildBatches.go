// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the identifiers of your build batches in the current region.
func (c *Client) ListBuildBatches(ctx context.Context, params *ListBuildBatchesInput, optFns ...func(*Options)) (*ListBuildBatchesOutput, error) {
	if params == nil {
		params = &ListBuildBatchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBuildBatches", params, optFns, c.addOperationListBuildBatchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBuildBatchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBuildBatchesInput struct {

	// A BuildBatchFilter object that specifies the filters for the search.
	Filter *types.BuildBatchFilter

	// The maximum number of results to return.
	MaxResults *int32

	// The nextToken value returned from a previous call to ListBuildBatches . This
	// specifies the next item to return. To return the beginning of the list, exclude
	// this parameter.
	NextToken *string

	// Specifies the sort order of the returned items. Valid values include:
	//
	//   - ASCENDING : List the batch build identifiers in ascending order by
	//   identifier.
	//
	//   - DESCENDING : List the batch build identifiers in descending order by
	//   identifier.
	SortOrder types.SortOrderType

	noSmithyDocumentSerde
}

type ListBuildBatchesOutput struct {

	// An array of strings that contains the batch build identifiers.
	Ids []string

	// If there are more items to return, this contains a token that is passed to a
	// subsequent call to ListBuildBatches to retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBuildBatchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBuildBatches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBuildBatches{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBuildBatches"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBuildBatches(options.Region), middleware.Before); err != nil {
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

// ListBuildBatchesPaginatorOptions is the paginator options for ListBuildBatches
type ListBuildBatchesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBuildBatchesPaginator is a paginator for ListBuildBatches
type ListBuildBatchesPaginator struct {
	options   ListBuildBatchesPaginatorOptions
	client    ListBuildBatchesAPIClient
	params    *ListBuildBatchesInput
	nextToken *string
	firstPage bool
}

// NewListBuildBatchesPaginator returns a new ListBuildBatchesPaginator
func NewListBuildBatchesPaginator(client ListBuildBatchesAPIClient, params *ListBuildBatchesInput, optFns ...func(*ListBuildBatchesPaginatorOptions)) *ListBuildBatchesPaginator {
	if params == nil {
		params = &ListBuildBatchesInput{}
	}

	options := ListBuildBatchesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBuildBatchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBuildBatchesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBuildBatches page.
func (p *ListBuildBatchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBuildBatchesOutput, error) {
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
	result, err := p.client.ListBuildBatches(ctx, &params, optFns...)
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

// ListBuildBatchesAPIClient is a client that implements the ListBuildBatches
// operation.
type ListBuildBatchesAPIClient interface {
	ListBuildBatches(context.Context, *ListBuildBatchesInput, ...func(*Options)) (*ListBuildBatchesOutput, error)
}

var _ ListBuildBatchesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListBuildBatches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBuildBatches",
	}
}
