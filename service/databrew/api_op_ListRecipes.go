// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the DataBrew recipes that are defined.
func (c *Client) ListRecipes(ctx context.Context, params *ListRecipesInput, optFns ...func(*Options)) (*ListRecipesOutput, error) {
	if params == nil {
		params = &ListRecipesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecipes", params, optFns, c.addOperationListRecipesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecipesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecipesInput struct {

	// The maximum number of results to return in this request.
	MaxResults *int32

	// The token returned by a previous call to retrieve the next set of results.
	NextToken *string

	// Return only those recipes with a version identifier of LATEST_WORKING or
	// LATEST_PUBLISHED . If RecipeVersion is omitted, ListRecipes returns all of the
	// LATEST_PUBLISHED recipe versions. Valid values: LATEST_WORKING |
	// LATEST_PUBLISHED
	RecipeVersion *string

	noSmithyDocumentSerde
}

type ListRecipesOutput struct {

	// A list of recipes that are defined.
	//
	// This member is required.
	Recipes []types.Recipe

	// A token that you can use in a subsequent call to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecipesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecipes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecipes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRecipes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecipes(options.Region), middleware.Before); err != nil {
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

// ListRecipesAPIClient is a client that implements the ListRecipes operation.
type ListRecipesAPIClient interface {
	ListRecipes(context.Context, *ListRecipesInput, ...func(*Options)) (*ListRecipesOutput, error)
}

var _ ListRecipesAPIClient = (*Client)(nil)

// ListRecipesPaginatorOptions is the paginator options for ListRecipes
type ListRecipesPaginatorOptions struct {
	// The maximum number of results to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecipesPaginator is a paginator for ListRecipes
type ListRecipesPaginator struct {
	options   ListRecipesPaginatorOptions
	client    ListRecipesAPIClient
	params    *ListRecipesInput
	nextToken *string
	firstPage bool
}

// NewListRecipesPaginator returns a new ListRecipesPaginator
func NewListRecipesPaginator(client ListRecipesAPIClient, params *ListRecipesInput, optFns ...func(*ListRecipesPaginatorOptions)) *ListRecipesPaginator {
	if params == nil {
		params = &ListRecipesInput{}
	}

	options := ListRecipesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecipesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecipesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecipes page.
func (p *ListRecipesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecipesOutput, error) {
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

	result, err := p.client.ListRecipes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecipes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRecipes",
	}
}
