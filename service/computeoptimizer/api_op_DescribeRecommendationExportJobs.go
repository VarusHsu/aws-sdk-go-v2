// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes recommendation export jobs created in the last seven days. Use the
// ExportAutoScalingGroupRecommendations or ExportEC2InstanceRecommendations
// actions to request an export of your recommendations. Then use the
// DescribeRecommendationExportJobs action to view your export jobs.
func (c *Client) DescribeRecommendationExportJobs(ctx context.Context, params *DescribeRecommendationExportJobsInput, optFns ...func(*Options)) (*DescribeRecommendationExportJobsOutput, error) {
	if params == nil {
		params = &DescribeRecommendationExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecommendationExportJobs", params, optFns, c.addOperationDescribeRecommendationExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecommendationExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecommendationExportJobsInput struct {

	// An array of objects to specify a filter that returns a more specific list of
	// export jobs.
	Filters []types.JobFilter

	// The identification numbers of the export jobs to return. An export job ID is
	// returned when you create an export using the
	// ExportAutoScalingGroupRecommendations or ExportEC2InstanceRecommendations
	// actions. All export jobs created in the last seven days are returned if this
	// parameter is omitted.
	JobIds []string

	// The maximum number of export jobs to return with a single request. To retrieve
	// the remaining results, make another request with the returned nextToken value.
	MaxResults *int32

	// The token to advance to the next page of export jobs.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeRecommendationExportJobsOutput struct {

	// The token to use to advance to the next page of export jobs. This value is null
	// when there are no more pages of export jobs to return.
	NextToken *string

	// An array of objects that describe recommendation export jobs.
	RecommendationExportJobs []types.RecommendationExportJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRecommendationExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeRecommendationExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeRecommendationExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeRecommendationExportJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecommendationExportJobs(options.Region), middleware.Before); err != nil {
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

// DescribeRecommendationExportJobsAPIClient is a client that implements the
// DescribeRecommendationExportJobs operation.
type DescribeRecommendationExportJobsAPIClient interface {
	DescribeRecommendationExportJobs(context.Context, *DescribeRecommendationExportJobsInput, ...func(*Options)) (*DescribeRecommendationExportJobsOutput, error)
}

var _ DescribeRecommendationExportJobsAPIClient = (*Client)(nil)

// DescribeRecommendationExportJobsPaginatorOptions is the paginator options for
// DescribeRecommendationExportJobs
type DescribeRecommendationExportJobsPaginatorOptions struct {
	// The maximum number of export jobs to return with a single request. To retrieve
	// the remaining results, make another request with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRecommendationExportJobsPaginator is a paginator for
// DescribeRecommendationExportJobs
type DescribeRecommendationExportJobsPaginator struct {
	options   DescribeRecommendationExportJobsPaginatorOptions
	client    DescribeRecommendationExportJobsAPIClient
	params    *DescribeRecommendationExportJobsInput
	nextToken *string
	firstPage bool
}

// NewDescribeRecommendationExportJobsPaginator returns a new
// DescribeRecommendationExportJobsPaginator
func NewDescribeRecommendationExportJobsPaginator(client DescribeRecommendationExportJobsAPIClient, params *DescribeRecommendationExportJobsInput, optFns ...func(*DescribeRecommendationExportJobsPaginatorOptions)) *DescribeRecommendationExportJobsPaginator {
	if params == nil {
		params = &DescribeRecommendationExportJobsInput{}
	}

	options := DescribeRecommendationExportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRecommendationExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRecommendationExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRecommendationExportJobs page.
func (p *DescribeRecommendationExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRecommendationExportJobsOutput, error) {
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

	result, err := p.client.DescribeRecommendationExportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRecommendationExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeRecommendationExportJobs",
	}
}
