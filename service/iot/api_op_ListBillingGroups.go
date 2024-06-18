// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the billing groups you have created.
//
// Requires permission to access the [ListBillingGroups] action.
//
// [ListBillingGroups]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListBillingGroups(ctx context.Context, params *ListBillingGroupsInput, optFns ...func(*Options)) (*ListBillingGroupsOutput, error) {
	if params == nil {
		params = &ListBillingGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBillingGroups", params, optFns, c.addOperationListBillingGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBillingGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBillingGroupsInput struct {

	// The maximum number of results to return per request.
	MaxResults *int32

	// Limit the results to billing groups whose names have the given prefix.
	NamePrefixFilter *string

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBillingGroupsOutput struct {

	// The list of billing groups.
	BillingGroups []types.GroupNameAndArn

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBillingGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBillingGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBillingGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBillingGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBillingGroups(options.Region), middleware.Before); err != nil {
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

// ListBillingGroupsPaginatorOptions is the paginator options for ListBillingGroups
type ListBillingGroupsPaginatorOptions struct {
	// The maximum number of results to return per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBillingGroupsPaginator is a paginator for ListBillingGroups
type ListBillingGroupsPaginator struct {
	options   ListBillingGroupsPaginatorOptions
	client    ListBillingGroupsAPIClient
	params    *ListBillingGroupsInput
	nextToken *string
	firstPage bool
}

// NewListBillingGroupsPaginator returns a new ListBillingGroupsPaginator
func NewListBillingGroupsPaginator(client ListBillingGroupsAPIClient, params *ListBillingGroupsInput, optFns ...func(*ListBillingGroupsPaginatorOptions)) *ListBillingGroupsPaginator {
	if params == nil {
		params = &ListBillingGroupsInput{}
	}

	options := ListBillingGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBillingGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBillingGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBillingGroups page.
func (p *ListBillingGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBillingGroupsOutput, error) {
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
	result, err := p.client.ListBillingGroups(ctx, &params, optFns...)
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

// ListBillingGroupsAPIClient is a client that implements the ListBillingGroups
// operation.
type ListBillingGroupsAPIClient interface {
	ListBillingGroups(context.Context, *ListBillingGroupsInput, ...func(*Options)) (*ListBillingGroupsOutput, error)
}

var _ ListBillingGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListBillingGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBillingGroups",
	}
}
