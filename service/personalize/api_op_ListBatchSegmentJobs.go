// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the batch segment jobs that have been performed off of a
// solution version that you specify.
func (c *Client) ListBatchSegmentJobs(ctx context.Context, params *ListBatchSegmentJobsInput, optFns ...func(*Options)) (*ListBatchSegmentJobsOutput, error) {
	if params == nil {
		params = &ListBatchSegmentJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBatchSegmentJobs", params, optFns, c.addOperationListBatchSegmentJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBatchSegmentJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBatchSegmentJobsInput struct {

	// The maximum number of batch segment job results to return in each page. The
	// default value is 100.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The Amazon Resource Name (ARN) of the solution version that the batch segment
	// jobs used to generate batch segments.
	SolutionVersionArn *string

	noSmithyDocumentSerde
}

type ListBatchSegmentJobsOutput struct {

	// A list containing information on each job that is returned.
	BatchSegmentJobs []types.BatchSegmentJobSummary

	// The token to use to retrieve the next page of results. The value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBatchSegmentJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBatchSegmentJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBatchSegmentJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBatchSegmentJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBatchSegmentJobs(options.Region), middleware.Before); err != nil {
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

// ListBatchSegmentJobsPaginatorOptions is the paginator options for
// ListBatchSegmentJobs
type ListBatchSegmentJobsPaginatorOptions struct {
	// The maximum number of batch segment job results to return in each page. The
	// default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBatchSegmentJobsPaginator is a paginator for ListBatchSegmentJobs
type ListBatchSegmentJobsPaginator struct {
	options   ListBatchSegmentJobsPaginatorOptions
	client    ListBatchSegmentJobsAPIClient
	params    *ListBatchSegmentJobsInput
	nextToken *string
	firstPage bool
}

// NewListBatchSegmentJobsPaginator returns a new ListBatchSegmentJobsPaginator
func NewListBatchSegmentJobsPaginator(client ListBatchSegmentJobsAPIClient, params *ListBatchSegmentJobsInput, optFns ...func(*ListBatchSegmentJobsPaginatorOptions)) *ListBatchSegmentJobsPaginator {
	if params == nil {
		params = &ListBatchSegmentJobsInput{}
	}

	options := ListBatchSegmentJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBatchSegmentJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBatchSegmentJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBatchSegmentJobs page.
func (p *ListBatchSegmentJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBatchSegmentJobsOutput, error) {
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
	result, err := p.client.ListBatchSegmentJobs(ctx, &params, optFns...)
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

// ListBatchSegmentJobsAPIClient is a client that implements the
// ListBatchSegmentJobs operation.
type ListBatchSegmentJobsAPIClient interface {
	ListBatchSegmentJobs(context.Context, *ListBatchSegmentJobsInput, ...func(*Options)) (*ListBatchSegmentJobsOutput, error)
}

var _ ListBatchSegmentJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListBatchSegmentJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBatchSegmentJobs",
	}
}
