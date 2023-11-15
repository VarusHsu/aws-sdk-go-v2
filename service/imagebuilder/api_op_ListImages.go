// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of images that you have access to. Newly created images can
// take up to two minutes to appear in the ListImages API Results.
func (c *Client) ListImages(ctx context.Context, params *ListImagesInput, optFns ...func(*Options)) (*ListImagesOutput, error) {
	if params == nil {
		params = &ListImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImages", params, optFns, c.addOperationListImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImagesInput struct {

	// Requests a list of images with a specific recipe name.
	ByName bool

	// Use the following filters to streamline results:
	//   - name
	//   - osVersion
	//   - platform
	//   - type
	//   - version
	Filters []types.Filter

	// Includes deprecated images in the response list.
	IncludeDeprecated *bool

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	// The owner defines which images you want to list. By default, this request will
	// only show images owned by your account. You can use this field to specify if you
	// want to view images owned by yourself, by Amazon, or those images that have been
	// shared with you by other customers.
	Owner types.Ownership

	noSmithyDocumentSerde
}

type ListImagesOutput struct {

	// The list of image semantic versions. The semantic version has four nodes: ../.
	// You can assign values for the first three, and can filter on all of them.
	// Filtering: With semantic versioning, you have the flexibility to use wildcards
	// (x) to specify the most recent versions or nodes when selecting the base image
	// or components for your recipe. When you use a wildcard in any node, all nodes to
	// the right of the first wildcard must also be wildcards.
	ImageVersionList []types.ImageVersion

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service has'ot included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImages{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListImages"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImages(options.Region), middleware.Before); err != nil {
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

// ListImagesAPIClient is a client that implements the ListImages operation.
type ListImagesAPIClient interface {
	ListImages(context.Context, *ListImagesInput, ...func(*Options)) (*ListImagesOutput, error)
}

var _ ListImagesAPIClient = (*Client)(nil)

// ListImagesPaginatorOptions is the paginator options for ListImages
type ListImagesPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImagesPaginator is a paginator for ListImages
type ListImagesPaginator struct {
	options   ListImagesPaginatorOptions
	client    ListImagesAPIClient
	params    *ListImagesInput
	nextToken *string
	firstPage bool
}

// NewListImagesPaginator returns a new ListImagesPaginator
func NewListImagesPaginator(client ListImagesAPIClient, params *ListImagesInput, optFns ...func(*ListImagesPaginatorOptions)) *ListImagesPaginator {
	if params == nil {
		params = &ListImagesInput{}
	}

	options := ListImagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImages page.
func (p *ListImagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImagesOutput, error) {
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

	result, err := p.client.ListImages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListImages",
	}
}
