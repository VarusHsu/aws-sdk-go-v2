// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the latest analytics data for controls within a specific control domain
// across all active assessments.
//
// Control insights are listed only if the control belongs to the control domain
// that was specified and the control collected evidence on the lastUpdated date
// of controlInsightsMetadata . If neither of these conditions are met, no data is
// listed for that control.
func (c *Client) ListControlInsightsByControlDomain(ctx context.Context, params *ListControlInsightsByControlDomainInput, optFns ...func(*Options)) (*ListControlInsightsByControlDomainOutput, error) {
	if params == nil {
		params = &ListControlInsightsByControlDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListControlInsightsByControlDomain", params, optFns, c.addOperationListControlInsightsByControlDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListControlInsightsByControlDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListControlInsightsByControlDomainInput struct {

	// The unique identifier for the control domain.
	//
	// Audit Manager supports the control domains that are provided by Amazon Web
	// Services Control Catalog. For information about how to find a list of available
	// control domains, see [ListDomains]ListDomains in the Amazon Web Services Control Catalog API
	// Reference.
	//
	// [ListDomains]: https://docs.aws.amazon.com/controlcatalog/latest/APIReference/API_ListDomains.html
	//
	// This member is required.
	ControlDomainId *string

	// Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListControlInsightsByControlDomainOutput struct {

	// The control analytics data that the ListControlInsightsByControlDomain API
	// returned.
	ControlInsightsMetadata []types.ControlInsightsMetadataItem

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListControlInsightsByControlDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListControlInsightsByControlDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListControlInsightsByControlDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListControlInsightsByControlDomain"); err != nil {
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
	if err = addOpListControlInsightsByControlDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListControlInsightsByControlDomain(options.Region), middleware.Before); err != nil {
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

// ListControlInsightsByControlDomainPaginatorOptions is the paginator options for
// ListControlInsightsByControlDomain
type ListControlInsightsByControlDomainPaginatorOptions struct {
	// Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListControlInsightsByControlDomainPaginator is a paginator for
// ListControlInsightsByControlDomain
type ListControlInsightsByControlDomainPaginator struct {
	options   ListControlInsightsByControlDomainPaginatorOptions
	client    ListControlInsightsByControlDomainAPIClient
	params    *ListControlInsightsByControlDomainInput
	nextToken *string
	firstPage bool
}

// NewListControlInsightsByControlDomainPaginator returns a new
// ListControlInsightsByControlDomainPaginator
func NewListControlInsightsByControlDomainPaginator(client ListControlInsightsByControlDomainAPIClient, params *ListControlInsightsByControlDomainInput, optFns ...func(*ListControlInsightsByControlDomainPaginatorOptions)) *ListControlInsightsByControlDomainPaginator {
	if params == nil {
		params = &ListControlInsightsByControlDomainInput{}
	}

	options := ListControlInsightsByControlDomainPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListControlInsightsByControlDomainPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListControlInsightsByControlDomainPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListControlInsightsByControlDomain page.
func (p *ListControlInsightsByControlDomainPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListControlInsightsByControlDomainOutput, error) {
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
	result, err := p.client.ListControlInsightsByControlDomain(ctx, &params, optFns...)
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

// ListControlInsightsByControlDomainAPIClient is a client that implements the
// ListControlInsightsByControlDomain operation.
type ListControlInsightsByControlDomainAPIClient interface {
	ListControlInsightsByControlDomain(context.Context, *ListControlInsightsByControlDomainInput, ...func(*Options)) (*ListControlInsightsByControlDomainOutput, error)
}

var _ ListControlInsightsByControlDomainAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListControlInsightsByControlDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListControlInsightsByControlDomain",
	}
}
