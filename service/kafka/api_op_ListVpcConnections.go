// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the VPC connections in this Region.
func (c *Client) ListVpcConnections(ctx context.Context, params *ListVpcConnectionsInput, optFns ...func(*Options)) (*ListVpcConnectionsOutput, error) {
	if params == nil {
		params = &ListVpcConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVpcConnections", params, optFns, c.addOperationListVpcConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVpcConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVpcConnectionsInput struct {

	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	MaxResults *int32

	// The paginated results marker. When the result of the operation is truncated,
	// the call returns NextToken in the response. To get the next batch, provide this
	// token in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVpcConnectionsOutput struct {

	// The paginated results marker. When the result of a ListClientVpcConnections
	// operation is truncated, the call returns NextToken in the response. To get
	// another batch of configurations, provide this token in your next request.
	NextToken *string

	// List of VPC connections.
	VpcConnections []types.VpcConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVpcConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVpcConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVpcConnections{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListVpcConnections"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVpcConnections(options.Region), middleware.Before); err != nil {
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

// ListVpcConnectionsAPIClient is a client that implements the ListVpcConnections
// operation.
type ListVpcConnectionsAPIClient interface {
	ListVpcConnections(context.Context, *ListVpcConnectionsInput, ...func(*Options)) (*ListVpcConnectionsOutput, error)
}

var _ ListVpcConnectionsAPIClient = (*Client)(nil)

// ListVpcConnectionsPaginatorOptions is the paginator options for
// ListVpcConnections
type ListVpcConnectionsPaginatorOptions struct {
	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVpcConnectionsPaginator is a paginator for ListVpcConnections
type ListVpcConnectionsPaginator struct {
	options   ListVpcConnectionsPaginatorOptions
	client    ListVpcConnectionsAPIClient
	params    *ListVpcConnectionsInput
	nextToken *string
	firstPage bool
}

// NewListVpcConnectionsPaginator returns a new ListVpcConnectionsPaginator
func NewListVpcConnectionsPaginator(client ListVpcConnectionsAPIClient, params *ListVpcConnectionsInput, optFns ...func(*ListVpcConnectionsPaginatorOptions)) *ListVpcConnectionsPaginator {
	if params == nil {
		params = &ListVpcConnectionsInput{}
	}

	options := ListVpcConnectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVpcConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVpcConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVpcConnections page.
func (p *ListVpcConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVpcConnectionsOutput, error) {
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

	result, err := p.client.ListVpcConnections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVpcConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListVpcConnections",
	}
}
