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

// Lists all in-progress executions for the specified workflow.
//
// If the specified workflow ID cannot be found, ListExecutions returns a
// ResourceNotFound exception.
func (c *Client) ListExecutions(ctx context.Context, params *ListExecutionsInput, optFns ...func(*Options)) (*ListExecutionsOutput, error) {
	if params == nil {
		params = &ListExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExecutions", params, optFns, c.addOperationListExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExecutionsInput struct {

	// A unique identifier for the workflow.
	//
	// This member is required.
	WorkflowId *string

	// Specifies the maximum number of executions to return.
	MaxResults *int32

	// ListExecutions returns the NextToken parameter in the output. You can then pass
	// the NextToken parameter in a subsequent command to continue listing additional
	// executions.
	//
	// This is useful for pagination, for instance. If you have 100 executions for a
	// workflow, you might only want to list first 10. If so, call the API by
	// specifying the max-results :
	//
	//     aws transfer list-executions --max-results 10
	//
	// This returns details for the first 10 executions, as well as the pointer (
	// NextToken ) to the eleventh execution. You can now call the API again, supplying
	// the NextToken value you received:
	//
	//     aws transfer list-executions --max-results 10 --next-token
	//     $somePointerReturnedFromPreviousListResult
	//
	// This call returns the next 10 executions, the 11th through the 20th. You can
	// then repeat the call until the details for all 100 executions have been
	// returned.
	NextToken *string

	noSmithyDocumentSerde
}

type ListExecutionsOutput struct {

	// Returns the details for each execution, in a ListedExecution array.
	//
	// This member is required.
	Executions []types.ListedExecution

	// A unique identifier for the workflow.
	//
	// This member is required.
	WorkflowId *string

	// ListExecutions returns the NextToken parameter in the output. You can then pass
	// the NextToken parameter in a subsequent command to continue listing additional
	// executions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListExecutions"); err != nil {
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
	if err = addOpListExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExecutions(options.Region), middleware.Before); err != nil {
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

// ListExecutionsPaginatorOptions is the paginator options for ListExecutions
type ListExecutionsPaginatorOptions struct {
	// Specifies the maximum number of executions to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExecutionsPaginator is a paginator for ListExecutions
type ListExecutionsPaginator struct {
	options   ListExecutionsPaginatorOptions
	client    ListExecutionsAPIClient
	params    *ListExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListExecutionsPaginator returns a new ListExecutionsPaginator
func NewListExecutionsPaginator(client ListExecutionsAPIClient, params *ListExecutionsInput, optFns ...func(*ListExecutionsPaginatorOptions)) *ListExecutionsPaginator {
	if params == nil {
		params = &ListExecutionsInput{}
	}

	options := ListExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExecutions page.
func (p *ListExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExecutionsOutput, error) {
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
	result, err := p.client.ListExecutions(ctx, &params, optFns...)
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

// ListExecutionsAPIClient is a client that implements the ListExecutions
// operation.
type ListExecutionsAPIClient interface {
	ListExecutions(context.Context, *ListExecutionsInput, ...func(*Options)) (*ListExecutionsOutput, error)
}

var _ ListExecutionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListExecutions",
	}
}
