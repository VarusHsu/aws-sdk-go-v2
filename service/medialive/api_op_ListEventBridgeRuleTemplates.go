// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists eventbridge rule templates.
func (c *Client) ListEventBridgeRuleTemplates(ctx context.Context, params *ListEventBridgeRuleTemplatesInput, optFns ...func(*Options)) (*ListEventBridgeRuleTemplatesOutput, error) {
	if params == nil {
		params = &ListEventBridgeRuleTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventBridgeRuleTemplates", params, optFns, c.addOperationListEventBridgeRuleTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventBridgeRuleTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListEventBridgeRuleTemplatesRequest
type ListEventBridgeRuleTemplatesInput struct {

	// An eventbridge rule template group's identifier. Can be either be its id or
	// current name.
	GroupIdentifier *string

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// A token used to retrieve the next set of results in paginated list responses.
	NextToken *string

	// A signal map's identifier. Can be either be its id or current name.
	SignalMapIdentifier *string

	noSmithyDocumentSerde
}

// Placeholder documentation for ListEventBridgeRuleTemplatesResponse
type ListEventBridgeRuleTemplatesOutput struct {

	// Placeholder documentation for __listOfEventBridgeRuleTemplateSummary
	EventBridgeRuleTemplates []types.EventBridgeRuleTemplateSummary

	// A token used to retrieve the next set of results in paginated list responses.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEventBridgeRuleTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEventBridgeRuleTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEventBridgeRuleTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEventBridgeRuleTemplates"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventBridgeRuleTemplates(options.Region), middleware.Before); err != nil {
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

// ListEventBridgeRuleTemplatesPaginatorOptions is the paginator options for
// ListEventBridgeRuleTemplates
type ListEventBridgeRuleTemplatesPaginatorOptions struct {
	// Placeholder documentation for MaxResults
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventBridgeRuleTemplatesPaginator is a paginator for
// ListEventBridgeRuleTemplates
type ListEventBridgeRuleTemplatesPaginator struct {
	options   ListEventBridgeRuleTemplatesPaginatorOptions
	client    ListEventBridgeRuleTemplatesAPIClient
	params    *ListEventBridgeRuleTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListEventBridgeRuleTemplatesPaginator returns a new
// ListEventBridgeRuleTemplatesPaginator
func NewListEventBridgeRuleTemplatesPaginator(client ListEventBridgeRuleTemplatesAPIClient, params *ListEventBridgeRuleTemplatesInput, optFns ...func(*ListEventBridgeRuleTemplatesPaginatorOptions)) *ListEventBridgeRuleTemplatesPaginator {
	if params == nil {
		params = &ListEventBridgeRuleTemplatesInput{}
	}

	options := ListEventBridgeRuleTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventBridgeRuleTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventBridgeRuleTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEventBridgeRuleTemplates page.
func (p *ListEventBridgeRuleTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventBridgeRuleTemplatesOutput, error) {
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
	result, err := p.client.ListEventBridgeRuleTemplates(ctx, &params, optFns...)
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

// ListEventBridgeRuleTemplatesAPIClient is a client that implements the
// ListEventBridgeRuleTemplates operation.
type ListEventBridgeRuleTemplatesAPIClient interface {
	ListEventBridgeRuleTemplates(context.Context, *ListEventBridgeRuleTemplatesInput, ...func(*Options)) (*ListEventBridgeRuleTemplatesOutput, error)
}

var _ ListEventBridgeRuleTemplatesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListEventBridgeRuleTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEventBridgeRuleTemplates",
	}
}
