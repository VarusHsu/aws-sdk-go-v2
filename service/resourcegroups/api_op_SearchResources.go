// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Amazon Web Services resource identifiers that matches the
// specified query. The query uses the same format as a resource query in a CreateGroupor UpdateGroupQuery
// operation.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
//   - resource-groups:SearchResources
//
//   - cloudformation:DescribeStacks
//
//   - cloudformation:ListStackResources
//
//   - tag:GetResources
func (c *Client) SearchResources(ctx context.Context, params *SearchResourcesInput, optFns ...func(*Options)) (*SearchResourcesOutput, error) {
	if params == nil {
		params = &SearchResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchResources", params, optFns, c.addOperationSearchResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchResourcesInput struct {

	// The search query, using the same formats that are supported for resource group
	// definition. For more information, see CreateGroup.
	//
	// This member is required.
	ResourceQuery *types.ResourceQuery

	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value provided by a previous call's
	// NextToken response to indicate where the output should continue from.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchResourcesOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string

	// A list of QueryError objects. Each error contains an ErrorCode and Message .
	//
	// Possible values for ErrorCode :
	//
	//   - CLOUDFORMATION_STACK_INACTIVE
	//
	//   - CLOUDFORMATION_STACK_NOT_EXISTING
	//
	//   - CLOUDFORMATION_STACK_UNASSUMABLE_ROLE
	QueryErrors []types.QueryError

	// The ARNs and resource types of resources that are members of the group that you
	// specified.
	ResourceIdentifiers []types.ResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchResources"); err != nil {
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
	if err = addOpSearchResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchResources(options.Region), middleware.Before); err != nil {
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

// SearchResourcesPaginatorOptions is the paginator options for SearchResources
type SearchResourcesPaginatorOptions struct {
	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchResourcesPaginator is a paginator for SearchResources
type SearchResourcesPaginator struct {
	options   SearchResourcesPaginatorOptions
	client    SearchResourcesAPIClient
	params    *SearchResourcesInput
	nextToken *string
	firstPage bool
}

// NewSearchResourcesPaginator returns a new SearchResourcesPaginator
func NewSearchResourcesPaginator(client SearchResourcesAPIClient, params *SearchResourcesInput, optFns ...func(*SearchResourcesPaginatorOptions)) *SearchResourcesPaginator {
	if params == nil {
		params = &SearchResourcesInput{}
	}

	options := SearchResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchResources page.
func (p *SearchResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchResourcesOutput, error) {
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
	result, err := p.client.SearchResources(ctx, &params, optFns...)
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

// SearchResourcesAPIClient is a client that implements the SearchResources
// operation.
type SearchResourcesAPIClient interface {
	SearchResources(context.Context, *SearchResourcesInput, ...func(*Options)) (*SearchResourcesOutput, error)
}

var _ SearchResourcesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opSearchResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchResources",
	}
}
