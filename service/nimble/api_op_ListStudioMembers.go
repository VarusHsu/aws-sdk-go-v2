// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get all users in a given studio membership.
//
// ListStudioMembers only returns admin members.
func (c *Client) ListStudioMembers(ctx context.Context, params *ListStudioMembersInput, optFns ...func(*Options)) (*ListStudioMembersOutput, error) {
	if params == nil {
		params = &ListStudioMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStudioMembers", params, optFns, c.addOperationListStudioMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStudioMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStudioMembersInput struct {

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// The max number of results to return in the response.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStudioMembersOutput struct {

	// A list of admin members.
	Members []types.StudioMembership

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStudioMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStudioMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStudioMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListStudioMembers"); err != nil {
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
	if err = addOpListStudioMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStudioMembers(options.Region), middleware.Before); err != nil {
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

// ListStudioMembersPaginatorOptions is the paginator options for ListStudioMembers
type ListStudioMembersPaginatorOptions struct {
	// The max number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStudioMembersPaginator is a paginator for ListStudioMembers
type ListStudioMembersPaginator struct {
	options   ListStudioMembersPaginatorOptions
	client    ListStudioMembersAPIClient
	params    *ListStudioMembersInput
	nextToken *string
	firstPage bool
}

// NewListStudioMembersPaginator returns a new ListStudioMembersPaginator
func NewListStudioMembersPaginator(client ListStudioMembersAPIClient, params *ListStudioMembersInput, optFns ...func(*ListStudioMembersPaginatorOptions)) *ListStudioMembersPaginator {
	if params == nil {
		params = &ListStudioMembersInput{}
	}

	options := ListStudioMembersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStudioMembersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStudioMembersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStudioMembers page.
func (p *ListStudioMembersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStudioMembersOutput, error) {
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
	result, err := p.client.ListStudioMembers(ctx, &params, optFns...)
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

// ListStudioMembersAPIClient is a client that implements the ListStudioMembers
// operation.
type ListStudioMembersAPIClient interface {
	ListStudioMembers(context.Context, *ListStudioMembersInput, ...func(*Options)) (*ListStudioMembersOutput, error)
}

var _ ListStudioMembersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListStudioMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListStudioMembers",
	}
}
