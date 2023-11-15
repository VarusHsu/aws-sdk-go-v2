// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists training jobs. When StatusEquals and MaxResults are set at the same time,
// the MaxResults number of training jobs are first retrieved ignoring the
// StatusEquals parameter and then they are filtered by the StatusEquals
// parameter, which is returned as a response. For example, if ListTrainingJobs is
// invoked with the following parameters: { ... MaxResults: 100, StatusEquals:
// InProgress ... } First, 100 trainings jobs with any status, including those
// other than InProgress , are selected (sorted according to the creation time,
// from the most current to the oldest). Next, those with a status of InProgress
// are returned. You can quickly test the API using the following Amazon Web
// Services CLI code. aws sagemaker list-training-jobs --max-results 100
// --status-equals InProgress
func (c *Client) ListTrainingJobs(ctx context.Context, params *ListTrainingJobsInput, optFns ...func(*Options)) (*ListTrainingJobsOutput, error) {
	if params == nil {
		params = &ListTrainingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrainingJobs", params, optFns, c.addOperationListTrainingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrainingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrainingJobsInput struct {

	// A filter that returns only training jobs created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only training jobs created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// A filter that returns only training jobs modified after the specified time
	// (timestamp).
	LastModifiedTimeAfter *time.Time

	// A filter that returns only training jobs modified before the specified time
	// (timestamp).
	LastModifiedTimeBefore *time.Time

	// The maximum number of training jobs to return in the response.
	MaxResults *int32

	// A string in the training job name. This filter returns only training jobs whose
	// name contains the specified string.
	NameContains *string

	// If the result of the previous ListTrainingJobs request was truncated, the
	// response includes a NextToken . To retrieve the next set of training jobs, use
	// the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime .
	SortBy types.SortBy

	// The sort order for results. The default is Ascending .
	SortOrder types.SortOrder

	// A filter that retrieves only training jobs with a specific status.
	StatusEquals types.TrainingJobStatus

	// A filter that retrieves only training jobs with a specific warm pool status.
	WarmPoolStatusEquals types.WarmPoolResourceStatus

	noSmithyDocumentSerde
}

type ListTrainingJobsOutput struct {

	// An array of TrainingJobSummary objects, each listing a training job.
	//
	// This member is required.
	TrainingJobSummaries []types.TrainingJobSummary

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of training jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrainingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTrainingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrainingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTrainingJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrainingJobs(options.Region), middleware.Before); err != nil {
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

// ListTrainingJobsAPIClient is a client that implements the ListTrainingJobs
// operation.
type ListTrainingJobsAPIClient interface {
	ListTrainingJobs(context.Context, *ListTrainingJobsInput, ...func(*Options)) (*ListTrainingJobsOutput, error)
}

var _ ListTrainingJobsAPIClient = (*Client)(nil)

// ListTrainingJobsPaginatorOptions is the paginator options for ListTrainingJobs
type ListTrainingJobsPaginatorOptions struct {
	// The maximum number of training jobs to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTrainingJobsPaginator is a paginator for ListTrainingJobs
type ListTrainingJobsPaginator struct {
	options   ListTrainingJobsPaginatorOptions
	client    ListTrainingJobsAPIClient
	params    *ListTrainingJobsInput
	nextToken *string
	firstPage bool
}

// NewListTrainingJobsPaginator returns a new ListTrainingJobsPaginator
func NewListTrainingJobsPaginator(client ListTrainingJobsAPIClient, params *ListTrainingJobsInput, optFns ...func(*ListTrainingJobsPaginatorOptions)) *ListTrainingJobsPaginator {
	if params == nil {
		params = &ListTrainingJobsInput{}
	}

	options := ListTrainingJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTrainingJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTrainingJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTrainingJobs page.
func (p *ListTrainingJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTrainingJobsOutput, error) {
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

	result, err := p.client.ListTrainingJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTrainingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTrainingJobs",
	}
}
