// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists user profiles.
func (c *Client) ListUserProfiles(ctx context.Context, params *ListUserProfilesInput, optFns ...func(*Options)) (*ListUserProfilesOutput, error) {
	if params == nil {
		params = &ListUserProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserProfiles", params, optFns, c.addOperationListUserProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserProfilesInput struct {

	// A parameter by which to filter the results.
	DomainIdEquals *string

	// This parameter defines the maximum number of results that can be return in a
	// single response. The MaxResults parameter is an upper bound, not a target. If
	// there are more results available than the value specified, a NextToken is
	// provided in the response. The NextToken indicates that the user should get the
	// next set of results by providing this token as a part of a subsequent call. The
	// default value for MaxResults is 10.
	MaxResults *int32

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// The parameter by which to sort the results. The default is CreationTime.
	SortBy types.UserProfileSortKey

	// The sort order for the results. The default is Ascending.
	SortOrder types.SortOrder

	// A parameter by which to filter the results.
	UserProfileNameContains *string

	noSmithyDocumentSerde
}

type ListUserProfilesOutput struct {

	// If the previous response was truncated, you will receive this token. Use it in
	// your next request to receive the next set of results.
	NextToken *string

	// The list of user profiles.
	UserProfiles []types.UserProfileDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUserProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUserProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUserProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListUserProfiles"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUserProfiles(options.Region), middleware.Before); err != nil {
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

// ListUserProfilesPaginatorOptions is the paginator options for ListUserProfiles
type ListUserProfilesPaginatorOptions struct {
	// This parameter defines the maximum number of results that can be return in a
	// single response. The MaxResults parameter is an upper bound, not a target. If
	// there are more results available than the value specified, a NextToken is
	// provided in the response. The NextToken indicates that the user should get the
	// next set of results by providing this token as a part of a subsequent call. The
	// default value for MaxResults is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUserProfilesPaginator is a paginator for ListUserProfiles
type ListUserProfilesPaginator struct {
	options   ListUserProfilesPaginatorOptions
	client    ListUserProfilesAPIClient
	params    *ListUserProfilesInput
	nextToken *string
	firstPage bool
}

// NewListUserProfilesPaginator returns a new ListUserProfilesPaginator
func NewListUserProfilesPaginator(client ListUserProfilesAPIClient, params *ListUserProfilesInput, optFns ...func(*ListUserProfilesPaginatorOptions)) *ListUserProfilesPaginator {
	if params == nil {
		params = &ListUserProfilesInput{}
	}

	options := ListUserProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUserProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUserProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUserProfiles page.
func (p *ListUserProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUserProfilesOutput, error) {
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
	result, err := p.client.ListUserProfiles(ctx, &params, optFns...)
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

// ListUserProfilesAPIClient is a client that implements the ListUserProfiles
// operation.
type ListUserProfilesAPIClient interface {
	ListUserProfiles(context.Context, *ListUserProfilesInput, ...func(*Options)) (*ListUserProfilesOutput, error)
}

var _ ListUserProfilesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListUserProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListUserProfiles",
	}
}
