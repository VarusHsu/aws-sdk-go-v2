// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of snapshot copy configurations.
func (c *Client) ListSnapshotCopyConfigurations(ctx context.Context, params *ListSnapshotCopyConfigurationsInput, optFns ...func(*Options)) (*ListSnapshotCopyConfigurationsOutput, error) {
	if params == nil {
		params = &ListSnapshotCopyConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSnapshotCopyConfigurations", params, optFns, c.addOperationListSnapshotCopyConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSnapshotCopyConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSnapshotCopyConfigurationsInput struct {

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to display the next page of results.
	MaxResults *int32

	// The namespace from which to list all snapshot copy configurations.
	NamespaceName *string

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSnapshotCopyConfigurationsOutput struct {

	// All of the returned snapshot copy configurations.
	//
	// This member is required.
	SnapshotCopyConfigurations []types.SnapshotCopyConfiguration

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSnapshotCopyConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSnapshotCopyConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSnapshotCopyConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSnapshotCopyConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSnapshotCopyConfigurations(options.Region), middleware.Before); err != nil {
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

// ListSnapshotCopyConfigurationsPaginatorOptions is the paginator options for
// ListSnapshotCopyConfigurations
type ListSnapshotCopyConfigurationsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to display the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSnapshotCopyConfigurationsPaginator is a paginator for
// ListSnapshotCopyConfigurations
type ListSnapshotCopyConfigurationsPaginator struct {
	options   ListSnapshotCopyConfigurationsPaginatorOptions
	client    ListSnapshotCopyConfigurationsAPIClient
	params    *ListSnapshotCopyConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListSnapshotCopyConfigurationsPaginator returns a new
// ListSnapshotCopyConfigurationsPaginator
func NewListSnapshotCopyConfigurationsPaginator(client ListSnapshotCopyConfigurationsAPIClient, params *ListSnapshotCopyConfigurationsInput, optFns ...func(*ListSnapshotCopyConfigurationsPaginatorOptions)) *ListSnapshotCopyConfigurationsPaginator {
	if params == nil {
		params = &ListSnapshotCopyConfigurationsInput{}
	}

	options := ListSnapshotCopyConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSnapshotCopyConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSnapshotCopyConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSnapshotCopyConfigurations page.
func (p *ListSnapshotCopyConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSnapshotCopyConfigurationsOutput, error) {
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
	result, err := p.client.ListSnapshotCopyConfigurations(ctx, &params, optFns...)
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

// ListSnapshotCopyConfigurationsAPIClient is a client that implements the
// ListSnapshotCopyConfigurations operation.
type ListSnapshotCopyConfigurationsAPIClient interface {
	ListSnapshotCopyConfigurations(context.Context, *ListSnapshotCopyConfigurationsInput, ...func(*Options)) (*ListSnapshotCopyConfigurationsOutput, error)
}

var _ ListSnapshotCopyConfigurationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSnapshotCopyConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSnapshotCopyConfigurations",
	}
}
