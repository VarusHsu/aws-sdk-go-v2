// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists private endpoints for a specified Neptune Analytics graph.
func (c *Client) ListPrivateGraphEndpoints(ctx context.Context, params *ListPrivateGraphEndpointsInput, optFns ...func(*Options)) (*ListPrivateGraphEndpointsOutput, error) {
	if params == nil {
		params = &ListPrivateGraphEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPrivateGraphEndpoints", params, optFns, c.addOperationListPrivateGraphEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPrivateGraphEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrivateGraphEndpointsInput struct {

	// The unique identifier of the Neptune Analytics graph.
	//
	// This member is required.
	GraphIdentifier *string

	// The total number of records to return in the command's output.
	//
	// If the total number of records available is more than the value specified,
	// nextToken is provided in the command's output. To resume pagination, provide the
	// nextToken output value in the nextToken argument of a subsequent command. Do
	// not use the nextToken response element directly outside of the Amazon CLI.
	MaxResults *int32

	// Pagination token used to paginate output.
	//
	// When this value is provided as input, the service returns results from where
	// the previous response left off. When this value is present in output, it
	// indicates that there are more results to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

func (in *ListPrivateGraphEndpointsInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type ListPrivateGraphEndpointsOutput struct {

	// A list of private endpoints for the specified Neptune Analytics graph.
	//
	// This member is required.
	PrivateGraphEndpoints []types.PrivateGraphEndpointSummary

	// Pagination token used to paginate output.
	//
	// When this value is provided as input, the service returns results from where
	// the previous response left off. When this value is present in output, it
	// indicates that there are more results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPrivateGraphEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPrivateGraphEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrivateGraphEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPrivateGraphEndpoints"); err != nil {
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
	if err = addOpListPrivateGraphEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrivateGraphEndpoints(options.Region), middleware.Before); err != nil {
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

// ListPrivateGraphEndpointsPaginatorOptions is the paginator options for
// ListPrivateGraphEndpoints
type ListPrivateGraphEndpointsPaginatorOptions struct {
	// The total number of records to return in the command's output.
	//
	// If the total number of records available is more than the value specified,
	// nextToken is provided in the command's output. To resume pagination, provide the
	// nextToken output value in the nextToken argument of a subsequent command. Do
	// not use the nextToken response element directly outside of the Amazon CLI.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPrivateGraphEndpointsPaginator is a paginator for ListPrivateGraphEndpoints
type ListPrivateGraphEndpointsPaginator struct {
	options   ListPrivateGraphEndpointsPaginatorOptions
	client    ListPrivateGraphEndpointsAPIClient
	params    *ListPrivateGraphEndpointsInput
	nextToken *string
	firstPage bool
}

// NewListPrivateGraphEndpointsPaginator returns a new
// ListPrivateGraphEndpointsPaginator
func NewListPrivateGraphEndpointsPaginator(client ListPrivateGraphEndpointsAPIClient, params *ListPrivateGraphEndpointsInput, optFns ...func(*ListPrivateGraphEndpointsPaginatorOptions)) *ListPrivateGraphEndpointsPaginator {
	if params == nil {
		params = &ListPrivateGraphEndpointsInput{}
	}

	options := ListPrivateGraphEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPrivateGraphEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPrivateGraphEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPrivateGraphEndpoints page.
func (p *ListPrivateGraphEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPrivateGraphEndpointsOutput, error) {
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
	result, err := p.client.ListPrivateGraphEndpoints(ctx, &params, optFns...)
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

// ListPrivateGraphEndpointsAPIClient is a client that implements the
// ListPrivateGraphEndpoints operation.
type ListPrivateGraphEndpointsAPIClient interface {
	ListPrivateGraphEndpoints(context.Context, *ListPrivateGraphEndpointsInput, ...func(*Options)) (*ListPrivateGraphEndpointsOutput, error)
}

var _ ListPrivateGraphEndpointsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPrivateGraphEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPrivateGraphEndpoints",
	}
}
