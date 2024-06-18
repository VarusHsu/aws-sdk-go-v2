// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Device Defender security profiles attached to a target (thing group).
//
// Requires permission to access the [ListSecurityProfilesForTarget] action.
//
// [ListSecurityProfilesForTarget]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListSecurityProfilesForTarget(ctx context.Context, params *ListSecurityProfilesForTargetInput, optFns ...func(*Options)) (*ListSecurityProfilesForTargetOutput, error) {
	if params == nil {
		params = &ListSecurityProfilesForTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityProfilesForTarget", params, optFns, c.addOperationListSecurityProfilesForTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityProfilesForTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityProfilesForTargetInput struct {

	// The ARN of the target (thing group) whose attached security profiles you want
	// to get.
	//
	// This member is required.
	SecurityProfileTargetArn *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// If true, return child groups too.
	Recursive bool

	noSmithyDocumentSerde
}

type ListSecurityProfilesForTargetOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// A list of security profiles and their associated targets.
	SecurityProfileTargetMappings []types.SecurityProfileTargetMapping

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecurityProfilesForTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSecurityProfilesForTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSecurityProfilesForTarget{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSecurityProfilesForTarget"); err != nil {
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
	if err = addOpListSecurityProfilesForTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityProfilesForTarget(options.Region), middleware.Before); err != nil {
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

// ListSecurityProfilesForTargetPaginatorOptions is the paginator options for
// ListSecurityProfilesForTarget
type ListSecurityProfilesForTargetPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecurityProfilesForTargetPaginator is a paginator for
// ListSecurityProfilesForTarget
type ListSecurityProfilesForTargetPaginator struct {
	options   ListSecurityProfilesForTargetPaginatorOptions
	client    ListSecurityProfilesForTargetAPIClient
	params    *ListSecurityProfilesForTargetInput
	nextToken *string
	firstPage bool
}

// NewListSecurityProfilesForTargetPaginator returns a new
// ListSecurityProfilesForTargetPaginator
func NewListSecurityProfilesForTargetPaginator(client ListSecurityProfilesForTargetAPIClient, params *ListSecurityProfilesForTargetInput, optFns ...func(*ListSecurityProfilesForTargetPaginatorOptions)) *ListSecurityProfilesForTargetPaginator {
	if params == nil {
		params = &ListSecurityProfilesForTargetInput{}
	}

	options := ListSecurityProfilesForTargetPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecurityProfilesForTargetPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecurityProfilesForTargetPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecurityProfilesForTarget page.
func (p *ListSecurityProfilesForTargetPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecurityProfilesForTargetOutput, error) {
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
	result, err := p.client.ListSecurityProfilesForTarget(ctx, &params, optFns...)
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

// ListSecurityProfilesForTargetAPIClient is a client that implements the
// ListSecurityProfilesForTarget operation.
type ListSecurityProfilesForTargetAPIClient interface {
	ListSecurityProfilesForTarget(context.Context, *ListSecurityProfilesForTargetInput, ...func(*Options)) (*ListSecurityProfilesForTargetOutput, error)
}

var _ ListSecurityProfilesForTargetAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListSecurityProfilesForTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSecurityProfilesForTarget",
	}
}
