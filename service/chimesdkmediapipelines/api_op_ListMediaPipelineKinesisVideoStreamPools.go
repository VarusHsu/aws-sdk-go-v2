// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the video stream pools in the media pipeline.
func (c *Client) ListMediaPipelineKinesisVideoStreamPools(ctx context.Context, params *ListMediaPipelineKinesisVideoStreamPoolsInput, optFns ...func(*Options)) (*ListMediaPipelineKinesisVideoStreamPoolsOutput, error) {
	if params == nil {
		params = &ListMediaPipelineKinesisVideoStreamPoolsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMediaPipelineKinesisVideoStreamPools", params, optFns, c.addOperationListMediaPipelineKinesisVideoStreamPoolsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMediaPipelineKinesisVideoStreamPoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMediaPipelineKinesisVideoStreamPoolsInput struct {

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token used to return the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMediaPipelineKinesisVideoStreamPoolsOutput struct {

	// The list of video stream pools.
	KinesisVideoStreamPools []types.KinesisVideoStreamPoolSummary

	// The token used to return the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMediaPipelineKinesisVideoStreamPoolsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMediaPipelineKinesisVideoStreamPools{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMediaPipelineKinesisVideoStreamPools{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMediaPipelineKinesisVideoStreamPools"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMediaPipelineKinesisVideoStreamPools(options.Region), middleware.Before); err != nil {
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

// ListMediaPipelineKinesisVideoStreamPoolsPaginatorOptions is the paginator
// options for ListMediaPipelineKinesisVideoStreamPools
type ListMediaPipelineKinesisVideoStreamPoolsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMediaPipelineKinesisVideoStreamPoolsPaginator is a paginator for
// ListMediaPipelineKinesisVideoStreamPools
type ListMediaPipelineKinesisVideoStreamPoolsPaginator struct {
	options   ListMediaPipelineKinesisVideoStreamPoolsPaginatorOptions
	client    ListMediaPipelineKinesisVideoStreamPoolsAPIClient
	params    *ListMediaPipelineKinesisVideoStreamPoolsInput
	nextToken *string
	firstPage bool
}

// NewListMediaPipelineKinesisVideoStreamPoolsPaginator returns a new
// ListMediaPipelineKinesisVideoStreamPoolsPaginator
func NewListMediaPipelineKinesisVideoStreamPoolsPaginator(client ListMediaPipelineKinesisVideoStreamPoolsAPIClient, params *ListMediaPipelineKinesisVideoStreamPoolsInput, optFns ...func(*ListMediaPipelineKinesisVideoStreamPoolsPaginatorOptions)) *ListMediaPipelineKinesisVideoStreamPoolsPaginator {
	if params == nil {
		params = &ListMediaPipelineKinesisVideoStreamPoolsInput{}
	}

	options := ListMediaPipelineKinesisVideoStreamPoolsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMediaPipelineKinesisVideoStreamPoolsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMediaPipelineKinesisVideoStreamPoolsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMediaPipelineKinesisVideoStreamPools page.
func (p *ListMediaPipelineKinesisVideoStreamPoolsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMediaPipelineKinesisVideoStreamPoolsOutput, error) {
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
	result, err := p.client.ListMediaPipelineKinesisVideoStreamPools(ctx, &params, optFns...)
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

// ListMediaPipelineKinesisVideoStreamPoolsAPIClient is a client that implements
// the ListMediaPipelineKinesisVideoStreamPools operation.
type ListMediaPipelineKinesisVideoStreamPoolsAPIClient interface {
	ListMediaPipelineKinesisVideoStreamPools(context.Context, *ListMediaPipelineKinesisVideoStreamPoolsInput, ...func(*Options)) (*ListMediaPipelineKinesisVideoStreamPoolsOutput, error)
}

var _ ListMediaPipelineKinesisVideoStreamPoolsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListMediaPipelineKinesisVideoStreamPools(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMediaPipelineKinesisVideoStreamPools",
	}
}
