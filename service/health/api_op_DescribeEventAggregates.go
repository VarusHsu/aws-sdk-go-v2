// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the number of events of each event type (issue, scheduled change, and
// account notification). If no filter is specified, the counts of all events in
// each category are returned. This API operation uses pagination. Specify the
// nextToken parameter in the next request to return more results.
func (c *Client) DescribeEventAggregates(ctx context.Context, params *DescribeEventAggregatesInput, optFns ...func(*Options)) (*DescribeEventAggregatesOutput, error) {
	if params == nil {
		params = &DescribeEventAggregatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventAggregates", params, optFns, c.addOperationDescribeEventAggregatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventAggregatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventAggregatesInput struct {

	// The only currently supported value is eventTypeCategory .
	//
	// This member is required.
	AggregateField types.EventAggregateField

	// Values to narrow the results returned.
	Filter *types.EventFilter

	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	MaxResults *int32

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeEventAggregatesOutput struct {

	// The number of events in each category that meet the optional filter criteria.
	EventAggregates []types.EventAggregate

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventAggregatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventAggregates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventAggregates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEventAggregates"); err != nil {
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
	if err = addOpDescribeEventAggregatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventAggregates(options.Region), middleware.Before); err != nil {
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

// DescribeEventAggregatesAPIClient is a client that implements the
// DescribeEventAggregates operation.
type DescribeEventAggregatesAPIClient interface {
	DescribeEventAggregates(context.Context, *DescribeEventAggregatesInput, ...func(*Options)) (*DescribeEventAggregatesOutput, error)
}

var _ DescribeEventAggregatesAPIClient = (*Client)(nil)

// DescribeEventAggregatesPaginatorOptions is the paginator options for
// DescribeEventAggregates
type DescribeEventAggregatesPaginatorOptions struct {
	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEventAggregatesPaginator is a paginator for DescribeEventAggregates
type DescribeEventAggregatesPaginator struct {
	options   DescribeEventAggregatesPaginatorOptions
	client    DescribeEventAggregatesAPIClient
	params    *DescribeEventAggregatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeEventAggregatesPaginator returns a new
// DescribeEventAggregatesPaginator
func NewDescribeEventAggregatesPaginator(client DescribeEventAggregatesAPIClient, params *DescribeEventAggregatesInput, optFns ...func(*DescribeEventAggregatesPaginatorOptions)) *DescribeEventAggregatesPaginator {
	if params == nil {
		params = &DescribeEventAggregatesInput{}
	}

	options := DescribeEventAggregatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEventAggregatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEventAggregatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEventAggregates page.
func (p *DescribeEventAggregatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEventAggregatesOutput, error) {
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

	result, err := p.client.DescribeEventAggregates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEventAggregates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEventAggregates",
	}
}
