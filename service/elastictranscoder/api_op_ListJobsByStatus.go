// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListJobsByStatus operation gets a list of jobs that have a specified
// status. The response body contains one element for each job that satisfies the
// search criteria.
func (c *Client) ListJobsByStatus(ctx context.Context, params *ListJobsByStatusInput, optFns ...func(*Options)) (*ListJobsByStatusOutput, error) {
	if params == nil {
		params = &ListJobsByStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobsByStatus", params, optFns, c.addOperationListJobsByStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobsByStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListJobsByStatusRequest structure.
type ListJobsByStatusInput struct {

	// To get information about all of the jobs associated with the current AWS
	// account that have a given status, specify the following status: Submitted ,
	// Progressing , Complete , Canceled , or Error .
	//
	// This member is required.
	Status *string

	//  To list jobs in chronological order by the date and time that they were
	// submitted, enter true . To list jobs in reverse chronological order, enter false
	// .
	Ascending *string

	//  When Elastic Transcoder returns more than one page of results, use pageToken
	// in subsequent GET requests to get each successive page of results.
	PageToken *string

	noSmithyDocumentSerde
}

// The ListJobsByStatusResponse structure.
type ListJobsByStatusOutput struct {

	// An array of Job objects that have the specified status.
	Jobs []types.Job

	//  A value that you use to access the second and subsequent pages of results, if
	// any. When the jobs in the specified pipeline fit on one page or when you've
	// reached the last page of results, the value of NextPageToken is null .
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJobsByStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobsByStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobsByStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListJobsByStatus"); err != nil {
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
	if err = addOpListJobsByStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobsByStatus(options.Region), middleware.Before); err != nil {
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

// ListJobsByStatusPaginatorOptions is the paginator options for ListJobsByStatus
type ListJobsByStatusPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobsByStatusPaginator is a paginator for ListJobsByStatus
type ListJobsByStatusPaginator struct {
	options   ListJobsByStatusPaginatorOptions
	client    ListJobsByStatusAPIClient
	params    *ListJobsByStatusInput
	nextToken *string
	firstPage bool
}

// NewListJobsByStatusPaginator returns a new ListJobsByStatusPaginator
func NewListJobsByStatusPaginator(client ListJobsByStatusAPIClient, params *ListJobsByStatusInput, optFns ...func(*ListJobsByStatusPaginatorOptions)) *ListJobsByStatusPaginator {
	if params == nil {
		params = &ListJobsByStatusInput{}
	}

	options := ListJobsByStatusPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJobsByStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobsByStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListJobsByStatus page.
func (p *ListJobsByStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobsByStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListJobsByStatus(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListJobsByStatusAPIClient is a client that implements the ListJobsByStatus
// operation.
type ListJobsByStatusAPIClient interface {
	ListJobsByStatus(context.Context, *ListJobsByStatusInput, ...func(*Options)) (*ListJobsByStatusOutput, error)
}

var _ ListJobsByStatusAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListJobsByStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListJobsByStatus",
	}
}
