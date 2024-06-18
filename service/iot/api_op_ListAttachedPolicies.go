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

// Lists the policies attached to the specified thing group.
//
// Requires permission to access the [ListAttachedPolicies] action.
//
// [ListAttachedPolicies]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListAttachedPolicies(ctx context.Context, params *ListAttachedPoliciesInput, optFns ...func(*Options)) (*ListAttachedPoliciesOutput, error) {
	if params == nil {
		params = &ListAttachedPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttachedPolicies", params, optFns, c.addOperationListAttachedPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttachedPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttachedPoliciesInput struct {

	// The group or principal for which the policies will be listed. Valid principals
	// are CertificateArn (arn:aws:iot:region:accountId:cert/certificateId),
	// thingGroupArn (arn:aws:iot:region:accountId:thinggroup/groupName) and CognitoId
	// (region:id).
	//
	// This member is required.
	Target *string

	// The token to retrieve the next set of results.
	Marker *string

	// The maximum number of results to be returned per request.
	PageSize *int32

	// When true, recursively list attached policies.
	Recursive bool

	noSmithyDocumentSerde
}

type ListAttachedPoliciesOutput struct {

	// The token to retrieve the next set of results, or ``null`` if there are no more
	// results.
	NextMarker *string

	// The policies.
	Policies []types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAttachedPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAttachedPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAttachedPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAttachedPolicies"); err != nil {
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
	if err = addOpListAttachedPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttachedPolicies(options.Region), middleware.Before); err != nil {
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

// ListAttachedPoliciesPaginatorOptions is the paginator options for
// ListAttachedPolicies
type ListAttachedPoliciesPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttachedPoliciesPaginator is a paginator for ListAttachedPolicies
type ListAttachedPoliciesPaginator struct {
	options   ListAttachedPoliciesPaginatorOptions
	client    ListAttachedPoliciesAPIClient
	params    *ListAttachedPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListAttachedPoliciesPaginator returns a new ListAttachedPoliciesPaginator
func NewListAttachedPoliciesPaginator(client ListAttachedPoliciesAPIClient, params *ListAttachedPoliciesInput, optFns ...func(*ListAttachedPoliciesPaginatorOptions)) *ListAttachedPoliciesPaginator {
	if params == nil {
		params = &ListAttachedPoliciesInput{}
	}

	options := ListAttachedPoliciesPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttachedPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttachedPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAttachedPolicies page.
func (p *ListAttachedPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttachedPoliciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListAttachedPolicies(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListAttachedPoliciesAPIClient is a client that implements the
// ListAttachedPolicies operation.
type ListAttachedPoliciesAPIClient interface {
	ListAttachedPolicies(context.Context, *ListAttachedPoliciesInput, ...func(*Options)) (*ListAttachedPoliciesOutput, error)
}

var _ ListAttachedPoliciesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAttachedPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAttachedPolicies",
	}
}
