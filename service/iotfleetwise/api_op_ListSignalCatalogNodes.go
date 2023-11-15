// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists of information about the signals (nodes) specified in a signal catalog.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func (c *Client) ListSignalCatalogNodes(ctx context.Context, params *ListSignalCatalogNodesInput, optFns ...func(*Options)) (*ListSignalCatalogNodesOutput, error) {
	if params == nil {
		params = &ListSignalCatalogNodesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSignalCatalogNodes", params, optFns, c.addOperationListSignalCatalogNodesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSignalCatalogNodesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSignalCatalogNodesInput struct {

	// The name of the signal catalog to list information about.
	//
	// This member is required.
	Name *string

	// The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

	// A pagination token for the next set of results. If the results of a search are
	// large, only a portion of the results are returned, and a nextToken pagination
	// token is returned in the response. To retrieve the next set of results, reissue
	// the search request and include the returned token. When all results have been
	// returned, the response does not contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSignalCatalogNodesOutput struct {

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	// A list of information about nodes.
	Nodes []types.Node

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSignalCatalogNodesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListSignalCatalogNodes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListSignalCatalogNodes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSignalCatalogNodes"); err != nil {
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
	if err = addOpListSignalCatalogNodesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSignalCatalogNodes(options.Region), middleware.Before); err != nil {
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

// ListSignalCatalogNodesAPIClient is a client that implements the
// ListSignalCatalogNodes operation.
type ListSignalCatalogNodesAPIClient interface {
	ListSignalCatalogNodes(context.Context, *ListSignalCatalogNodesInput, ...func(*Options)) (*ListSignalCatalogNodesOutput, error)
}

var _ ListSignalCatalogNodesAPIClient = (*Client)(nil)

// ListSignalCatalogNodesPaginatorOptions is the paginator options for
// ListSignalCatalogNodes
type ListSignalCatalogNodesPaginatorOptions struct {
	// The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSignalCatalogNodesPaginator is a paginator for ListSignalCatalogNodes
type ListSignalCatalogNodesPaginator struct {
	options   ListSignalCatalogNodesPaginatorOptions
	client    ListSignalCatalogNodesAPIClient
	params    *ListSignalCatalogNodesInput
	nextToken *string
	firstPage bool
}

// NewListSignalCatalogNodesPaginator returns a new ListSignalCatalogNodesPaginator
func NewListSignalCatalogNodesPaginator(client ListSignalCatalogNodesAPIClient, params *ListSignalCatalogNodesInput, optFns ...func(*ListSignalCatalogNodesPaginatorOptions)) *ListSignalCatalogNodesPaginator {
	if params == nil {
		params = &ListSignalCatalogNodesInput{}
	}

	options := ListSignalCatalogNodesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSignalCatalogNodesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSignalCatalogNodesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSignalCatalogNodes page.
func (p *ListSignalCatalogNodesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSignalCatalogNodesOutput, error) {
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

	result, err := p.client.ListSignalCatalogNodes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSignalCatalogNodes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSignalCatalogNodes",
	}
}
