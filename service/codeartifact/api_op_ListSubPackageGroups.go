// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of direct children of the specified package group.
//
// For information package group hierarchy, see [Package group definition syntax and matching behavior] in the CodeArtifact User Guide.
//
// [Package group definition syntax and matching behavior]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-group-definition-syntax-matching-behavior.html
func (c *Client) ListSubPackageGroups(ctx context.Context, params *ListSubPackageGroupsInput, optFns ...func(*Options)) (*ListSubPackageGroupsOutput, error) {
	if params == nil {
		params = &ListSubPackageGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubPackageGroups", params, optFns, c.addOperationListSubPackageGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubPackageGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSubPackageGroupsInput struct {

	//  The name of the domain which contains the package group from which to list sub
	// package groups.
	//
	// This member is required.
	Domain *string

	//  The pattern of the package group from which to list sub package groups.
	//
	// This member is required.
	PackageGroup *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	//  The maximum number of results to return per page.
	MaxResults *int32

	//  The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSubPackageGroupsOutput struct {

	//  If there are additional results, this is the token for the next set of
	// results.
	NextToken *string

	//  A list of sub package groups for the requested package group.
	PackageGroups []types.PackageGroupSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSubPackageGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSubPackageGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSubPackageGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSubPackageGroups"); err != nil {
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
	if err = addOpListSubPackageGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubPackageGroups(options.Region), middleware.Before); err != nil {
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

// ListSubPackageGroupsPaginatorOptions is the paginator options for
// ListSubPackageGroups
type ListSubPackageGroupsPaginatorOptions struct {
	//  The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSubPackageGroupsPaginator is a paginator for ListSubPackageGroups
type ListSubPackageGroupsPaginator struct {
	options   ListSubPackageGroupsPaginatorOptions
	client    ListSubPackageGroupsAPIClient
	params    *ListSubPackageGroupsInput
	nextToken *string
	firstPage bool
}

// NewListSubPackageGroupsPaginator returns a new ListSubPackageGroupsPaginator
func NewListSubPackageGroupsPaginator(client ListSubPackageGroupsAPIClient, params *ListSubPackageGroupsInput, optFns ...func(*ListSubPackageGroupsPaginatorOptions)) *ListSubPackageGroupsPaginator {
	if params == nil {
		params = &ListSubPackageGroupsInput{}
	}

	options := ListSubPackageGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSubPackageGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSubPackageGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSubPackageGroups page.
func (p *ListSubPackageGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSubPackageGroupsOutput, error) {
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
	result, err := p.client.ListSubPackageGroups(ctx, &params, optFns...)
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

// ListSubPackageGroupsAPIClient is a client that implements the
// ListSubPackageGroups operation.
type ListSubPackageGroupsAPIClient interface {
	ListSubPackageGroups(context.Context, *ListSubPackageGroupsInput, ...func(*Options)) (*ListSubPackageGroupsOutput, error)
}

var _ ListSubPackageGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSubPackageGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSubPackageGroups",
	}
}
