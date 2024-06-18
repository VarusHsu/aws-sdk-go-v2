// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a service graph for one or more specific trace IDs.
func (c *Client) GetTraceGraph(ctx context.Context, params *GetTraceGraphInput, optFns ...func(*Options)) (*GetTraceGraphOutput, error) {
	if params == nil {
		params = &GetTraceGraphInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTraceGraph", params, optFns, c.addOperationGetTraceGraphMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTraceGraphOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTraceGraphInput struct {

	// Trace IDs of requests for which to generate a service graph.
	//
	// This member is required.
	TraceIds []string

	// Pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type GetTraceGraphOutput struct {

	// Pagination token.
	NextToken *string

	// The services that have processed one of the specified requests.
	Services []types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTraceGraphMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTraceGraph{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTraceGraph{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTraceGraph"); err != nil {
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
	if err = addOpGetTraceGraphValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTraceGraph(options.Region), middleware.Before); err != nil {
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

// GetTraceGraphPaginatorOptions is the paginator options for GetTraceGraph
type GetTraceGraphPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTraceGraphPaginator is a paginator for GetTraceGraph
type GetTraceGraphPaginator struct {
	options   GetTraceGraphPaginatorOptions
	client    GetTraceGraphAPIClient
	params    *GetTraceGraphInput
	nextToken *string
	firstPage bool
}

// NewGetTraceGraphPaginator returns a new GetTraceGraphPaginator
func NewGetTraceGraphPaginator(client GetTraceGraphAPIClient, params *GetTraceGraphInput, optFns ...func(*GetTraceGraphPaginatorOptions)) *GetTraceGraphPaginator {
	if params == nil {
		params = &GetTraceGraphInput{}
	}

	options := GetTraceGraphPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTraceGraphPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTraceGraphPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTraceGraph page.
func (p *GetTraceGraphPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTraceGraphOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetTraceGraph(ctx, &params, optFns...)
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

// GetTraceGraphAPIClient is a client that implements the GetTraceGraph operation.
type GetTraceGraphAPIClient interface {
	GetTraceGraph(context.Context, *GetTraceGraphInput, ...func(*Options)) (*GetTraceGraphOutput, error)
}

var _ GetTraceGraphAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetTraceGraph(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTraceGraph",
	}
}
