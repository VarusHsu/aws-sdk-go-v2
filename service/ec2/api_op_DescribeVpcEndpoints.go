// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes your VPC endpoints.
func (c *Client) DescribeVpcEndpoints(ctx context.Context, params *DescribeVpcEndpointsInput, optFns ...func(*Options)) (*DescribeVpcEndpointsOutput, error) {
	if params == nil {
		params = &DescribeVpcEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcEndpoints", params, optFns, c.addOperationDescribeVpcEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcEndpointsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//   - ip-address-type - The IP address type ( ipv4 | ipv6 ).
	//   - service-name - The name of the service.
	//   - tag : - The key/value combination of a tag assigned to the resource. Use the
	//   tag key in the filter name and the tag value as the filter value. For example,
	//   to find all resources that have a tag with the key Owner and the value TeamA ,
	//   specify tag:Owner for the filter name and TeamA for the filter value.
	//   - tag-key - The key of a tag assigned to the resource. Use this filter to find
	//   all resources assigned a tag with a specific key, regardless of the tag value.
	//   - vpc-id - The ID of the VPC in which the endpoint resides.
	//   - vpc-endpoint-id - The ID of the endpoint.
	//   - vpc-endpoint-state - The state of the endpoint ( pendingAcceptance | pending
	//   | available | deleting | deleted | rejected | failed ).
	//   - vpc-endpoint-type - The type of VPC endpoint ( Interface | Gateway |
	//   GatewayLoadBalancer ).
	Filters []types.Filter

	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// Constraint: If the value is greater than 1,000, we return only 1,000 items.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// prior call.)
	NextToken *string

	// The IDs of the VPC endpoints.
	VpcEndpointIds []string

	noSmithyDocumentSerde
}

type DescribeVpcEndpointsOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the endpoints.
	VpcEndpoints []types.VpcEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpcEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeVpcEndpoints"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpoints(options.Region), middleware.Before); err != nil {
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

// DescribeVpcEndpointsAPIClient is a client that implements the
// DescribeVpcEndpoints operation.
type DescribeVpcEndpointsAPIClient interface {
	DescribeVpcEndpoints(context.Context, *DescribeVpcEndpointsInput, ...func(*Options)) (*DescribeVpcEndpointsOutput, error)
}

var _ DescribeVpcEndpointsAPIClient = (*Client)(nil)

// DescribeVpcEndpointsPaginatorOptions is the paginator options for
// DescribeVpcEndpoints
type DescribeVpcEndpointsPaginatorOptions struct {
	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// Constraint: If the value is greater than 1,000, we return only 1,000 items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVpcEndpointsPaginator is a paginator for DescribeVpcEndpoints
type DescribeVpcEndpointsPaginator struct {
	options   DescribeVpcEndpointsPaginatorOptions
	client    DescribeVpcEndpointsAPIClient
	params    *DescribeVpcEndpointsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVpcEndpointsPaginator returns a new DescribeVpcEndpointsPaginator
func NewDescribeVpcEndpointsPaginator(client DescribeVpcEndpointsAPIClient, params *DescribeVpcEndpointsInput, optFns ...func(*DescribeVpcEndpointsPaginatorOptions)) *DescribeVpcEndpointsPaginator {
	if params == nil {
		params = &DescribeVpcEndpointsInput{}
	}

	options := DescribeVpcEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVpcEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVpcEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVpcEndpoints page.
func (p *DescribeVpcEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVpcEndpointsOutput, error) {
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

	result, err := p.client.DescribeVpcEndpoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeVpcEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeVpcEndpoints",
	}
}
