// Code generated by smithy-go-codegen DO NOT EDIT.

package apptest

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apptest/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists test runs.
func (c *Client) ListTestRuns(ctx context.Context, params *ListTestRunsInput, optFns ...func(*Options)) (*ListTestRunsOutput, error) {
	if params == nil {
		params = &ListTestRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTestRuns", params, optFns, c.addOperationListTestRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTestRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestRunsInput struct {

	// The maximum number of test runs to return in one page of results.
	MaxResults *int32

	// The token from the previous request to retrieve the next page of test run
	// results.
	NextToken *string

	// The test run IDs of the test runs.
	TestRunIds []string

	// The test suite ID of the test runs.
	TestSuiteId *string

	noSmithyDocumentSerde
}

type ListTestRunsOutput struct {

	// The test runs of the response query.
	//
	// This member is required.
	TestRuns []types.TestRunSummary

	// The token from the previous request to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTestRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTestRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTestRuns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTestRuns"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTestRuns(options.Region), middleware.Before); err != nil {
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

// ListTestRunsPaginatorOptions is the paginator options for ListTestRuns
type ListTestRunsPaginatorOptions struct {
	// The maximum number of test runs to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTestRunsPaginator is a paginator for ListTestRuns
type ListTestRunsPaginator struct {
	options   ListTestRunsPaginatorOptions
	client    ListTestRunsAPIClient
	params    *ListTestRunsInput
	nextToken *string
	firstPage bool
}

// NewListTestRunsPaginator returns a new ListTestRunsPaginator
func NewListTestRunsPaginator(client ListTestRunsAPIClient, params *ListTestRunsInput, optFns ...func(*ListTestRunsPaginatorOptions)) *ListTestRunsPaginator {
	if params == nil {
		params = &ListTestRunsInput{}
	}

	options := ListTestRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTestRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTestRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTestRuns page.
func (p *ListTestRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTestRunsOutput, error) {
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
	result, err := p.client.ListTestRuns(ctx, &params, optFns...)
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

// ListTestRunsAPIClient is a client that implements the ListTestRuns operation.
type ListTestRunsAPIClient interface {
	ListTestRuns(context.Context, *ListTestRunsInput, ...func(*Options)) (*ListTestRunsOutput, error)
}

var _ ListTestRunsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTestRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTestRuns",
	}
}
