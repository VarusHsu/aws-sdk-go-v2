// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the accessors and their properties. Accessor objects are
// containers that have the information required for token based access to your
// Ethereum nodes.
func (c *Client) ListAccessors(ctx context.Context, params *ListAccessorsInput, optFns ...func(*Options)) (*ListAccessorsOutput, error) {
	if params == nil {
		params = &ListAccessorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessors", params, optFns, c.addOperationListAccessorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessorsInput struct {

	//  The maximum number of accessors to list.
	MaxResults *int32

	// The blockchain network that the Accessor token is created for.
	//
	// Use the value ETHEREUM_MAINNET_AND_GOERLI for all existing Accessors tokens
	// that were created before the networkType property was introduced.
	NetworkType types.AccessorNetworkType

	//  The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAccessorsOutput struct {

	// An array of AccessorSummary objects that contain configuration properties for
	// each accessor.
	Accessors []types.AccessorSummary

	//  The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccessorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAccessors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAccessors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccessors"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessors(options.Region), middleware.Before); err != nil {
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

// ListAccessorsPaginatorOptions is the paginator options for ListAccessors
type ListAccessorsPaginatorOptions struct {
	//  The maximum number of accessors to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessorsPaginator is a paginator for ListAccessors
type ListAccessorsPaginator struct {
	options   ListAccessorsPaginatorOptions
	client    ListAccessorsAPIClient
	params    *ListAccessorsInput
	nextToken *string
	firstPage bool
}

// NewListAccessorsPaginator returns a new ListAccessorsPaginator
func NewListAccessorsPaginator(client ListAccessorsAPIClient, params *ListAccessorsInput, optFns ...func(*ListAccessorsPaginatorOptions)) *ListAccessorsPaginator {
	if params == nil {
		params = &ListAccessorsInput{}
	}

	options := ListAccessorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccessors page.
func (p *ListAccessorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessorsOutput, error) {
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
	result, err := p.client.ListAccessors(ctx, &params, optFns...)
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

// ListAccessorsAPIClient is a client that implements the ListAccessors operation.
type ListAccessorsAPIClient interface {
	ListAccessors(context.Context, *ListAccessorsInput, ...func(*Options)) (*ListAccessorsOutput, error)
}

var _ ListAccessorsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAccessors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccessors",
	}
}
