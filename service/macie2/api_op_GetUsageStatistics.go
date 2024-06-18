// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves (queries) quotas and aggregated usage data for one or more accounts.
func (c *Client) GetUsageStatistics(ctx context.Context, params *GetUsageStatisticsInput, optFns ...func(*Options)) (*GetUsageStatisticsOutput, error) {
	if params == nil {
		params = &GetUsageStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsageStatistics", params, optFns, c.addOperationGetUsageStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsageStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUsageStatisticsInput struct {

	// An array of objects, one for each condition to use to filter the query results.
	// If you specify more than one condition, Amazon Macie uses an AND operator to
	// join the conditions.
	FilterBy []types.UsageStatisticsFilter

	// The maximum number of items to include in each page of the response.
	MaxResults *int32

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string

	// The criteria to use to sort the query results.
	SortBy *types.UsageStatisticsSortBy

	// The inclusive time period to query usage data for. Valid values are:
	// MONTH_TO_DATE, for the current calendar month to date; and, PAST_30_DAYS, for
	// the preceding 30 days. If you don't specify a value, Amazon Macie provides usage
	// data for the preceding 30 days.
	TimeRange types.TimeRange

	noSmithyDocumentSerde
}

type GetUsageStatisticsOutput struct {

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// An array of objects that contains the results of the query. Each object
	// contains the data for an account that matches the filter criteria specified in
	// the request.
	Records []types.UsageRecord

	// The inclusive time period that the usage data applies to. Possible values are:
	// MONTH_TO_DATE, for the current calendar month to date; and, PAST_30_DAYS, for
	// the preceding 30 days.
	TimeRange types.TimeRange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUsageStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsageStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsageStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUsageStatistics"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsageStatistics(options.Region), middleware.Before); err != nil {
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

// GetUsageStatisticsPaginatorOptions is the paginator options for
// GetUsageStatistics
type GetUsageStatisticsPaginatorOptions struct {
	// The maximum number of items to include in each page of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetUsageStatisticsPaginator is a paginator for GetUsageStatistics
type GetUsageStatisticsPaginator struct {
	options   GetUsageStatisticsPaginatorOptions
	client    GetUsageStatisticsAPIClient
	params    *GetUsageStatisticsInput
	nextToken *string
	firstPage bool
}

// NewGetUsageStatisticsPaginator returns a new GetUsageStatisticsPaginator
func NewGetUsageStatisticsPaginator(client GetUsageStatisticsAPIClient, params *GetUsageStatisticsInput, optFns ...func(*GetUsageStatisticsPaginatorOptions)) *GetUsageStatisticsPaginator {
	if params == nil {
		params = &GetUsageStatisticsInput{}
	}

	options := GetUsageStatisticsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetUsageStatisticsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetUsageStatisticsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetUsageStatistics page.
func (p *GetUsageStatisticsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetUsageStatisticsOutput, error) {
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
	result, err := p.client.GetUsageStatistics(ctx, &params, optFns...)
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

// GetUsageStatisticsAPIClient is a client that implements the GetUsageStatistics
// operation.
type GetUsageStatisticsAPIClient interface {
	GetUsageStatistics(context.Context, *GetUsageStatisticsInput, ...func(*Options)) (*GetUsageStatisticsOutput, error)
}

var _ GetUsageStatisticsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetUsageStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUsageStatistics",
	}
}
