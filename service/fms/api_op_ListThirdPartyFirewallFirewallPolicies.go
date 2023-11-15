// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of all of the third-party firewall policies that are
// associated with the third-party firewall administrator's account.
func (c *Client) ListThirdPartyFirewallFirewallPolicies(ctx context.Context, params *ListThirdPartyFirewallFirewallPoliciesInput, optFns ...func(*Options)) (*ListThirdPartyFirewallFirewallPoliciesOutput, error) {
	if params == nil {
		params = &ListThirdPartyFirewallFirewallPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThirdPartyFirewallFirewallPolicies", params, optFns, c.addOperationListThirdPartyFirewallFirewallPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThirdPartyFirewallFirewallPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThirdPartyFirewallFirewallPoliciesInput struct {

	// The maximum number of third-party firewall policies that you want Firewall
	// Manager to return. If the specified third-party firewall vendor is associated
	// with more than MaxResults firewall policies, the response includes a NextToken
	// element. NextToken contains an encrypted token that identifies the first
	// third-party firewall policies that Firewall Manager will return if you submit
	// another request.
	//
	// This member is required.
	MaxResults *int32

	// The name of the third-party firewall vendor.
	//
	// This member is required.
	ThirdPartyFirewall types.ThirdPartyFirewall

	// If the previous response included a NextToken element, the specified
	// third-party firewall vendor is associated with more third-party firewall
	// policies. To get more third-party firewall policies, submit another
	// ListThirdPartyFirewallFirewallPoliciesRequest request. For the value of
	// NextToken , specify the value of NextToken from the previous response. If the
	// previous response didn't include a NextToken element, there are no more
	// third-party firewall policies to get.
	NextToken *string

	noSmithyDocumentSerde
}

type ListThirdPartyFirewallFirewallPoliciesOutput struct {

	// The value that you will use for NextToken in the next
	// ListThirdPartyFirewallFirewallPolicies request.
	NextToken *string

	// A list that contains one ThirdPartyFirewallFirewallPolicies element for each
	// third-party firewall policies that the specified third-party firewall vendor is
	// associated with. Each ThirdPartyFirewallFirewallPolicies element contains the
	// firewall policy name and ID.
	ThirdPartyFirewallFirewallPolicies []types.ThirdPartyFirewallFirewallPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThirdPartyFirewallFirewallPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListThirdPartyFirewallFirewallPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListThirdPartyFirewallFirewallPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListThirdPartyFirewallFirewallPolicies"); err != nil {
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
	if err = addOpListThirdPartyFirewallFirewallPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThirdPartyFirewallFirewallPolicies(options.Region), middleware.Before); err != nil {
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

// ListThirdPartyFirewallFirewallPoliciesAPIClient is a client that implements the
// ListThirdPartyFirewallFirewallPolicies operation.
type ListThirdPartyFirewallFirewallPoliciesAPIClient interface {
	ListThirdPartyFirewallFirewallPolicies(context.Context, *ListThirdPartyFirewallFirewallPoliciesInput, ...func(*Options)) (*ListThirdPartyFirewallFirewallPoliciesOutput, error)
}

var _ ListThirdPartyFirewallFirewallPoliciesAPIClient = (*Client)(nil)

// ListThirdPartyFirewallFirewallPoliciesPaginatorOptions is the paginator options
// for ListThirdPartyFirewallFirewallPolicies
type ListThirdPartyFirewallFirewallPoliciesPaginatorOptions struct {
	// The maximum number of third-party firewall policies that you want Firewall
	// Manager to return. If the specified third-party firewall vendor is associated
	// with more than MaxResults firewall policies, the response includes a NextToken
	// element. NextToken contains an encrypted token that identifies the first
	// third-party firewall policies that Firewall Manager will return if you submit
	// another request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThirdPartyFirewallFirewallPoliciesPaginator is a paginator for
// ListThirdPartyFirewallFirewallPolicies
type ListThirdPartyFirewallFirewallPoliciesPaginator struct {
	options   ListThirdPartyFirewallFirewallPoliciesPaginatorOptions
	client    ListThirdPartyFirewallFirewallPoliciesAPIClient
	params    *ListThirdPartyFirewallFirewallPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListThirdPartyFirewallFirewallPoliciesPaginator returns a new
// ListThirdPartyFirewallFirewallPoliciesPaginator
func NewListThirdPartyFirewallFirewallPoliciesPaginator(client ListThirdPartyFirewallFirewallPoliciesAPIClient, params *ListThirdPartyFirewallFirewallPoliciesInput, optFns ...func(*ListThirdPartyFirewallFirewallPoliciesPaginatorOptions)) *ListThirdPartyFirewallFirewallPoliciesPaginator {
	if params == nil {
		params = &ListThirdPartyFirewallFirewallPoliciesInput{}
	}

	options := ListThirdPartyFirewallFirewallPoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThirdPartyFirewallFirewallPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThirdPartyFirewallFirewallPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThirdPartyFirewallFirewallPolicies page.
func (p *ListThirdPartyFirewallFirewallPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThirdPartyFirewallFirewallPoliciesOutput, error) {
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

	result, err := p.client.ListThirdPartyFirewallFirewallPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListThirdPartyFirewallFirewallPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListThirdPartyFirewallFirewallPolicies",
	}
}
