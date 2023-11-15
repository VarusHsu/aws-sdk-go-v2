// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the Auto Scaling instances in the account and Region.
func (c *Client) DescribeAutoScalingInstances(ctx context.Context, params *DescribeAutoScalingInstancesInput, optFns ...func(*Options)) (*DescribeAutoScalingInstancesOutput, error) {
	if params == nil {
		params = &DescribeAutoScalingInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutoScalingInstances", params, optFns, c.addOperationDescribeAutoScalingInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutoScalingInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutoScalingInstancesInput struct {

	// The IDs of the instances. If you omit this property, all Auto Scaling instances
	// are described. If you specify an ID that does not exist, it is ignored with no
	// error. Array Members: Maximum number of 50 items.
	InstanceIds []string

	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 50 .
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAutoScalingInstancesOutput struct {

	// The instances.
	AutoScalingInstances []types.AutoScalingInstanceDetails

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAutoScalingInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAutoScalingInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAutoScalingInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAutoScalingInstances"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutoScalingInstances(options.Region), middleware.Before); err != nil {
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

// DescribeAutoScalingInstancesAPIClient is a client that implements the
// DescribeAutoScalingInstances operation.
type DescribeAutoScalingInstancesAPIClient interface {
	DescribeAutoScalingInstances(context.Context, *DescribeAutoScalingInstancesInput, ...func(*Options)) (*DescribeAutoScalingInstancesOutput, error)
}

var _ DescribeAutoScalingInstancesAPIClient = (*Client)(nil)

// DescribeAutoScalingInstancesPaginatorOptions is the paginator options for
// DescribeAutoScalingInstances
type DescribeAutoScalingInstancesPaginatorOptions struct {
	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 50 .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAutoScalingInstancesPaginator is a paginator for
// DescribeAutoScalingInstances
type DescribeAutoScalingInstancesPaginator struct {
	options   DescribeAutoScalingInstancesPaginatorOptions
	client    DescribeAutoScalingInstancesAPIClient
	params    *DescribeAutoScalingInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeAutoScalingInstancesPaginator returns a new
// DescribeAutoScalingInstancesPaginator
func NewDescribeAutoScalingInstancesPaginator(client DescribeAutoScalingInstancesAPIClient, params *DescribeAutoScalingInstancesInput, optFns ...func(*DescribeAutoScalingInstancesPaginatorOptions)) *DescribeAutoScalingInstancesPaginator {
	if params == nil {
		params = &DescribeAutoScalingInstancesInput{}
	}

	options := DescribeAutoScalingInstancesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAutoScalingInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAutoScalingInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAutoScalingInstances page.
func (p *DescribeAutoScalingInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAutoScalingInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeAutoScalingInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAutoScalingInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAutoScalingInstances",
	}
}
