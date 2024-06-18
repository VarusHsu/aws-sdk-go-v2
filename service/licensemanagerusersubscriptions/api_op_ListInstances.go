// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerusersubscriptions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the EC2 instances providing user-based subscriptions.
func (c *Client) ListInstances(ctx context.Context, params *ListInstancesInput, optFns ...func(*Options)) (*ListInstancesOutput, error) {
	if params == nil {
		params = &ListInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstances", params, optFns, c.addOperationListInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstancesInput struct {

	// An array of structures that you can use to filter the results to those that
	// match one or more sets of key-value pairs that you specify.
	Filters []types.Filter

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListInstancesOutput struct {

	// Metadata that describes the list instances operation.
	InstanceSummaries []types.InstanceSummary

	// Token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInstances"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstances(options.Region), middleware.Before); err != nil {
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

// ListInstancesPaginatorOptions is the paginator options for ListInstances
type ListInstancesPaginatorOptions struct {
	// Maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstancesPaginator is a paginator for ListInstances
type ListInstancesPaginator struct {
	options   ListInstancesPaginatorOptions
	client    ListInstancesAPIClient
	params    *ListInstancesInput
	nextToken *string
	firstPage bool
}

// NewListInstancesPaginator returns a new ListInstancesPaginator
func NewListInstancesPaginator(client ListInstancesAPIClient, params *ListInstancesInput, optFns ...func(*ListInstancesPaginatorOptions)) *ListInstancesPaginator {
	if params == nil {
		params = &ListInstancesInput{}
	}

	options := ListInstancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInstances page.
func (p *ListInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstancesOutput, error) {
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
	result, err := p.client.ListInstances(ctx, &params, optFns...)
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

// ListInstancesAPIClient is a client that implements the ListInstances operation.
type ListInstancesAPIClient interface {
	ListInstances(context.Context, *ListInstancesInput, ...func(*Options)) (*ListInstancesOutput, error)
}

var _ ListInstancesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInstances",
	}
}
