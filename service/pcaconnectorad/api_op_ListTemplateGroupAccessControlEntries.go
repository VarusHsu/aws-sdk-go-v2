// Code generated by smithy-go-codegen DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorad/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists group access control entries you created.
func (c *Client) ListTemplateGroupAccessControlEntries(ctx context.Context, params *ListTemplateGroupAccessControlEntriesInput, optFns ...func(*Options)) (*ListTemplateGroupAccessControlEntriesOutput, error) {
	if params == nil {
		params = &ListTemplateGroupAccessControlEntriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTemplateGroupAccessControlEntries", params, optFns, c.addOperationListTemplateGroupAccessControlEntriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTemplateGroupAccessControlEntriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTemplateGroupAccessControlEntriesInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called CreateTemplate (https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html)
	// .
	//
	// This member is required.
	TemplateArn *string

	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response on each page. If additional items exist beyond
	// the number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	MaxResults *int32

	// Use this parameter when paginating results in a subsequent request after you
	// receive a response with truncated results. Set it to the value of the NextToken
	// parameter from the response you just received.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTemplateGroupAccessControlEntriesOutput struct {

	// An access control entry grants or denies permission to an Active Directory
	// group to enroll certificates for a template.
	AccessControlEntries []types.AccessControlEntrySummary

	// Use this parameter when paginating results in a subsequent request after you
	// receive a response with truncated results. Set it to the value of the NextToken
	// parameter from the response you just received.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTemplateGroupAccessControlEntriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTemplateGroupAccessControlEntries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTemplateGroupAccessControlEntries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTemplateGroupAccessControlEntries"); err != nil {
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
	if err = addOpListTemplateGroupAccessControlEntriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTemplateGroupAccessControlEntries(options.Region), middleware.Before); err != nil {
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

// ListTemplateGroupAccessControlEntriesAPIClient is a client that implements the
// ListTemplateGroupAccessControlEntries operation.
type ListTemplateGroupAccessControlEntriesAPIClient interface {
	ListTemplateGroupAccessControlEntries(context.Context, *ListTemplateGroupAccessControlEntriesInput, ...func(*Options)) (*ListTemplateGroupAccessControlEntriesOutput, error)
}

var _ ListTemplateGroupAccessControlEntriesAPIClient = (*Client)(nil)

// ListTemplateGroupAccessControlEntriesPaginatorOptions is the paginator options
// for ListTemplateGroupAccessControlEntries
type ListTemplateGroupAccessControlEntriesPaginatorOptions struct {
	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response on each page. If additional items exist beyond
	// the number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTemplateGroupAccessControlEntriesPaginator is a paginator for
// ListTemplateGroupAccessControlEntries
type ListTemplateGroupAccessControlEntriesPaginator struct {
	options   ListTemplateGroupAccessControlEntriesPaginatorOptions
	client    ListTemplateGroupAccessControlEntriesAPIClient
	params    *ListTemplateGroupAccessControlEntriesInput
	nextToken *string
	firstPage bool
}

// NewListTemplateGroupAccessControlEntriesPaginator returns a new
// ListTemplateGroupAccessControlEntriesPaginator
func NewListTemplateGroupAccessControlEntriesPaginator(client ListTemplateGroupAccessControlEntriesAPIClient, params *ListTemplateGroupAccessControlEntriesInput, optFns ...func(*ListTemplateGroupAccessControlEntriesPaginatorOptions)) *ListTemplateGroupAccessControlEntriesPaginator {
	if params == nil {
		params = &ListTemplateGroupAccessControlEntriesInput{}
	}

	options := ListTemplateGroupAccessControlEntriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTemplateGroupAccessControlEntriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTemplateGroupAccessControlEntriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTemplateGroupAccessControlEntries page.
func (p *ListTemplateGroupAccessControlEntriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTemplateGroupAccessControlEntriesOutput, error) {
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

	result, err := p.client.ListTemplateGroupAccessControlEntries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTemplateGroupAccessControlEntries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTemplateGroupAccessControlEntries",
	}
}
