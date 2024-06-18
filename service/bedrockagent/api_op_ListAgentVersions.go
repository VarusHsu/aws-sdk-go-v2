// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagent

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the versions of an agent and information about each version.
func (c *Client) ListAgentVersions(ctx context.Context, params *ListAgentVersionsInput, optFns ...func(*Options)) (*ListAgentVersionsOutput, error) {
	if params == nil {
		params = &ListAgentVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAgentVersions", params, optFns, c.addOperationListAgentVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAgentVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAgentVersionsInput struct {

	// The unique identifier of the agent.
	//
	// This member is required.
	AgentId *string

	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	MaxResults *int32

	// If the total number of results is greater than the maxResults value provided in
	// the request, enter the token returned in the nextToken field in the response in
	// this field to return the next batch of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAgentVersionsOutput struct {

	// A list of objects, each of which contains information about a version of the
	// agent.
	//
	// This member is required.
	AgentVersionSummaries []types.AgentVersionSummary

	// If the total number of results is greater than the maxResults value provided in
	// the request, use this token when making another request in the nextToken field
	// to return the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAgentVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAgentVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAgentVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAgentVersions"); err != nil {
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
	if err = addOpListAgentVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAgentVersions(options.Region), middleware.Before); err != nil {
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

// ListAgentVersionsPaginatorOptions is the paginator options for ListAgentVersions
type ListAgentVersionsPaginatorOptions struct {
	// The maximum number of results to return in the response. If the total number of
	// results is greater than this value, use the token returned in the response in
	// the nextToken field when making another request to return the next batch of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAgentVersionsPaginator is a paginator for ListAgentVersions
type ListAgentVersionsPaginator struct {
	options   ListAgentVersionsPaginatorOptions
	client    ListAgentVersionsAPIClient
	params    *ListAgentVersionsInput
	nextToken *string
	firstPage bool
}

// NewListAgentVersionsPaginator returns a new ListAgentVersionsPaginator
func NewListAgentVersionsPaginator(client ListAgentVersionsAPIClient, params *ListAgentVersionsInput, optFns ...func(*ListAgentVersionsPaginatorOptions)) *ListAgentVersionsPaginator {
	if params == nil {
		params = &ListAgentVersionsInput{}
	}

	options := ListAgentVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAgentVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAgentVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAgentVersions page.
func (p *ListAgentVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAgentVersionsOutput, error) {
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
	result, err := p.client.ListAgentVersions(ctx, &params, optFns...)
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

// ListAgentVersionsAPIClient is a client that implements the ListAgentVersions
// operation.
type ListAgentVersionsAPIClient interface {
	ListAgentVersions(context.Context, *ListAgentVersionsInput, ...func(*Options)) (*ListAgentVersionsOutput, error)
}

var _ ListAgentVersionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAgentVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAgentVersions",
	}
}
