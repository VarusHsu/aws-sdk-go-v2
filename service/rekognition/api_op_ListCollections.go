// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns list of collection IDs in your account. If the result is truncated, the
// response also provides a NextToken that you can use in the subsequent request
// to fetch the next set of collection IDs.
//
// For an example, see Listing collections in the Amazon Rekognition Developer
// Guide.
//
// This operation requires permissions to perform the rekognition:ListCollections
// action.
func (c *Client) ListCollections(ctx context.Context, params *ListCollectionsInput, optFns ...func(*Options)) (*ListCollectionsOutput, error) {
	if params == nil {
		params = &ListCollectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCollections", params, optFns, c.addOperationListCollectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCollectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCollectionsInput struct {

	// Maximum number of collection IDs to return.
	MaxResults *int32

	// Pagination token from the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCollectionsOutput struct {

	// An array of collection IDs.
	CollectionIds []string

	// Version numbers of the face detection models associated with the collections in
	// the array CollectionIds . For example, the value of FaceModelVersions[2] is the
	// version number for the face detection model used by the collection in
	// CollectionId[2] .
	FaceModelVersions []string

	// If the result is truncated, the response provides a NextToken that you can use
	// in the subsequent request to fetch the next set of collection IDs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCollectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCollections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCollections{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCollections"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCollections(options.Region), middleware.Before); err != nil {
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

// ListCollectionsPaginatorOptions is the paginator options for ListCollections
type ListCollectionsPaginatorOptions struct {
	// Maximum number of collection IDs to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCollectionsPaginator is a paginator for ListCollections
type ListCollectionsPaginator struct {
	options   ListCollectionsPaginatorOptions
	client    ListCollectionsAPIClient
	params    *ListCollectionsInput
	nextToken *string
	firstPage bool
}

// NewListCollectionsPaginator returns a new ListCollectionsPaginator
func NewListCollectionsPaginator(client ListCollectionsAPIClient, params *ListCollectionsInput, optFns ...func(*ListCollectionsPaginatorOptions)) *ListCollectionsPaginator {
	if params == nil {
		params = &ListCollectionsInput{}
	}

	options := ListCollectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCollectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCollectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCollections page.
func (p *ListCollectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCollectionsOutput, error) {
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
	result, err := p.client.ListCollections(ctx, &params, optFns...)
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

// ListCollectionsAPIClient is a client that implements the ListCollections
// operation.
type ListCollectionsAPIClient interface {
	ListCollections(context.Context, *ListCollectionsInput, ...func(*Options)) (*ListCollectionsOutput, error)
}

var _ ListCollectionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCollections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCollections",
	}
}
