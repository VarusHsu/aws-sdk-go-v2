// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backupgateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your hypervisors.
func (c *Client) ListHypervisors(ctx context.Context, params *ListHypervisorsInput, optFns ...func(*Options)) (*ListHypervisorsOutput, error) {
	if params == nil {
		params = &ListHypervisorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHypervisors", params, optFns, c.addOperationListHypervisorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHypervisorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHypervisorsInput struct {

	// The maximum number of hypervisors to list.
	MaxResults *int32

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return maxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListHypervisorsOutput struct {

	// A list of your Hypervisor objects, ordered by their Amazon Resource Names
	// (ARNs).
	Hypervisors []types.Hypervisor

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return maxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHypervisorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListHypervisors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListHypervisors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListHypervisors"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHypervisors(options.Region), middleware.Before); err != nil {
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

// ListHypervisorsPaginatorOptions is the paginator options for ListHypervisors
type ListHypervisorsPaginatorOptions struct {
	// The maximum number of hypervisors to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHypervisorsPaginator is a paginator for ListHypervisors
type ListHypervisorsPaginator struct {
	options   ListHypervisorsPaginatorOptions
	client    ListHypervisorsAPIClient
	params    *ListHypervisorsInput
	nextToken *string
	firstPage bool
}

// NewListHypervisorsPaginator returns a new ListHypervisorsPaginator
func NewListHypervisorsPaginator(client ListHypervisorsAPIClient, params *ListHypervisorsInput, optFns ...func(*ListHypervisorsPaginatorOptions)) *ListHypervisorsPaginator {
	if params == nil {
		params = &ListHypervisorsInput{}
	}

	options := ListHypervisorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHypervisorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHypervisorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHypervisors page.
func (p *ListHypervisorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHypervisorsOutput, error) {
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
	result, err := p.client.ListHypervisors(ctx, &params, optFns...)
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

// ListHypervisorsAPIClient is a client that implements the ListHypervisors
// operation.
type ListHypervisorsAPIClient interface {
	ListHypervisors(context.Context, *ListHypervisorsInput, ...func(*Options)) (*ListHypervisorsOutput, error)
}

var _ ListHypervisorsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListHypervisors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListHypervisors",
	}
}
