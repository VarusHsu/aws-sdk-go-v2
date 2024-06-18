// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns status information for sources within an aggregator. The status
// includes information about the last time Config verified authorization between
// the source account and an aggregator account. In case of a failure, the status
// contains the related error code or message.
func (c *Client) DescribeConfigurationAggregatorSourcesStatus(ctx context.Context, params *DescribeConfigurationAggregatorSourcesStatusInput, optFns ...func(*Options)) (*DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	if params == nil {
		params = &DescribeConfigurationAggregatorSourcesStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationAggregatorSourcesStatus", params, optFns, c.addOperationDescribeConfigurationAggregatorSourcesStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationAggregatorSourcesStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationAggregatorSourcesStatusInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The maximum number of AggregatorSourceStatus returned on each page. The default
	// is maximum. If you specify 0, Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Filters the status type.
	//
	//   - Valid value FAILED indicates errors while moving data.
	//
	//   - Valid value SUCCEEDED indicates the data was successfully moved.
	//
	//   - Valid value OUTDATED indicates the data is not the most recent.
	UpdateStatus []types.AggregatedSourceStatusType

	noSmithyDocumentSerde
}

type DescribeConfigurationAggregatorSourcesStatusOutput struct {

	// Returns an AggregatedSourceStatus object.
	AggregatedSourceStatusList []types.AggregatedSourceStatus

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationAggregatorSourcesStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurationAggregatorSourcesStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurationAggregatorSourcesStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeConfigurationAggregatorSourcesStatus"); err != nil {
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
	if err = addOpDescribeConfigurationAggregatorSourcesStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationAggregatorSourcesStatus(options.Region), middleware.Before); err != nil {
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

// DescribeConfigurationAggregatorSourcesStatusPaginatorOptions is the paginator
// options for DescribeConfigurationAggregatorSourcesStatus
type DescribeConfigurationAggregatorSourcesStatusPaginatorOptions struct {
	// The maximum number of AggregatorSourceStatus returned on each page. The default
	// is maximum. If you specify 0, Config uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConfigurationAggregatorSourcesStatusPaginator is a paginator for
// DescribeConfigurationAggregatorSourcesStatus
type DescribeConfigurationAggregatorSourcesStatusPaginator struct {
	options   DescribeConfigurationAggregatorSourcesStatusPaginatorOptions
	client    DescribeConfigurationAggregatorSourcesStatusAPIClient
	params    *DescribeConfigurationAggregatorSourcesStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeConfigurationAggregatorSourcesStatusPaginator returns a new
// DescribeConfigurationAggregatorSourcesStatusPaginator
func NewDescribeConfigurationAggregatorSourcesStatusPaginator(client DescribeConfigurationAggregatorSourcesStatusAPIClient, params *DescribeConfigurationAggregatorSourcesStatusInput, optFns ...func(*DescribeConfigurationAggregatorSourcesStatusPaginatorOptions)) *DescribeConfigurationAggregatorSourcesStatusPaginator {
	if params == nil {
		params = &DescribeConfigurationAggregatorSourcesStatusInput{}
	}

	options := DescribeConfigurationAggregatorSourcesStatusPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConfigurationAggregatorSourcesStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConfigurationAggregatorSourcesStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeConfigurationAggregatorSourcesStatus page.
func (p *DescribeConfigurationAggregatorSourcesStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeConfigurationAggregatorSourcesStatus(ctx, &params, optFns...)
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

// DescribeConfigurationAggregatorSourcesStatusAPIClient is a client that
// implements the DescribeConfigurationAggregatorSourcesStatus operation.
type DescribeConfigurationAggregatorSourcesStatusAPIClient interface {
	DescribeConfigurationAggregatorSourcesStatus(context.Context, *DescribeConfigurationAggregatorSourcesStatusInput, ...func(*Options)) (*DescribeConfigurationAggregatorSourcesStatusOutput, error)
}

var _ DescribeConfigurationAggregatorSourcesStatusAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeConfigurationAggregatorSourcesStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeConfigurationAggregatorSourcesStatus",
	}
}
