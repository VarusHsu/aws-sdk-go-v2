// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns detailed information about the privacy budget templates in a specified
// membership.
func (c *Client) ListPrivacyBudgetTemplates(ctx context.Context, params *ListPrivacyBudgetTemplatesInput, optFns ...func(*Options)) (*ListPrivacyBudgetTemplatesOutput, error) {
	if params == nil {
		params = &ListPrivacyBudgetTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPrivacyBudgetTemplates", params, optFns, c.addOperationListPrivacyBudgetTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPrivacyBudgetTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrivacyBudgetTemplatesInput struct {

	// A unique identifier for one of your memberships for a collaboration. The
	// privacy budget templates are retrieved from the collaboration that this
	// membership belongs to. Accepts a membership ID.
	//
	// This member is required.
	MembershipIdentifier *string

	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service may return a nextToken even if the
	// maximum results has not been met.
	MaxResults *int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPrivacyBudgetTemplatesOutput struct {

	// An array that summarizes the privacy budget templates. The summary includes
	// collaboration information, creation information, and privacy budget type.
	//
	// This member is required.
	PrivacyBudgetTemplateSummaries []types.PrivacyBudgetTemplateSummary

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPrivacyBudgetTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPrivacyBudgetTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrivacyBudgetTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPrivacyBudgetTemplates"); err != nil {
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
	if err = addOpListPrivacyBudgetTemplatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrivacyBudgetTemplates(options.Region), middleware.Before); err != nil {
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

// ListPrivacyBudgetTemplatesPaginatorOptions is the paginator options for
// ListPrivacyBudgetTemplates
type ListPrivacyBudgetTemplatesPaginatorOptions struct {
	// The maximum size of the results that is returned per call. Service chooses a
	// default if it has not been set. Service may return a nextToken even if the
	// maximum results has not been met.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPrivacyBudgetTemplatesPaginator is a paginator for
// ListPrivacyBudgetTemplates
type ListPrivacyBudgetTemplatesPaginator struct {
	options   ListPrivacyBudgetTemplatesPaginatorOptions
	client    ListPrivacyBudgetTemplatesAPIClient
	params    *ListPrivacyBudgetTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListPrivacyBudgetTemplatesPaginator returns a new
// ListPrivacyBudgetTemplatesPaginator
func NewListPrivacyBudgetTemplatesPaginator(client ListPrivacyBudgetTemplatesAPIClient, params *ListPrivacyBudgetTemplatesInput, optFns ...func(*ListPrivacyBudgetTemplatesPaginatorOptions)) *ListPrivacyBudgetTemplatesPaginator {
	if params == nil {
		params = &ListPrivacyBudgetTemplatesInput{}
	}

	options := ListPrivacyBudgetTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPrivacyBudgetTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPrivacyBudgetTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPrivacyBudgetTemplates page.
func (p *ListPrivacyBudgetTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPrivacyBudgetTemplatesOutput, error) {
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
	result, err := p.client.ListPrivacyBudgetTemplates(ctx, &params, optFns...)
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

// ListPrivacyBudgetTemplatesAPIClient is a client that implements the
// ListPrivacyBudgetTemplates operation.
type ListPrivacyBudgetTemplatesAPIClient interface {
	ListPrivacyBudgetTemplates(context.Context, *ListPrivacyBudgetTemplatesInput, ...func(*Options)) (*ListPrivacyBudgetTemplatesOutput, error)
}

var _ ListPrivacyBudgetTemplatesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPrivacyBudgetTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPrivacyBudgetTemplates",
	}
}
