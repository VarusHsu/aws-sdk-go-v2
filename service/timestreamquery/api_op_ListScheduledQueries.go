// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of all scheduled queries in the caller's Amazon account and Region.
// ListScheduledQueries is eventually consistent.
func (c *Client) ListScheduledQueries(ctx context.Context, params *ListScheduledQueriesInput, optFns ...func(*Options)) (*ListScheduledQueriesOutput, error) {
	if params == nil {
		params = &ListScheduledQueriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListScheduledQueries", params, optFns, c.addOperationListScheduledQueriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListScheduledQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListScheduledQueriesInput struct {

	// The maximum number of items to return in the output. If the total number of
	// items available is more than the value specified, a NextToken is provided in
	// the output. To resume pagination, provide the NextToken value as the argument
	// to the subsequent call to ListScheduledQueriesRequest .
	MaxResults *int32

	//  A pagination token to resume pagination.
	NextToken *string

	noSmithyDocumentSerde
}

type ListScheduledQueriesOutput struct {

	// A list of scheduled queries.
	//
	// This member is required.
	ScheduledQueries []types.ScheduledQuery

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListScheduledQueriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListScheduledQueries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListScheduledQueries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListScheduledQueries"); err != nil {
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
	if err = addOpListScheduledQueriesDiscoverEndpointMiddleware(stack, options, c); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListScheduledQueries(options.Region), middleware.Before); err != nil {
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

// ListScheduledQueriesPaginatorOptions is the paginator options for
// ListScheduledQueries
type ListScheduledQueriesPaginatorOptions struct {
	// The maximum number of items to return in the output. If the total number of
	// items available is more than the value specified, a NextToken is provided in
	// the output. To resume pagination, provide the NextToken value as the argument
	// to the subsequent call to ListScheduledQueriesRequest .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListScheduledQueriesPaginator is a paginator for ListScheduledQueries
type ListScheduledQueriesPaginator struct {
	options   ListScheduledQueriesPaginatorOptions
	client    ListScheduledQueriesAPIClient
	params    *ListScheduledQueriesInput
	nextToken *string
	firstPage bool
}

// NewListScheduledQueriesPaginator returns a new ListScheduledQueriesPaginator
func NewListScheduledQueriesPaginator(client ListScheduledQueriesAPIClient, params *ListScheduledQueriesInput, optFns ...func(*ListScheduledQueriesPaginatorOptions)) *ListScheduledQueriesPaginator {
	if params == nil {
		params = &ListScheduledQueriesInput{}
	}

	options := ListScheduledQueriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListScheduledQueriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListScheduledQueriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListScheduledQueries page.
func (p *ListScheduledQueriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListScheduledQueriesOutput, error) {
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
	result, err := p.client.ListScheduledQueries(ctx, &params, optFns...)
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

func addOpListScheduledQueriesDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
				opt.EndpointResolverUsedForDiscovery = o.EndpointDiscovery.EndpointResolverUsedForDiscovery
			},
		},
		DiscoverOperation:            c.fetchOpListScheduledQueriesDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    true,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpListScheduledQueriesDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*ListScheduledQueriesInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)
	identifierMap["sdk#Region"] = region

	key := fmt.Sprintf("Timestream Query.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	endpoint, err := c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, region, key, opt)
	if err != nil {
		return internalEndpointDiscovery.WeightedAddress{}, err
	}

	weighted, ok := endpoint.GetValidAddress()
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("no valid endpoint address returned by the endpoint discovery api")
	}
	return weighted, nil
}

// ListScheduledQueriesAPIClient is a client that implements the
// ListScheduledQueries operation.
type ListScheduledQueriesAPIClient interface {
	ListScheduledQueries(context.Context, *ListScheduledQueriesInput, ...func(*Options)) (*ListScheduledQueriesOutput, error)
}

var _ ListScheduledQueriesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListScheduledQueries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListScheduledQueries",
	}
}
