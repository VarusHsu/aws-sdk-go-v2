// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists protected queries, sorted by the most recent query.
func (c *Client) ListProtectedQueries(ctx context.Context, params *ListProtectedQueriesInput, optFns ...func(*Options)) (*ListProtectedQueriesOutput, error) {
	if params == nil {
		params = &ListProtectedQueriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProtectedQueries", params, optFns, c.addOperationListProtectedQueriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProtectedQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProtectedQueriesInput struct {

	// The identifier for the membership in the collaboration.
	//
	// This member is required.
	MembershipIdentifier *string

	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service can return a nextToken even if the
	// maximum results has not been met.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// A filter on the status of the protected query.
	Status types.ProtectedQueryStatus

	noSmithyDocumentSerde
}

type ListProtectedQueriesOutput struct {

	// A list of protected queries.
	//
	// This member is required.
	ProtectedQueries []types.ProtectedQuerySummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProtectedQueriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProtectedQueries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProtectedQueries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProtectedQueries"); err != nil {
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
	if err = addOpListProtectedQueriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProtectedQueries(options.Region), middleware.Before); err != nil {
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

// ListProtectedQueriesAPIClient is a client that implements the
// ListProtectedQueries operation.
type ListProtectedQueriesAPIClient interface {
	ListProtectedQueries(context.Context, *ListProtectedQueriesInput, ...func(*Options)) (*ListProtectedQueriesOutput, error)
}

var _ ListProtectedQueriesAPIClient = (*Client)(nil)

// ListProtectedQueriesPaginatorOptions is the paginator options for
// ListProtectedQueries
type ListProtectedQueriesPaginatorOptions struct {
	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service can return a nextToken even if the
	// maximum results has not been met.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProtectedQueriesPaginator is a paginator for ListProtectedQueries
type ListProtectedQueriesPaginator struct {
	options   ListProtectedQueriesPaginatorOptions
	client    ListProtectedQueriesAPIClient
	params    *ListProtectedQueriesInput
	nextToken *string
	firstPage bool
}

// NewListProtectedQueriesPaginator returns a new ListProtectedQueriesPaginator
func NewListProtectedQueriesPaginator(client ListProtectedQueriesAPIClient, params *ListProtectedQueriesInput, optFns ...func(*ListProtectedQueriesPaginatorOptions)) *ListProtectedQueriesPaginator {
	if params == nil {
		params = &ListProtectedQueriesInput{}
	}

	options := ListProtectedQueriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProtectedQueriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProtectedQueriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProtectedQueries page.
func (p *ListProtectedQueriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProtectedQueriesOutput, error) {
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

	result, err := p.client.ListProtectedQueries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProtectedQueries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProtectedQueries",
	}
}
