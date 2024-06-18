// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the audience export jobs.
func (c *Client) ListAudienceExportJobs(ctx context.Context, params *ListAudienceExportJobsInput, optFns ...func(*Options)) (*ListAudienceExportJobsOutput, error) {
	if params == nil {
		params = &ListAudienceExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAudienceExportJobs", params, optFns, c.addOperationListAudienceExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAudienceExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAudienceExportJobsInput struct {

	// The Amazon Resource Name (ARN) of the audience generation job that you are
	// interested in.
	AudienceGenerationJobArn *string

	// The maximum size of the results that is returned per call.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAudienceExportJobsOutput struct {

	// The audience export jobs that match the request.
	//
	// This member is required.
	AudienceExportJobs []types.AudienceExportJobSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAudienceExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAudienceExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAudienceExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAudienceExportJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAudienceExportJobs(options.Region), middleware.Before); err != nil {
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

// ListAudienceExportJobsPaginatorOptions is the paginator options for
// ListAudienceExportJobs
type ListAudienceExportJobsPaginatorOptions struct {
	// The maximum size of the results that is returned per call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAudienceExportJobsPaginator is a paginator for ListAudienceExportJobs
type ListAudienceExportJobsPaginator struct {
	options   ListAudienceExportJobsPaginatorOptions
	client    ListAudienceExportJobsAPIClient
	params    *ListAudienceExportJobsInput
	nextToken *string
	firstPage bool
}

// NewListAudienceExportJobsPaginator returns a new ListAudienceExportJobsPaginator
func NewListAudienceExportJobsPaginator(client ListAudienceExportJobsAPIClient, params *ListAudienceExportJobsInput, optFns ...func(*ListAudienceExportJobsPaginatorOptions)) *ListAudienceExportJobsPaginator {
	if params == nil {
		params = &ListAudienceExportJobsInput{}
	}

	options := ListAudienceExportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAudienceExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAudienceExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAudienceExportJobs page.
func (p *ListAudienceExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAudienceExportJobsOutput, error) {
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
	result, err := p.client.ListAudienceExportJobs(ctx, &params, optFns...)
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

// ListAudienceExportJobsAPIClient is a client that implements the
// ListAudienceExportJobs operation.
type ListAudienceExportJobsAPIClient interface {
	ListAudienceExportJobs(context.Context, *ListAudienceExportJobsInput, ...func(*Options)) (*ListAudienceExportJobsOutput, error)
}

var _ ListAudienceExportJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAudienceExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAudienceExportJobs",
	}
}
