// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches contacts in an Amazon Connect instance.
func (c *Client) SearchContacts(ctx context.Context, params *SearchContactsInput, optFns ...func(*Options)) (*SearchContactsOutput, error) {
	if params == nil {
		params = &SearchContactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchContacts", params, optFns, c.addOperationSearchContactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchContactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchContactsInput struct {

	// The identifier of Amazon Connect instance. You can find the instance ID in the
	// Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// Time range that you want to search results.
	//
	// This member is required.
	TimeRange *types.SearchContactsTimeRange

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The search criteria to be used to return contacts.
	SearchCriteria *types.SearchCriteria

	// Specifies a field to sort by and a sort order.
	Sort *types.Sort

	noSmithyDocumentSerde
}

type SearchContactsOutput struct {

	// Information about the contacts.
	//
	// This member is required.
	Contacts []types.ContactSearchSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The total number of contacts which matched your search query.
	TotalCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchContactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchContacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchContacts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchContacts"); err != nil {
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
	if err = addOpSearchContactsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchContacts(options.Region), middleware.Before); err != nil {
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

// SearchContactsPaginatorOptions is the paginator options for SearchContacts
type SearchContactsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchContactsPaginator is a paginator for SearchContacts
type SearchContactsPaginator struct {
	options   SearchContactsPaginatorOptions
	client    SearchContactsAPIClient
	params    *SearchContactsInput
	nextToken *string
	firstPage bool
}

// NewSearchContactsPaginator returns a new SearchContactsPaginator
func NewSearchContactsPaginator(client SearchContactsAPIClient, params *SearchContactsInput, optFns ...func(*SearchContactsPaginatorOptions)) *SearchContactsPaginator {
	if params == nil {
		params = &SearchContactsInput{}
	}

	options := SearchContactsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchContactsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchContactsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchContacts page.
func (p *SearchContactsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchContactsOutput, error) {
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
	result, err := p.client.SearchContacts(ctx, &params, optFns...)
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

// SearchContactsAPIClient is a client that implements the SearchContacts
// operation.
type SearchContactsAPIClient interface {
	SearchContacts(context.Context, *SearchContactsInput, ...func(*Options)) (*SearchContactsOutput, error)
}

var _ SearchContactsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opSearchContacts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchContacts",
	}
}
