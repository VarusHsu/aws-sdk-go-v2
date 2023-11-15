// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List purchased reservations.
func (c *Client) ListReservations(ctx context.Context, params *ListReservationsInput, optFns ...func(*Options)) (*ListReservationsOutput, error) {
	if params == nil {
		params = &ListReservationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReservations", params, optFns, c.addOperationListReservationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReservationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListReservationsRequest
type ListReservationsInput struct {

	// Filter by channel class, 'STANDARD' or 'SINGLE_PIPELINE'
	ChannelClass *string

	// Filter by codec, 'AVC', 'HEVC', 'MPEG2', 'AUDIO', or 'LINK'
	Codec *string

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// Filter by bitrate, 'MAX_10_MBPS', 'MAX_20_MBPS', or 'MAX_50_MBPS'
	MaximumBitrate *string

	// Filter by framerate, 'MAX_30_FPS' or 'MAX_60_FPS'
	MaximumFramerate *string

	// Placeholder documentation for __string
	NextToken *string

	// Filter by resolution, 'SD', 'HD', 'FHD', or 'UHD'
	Resolution *string

	// Filter by resource type, 'INPUT', 'OUTPUT', 'MULTIPLEX', or 'CHANNEL'
	ResourceType *string

	// Filter by special feature, 'ADVANCED_AUDIO' or 'AUDIO_NORMALIZATION'
	SpecialFeature *string

	// Filter by video quality, 'STANDARD', 'ENHANCED', or 'PREMIUM'
	VideoQuality *string

	noSmithyDocumentSerde
}

// Placeholder documentation for ListReservationsResponse
type ListReservationsOutput struct {

	// Token to retrieve the next page of results
	NextToken *string

	// List of reservations
	Reservations []types.Reservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReservationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReservations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReservations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListReservations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReservations(options.Region), middleware.Before); err != nil {
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

// ListReservationsAPIClient is a client that implements the ListReservations
// operation.
type ListReservationsAPIClient interface {
	ListReservations(context.Context, *ListReservationsInput, ...func(*Options)) (*ListReservationsOutput, error)
}

var _ ListReservationsAPIClient = (*Client)(nil)

// ListReservationsPaginatorOptions is the paginator options for ListReservations
type ListReservationsPaginatorOptions struct {
	// Placeholder documentation for MaxResults
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReservationsPaginator is a paginator for ListReservations
type ListReservationsPaginator struct {
	options   ListReservationsPaginatorOptions
	client    ListReservationsAPIClient
	params    *ListReservationsInput
	nextToken *string
	firstPage bool
}

// NewListReservationsPaginator returns a new ListReservationsPaginator
func NewListReservationsPaginator(client ListReservationsAPIClient, params *ListReservationsInput, optFns ...func(*ListReservationsPaginatorOptions)) *ListReservationsPaginator {
	if params == nil {
		params = &ListReservationsInput{}
	}

	options := ListReservationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReservationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReservationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReservations page.
func (p *ListReservationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReservationsOutput, error) {
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

	result, err := p.client.ListReservations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReservations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListReservations",
	}
}
