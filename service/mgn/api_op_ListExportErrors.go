// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List export errors.
func (c *Client) ListExportErrors(ctx context.Context, params *ListExportErrorsInput, optFns ...func(*Options)) (*ListExportErrorsOutput, error) {
	if params == nil {
		params = &ListExportErrorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExportErrors", params, optFns, c.addOperationListExportErrorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExportErrorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// List export errors request.
type ListExportErrorsInput struct {

	// List export errors request export id.
	//
	// This member is required.
	ExportID *string

	// List export errors request max results.
	MaxResults int32

	// List export errors request next token.
	NextToken *string

	noSmithyDocumentSerde
}

// List export errors response.
type ListExportErrorsOutput struct {

	// List export errors response items.
	Items []types.ExportTaskError

	// List export errors response next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExportErrorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListExportErrors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListExportErrors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListExportErrors"); err != nil {
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
	if err = addOpListExportErrorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExportErrors(options.Region), middleware.Before); err != nil {
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

// ListExportErrorsAPIClient is a client that implements the ListExportErrors
// operation.
type ListExportErrorsAPIClient interface {
	ListExportErrors(context.Context, *ListExportErrorsInput, ...func(*Options)) (*ListExportErrorsOutput, error)
}

var _ ListExportErrorsAPIClient = (*Client)(nil)

// ListExportErrorsPaginatorOptions is the paginator options for ListExportErrors
type ListExportErrorsPaginatorOptions struct {
	// List export errors request max results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExportErrorsPaginator is a paginator for ListExportErrors
type ListExportErrorsPaginator struct {
	options   ListExportErrorsPaginatorOptions
	client    ListExportErrorsAPIClient
	params    *ListExportErrorsInput
	nextToken *string
	firstPage bool
}

// NewListExportErrorsPaginator returns a new ListExportErrorsPaginator
func NewListExportErrorsPaginator(client ListExportErrorsAPIClient, params *ListExportErrorsInput, optFns ...func(*ListExportErrorsPaginatorOptions)) *ListExportErrorsPaginator {
	if params == nil {
		params = &ListExportErrorsInput{}
	}

	options := ListExportErrorsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExportErrorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExportErrorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExportErrors page.
func (p *ListExportErrorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExportErrorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListExportErrors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExportErrors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListExportErrors",
	}
}
