// Code generated by smithy-go-codegen DO NOT EDIT.

package medicalimaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medicalimaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List image set versions.
func (c *Client) ListImageSetVersions(ctx context.Context, params *ListImageSetVersionsInput, optFns ...func(*Options)) (*ListImageSetVersionsOutput, error) {
	if params == nil {
		params = &ListImageSetVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImageSetVersions", params, optFns, c.addOperationListImageSetVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImageSetVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImageSetVersionsInput struct {

	// The data store identifier.
	//
	// This member is required.
	DatastoreId *string

	// The image set identifier.
	//
	// This member is required.
	ImageSetId *string

	// The max results count.
	MaxResults *int32

	// The pagination token used to request the list of image set versions on the next
	// page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListImageSetVersionsOutput struct {

	// Lists all properties associated with an image set.
	//
	// This member is required.
	ImageSetPropertiesList []types.ImageSetProperties

	// The pagination token used to retrieve the list of image set versions on the
	// next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImageSetVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImageSetVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImageSetVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListImageSetVersions"); err != nil {
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
	if err = addEndpointPrefix_opListImageSetVersionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListImageSetVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImageSetVersions(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListImageSetVersionsMiddleware struct {
}

func (*endpointPrefix_opListImageSetVersionsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListImageSetVersionsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "runtime-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListImageSetVersionsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListImageSetVersionsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListImageSetVersionsAPIClient is a client that implements the
// ListImageSetVersions operation.
type ListImageSetVersionsAPIClient interface {
	ListImageSetVersions(context.Context, *ListImageSetVersionsInput, ...func(*Options)) (*ListImageSetVersionsOutput, error)
}

var _ ListImageSetVersionsAPIClient = (*Client)(nil)

// ListImageSetVersionsPaginatorOptions is the paginator options for
// ListImageSetVersions
type ListImageSetVersionsPaginatorOptions struct {
	// The max results count.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImageSetVersionsPaginator is a paginator for ListImageSetVersions
type ListImageSetVersionsPaginator struct {
	options   ListImageSetVersionsPaginatorOptions
	client    ListImageSetVersionsAPIClient
	params    *ListImageSetVersionsInput
	nextToken *string
	firstPage bool
}

// NewListImageSetVersionsPaginator returns a new ListImageSetVersionsPaginator
func NewListImageSetVersionsPaginator(client ListImageSetVersionsAPIClient, params *ListImageSetVersionsInput, optFns ...func(*ListImageSetVersionsPaginatorOptions)) *ListImageSetVersionsPaginator {
	if params == nil {
		params = &ListImageSetVersionsInput{}
	}

	options := ListImageSetVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImageSetVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImageSetVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImageSetVersions page.
func (p *ListImageSetVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImageSetVersionsOutput, error) {
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

	result, err := p.client.ListImageSetVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImageSetVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListImageSetVersions",
	}
}
