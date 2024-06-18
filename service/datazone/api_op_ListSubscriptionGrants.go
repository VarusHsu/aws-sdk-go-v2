// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists subscription grants.
func (c *Client) ListSubscriptionGrants(ctx context.Context, params *ListSubscriptionGrantsInput, optFns ...func(*Options)) (*ListSubscriptionGrantsOutput, error) {
	if params == nil {
		params = &ListSubscriptionGrantsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubscriptionGrants", params, optFns, c.addOperationListSubscriptionGrantsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubscriptionGrantsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSubscriptionGrantsInput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the Amazon DataZone environment.
	EnvironmentId *string

	// The maximum number of subscription grants to return in a single call to
	// ListSubscriptionGrants . When the number of subscription grants to be listed is
	// greater than the value of MaxResults , the response contains a NextToken value
	// that you can use in a subsequent call to ListSubscriptionGrants to list the
	// next set of subscription grants.
	MaxResults *int32

	// When the number of subscription grants is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of subscription grants, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListSubscriptionGrants to list the next set of subscription
	// grants.
	NextToken *string

	// Specifies the way of sorting the results of this action.
	SortBy types.SortKey

	// Specifies the sort order of this action.
	SortOrder types.SortOrder

	// The identifier of the subscribed listing.
	SubscribedListingId *string

	// The identifier of the subscription.
	SubscriptionId *string

	// The identifier of the subscription target.
	SubscriptionTargetId *string

	noSmithyDocumentSerde
}

type ListSubscriptionGrantsOutput struct {

	// The results of the ListSubscriptionGrants action.
	//
	// This member is required.
	Items []types.SubscriptionGrantSummary

	// When the number of subscription grants is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of subscription grants, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListSubscriptionGrants to list the next set of subscription
	// grants.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSubscriptionGrantsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSubscriptionGrants{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSubscriptionGrants{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSubscriptionGrants"); err != nil {
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
	if err = addOpListSubscriptionGrantsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscriptionGrants(options.Region), middleware.Before); err != nil {
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

// ListSubscriptionGrantsPaginatorOptions is the paginator options for
// ListSubscriptionGrants
type ListSubscriptionGrantsPaginatorOptions struct {
	// The maximum number of subscription grants to return in a single call to
	// ListSubscriptionGrants . When the number of subscription grants to be listed is
	// greater than the value of MaxResults , the response contains a NextToken value
	// that you can use in a subsequent call to ListSubscriptionGrants to list the
	// next set of subscription grants.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSubscriptionGrantsPaginator is a paginator for ListSubscriptionGrants
type ListSubscriptionGrantsPaginator struct {
	options   ListSubscriptionGrantsPaginatorOptions
	client    ListSubscriptionGrantsAPIClient
	params    *ListSubscriptionGrantsInput
	nextToken *string
	firstPage bool
}

// NewListSubscriptionGrantsPaginator returns a new ListSubscriptionGrantsPaginator
func NewListSubscriptionGrantsPaginator(client ListSubscriptionGrantsAPIClient, params *ListSubscriptionGrantsInput, optFns ...func(*ListSubscriptionGrantsPaginatorOptions)) *ListSubscriptionGrantsPaginator {
	if params == nil {
		params = &ListSubscriptionGrantsInput{}
	}

	options := ListSubscriptionGrantsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSubscriptionGrantsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSubscriptionGrantsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSubscriptionGrants page.
func (p *ListSubscriptionGrantsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSubscriptionGrantsOutput, error) {
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
	result, err := p.client.ListSubscriptionGrants(ctx, &params, optFns...)
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

// ListSubscriptionGrantsAPIClient is a client that implements the
// ListSubscriptionGrants operation.
type ListSubscriptionGrantsAPIClient interface {
	ListSubscriptionGrants(context.Context, *ListSubscriptionGrantsInput, ...func(*Options)) (*ListSubscriptionGrantsOutput, error)
}

var _ ListSubscriptionGrantsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSubscriptionGrants(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSubscriptionGrants",
	}
}
