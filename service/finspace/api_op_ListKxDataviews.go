// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the dataviews in the database.
func (c *Client) ListKxDataviews(ctx context.Context, params *ListKxDataviewsInput, optFns ...func(*Options)) (*ListKxDataviewsOutput, error) {
	if params == nil {
		params = &ListKxDataviewsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKxDataviews", params, optFns, c.addOperationListKxDataviewsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKxDataviewsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKxDataviewsInput struct {

	//  The name of the database where the dataviews were created.
	//
	// This member is required.
	DatabaseName *string

	// A unique identifier for the kdb environment, for which you want to retrieve a
	// list of dataviews.
	//
	// This member is required.
	EnvironmentId *string

	// The maximum number of results to return in this request.
	MaxResults int32

	//  A token that indicates where a results page should begin.
	NextToken *string

	noSmithyDocumentSerde
}

type ListKxDataviewsOutput struct {

	//  The list of kdb dataviews that are currently active for the given database.
	KxDataviews []types.KxDataviewListEntry

	//  A token that indicates where a results page should begin.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKxDataviewsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKxDataviews{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKxDataviews{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListKxDataviews"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpListKxDataviewsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKxDataviews(options.Region), middleware.Before); err != nil {
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

// ListKxDataviewsPaginatorOptions is the paginator options for ListKxDataviews
type ListKxDataviewsPaginatorOptions struct {
	// The maximum number of results to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKxDataviewsPaginator is a paginator for ListKxDataviews
type ListKxDataviewsPaginator struct {
	options   ListKxDataviewsPaginatorOptions
	client    ListKxDataviewsAPIClient
	params    *ListKxDataviewsInput
	nextToken *string
	firstPage bool
}

// NewListKxDataviewsPaginator returns a new ListKxDataviewsPaginator
func NewListKxDataviewsPaginator(client ListKxDataviewsAPIClient, params *ListKxDataviewsInput, optFns ...func(*ListKxDataviewsPaginatorOptions)) *ListKxDataviewsPaginator {
	if params == nil {
		params = &ListKxDataviewsInput{}
	}

	options := ListKxDataviewsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKxDataviewsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKxDataviewsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKxDataviews page.
func (p *ListKxDataviewsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKxDataviewsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListKxDataviews(ctx, &params, optFns...)
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

// ListKxDataviewsAPIClient is a client that implements the ListKxDataviews
// operation.
type ListKxDataviewsAPIClient interface {
	ListKxDataviews(context.Context, *ListKxDataviewsInput, ...func(*Options)) (*ListKxDataviewsOutput, error)
}

var _ ListKxDataviewsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListKxDataviews(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListKxDataviews",
	}
}
