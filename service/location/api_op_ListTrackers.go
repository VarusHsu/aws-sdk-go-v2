// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists tracker resources in your Amazon Web Services account.
func (c *Client) ListTrackers(ctx context.Context, params *ListTrackersInput, optFns ...func(*Options)) (*ListTrackersOutput, error) {
	if params == nil {
		params = &ListTrackersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrackers", params, optFns, c.addOperationListTrackersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrackersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrackersInput struct {

	// An optional limit for the number of resources returned in a single call.
	//
	// Default value: 100
	MaxResults *int32

	// The pagination token specifying which page of results to return in the
	// response. If no token is provided, the default page is the first page.
	//
	// Default value: null
	NextToken *string

	noSmithyDocumentSerde
}

type ListTrackersOutput struct {

	// Contains tracker resources in your Amazon Web Services account. Details include
	// tracker name, description and timestamps for when the tracker was created and
	// last updated.
	//
	// This member is required.
	Entries []types.ListTrackersResponseEntry

	// A pagination token indicating there are additional pages available. You can use
	// the token in a following request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrackersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTrackers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTrackers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTrackers"); err != nil {
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
	if err = addEndpointPrefix_opListTrackersMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrackers(options.Region), middleware.Before); err != nil {
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

// ListTrackersPaginatorOptions is the paginator options for ListTrackers
type ListTrackersPaginatorOptions struct {
	// An optional limit for the number of resources returned in a single call.
	//
	// Default value: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTrackersPaginator is a paginator for ListTrackers
type ListTrackersPaginator struct {
	options   ListTrackersPaginatorOptions
	client    ListTrackersAPIClient
	params    *ListTrackersInput
	nextToken *string
	firstPage bool
}

// NewListTrackersPaginator returns a new ListTrackersPaginator
func NewListTrackersPaginator(client ListTrackersAPIClient, params *ListTrackersInput, optFns ...func(*ListTrackersPaginatorOptions)) *ListTrackersPaginator {
	if params == nil {
		params = &ListTrackersInput{}
	}

	options := ListTrackersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTrackersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTrackersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTrackers page.
func (p *ListTrackersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTrackersOutput, error) {
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
	result, err := p.client.ListTrackers(ctx, &params, optFns...)
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

type endpointPrefix_opListTrackersMiddleware struct {
}

func (*endpointPrefix_opListTrackersMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListTrackersMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "cp.tracking." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListTrackersMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListTrackersMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListTrackersAPIClient is a client that implements the ListTrackers operation.
type ListTrackersAPIClient interface {
	ListTrackers(context.Context, *ListTrackersInput, ...func(*Options)) (*ListTrackersOutput, error)
}

var _ ListTrackersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTrackers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTrackers",
	}
}
