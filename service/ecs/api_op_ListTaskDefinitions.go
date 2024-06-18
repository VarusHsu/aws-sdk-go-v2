// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of task definitions that are registered to your account. You can
// filter the results by family name with the familyPrefix parameter or by status
// with the status parameter.
func (c *Client) ListTaskDefinitions(ctx context.Context, params *ListTaskDefinitionsInput, optFns ...func(*Options)) (*ListTaskDefinitionsOutput, error) {
	if params == nil {
		params = &ListTaskDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTaskDefinitions", params, optFns, c.addOperationListTaskDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTaskDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTaskDefinitionsInput struct {

	// The full family name to filter the ListTaskDefinitions results with. Specifying
	// a familyPrefix limits the listed task definitions to task definition revisions
	// that belong to that family.
	FamilyPrefix *string

	// The maximum number of task definition results that ListTaskDefinitions returned
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitions request with the returned nextToken value. This
	// value can be between 1 and 100. If this parameter isn't used, then
	// ListTaskDefinitions returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// The nextToken value returned from a ListTaskDefinitions request indicating that
	// more results are available to fulfill the request and further calls will be
	// needed. If maxResults was provided, it is possible the number of results to be
	// fewer than maxResults .
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string

	// The order to sort the results in. Valid values are ASC and DESC . By default, (
	// ASC ) task definitions are listed lexicographically by family name and in
	// ascending numerical order by revision so that the newest task definitions in a
	// family are listed last. Setting this parameter to DESC reverses the sort order
	// on family name and revision. This is so that the newest task definitions in a
	// family are listed first.
	Sort types.SortOrder

	// The task definition status to filter the ListTaskDefinitions results with. By
	// default, only ACTIVE task definitions are listed. By setting this parameter to
	// INACTIVE , you can view task definitions that are INACTIVE as long as an active
	// task or service still references them. If you paginate the resulting output, be
	// sure to keep the status value constant in each subsequent request.
	Status types.TaskDefinitionStatus

	noSmithyDocumentSerde
}

type ListTaskDefinitionsOutput struct {

	// The nextToken value to include in a future ListTaskDefinitions request. When
	// the results of a ListTaskDefinitions request exceed maxResults , this value can
	// be used to retrieve the next page of results. This value is null when there are
	// no more results to return.
	NextToken *string

	// The list of task definition Amazon Resource Name (ARN) entries for the
	// ListTaskDefinitions request.
	TaskDefinitionArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTaskDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTaskDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTaskDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTaskDefinitions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTaskDefinitions(options.Region), middleware.Before); err != nil {
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

// ListTaskDefinitionsPaginatorOptions is the paginator options for
// ListTaskDefinitions
type ListTaskDefinitionsPaginatorOptions struct {
	// The maximum number of task definition results that ListTaskDefinitions returned
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitions request with the returned nextToken value. This
	// value can be between 1 and 100. If this parameter isn't used, then
	// ListTaskDefinitions returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTaskDefinitionsPaginator is a paginator for ListTaskDefinitions
type ListTaskDefinitionsPaginator struct {
	options   ListTaskDefinitionsPaginatorOptions
	client    ListTaskDefinitionsAPIClient
	params    *ListTaskDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListTaskDefinitionsPaginator returns a new ListTaskDefinitionsPaginator
func NewListTaskDefinitionsPaginator(client ListTaskDefinitionsAPIClient, params *ListTaskDefinitionsInput, optFns ...func(*ListTaskDefinitionsPaginatorOptions)) *ListTaskDefinitionsPaginator {
	if params == nil {
		params = &ListTaskDefinitionsInput{}
	}

	options := ListTaskDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTaskDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTaskDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTaskDefinitions page.
func (p *ListTaskDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTaskDefinitionsOutput, error) {
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
	result, err := p.client.ListTaskDefinitions(ctx, &params, optFns...)
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

// ListTaskDefinitionsAPIClient is a client that implements the
// ListTaskDefinitions operation.
type ListTaskDefinitionsAPIClient interface {
	ListTaskDefinitions(context.Context, *ListTaskDefinitionsInput, ...func(*Options)) (*ListTaskDefinitionsOutput, error)
}

var _ ListTaskDefinitionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTaskDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTaskDefinitions",
	}
}
