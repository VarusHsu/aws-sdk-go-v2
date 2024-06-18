// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List CloudFront distributions.
func (c *Client) ListDistributions(ctx context.Context, params *ListDistributionsInput, optFns ...func(*Options)) (*ListDistributionsOutput, error) {
	if params == nil {
		params = &ListDistributionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDistributions", params, optFns, c.addOperationListDistributionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDistributionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list your distributions.
type ListDistributionsInput struct {

	// Use this when paginating results to indicate where to begin in your list of
	// distributions. The results include distributions in the list that occur after
	// the marker. To get the next page of results, set the Marker to the value of the
	// NextMarker from the current page's response (which is also the ID of the last
	// distribution on that page).
	Marker *string

	// The maximum number of distributions you want in the response body.
	MaxItems *int32

	noSmithyDocumentSerde
}

// The returned result of the corresponding request.
type ListDistributionsOutput struct {

	// The DistributionList type.
	DistributionList *types.DistributionList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDistributionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListDistributions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListDistributions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDistributions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDistributions(options.Region), middleware.Before); err != nil {
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

// ListDistributionsPaginatorOptions is the paginator options for ListDistributions
type ListDistributionsPaginatorOptions struct {
	// The maximum number of distributions you want in the response body.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDistributionsPaginator is a paginator for ListDistributions
type ListDistributionsPaginator struct {
	options   ListDistributionsPaginatorOptions
	client    ListDistributionsAPIClient
	params    *ListDistributionsInput
	nextToken *string
	firstPage bool
}

// NewListDistributionsPaginator returns a new ListDistributionsPaginator
func NewListDistributionsPaginator(client ListDistributionsAPIClient, params *ListDistributionsInput, optFns ...func(*ListDistributionsPaginatorOptions)) *ListDistributionsPaginator {
	if params == nil {
		params = &ListDistributionsInput{}
	}

	options := ListDistributionsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDistributionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDistributionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDistributions page.
func (p *ListDistributionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDistributionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListDistributions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = nil
	if result.DistributionList != nil {
		p.nextToken = result.DistributionList.NextMarker
	}

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListDistributionsAPIClient is a client that implements the ListDistributions
// operation.
type ListDistributionsAPIClient interface {
	ListDistributions(context.Context, *ListDistributionsInput, ...func(*Options)) (*ListDistributionsOutput, error)
}

var _ ListDistributionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDistributions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDistributions",
	}
}
