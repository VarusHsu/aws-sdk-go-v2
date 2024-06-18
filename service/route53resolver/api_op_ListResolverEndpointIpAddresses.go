// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the IP addresses for a specified Resolver endpoint.
func (c *Client) ListResolverEndpointIpAddresses(ctx context.Context, params *ListResolverEndpointIpAddressesInput, optFns ...func(*Options)) (*ListResolverEndpointIpAddressesOutput, error) {
	if params == nil {
		params = &ListResolverEndpointIpAddressesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResolverEndpointIpAddresses", params, optFns, c.addOperationListResolverEndpointIpAddressesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResolverEndpointIpAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverEndpointIpAddressesInput struct {

	// The ID of the Resolver endpoint that you want to get IP addresses for.
	//
	// This member is required.
	ResolverEndpointId *string

	// The maximum number of IP addresses that you want to return in the response to a
	// ListResolverEndpointIpAddresses request. If you don't specify a value for
	// MaxResults , Resolver returns up to 100 IP addresses.
	MaxResults *int32

	// For the first ListResolverEndpointIpAddresses request, omit this value.
	//
	// If the specified Resolver endpoint has more than MaxResults IP addresses, you
	// can submit another ListResolverEndpointIpAddresses request to get the next
	// group of IP addresses. In the next request, specify the value of NextToken from
	// the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResolverEndpointIpAddressesOutput struct {

	// Information about the IP addresses in your VPC that DNS queries originate from
	// (for outbound endpoints) or that you forward DNS queries to (for inbound
	// endpoints).
	IpAddresses []types.IpAddressResponse

	// The value that you specified for MaxResults in the request.
	MaxResults *int32

	// If the specified endpoint has more than MaxResults IP addresses, you can submit
	// another ListResolverEndpointIpAddresses request to get the next group of IP
	// addresses. In the next request, specify the value of NextToken from the
	// previous response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResolverEndpointIpAddressesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverEndpointIpAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverEndpointIpAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResolverEndpointIpAddresses"); err != nil {
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
	if err = addOpListResolverEndpointIpAddressesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverEndpointIpAddresses(options.Region), middleware.Before); err != nil {
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

// ListResolverEndpointIpAddressesPaginatorOptions is the paginator options for
// ListResolverEndpointIpAddresses
type ListResolverEndpointIpAddressesPaginatorOptions struct {
	// The maximum number of IP addresses that you want to return in the response to a
	// ListResolverEndpointIpAddresses request. If you don't specify a value for
	// MaxResults , Resolver returns up to 100 IP addresses.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResolverEndpointIpAddressesPaginator is a paginator for
// ListResolverEndpointIpAddresses
type ListResolverEndpointIpAddressesPaginator struct {
	options   ListResolverEndpointIpAddressesPaginatorOptions
	client    ListResolverEndpointIpAddressesAPIClient
	params    *ListResolverEndpointIpAddressesInput
	nextToken *string
	firstPage bool
}

// NewListResolverEndpointIpAddressesPaginator returns a new
// ListResolverEndpointIpAddressesPaginator
func NewListResolverEndpointIpAddressesPaginator(client ListResolverEndpointIpAddressesAPIClient, params *ListResolverEndpointIpAddressesInput, optFns ...func(*ListResolverEndpointIpAddressesPaginatorOptions)) *ListResolverEndpointIpAddressesPaginator {
	if params == nil {
		params = &ListResolverEndpointIpAddressesInput{}
	}

	options := ListResolverEndpointIpAddressesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResolverEndpointIpAddressesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResolverEndpointIpAddressesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResolverEndpointIpAddresses page.
func (p *ListResolverEndpointIpAddressesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResolverEndpointIpAddressesOutput, error) {
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
	result, err := p.client.ListResolverEndpointIpAddresses(ctx, &params, optFns...)
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

// ListResolverEndpointIpAddressesAPIClient is a client that implements the
// ListResolverEndpointIpAddresses operation.
type ListResolverEndpointIpAddressesAPIClient interface {
	ListResolverEndpointIpAddresses(context.Context, *ListResolverEndpointIpAddressesInput, ...func(*Options)) (*ListResolverEndpointIpAddressesOutput, error)
}

var _ ListResolverEndpointIpAddressesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListResolverEndpointIpAddresses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResolverEndpointIpAddresses",
	}
}
