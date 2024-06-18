// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns exchange status details and associated metadata for a reserved-node
// exchange. Statuses include such values as in progress and requested.
func (c *Client) DescribeReservedNodeExchangeStatus(ctx context.Context, params *DescribeReservedNodeExchangeStatusInput, optFns ...func(*Options)) (*DescribeReservedNodeExchangeStatusOutput, error) {
	if params == nil {
		params = &DescribeReservedNodeExchangeStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedNodeExchangeStatus", params, optFns, c.addOperationDescribeReservedNodeExchangeStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedNodeExchangeStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReservedNodeExchangeStatusInput struct {

	// An optional pagination token provided by a previous
	// DescribeReservedNodeExchangeStatus request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// the MaxRecords parameter. You can retrieve the next set of response records by
	// providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a Marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	MaxRecords *int32

	// The identifier of the reserved-node exchange request.
	ReservedNodeExchangeRequestId *string

	// The identifier of the source reserved node in a reserved-node exchange request.
	ReservedNodeId *string

	noSmithyDocumentSerde
}

type DescribeReservedNodeExchangeStatusOutput struct {

	// A pagination token provided by a previous DescribeReservedNodeExchangeStatus
	// request.
	Marker *string

	// The details of the reserved-node exchange request, including the status,
	// request time, source reserved-node identifier, and additional details.
	ReservedNodeExchangeStatusDetails []types.ReservedNodeExchangeStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedNodeExchangeStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeReservedNodeExchangeStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeReservedNodeExchangeStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReservedNodeExchangeStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedNodeExchangeStatus(options.Region), middleware.Before); err != nil {
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

// DescribeReservedNodeExchangeStatusPaginatorOptions is the paginator options for
// DescribeReservedNodeExchangeStatus
type DescribeReservedNodeExchangeStatusPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a Marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedNodeExchangeStatusPaginator is a paginator for
// DescribeReservedNodeExchangeStatus
type DescribeReservedNodeExchangeStatusPaginator struct {
	options   DescribeReservedNodeExchangeStatusPaginatorOptions
	client    DescribeReservedNodeExchangeStatusAPIClient
	params    *DescribeReservedNodeExchangeStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedNodeExchangeStatusPaginator returns a new
// DescribeReservedNodeExchangeStatusPaginator
func NewDescribeReservedNodeExchangeStatusPaginator(client DescribeReservedNodeExchangeStatusAPIClient, params *DescribeReservedNodeExchangeStatusInput, optFns ...func(*DescribeReservedNodeExchangeStatusPaginatorOptions)) *DescribeReservedNodeExchangeStatusPaginator {
	if params == nil {
		params = &DescribeReservedNodeExchangeStatusInput{}
	}

	options := DescribeReservedNodeExchangeStatusPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedNodeExchangeStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedNodeExchangeStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedNodeExchangeStatus page.
func (p *DescribeReservedNodeExchangeStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedNodeExchangeStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeReservedNodeExchangeStatus(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeReservedNodeExchangeStatusAPIClient is a client that implements the
// DescribeReservedNodeExchangeStatus operation.
type DescribeReservedNodeExchangeStatusAPIClient interface {
	DescribeReservedNodeExchangeStatus(context.Context, *DescribeReservedNodeExchangeStatusInput, ...func(*Options)) (*DescribeReservedNodeExchangeStatusOutput, error)
}

var _ DescribeReservedNodeExchangeStatusAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeReservedNodeExchangeStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReservedNodeExchangeStatus",
	}
}
