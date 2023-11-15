// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerlinuxsubscriptions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Linux subscriptions that have been discovered. If you have linked
// your organization, the returned results will include data aggregated across your
// accounts in Organizations.
func (c *Client) ListLinuxSubscriptions(ctx context.Context, params *ListLinuxSubscriptionsInput, optFns ...func(*Options)) (*ListLinuxSubscriptionsOutput, error) {
	if params == nil {
		params = &ListLinuxSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLinuxSubscriptions", params, optFns, c.addOperationListLinuxSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLinuxSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// NextToken length limit is half of ddb accepted limit. Increase this limit if
// parameters in request increases.
type ListLinuxSubscriptionsInput struct {

	// An array of structures that you can use to filter the results to those that
	// match one or more sets of key-value pairs that you specify. For example, you can
	// filter by the name of Subscription with an optional operator to see
	// subscriptions that match, partially match, or don't match a certain
	// subscription's name. The valid names for this filter are:
	//   - Subscription
	// The valid Operators for this filter are:
	//   - contains
	//   - equals
	//   - Notequal
	Filters []types.Filter

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLinuxSubscriptionsOutput struct {

	// Token for the next set of results.
	NextToken *string

	// An array that contains subscription objects.
	Subscriptions []types.Subscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLinuxSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLinuxSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLinuxSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLinuxSubscriptions"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLinuxSubscriptions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

// ListLinuxSubscriptionsAPIClient is a client that implements the
// ListLinuxSubscriptions operation.
type ListLinuxSubscriptionsAPIClient interface {
	ListLinuxSubscriptions(context.Context, *ListLinuxSubscriptionsInput, ...func(*Options)) (*ListLinuxSubscriptionsOutput, error)
}

var _ ListLinuxSubscriptionsAPIClient = (*Client)(nil)

// ListLinuxSubscriptionsPaginatorOptions is the paginator options for
// ListLinuxSubscriptions
type ListLinuxSubscriptionsPaginatorOptions struct {
	// Maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLinuxSubscriptionsPaginator is a paginator for ListLinuxSubscriptions
type ListLinuxSubscriptionsPaginator struct {
	options   ListLinuxSubscriptionsPaginatorOptions
	client    ListLinuxSubscriptionsAPIClient
	params    *ListLinuxSubscriptionsInput
	nextToken *string
	firstPage bool
}

// NewListLinuxSubscriptionsPaginator returns a new ListLinuxSubscriptionsPaginator
func NewListLinuxSubscriptionsPaginator(client ListLinuxSubscriptionsAPIClient, params *ListLinuxSubscriptionsInput, optFns ...func(*ListLinuxSubscriptionsPaginatorOptions)) *ListLinuxSubscriptionsPaginator {
	if params == nil {
		params = &ListLinuxSubscriptionsInput{}
	}

	options := ListLinuxSubscriptionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLinuxSubscriptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLinuxSubscriptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLinuxSubscriptions page.
func (p *ListLinuxSubscriptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLinuxSubscriptionsOutput, error) {
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

	result, err := p.client.ListLinuxSubscriptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLinuxSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLinuxSubscriptions",
	}
}
