// Code generated by smithy-go-codegen DO NOT EDIT.

package supportapp

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/supportapp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Slack channel configurations for an Amazon Web Services account.
func (c *Client) ListSlackChannelConfigurations(ctx context.Context, params *ListSlackChannelConfigurationsInput, optFns ...func(*Options)) (*ListSlackChannelConfigurationsOutput, error) {
	if params == nil {
		params = &ListSlackChannelConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSlackChannelConfigurations", params, optFns, c.addOperationListSlackChannelConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSlackChannelConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSlackChannelConfigurationsInput struct {

	// If the results of a search are large, the API only returns a portion of the
	// results and includes a nextToken pagination token in the response. To retrieve
	// the next batch of results, reissue the search request and include the returned
	// token. When the API returns the last set of results, the response doesn't
	// include a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSlackChannelConfigurationsOutput struct {

	// The configurations for a Slack channel.
	//
	// This member is required.
	SlackChannelConfigurations []types.SlackChannelConfiguration

	// The point where pagination should resume when the response returns only partial
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSlackChannelConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSlackChannelConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSlackChannelConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSlackChannelConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSlackChannelConfigurations(options.Region), middleware.Before); err != nil {
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

// ListSlackChannelConfigurationsPaginatorOptions is the paginator options for
// ListSlackChannelConfigurations
type ListSlackChannelConfigurationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSlackChannelConfigurationsPaginator is a paginator for
// ListSlackChannelConfigurations
type ListSlackChannelConfigurationsPaginator struct {
	options   ListSlackChannelConfigurationsPaginatorOptions
	client    ListSlackChannelConfigurationsAPIClient
	params    *ListSlackChannelConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListSlackChannelConfigurationsPaginator returns a new
// ListSlackChannelConfigurationsPaginator
func NewListSlackChannelConfigurationsPaginator(client ListSlackChannelConfigurationsAPIClient, params *ListSlackChannelConfigurationsInput, optFns ...func(*ListSlackChannelConfigurationsPaginatorOptions)) *ListSlackChannelConfigurationsPaginator {
	if params == nil {
		params = &ListSlackChannelConfigurationsInput{}
	}

	options := ListSlackChannelConfigurationsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSlackChannelConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSlackChannelConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSlackChannelConfigurations page.
func (p *ListSlackChannelConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSlackChannelConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListSlackChannelConfigurations(ctx, &params, optFns...)
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

// ListSlackChannelConfigurationsAPIClient is a client that implements the
// ListSlackChannelConfigurations operation.
type ListSlackChannelConfigurationsAPIClient interface {
	ListSlackChannelConfigurations(context.Context, *ListSlackChannelConfigurationsInput, ...func(*Options)) (*ListSlackChannelConfigurationsOutput, error)
}

var _ ListSlackChannelConfigurationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSlackChannelConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSlackChannelConfigurations",
	}
}
