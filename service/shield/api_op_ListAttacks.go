// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all ongoing DDoS attacks or all DDoS attacks during a specified time
// period.
func (c *Client) ListAttacks(ctx context.Context, params *ListAttacksInput, optFns ...func(*Options)) (*ListAttacksOutput, error) {
	if params == nil {
		params = &ListAttacksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttacks", params, optFns, c.addOperationListAttacksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttacksInput struct {

	// The end of the time period for the attacks. This is a timestamp type. The
	// request syntax listing for this call indicates a number type, but you can
	// provide the time in any valid [timestamp format]setting.
	//
	// [timestamp format]: https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-parameters-types.html#parameter-type-timestamp
	EndTime *types.TimeRange

	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response.
	//
	// The default setting is 20.
	MaxResults *int32

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request.
	//
	// You can indicate the maximum number of objects that you want Shield Advanced to
	// return for a single call with the MaxResults setting. Shield Advanced will not
	// return more than MaxResults objects, but may return fewer, even if more objects
	// are still available.
	//
	// Whenever more objects remain that Shield Advanced has not yet returned to you,
	// the response will include a NextToken value.
	//
	// On your first call to a list operation, leave this setting empty.
	NextToken *string

	// The ARNs (Amazon Resource Names) of the resources that were attacked. If you
	// leave this blank, all applicable resources for this account will be included.
	ResourceArns []string

	// The start of the time period for the attacks. This is a timestamp type. The
	// request syntax listing for this call indicates a number type, but you can
	// provide the time in any valid [timestamp format]setting.
	//
	// [timestamp format]: https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-parameters-types.html#parameter-type-timestamp
	StartTime *types.TimeRange

	noSmithyDocumentSerde
}

type ListAttacksOutput struct {

	// The attack information for the specified time range.
	AttackSummaries []types.AttackSummary

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request.
	//
	// You can indicate the maximum number of objects that you want Shield Advanced to
	// return for a single call with the MaxResults setting. Shield Advanced will not
	// return more than MaxResults objects, but may return fewer, even if more objects
	// are still available.
	//
	// Whenever more objects remain that Shield Advanced has not yet returned to you,
	// the response will include a NextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAttacksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAttacks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAttacks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAttacks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttacks(options.Region), middleware.Before); err != nil {
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

// ListAttacksPaginatorOptions is the paginator options for ListAttacks
type ListAttacksPaginatorOptions struct {
	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response.
	//
	// The default setting is 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttacksPaginator is a paginator for ListAttacks
type ListAttacksPaginator struct {
	options   ListAttacksPaginatorOptions
	client    ListAttacksAPIClient
	params    *ListAttacksInput
	nextToken *string
	firstPage bool
}

// NewListAttacksPaginator returns a new ListAttacksPaginator
func NewListAttacksPaginator(client ListAttacksAPIClient, params *ListAttacksInput, optFns ...func(*ListAttacksPaginatorOptions)) *ListAttacksPaginator {
	if params == nil {
		params = &ListAttacksInput{}
	}

	options := ListAttacksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttacksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttacksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAttacks page.
func (p *ListAttacksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttacksOutput, error) {
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
	result, err := p.client.ListAttacks(ctx, &params, optFns...)
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

// ListAttacksAPIClient is a client that implements the ListAttacks operation.
type ListAttacksAPIClient interface {
	ListAttacks(context.Context, *ListAttacksInput, ...func(*Options)) (*ListAttacksOutput, error)
}

var _ ListAttacksAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAttacks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAttacks",
	}
}
