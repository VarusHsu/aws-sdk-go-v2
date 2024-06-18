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

// A batch request to retrieve all device positions.
func (c *Client) ListDevicePositions(ctx context.Context, params *ListDevicePositionsInput, optFns ...func(*Options)) (*ListDevicePositionsOutput, error) {
	if params == nil {
		params = &ListDevicePositionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevicePositions", params, optFns, c.addOperationListDevicePositionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevicePositionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDevicePositionsInput struct {

	// The tracker resource containing the requested devices.
	//
	// This member is required.
	TrackerName *string

	// The geometry used to filter device positions.
	FilterGeometry *types.TrackingFilterGeometry

	// An optional limit for the number of entries returned in a single call.
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

type ListDevicePositionsOutput struct {

	// Contains details about each device's last known position.
	//
	// This member is required.
	Entries []types.ListDevicePositionsResponseEntry

	// A pagination token indicating there are additional pages available. You can use
	// the token in a following request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDevicePositionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDevicePositions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDevicePositions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDevicePositions"); err != nil {
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
	if err = addEndpointPrefix_opListDevicePositionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListDevicePositionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDevicePositions(options.Region), middleware.Before); err != nil {
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

// ListDevicePositionsPaginatorOptions is the paginator options for
// ListDevicePositions
type ListDevicePositionsPaginatorOptions struct {
	// An optional limit for the number of entries returned in a single call.
	//
	// Default value: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDevicePositionsPaginator is a paginator for ListDevicePositions
type ListDevicePositionsPaginator struct {
	options   ListDevicePositionsPaginatorOptions
	client    ListDevicePositionsAPIClient
	params    *ListDevicePositionsInput
	nextToken *string
	firstPage bool
}

// NewListDevicePositionsPaginator returns a new ListDevicePositionsPaginator
func NewListDevicePositionsPaginator(client ListDevicePositionsAPIClient, params *ListDevicePositionsInput, optFns ...func(*ListDevicePositionsPaginatorOptions)) *ListDevicePositionsPaginator {
	if params == nil {
		params = &ListDevicePositionsInput{}
	}

	options := ListDevicePositionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDevicePositionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDevicePositionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDevicePositions page.
func (p *ListDevicePositionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDevicePositionsOutput, error) {
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
	result, err := p.client.ListDevicePositions(ctx, &params, optFns...)
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

type endpointPrefix_opListDevicePositionsMiddleware struct {
}

func (*endpointPrefix_opListDevicePositionsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListDevicePositionsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "tracking." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListDevicePositionsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListDevicePositionsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListDevicePositionsAPIClient is a client that implements the
// ListDevicePositions operation.
type ListDevicePositionsAPIClient interface {
	ListDevicePositions(context.Context, *ListDevicePositionsInput, ...func(*Options)) (*ListDevicePositionsOutput, error)
}

var _ ListDevicePositionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDevicePositions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDevicePositions",
	}
}
