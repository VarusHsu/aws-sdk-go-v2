// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the event subscriptions for a customer account. The description of a
// subscription includes SubscriptionName , SNSTopicARN , CustomerID , SourceType ,
// SourceID , CreationTime , and Status .
//
// If you specify SubscriptionName , this action lists the description for that
// subscription.
func (c *Client) DescribeEventSubscriptions(ctx context.Context, params *DescribeEventSubscriptionsInput, optFns ...func(*Options)) (*DescribeEventSubscriptionsOutput, error) {
	if params == nil {
		params = &DescribeEventSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventSubscriptions", params, optFns, c.addOperationDescribeEventSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventSubscriptionsInput struct {

	// Filters applied to event subscriptions.
	//
	// Valid filter names: event-subscription-arn | event-subscription-id
	Filters []types.Filter

	//  An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	//  The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The name of the DMS event subscription to be described.
	SubscriptionName *string

	noSmithyDocumentSerde
}

type DescribeEventSubscriptionsOutput struct {

	// A list of event subscriptions.
	EventSubscriptionsList []types.EventSubscription

	//  An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEventSubscriptions"); err != nil {
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
	if err = addOpDescribeEventSubscriptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventSubscriptions(options.Region), middleware.Before); err != nil {
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

// DescribeEventSubscriptionsPaginatorOptions is the paginator options for
// DescribeEventSubscriptions
type DescribeEventSubscriptionsPaginatorOptions struct {
	//  The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEventSubscriptionsPaginator is a paginator for
// DescribeEventSubscriptions
type DescribeEventSubscriptionsPaginator struct {
	options   DescribeEventSubscriptionsPaginatorOptions
	client    DescribeEventSubscriptionsAPIClient
	params    *DescribeEventSubscriptionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEventSubscriptionsPaginator returns a new
// DescribeEventSubscriptionsPaginator
func NewDescribeEventSubscriptionsPaginator(client DescribeEventSubscriptionsAPIClient, params *DescribeEventSubscriptionsInput, optFns ...func(*DescribeEventSubscriptionsPaginatorOptions)) *DescribeEventSubscriptionsPaginator {
	if params == nil {
		params = &DescribeEventSubscriptionsInput{}
	}

	options := DescribeEventSubscriptionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEventSubscriptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEventSubscriptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEventSubscriptions page.
func (p *DescribeEventSubscriptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEventSubscriptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeEventSubscriptions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeEventSubscriptionsAPIClient is a client that implements the
// DescribeEventSubscriptions operation.
type DescribeEventSubscriptionsAPIClient interface {
	DescribeEventSubscriptions(context.Context, *DescribeEventSubscriptionsInput, ...func(*Options)) (*DescribeEventSubscriptionsOutput, error)
}

var _ DescribeEventSubscriptionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeEventSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEventSubscriptions",
	}
}
