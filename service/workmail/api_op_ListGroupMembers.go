// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an overview of the members of a group. Users and groups can be members
// of a group.
func (c *Client) ListGroupMembers(ctx context.Context, params *ListGroupMembersInput, optFns ...func(*Options)) (*ListGroupMembersOutput, error) {
	if params == nil {
		params = &ListGroupMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroupMembers", params, optFns, c.addOperationListGroupMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupMembersInput struct {

	// The identifier for the group to which the members (users or groups) are
	// associated.
	//
	// The identifier can accept GroupId, Groupname, or email. The following identity
	// formats are available:
	//
	//   - Group ID: 12345678-1234-1234-1234-123456789012 or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//
	//   - Email address: group@domain.tld
	//
	//   - Group name: group
	//
	// This member is required.
	GroupId *string

	// The identifier for the organization under which the group exists.
	//
	// This member is required.
	OrganizationId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	//  The token to use to retrieve the next page of results. The first call does not
	// contain any tokens.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupMembersOutput struct {

	// The members associated to the group.
	Members []types.Member

	// The token to use to retrieve the next page of results. The first call does not
	// contain any tokens.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroupMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGroupMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGroupMembers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGroupMembers"); err != nil {
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
	if err = addOpListGroupMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupMembers(options.Region), middleware.Before); err != nil {
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

// ListGroupMembersPaginatorOptions is the paginator options for ListGroupMembers
type ListGroupMembersPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupMembersPaginator is a paginator for ListGroupMembers
type ListGroupMembersPaginator struct {
	options   ListGroupMembersPaginatorOptions
	client    ListGroupMembersAPIClient
	params    *ListGroupMembersInput
	nextToken *string
	firstPage bool
}

// NewListGroupMembersPaginator returns a new ListGroupMembersPaginator
func NewListGroupMembersPaginator(client ListGroupMembersAPIClient, params *ListGroupMembersInput, optFns ...func(*ListGroupMembersPaginatorOptions)) *ListGroupMembersPaginator {
	if params == nil {
		params = &ListGroupMembersInput{}
	}

	options := ListGroupMembersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupMembersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupMembersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroupMembers page.
func (p *ListGroupMembersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupMembersOutput, error) {
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
	result, err := p.client.ListGroupMembers(ctx, &params, optFns...)
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

// ListGroupMembersAPIClient is a client that implements the ListGroupMembers
// operation.
type ListGroupMembersAPIClient interface {
	ListGroupMembers(context.Context, *ListGroupMembersInput, ...func(*Options)) (*ListGroupMembersOutput, error)
}

var _ ListGroupMembersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListGroupMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGroupMembers",
	}
}
