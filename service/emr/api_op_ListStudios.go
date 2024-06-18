// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all Amazon EMR Studios associated with the Amazon Web
// Services account. The list includes details such as ID, Studio Access URL, and
// creation time for each Studio.
func (c *Client) ListStudios(ctx context.Context, params *ListStudiosInput, optFns ...func(*Options)) (*ListStudiosOutput, error) {
	if params == nil {
		params = &ListStudiosInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStudios", params, optFns, c.addOperationListStudiosMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStudiosOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStudiosInput struct {

	// The pagination token that indicates the set of results to retrieve.
	Marker *string

	noSmithyDocumentSerde
}

type ListStudiosOutput struct {

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// The list of Studio summary objects.
	Studios []types.StudioSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStudiosMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStudios{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStudios{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStudios"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStudios(options.Region), middleware.Before); err != nil {
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

// ListStudiosPaginatorOptions is the paginator options for ListStudios
type ListStudiosPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStudiosPaginator is a paginator for ListStudios
type ListStudiosPaginator struct {
	options   ListStudiosPaginatorOptions
	client    ListStudiosAPIClient
	params    *ListStudiosInput
	nextToken *string
	firstPage bool
}

// NewListStudiosPaginator returns a new ListStudiosPaginator
func NewListStudiosPaginator(client ListStudiosAPIClient, params *ListStudiosInput, optFns ...func(*ListStudiosPaginatorOptions)) *ListStudiosPaginator {
	if params == nil {
		params = &ListStudiosInput{}
	}

	options := ListStudiosPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStudiosPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStudiosPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStudios page.
func (p *ListStudiosPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStudiosOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListStudios(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListStudiosAPIClient is a client that implements the ListStudios operation.
type ListStudiosAPIClient interface {
	ListStudios(context.Context, *ListStudiosInput, ...func(*Options)) (*ListStudiosOutput, error)
}

var _ ListStudiosAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListStudios(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStudios",
	}
}
