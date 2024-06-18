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

// Lists traffic distribution groups.
func (c *Client) ListTrafficDistributionGroups(ctx context.Context, params *ListTrafficDistributionGroupsInput, optFns ...func(*Options)) (*ListTrafficDistributionGroupsOutput, error) {
	if params == nil {
		params = &ListTrafficDistributionGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrafficDistributionGroups", params, optFns, c.addOperationListTrafficDistributionGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrafficDistributionGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrafficDistributionGroupsInput struct {

	// The identifier of the Amazon Connect instance. You can [find the instance ID] in the Amazon Resource
	// Name (ARN) of the instance.
	//
	// [find the instance ID]: https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTrafficDistributionGroupsOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// A list of traffic distribution groups.
	TrafficDistributionGroupSummaryList []types.TrafficDistributionGroupSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrafficDistributionGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTrafficDistributionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTrafficDistributionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTrafficDistributionGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficDistributionGroups(options.Region), middleware.Before); err != nil {
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

// ListTrafficDistributionGroupsPaginatorOptions is the paginator options for
// ListTrafficDistributionGroups
type ListTrafficDistributionGroupsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTrafficDistributionGroupsPaginator is a paginator for
// ListTrafficDistributionGroups
type ListTrafficDistributionGroupsPaginator struct {
	options   ListTrafficDistributionGroupsPaginatorOptions
	client    ListTrafficDistributionGroupsAPIClient
	params    *ListTrafficDistributionGroupsInput
	nextToken *string
	firstPage bool
}

// NewListTrafficDistributionGroupsPaginator returns a new
// ListTrafficDistributionGroupsPaginator
func NewListTrafficDistributionGroupsPaginator(client ListTrafficDistributionGroupsAPIClient, params *ListTrafficDistributionGroupsInput, optFns ...func(*ListTrafficDistributionGroupsPaginatorOptions)) *ListTrafficDistributionGroupsPaginator {
	if params == nil {
		params = &ListTrafficDistributionGroupsInput{}
	}

	options := ListTrafficDistributionGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTrafficDistributionGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTrafficDistributionGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTrafficDistributionGroups page.
func (p *ListTrafficDistributionGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTrafficDistributionGroupsOutput, error) {
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
	result, err := p.client.ListTrafficDistributionGroups(ctx, &params, optFns...)
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

// ListTrafficDistributionGroupsAPIClient is a client that implements the
// ListTrafficDistributionGroups operation.
type ListTrafficDistributionGroupsAPIClient interface {
	ListTrafficDistributionGroups(context.Context, *ListTrafficDistributionGroupsInput, ...func(*Options)) (*ListTrafficDistributionGroupsOutput, error)
}

var _ ListTrafficDistributionGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTrafficDistributionGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTrafficDistributionGroups",
	}
}
