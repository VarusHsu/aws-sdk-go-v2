// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all workflows associated with your Amazon Web Services account for your
// current region.
func (c *Client) ListWorkflows(ctx context.Context, params *ListWorkflowsInput, optFns ...func(*Options)) (*ListWorkflowsOutput, error) {
	if params == nil {
		params = &ListWorkflowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkflows", params, optFns, c.addOperationListWorkflowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkflowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkflowsInput struct {

	// Specifies the maximum number of workflows to return.
	MaxResults *int32

	// ListWorkflows returns the NextToken parameter in the output. You can then pass
	// the NextToken parameter in a subsequent command to continue listing additional
	// workflows.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkflowsOutput struct {

	// Returns the Arn , WorkflowId , and Description for each workflow.
	//
	// This member is required.
	Workflows []types.ListedWorkflow

	// ListWorkflows returns the NextToken parameter in the output. You can then pass
	// the NextToken parameter in a subsequent command to continue listing additional
	// workflows.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkflowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkflows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkflows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkflows"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkflows(options.Region), middleware.Before); err != nil {
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

// ListWorkflowsPaginatorOptions is the paginator options for ListWorkflows
type ListWorkflowsPaginatorOptions struct {
	// Specifies the maximum number of workflows to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkflowsPaginator is a paginator for ListWorkflows
type ListWorkflowsPaginator struct {
	options   ListWorkflowsPaginatorOptions
	client    ListWorkflowsAPIClient
	params    *ListWorkflowsInput
	nextToken *string
	firstPage bool
}

// NewListWorkflowsPaginator returns a new ListWorkflowsPaginator
func NewListWorkflowsPaginator(client ListWorkflowsAPIClient, params *ListWorkflowsInput, optFns ...func(*ListWorkflowsPaginatorOptions)) *ListWorkflowsPaginator {
	if params == nil {
		params = &ListWorkflowsInput{}
	}

	options := ListWorkflowsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkflowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkflowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkflows page.
func (p *ListWorkflowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkflowsOutput, error) {
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
	result, err := p.client.ListWorkflows(ctx, &params, optFns...)
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

// ListWorkflowsAPIClient is a client that implements the ListWorkflows operation.
type ListWorkflowsAPIClient interface {
	ListWorkflows(context.Context, *ListWorkflowsInput, ...func(*Options)) (*ListWorkflowsOutput, error)
}

var _ ListWorkflowsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListWorkflows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkflows",
	}
}
