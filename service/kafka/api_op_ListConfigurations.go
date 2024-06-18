// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the MSK configurations in this Region.
func (c *Client) ListConfigurations(ctx context.Context, params *ListConfigurationsInput, optFns ...func(*Options)) (*ListConfigurationsOutput, error) {
	if params == nil {
		params = &ListConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurations", params, optFns, c.addOperationListConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationsInput struct {

	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	MaxResults *int32

	// The paginated results marker. When the result of the operation is truncated,
	// the call returns NextToken in the response. To get the next batch, provide this
	// token in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConfigurationsOutput struct {

	// An array of MSK configurations.
	Configurations []types.Configuration

	// The paginated results marker. When the result of a ListConfigurations operation
	// is truncated, the call returns NextToken in the response. To get another batch
	// of configurations, provide this token in your next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurations(options.Region), middleware.Before); err != nil {
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

// ListConfigurationsPaginatorOptions is the paginator options for
// ListConfigurations
type ListConfigurationsPaginatorOptions struct {
	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConfigurationsPaginator is a paginator for ListConfigurations
type ListConfigurationsPaginator struct {
	options   ListConfigurationsPaginatorOptions
	client    ListConfigurationsAPIClient
	params    *ListConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListConfigurationsPaginator returns a new ListConfigurationsPaginator
func NewListConfigurationsPaginator(client ListConfigurationsAPIClient, params *ListConfigurationsInput, optFns ...func(*ListConfigurationsPaginatorOptions)) *ListConfigurationsPaginator {
	if params == nil {
		params = &ListConfigurationsInput{}
	}

	options := ListConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConfigurations page.
func (p *ListConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConfigurationsOutput, error) {
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
	result, err := p.client.ListConfigurations(ctx, &params, optFns...)
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

// ListConfigurationsAPIClient is a client that implements the ListConfigurations
// operation.
type ListConfigurationsAPIClient interface {
	ListConfigurations(context.Context, *ListConfigurationsInput, ...func(*Options)) (*ListConfigurationsOutput, error)
}

var _ ListConfigurationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConfigurations",
	}
}
