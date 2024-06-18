// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list that describes the available WorkSpace bundles.
//
// You can filter the results using either bundle ID or owner, but not both.
func (c *Client) DescribeWorkspaceBundles(ctx context.Context, params *DescribeWorkspaceBundlesInput, optFns ...func(*Options)) (*DescribeWorkspaceBundlesOutput, error) {
	if params == nil {
		params = &DescribeWorkspaceBundlesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkspaceBundles", params, optFns, c.addOperationDescribeWorkspaceBundlesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkspaceBundlesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceBundlesInput struct {

	// The identifiers of the bundles. You cannot combine this parameter with any
	// other filter.
	BundleIds []string

	// The token for the next set of results. (You received this token from a previous
	// call.)
	NextToken *string

	// The owner of the bundles. You cannot combine this parameter with any other
	// filter.
	//
	// To describe the bundles provided by Amazon Web Services, specify AMAZON . To
	// describe the bundles that belong to your account, don't specify a value.
	Owner *string

	noSmithyDocumentSerde
}

type DescribeWorkspaceBundlesOutput struct {

	// Information about the bundles.
	Bundles []types.WorkspaceBundle

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return. This token is valid for one day and must be
	// used within that time frame.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWorkspaceBundlesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaceBundles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaceBundles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeWorkspaceBundles"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceBundles(options.Region), middleware.Before); err != nil {
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

// DescribeWorkspaceBundlesPaginatorOptions is the paginator options for
// DescribeWorkspaceBundles
type DescribeWorkspaceBundlesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeWorkspaceBundlesPaginator is a paginator for DescribeWorkspaceBundles
type DescribeWorkspaceBundlesPaginator struct {
	options   DescribeWorkspaceBundlesPaginatorOptions
	client    DescribeWorkspaceBundlesAPIClient
	params    *DescribeWorkspaceBundlesInput
	nextToken *string
	firstPage bool
}

// NewDescribeWorkspaceBundlesPaginator returns a new
// DescribeWorkspaceBundlesPaginator
func NewDescribeWorkspaceBundlesPaginator(client DescribeWorkspaceBundlesAPIClient, params *DescribeWorkspaceBundlesInput, optFns ...func(*DescribeWorkspaceBundlesPaginatorOptions)) *DescribeWorkspaceBundlesPaginator {
	if params == nil {
		params = &DescribeWorkspaceBundlesInput{}
	}

	options := DescribeWorkspaceBundlesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeWorkspaceBundlesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeWorkspaceBundlesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeWorkspaceBundles page.
func (p *DescribeWorkspaceBundlesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeWorkspaceBundlesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeWorkspaceBundles(ctx, &params, optFns...)
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

// DescribeWorkspaceBundlesAPIClient is a client that implements the
// DescribeWorkspaceBundles operation.
type DescribeWorkspaceBundlesAPIClient interface {
	DescribeWorkspaceBundles(context.Context, *DescribeWorkspaceBundlesInput, ...func(*Options)) (*DescribeWorkspaceBundlesOutput, error)
}

var _ DescribeWorkspaceBundlesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeWorkspaceBundles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeWorkspaceBundles",
	}
}
