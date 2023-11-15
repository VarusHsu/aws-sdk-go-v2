// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the history of an asset property's values. For more information, see
// Querying historical values (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#historical-values)
// in the IoT SiteWise User Guide. To identify an asset property, you must specify
// one of the following:
//   - The assetId and propertyId of an asset property.
//   - A propertyAlias , which is a data stream alias (for example,
//     /company/windfarm/3/turbine/7/temperature ). To define an asset property's
//     alias, see UpdateAssetProperty (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html)
//     .
func (c *Client) GetAssetPropertyValueHistory(ctx context.Context, params *GetAssetPropertyValueHistoryInput, optFns ...func(*Options)) (*GetAssetPropertyValueHistoryOutput, error) {
	if params == nil {
		params = &GetAssetPropertyValueHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssetPropertyValueHistory", params, optFns, c.addOperationGetAssetPropertyValueHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssetPropertyValueHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssetPropertyValueHistoryInput struct {

	// The ID of the asset.
	AssetId *string

	// The inclusive end of the range from which to query historical data, expressed
	// in seconds in Unix epoch time.
	EndDate *time.Time

	// The maximum number of results to return for each paginated request. A result
	// set is returned in the two cases, whichever occurs first.
	//   - The size of the result set is equal to 4 MB.
	//   - The number of data points in the result set is equal to the value of
	//   maxResults . The maximum value of maxResults is 20000.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	// The alias that identifies the property, such as an OPC-UA server data stream
	// path (for example, /company/windfarm/3/turbine/7/temperature ). For more
	// information, see Mapping industrial data streams to asset properties (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the IoT SiteWise User Guide.
	PropertyAlias *string

	// The ID of the asset property.
	PropertyId *string

	// The quality by which to filter asset data.
	Qualities []types.Quality

	// The exclusive start of the range from which to query historical data, expressed
	// in seconds in Unix epoch time.
	StartDate *time.Time

	// The chronological sorting order of the requested information. Default: ASCENDING
	TimeOrdering types.TimeOrdering

	noSmithyDocumentSerde
}

type GetAssetPropertyValueHistoryOutput struct {

	// The asset property's value history.
	//
	// This member is required.
	AssetPropertyValueHistory []types.AssetPropertyValue

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAssetPropertyValueHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAssetPropertyValueHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAssetPropertyValueHistory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAssetPropertyValueHistory"); err != nil {
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
	if err = addEndpointPrefix_opGetAssetPropertyValueHistoryMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssetPropertyValueHistory(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetAssetPropertyValueHistoryMiddleware struct {
}

func (*endpointPrefix_opGetAssetPropertyValueHistoryMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAssetPropertyValueHistoryMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetAssetPropertyValueHistoryMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetAssetPropertyValueHistoryMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// GetAssetPropertyValueHistoryAPIClient is a client that implements the
// GetAssetPropertyValueHistory operation.
type GetAssetPropertyValueHistoryAPIClient interface {
	GetAssetPropertyValueHistory(context.Context, *GetAssetPropertyValueHistoryInput, ...func(*Options)) (*GetAssetPropertyValueHistoryOutput, error)
}

var _ GetAssetPropertyValueHistoryAPIClient = (*Client)(nil)

// GetAssetPropertyValueHistoryPaginatorOptions is the paginator options for
// GetAssetPropertyValueHistory
type GetAssetPropertyValueHistoryPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. A result
	// set is returned in the two cases, whichever occurs first.
	//   - The size of the result set is equal to 4 MB.
	//   - The number of data points in the result set is equal to the value of
	//   maxResults . The maximum value of maxResults is 20000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetAssetPropertyValueHistoryPaginator is a paginator for
// GetAssetPropertyValueHistory
type GetAssetPropertyValueHistoryPaginator struct {
	options   GetAssetPropertyValueHistoryPaginatorOptions
	client    GetAssetPropertyValueHistoryAPIClient
	params    *GetAssetPropertyValueHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetAssetPropertyValueHistoryPaginator returns a new
// GetAssetPropertyValueHistoryPaginator
func NewGetAssetPropertyValueHistoryPaginator(client GetAssetPropertyValueHistoryAPIClient, params *GetAssetPropertyValueHistoryInput, optFns ...func(*GetAssetPropertyValueHistoryPaginatorOptions)) *GetAssetPropertyValueHistoryPaginator {
	if params == nil {
		params = &GetAssetPropertyValueHistoryInput{}
	}

	options := GetAssetPropertyValueHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetAssetPropertyValueHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetAssetPropertyValueHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetAssetPropertyValueHistory page.
func (p *GetAssetPropertyValueHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetAssetPropertyValueHistoryOutput, error) {
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

	result, err := p.client.GetAssetPropertyValueHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetAssetPropertyValueHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAssetPropertyValueHistory",
	}
}
