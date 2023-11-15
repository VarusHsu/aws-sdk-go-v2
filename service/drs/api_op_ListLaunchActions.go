// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists resource launch actions.
func (c *Client) ListLaunchActions(ctx context.Context, params *ListLaunchActionsInput, optFns ...func(*Options)) (*ListLaunchActionsOutput, error) {
	if params == nil {
		params = &ListLaunchActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLaunchActions", params, optFns, c.addOperationListLaunchActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLaunchActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLaunchActionsInput struct {

	// Launch configuration template Id or Source Server Id
	//
	// This member is required.
	ResourceId *string

	// Filters to apply when listing resource launch actions.
	Filters *types.LaunchActionsRequestFilters

	// Maximum amount of items to return when listing resource launch actions.
	MaxResults int32

	// Next token to use when listing resource launch actions.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLaunchActionsOutput struct {

	// List of resource launch actions.
	Items []types.LaunchAction

	// Next token returned when listing resource launch actions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLaunchActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLaunchActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLaunchActions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLaunchActions"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpListLaunchActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLaunchActions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

// ListLaunchActionsAPIClient is a client that implements the ListLaunchActions
// operation.
type ListLaunchActionsAPIClient interface {
	ListLaunchActions(context.Context, *ListLaunchActionsInput, ...func(*Options)) (*ListLaunchActionsOutput, error)
}

var _ ListLaunchActionsAPIClient = (*Client)(nil)

// ListLaunchActionsPaginatorOptions is the paginator options for ListLaunchActions
type ListLaunchActionsPaginatorOptions struct {
	// Maximum amount of items to return when listing resource launch actions.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLaunchActionsPaginator is a paginator for ListLaunchActions
type ListLaunchActionsPaginator struct {
	options   ListLaunchActionsPaginatorOptions
	client    ListLaunchActionsAPIClient
	params    *ListLaunchActionsInput
	nextToken *string
	firstPage bool
}

// NewListLaunchActionsPaginator returns a new ListLaunchActionsPaginator
func NewListLaunchActionsPaginator(client ListLaunchActionsAPIClient, params *ListLaunchActionsInput, optFns ...func(*ListLaunchActionsPaginatorOptions)) *ListLaunchActionsPaginator {
	if params == nil {
		params = &ListLaunchActionsInput{}
	}

	options := ListLaunchActionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLaunchActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLaunchActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLaunchActions page.
func (p *ListLaunchActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLaunchActionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListLaunchActions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLaunchActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLaunchActions",
	}
}
