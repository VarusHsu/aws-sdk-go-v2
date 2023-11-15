// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// An array of TransactionEvent objects. Each object contains details about the
// transaction event.
func (c *Client) ListTransactionEvents(ctx context.Context, params *ListTransactionEventsInput, optFns ...func(*Options)) (*ListTransactionEventsOutput, error) {
	if params == nil {
		params = &ListTransactionEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTransactionEvents", params, optFns, c.addOperationListTransactionEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTransactionEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTransactionEventsInput struct {

	// The blockchain network where the transaction events occurred.
	//
	// This member is required.
	Network types.QueryNetwork

	// The hash of the transaction. It is generated whenever a transaction is verified
	// and added to the blockchain.
	//
	// This member is required.
	TransactionHash *string

	// The maximum number of transaction events to list. Even if additional results
	// can be retrieved, the request can return less results than maxResults or an
	// empty array of results. To retrieve the next set of results, make another
	// request with the returned nextToken value. The value of nextToken is null when
	// there are no more results to return
	MaxResults *int32

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTransactionEventsOutput struct {

	// An array of TransactionEvent objects. Each object contains details about the
	// transaction events.
	//
	// This member is required.
	Events []types.TransactionEvent

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTransactionEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTransactionEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTransactionEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTransactionEvents"); err != nil {
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
	if err = addOpListTransactionEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTransactionEvents(options.Region), middleware.Before); err != nil {
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

// ListTransactionEventsAPIClient is a client that implements the
// ListTransactionEvents operation.
type ListTransactionEventsAPIClient interface {
	ListTransactionEvents(context.Context, *ListTransactionEventsInput, ...func(*Options)) (*ListTransactionEventsOutput, error)
}

var _ ListTransactionEventsAPIClient = (*Client)(nil)

// ListTransactionEventsPaginatorOptions is the paginator options for
// ListTransactionEvents
type ListTransactionEventsPaginatorOptions struct {
	// The maximum number of transaction events to list. Even if additional results
	// can be retrieved, the request can return less results than maxResults or an
	// empty array of results. To retrieve the next set of results, make another
	// request with the returned nextToken value. The value of nextToken is null when
	// there are no more results to return
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTransactionEventsPaginator is a paginator for ListTransactionEvents
type ListTransactionEventsPaginator struct {
	options   ListTransactionEventsPaginatorOptions
	client    ListTransactionEventsAPIClient
	params    *ListTransactionEventsInput
	nextToken *string
	firstPage bool
}

// NewListTransactionEventsPaginator returns a new ListTransactionEventsPaginator
func NewListTransactionEventsPaginator(client ListTransactionEventsAPIClient, params *ListTransactionEventsInput, optFns ...func(*ListTransactionEventsPaginatorOptions)) *ListTransactionEventsPaginator {
	if params == nil {
		params = &ListTransactionEventsInput{}
	}

	options := ListTransactionEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTransactionEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTransactionEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTransactionEvents page.
func (p *ListTransactionEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTransactionEventsOutput, error) {
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

	result, err := p.client.ListTransactionEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTransactionEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTransactionEvents",
	}
}
