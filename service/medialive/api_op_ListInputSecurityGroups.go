// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Produces a list of Input Security Groups for an account
func (c *Client) ListInputSecurityGroups(ctx context.Context, params *ListInputSecurityGroupsInput, optFns ...func(*Options)) (*ListInputSecurityGroupsOutput, error) {
	if params == nil {
		params = &ListInputSecurityGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInputSecurityGroups", params, optFns, c.addOperationListInputSecurityGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInputSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListInputSecurityGroupsRequest
type ListInputSecurityGroupsInput struct {

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// Placeholder documentation for __string
	NextToken *string

	noSmithyDocumentSerde
}

// Placeholder documentation for ListInputSecurityGroupsResponse
type ListInputSecurityGroupsOutput struct {

	// List of input security groups
	InputSecurityGroups []types.InputSecurityGroup

	// Placeholder documentation for __string
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInputSecurityGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInputSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInputSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInputSecurityGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInputSecurityGroups(options.Region), middleware.Before); err != nil {
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

// ListInputSecurityGroupsPaginatorOptions is the paginator options for
// ListInputSecurityGroups
type ListInputSecurityGroupsPaginatorOptions struct {
	// Placeholder documentation for MaxResults
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInputSecurityGroupsPaginator is a paginator for ListInputSecurityGroups
type ListInputSecurityGroupsPaginator struct {
	options   ListInputSecurityGroupsPaginatorOptions
	client    ListInputSecurityGroupsAPIClient
	params    *ListInputSecurityGroupsInput
	nextToken *string
	firstPage bool
}

// NewListInputSecurityGroupsPaginator returns a new
// ListInputSecurityGroupsPaginator
func NewListInputSecurityGroupsPaginator(client ListInputSecurityGroupsAPIClient, params *ListInputSecurityGroupsInput, optFns ...func(*ListInputSecurityGroupsPaginatorOptions)) *ListInputSecurityGroupsPaginator {
	if params == nil {
		params = &ListInputSecurityGroupsInput{}
	}

	options := ListInputSecurityGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInputSecurityGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInputSecurityGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInputSecurityGroups page.
func (p *ListInputSecurityGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInputSecurityGroupsOutput, error) {
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
	result, err := p.client.ListInputSecurityGroups(ctx, &params, optFns...)
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

// ListInputSecurityGroupsAPIClient is a client that implements the
// ListInputSecurityGroups operation.
type ListInputSecurityGroupsAPIClient interface {
	ListInputSecurityGroups(context.Context, *ListInputSecurityGroupsInput, ...func(*Options)) (*ListInputSecurityGroupsOutput, error)
}

var _ ListInputSecurityGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListInputSecurityGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInputSecurityGroups",
	}
}
