// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches existing Amazon Q in Connect quick responses in an Amazon Q in Connect
// knowledge base.
func (c *Client) SearchQuickResponses(ctx context.Context, params *SearchQuickResponsesInput, optFns ...func(*Options)) (*SearchQuickResponsesOutput, error) {
	if params == nil {
		params = &SearchQuickResponsesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchQuickResponses", params, optFns, c.addOperationSearchQuickResponsesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchQuickResponsesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchQuickResponsesInput struct {

	// The identifier of the knowledge base. This should be a QUICK_RESPONSES type
	// knowledge base. Can be either the ID or the ARN. URLs cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The search expression for querying the quick response.
	//
	// This member is required.
	SearchExpression *types.QuickResponseSearchExpression

	// The [user-defined Amazon Connect contact attributes] to be resolved when search results are returned.
	//
	// [user-defined Amazon Connect contact attributes]: https://docs.aws.amazon.com/connect/latest/adminguide/connect-attrib-list.html#user-defined-attributes
	Attributes map[string]string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchQuickResponsesOutput struct {

	// The results of the quick response search.
	//
	// This member is required.
	Results []types.QuickResponseSearchResultData

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchQuickResponsesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchQuickResponses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchQuickResponses{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchQuickResponses"); err != nil {
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
	if err = addOpSearchQuickResponsesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchQuickResponses(options.Region), middleware.Before); err != nil {
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

// SearchQuickResponsesPaginatorOptions is the paginator options for
// SearchQuickResponses
type SearchQuickResponsesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchQuickResponsesPaginator is a paginator for SearchQuickResponses
type SearchQuickResponsesPaginator struct {
	options   SearchQuickResponsesPaginatorOptions
	client    SearchQuickResponsesAPIClient
	params    *SearchQuickResponsesInput
	nextToken *string
	firstPage bool
}

// NewSearchQuickResponsesPaginator returns a new SearchQuickResponsesPaginator
func NewSearchQuickResponsesPaginator(client SearchQuickResponsesAPIClient, params *SearchQuickResponsesInput, optFns ...func(*SearchQuickResponsesPaginatorOptions)) *SearchQuickResponsesPaginator {
	if params == nil {
		params = &SearchQuickResponsesInput{}
	}

	options := SearchQuickResponsesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchQuickResponsesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchQuickResponsesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchQuickResponses page.
func (p *SearchQuickResponsesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchQuickResponsesOutput, error) {
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
	result, err := p.client.SearchQuickResponses(ctx, &params, optFns...)
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

// SearchQuickResponsesAPIClient is a client that implements the
// SearchQuickResponses operation.
type SearchQuickResponsesAPIClient interface {
	SearchQuickResponses(context.Context, *SearchQuickResponsesInput, ...func(*Options)) (*SearchQuickResponsesOutput, error)
}

var _ SearchQuickResponsesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opSearchQuickResponses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchQuickResponses",
	}
}
