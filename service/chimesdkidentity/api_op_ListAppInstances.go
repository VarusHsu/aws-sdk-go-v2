// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Amazon Chime AppInstance s created under a single AWS account.
func (c *Client) ListAppInstances(ctx context.Context, params *ListAppInstancesInput, optFns ...func(*Options)) (*ListAppInstancesOutput, error) {
	if params == nil {
		params = &ListAppInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppInstances", params, optFns, c.addOperationListAppInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppInstancesInput struct {

	// The maximum number of AppInstance s that you want to return.
	MaxResults *int32

	// The token passed by previous API requests until you reach the maximum number of
	// AppInstances .
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppInstancesOutput struct {

	// The information for each AppInstance .
	AppInstances []types.AppInstanceSummary

	// The token passed by previous API requests until the maximum number of
	// AppInstance s is reached.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAppInstances"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppInstances(options.Region), middleware.Before); err != nil {
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

// ListAppInstancesPaginatorOptions is the paginator options for ListAppInstances
type ListAppInstancesPaginatorOptions struct {
	// The maximum number of AppInstance s that you want to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppInstancesPaginator is a paginator for ListAppInstances
type ListAppInstancesPaginator struct {
	options   ListAppInstancesPaginatorOptions
	client    ListAppInstancesAPIClient
	params    *ListAppInstancesInput
	nextToken *string
	firstPage bool
}

// NewListAppInstancesPaginator returns a new ListAppInstancesPaginator
func NewListAppInstancesPaginator(client ListAppInstancesAPIClient, params *ListAppInstancesInput, optFns ...func(*ListAppInstancesPaginatorOptions)) *ListAppInstancesPaginator {
	if params == nil {
		params = &ListAppInstancesInput{}
	}

	options := ListAppInstancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppInstances page.
func (p *ListAppInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppInstancesOutput, error) {
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
	result, err := p.client.ListAppInstances(ctx, &params, optFns...)
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

// ListAppInstancesAPIClient is a client that implements the ListAppInstances
// operation.
type ListAppInstancesAPIClient interface {
	ListAppInstances(context.Context, *ListAppInstancesInput, ...func(*Options)) (*ListAppInstancesOutput, error)
}

var _ ListAppInstancesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAppInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAppInstances",
	}
}
