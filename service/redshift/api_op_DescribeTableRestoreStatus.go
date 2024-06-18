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

// Lists the status of one or more table restore requests made using the RestoreTableFromClusterSnapshot API
// action. If you don't specify a value for the TableRestoreRequestId parameter,
// then DescribeTableRestoreStatus returns the status of all table restore
// requests ordered by the date and time of the request in ascending order.
// Otherwise DescribeTableRestoreStatus returns the status of the table specified
// by TableRestoreRequestId .
func (c *Client) DescribeTableRestoreStatus(ctx context.Context, params *DescribeTableRestoreStatusInput, optFns ...func(*Options)) (*DescribeTableRestoreStatusOutput, error) {
	if params == nil {
		params = &DescribeTableRestoreStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTableRestoreStatus", params, optFns, c.addOperationDescribeTableRestoreStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTableRestoreStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTableRestoreStatusInput struct {

	// The Amazon Redshift cluster that the table is being restored to.
	ClusterIdentifier *string

	// An optional pagination token provided by a previous DescribeTableRestoreStatus
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by the MaxRecords parameter.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32

	// The identifier of the table restore request to return status for. If you don't
	// specify a TableRestoreRequestId value, then DescribeTableRestoreStatus returns
	// the status of all in-progress table restore requests.
	TableRestoreRequestId *string

	noSmithyDocumentSerde
}

type DescribeTableRestoreStatusOutput struct {

	// A pagination token that can be used in a subsequent DescribeTableRestoreStatus request.
	Marker *string

	// A list of status details for one or more table restore requests.
	TableRestoreStatusDetails []types.TableRestoreStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTableRestoreStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTableRestoreStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTableRestoreStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTableRestoreStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTableRestoreStatus(options.Region), middleware.Before); err != nil {
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

// DescribeTableRestoreStatusPaginatorOptions is the paginator options for
// DescribeTableRestoreStatus
type DescribeTableRestoreStatusPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTableRestoreStatusPaginator is a paginator for
// DescribeTableRestoreStatus
type DescribeTableRestoreStatusPaginator struct {
	options   DescribeTableRestoreStatusPaginatorOptions
	client    DescribeTableRestoreStatusAPIClient
	params    *DescribeTableRestoreStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeTableRestoreStatusPaginator returns a new
// DescribeTableRestoreStatusPaginator
func NewDescribeTableRestoreStatusPaginator(client DescribeTableRestoreStatusAPIClient, params *DescribeTableRestoreStatusInput, optFns ...func(*DescribeTableRestoreStatusPaginatorOptions)) *DescribeTableRestoreStatusPaginator {
	if params == nil {
		params = &DescribeTableRestoreStatusInput{}
	}

	options := DescribeTableRestoreStatusPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTableRestoreStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTableRestoreStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTableRestoreStatus page.
func (p *DescribeTableRestoreStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTableRestoreStatusOutput, error) {
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
	result, err := p.client.DescribeTableRestoreStatus(ctx, &params, optFns...)
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

// DescribeTableRestoreStatusAPIClient is a client that implements the
// DescribeTableRestoreStatus operation.
type DescribeTableRestoreStatusAPIClient interface {
	DescribeTableRestoreStatus(context.Context, *DescribeTableRestoreStatusInput, ...func(*Options)) (*DescribeTableRestoreStatusOutput, error)
}

var _ DescribeTableRestoreStatusAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeTableRestoreStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTableRestoreStatus",
	}
}
