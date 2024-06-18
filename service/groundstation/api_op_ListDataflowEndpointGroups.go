// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of DataflowEndpoint groups.
func (c *Client) ListDataflowEndpointGroups(ctx context.Context, params *ListDataflowEndpointGroupsInput, optFns ...func(*Options)) (*ListDataflowEndpointGroupsOutput, error) {
	if params == nil {
		params = &ListDataflowEndpointGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataflowEndpointGroups", params, optFns, c.addOperationListDataflowEndpointGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataflowEndpointGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataflowEndpointGroupsInput struct {

	// Maximum number of dataflow endpoint groups returned.
	MaxResults *int32

	// Next token returned in the request of a previous ListDataflowEndpointGroups
	// call. Used to get the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataflowEndpointGroupsOutput struct {

	// A list of dataflow endpoint groups.
	DataflowEndpointGroupList []types.DataflowEndpointListItem

	// Next token returned in the response of a previous ListDataflowEndpointGroups
	// call. Used to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataflowEndpointGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataflowEndpointGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataflowEndpointGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataflowEndpointGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataflowEndpointGroups(options.Region), middleware.Before); err != nil {
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

// ListDataflowEndpointGroupsPaginatorOptions is the paginator options for
// ListDataflowEndpointGroups
type ListDataflowEndpointGroupsPaginatorOptions struct {
	// Maximum number of dataflow endpoint groups returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataflowEndpointGroupsPaginator is a paginator for
// ListDataflowEndpointGroups
type ListDataflowEndpointGroupsPaginator struct {
	options   ListDataflowEndpointGroupsPaginatorOptions
	client    ListDataflowEndpointGroupsAPIClient
	params    *ListDataflowEndpointGroupsInput
	nextToken *string
	firstPage bool
}

// NewListDataflowEndpointGroupsPaginator returns a new
// ListDataflowEndpointGroupsPaginator
func NewListDataflowEndpointGroupsPaginator(client ListDataflowEndpointGroupsAPIClient, params *ListDataflowEndpointGroupsInput, optFns ...func(*ListDataflowEndpointGroupsPaginatorOptions)) *ListDataflowEndpointGroupsPaginator {
	if params == nil {
		params = &ListDataflowEndpointGroupsInput{}
	}

	options := ListDataflowEndpointGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataflowEndpointGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataflowEndpointGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataflowEndpointGroups page.
func (p *ListDataflowEndpointGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataflowEndpointGroupsOutput, error) {
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
	result, err := p.client.ListDataflowEndpointGroups(ctx, &params, optFns...)
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

// ListDataflowEndpointGroupsAPIClient is a client that implements the
// ListDataflowEndpointGroups operation.
type ListDataflowEndpointGroupsAPIClient interface {
	ListDataflowEndpointGroups(context.Context, *ListDataflowEndpointGroupsInput, ...func(*Options)) (*ListDataflowEndpointGroupsOutput, error)
}

var _ ListDataflowEndpointGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDataflowEndpointGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataflowEndpointGroups",
	}
}
