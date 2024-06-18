// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves utilization statistics for one or more fleets. Utilization data
// provides a snapshot of how the fleet's hosting resources are currently being
// used. For fleets with remote locations, this operation retrieves data for the
// fleet's home Region only. See [DescribeFleetLocationUtilization]to get utilization statistics for a fleet's
// remote locations.
//
// This operation can be used in the following ways:
//
//   - To get utilization data for one or more specific fleets, provide a list of
//     fleet IDs or fleet ARNs.
//
//   - To get utilization data for all fleets, do not provide a fleet identifier.
//
// When requesting multiple fleets, use the pagination parameters to retrieve
// results as a set of sequential pages.
//
// If successful, a [FleetUtilization] object is returned for each requested fleet ID, unless the
// fleet identifier is not found. Each fleet utilization object includes a Location
// property, which is set to the fleet's home Region.
//
// Some API operations may limit the number of fleet IDs allowed in one request.
// If a request exceeds this limit, the request fails and the error message
// includes the maximum allowed.
//
// # Learn more
//
// [Setting up Amazon GameLift Fleets]
//
// [GameLift Metrics for Fleets]
//
// [FleetUtilization]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_FleetUtilization.html
// [Setting up Amazon GameLift Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
// [DescribeFleetLocationUtilization]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationUtilization.html
// [GameLift Metrics for Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html#gamelift-metrics-fleet
func (c *Client) DescribeFleetUtilization(ctx context.Context, params *DescribeFleetUtilizationInput, optFns ...func(*Options)) (*DescribeFleetUtilizationOutput, error) {
	if params == nil {
		params = &DescribeFleetUtilizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetUtilization", params, optFns, c.addOperationDescribeFleetUtilizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetUtilizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetUtilizationInput struct {

	// A unique identifier for the fleet to retrieve utilization data for. You can use
	// either the fleet ID or ARN value. To retrieve attributes for all current fleets,
	// do not include this parameter.
	FleetIds []string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is ignored when the
	// request specifies one or a list of fleet IDs.
	Limit *int32

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value. This parameter is
	// ignored when the request specifies one or a list of fleet IDs.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetUtilizationOutput struct {

	// A collection of objects containing utilization information for each requested
	// fleet ID. Utilization objects are returned only for fleets that currently exist.
	FleetUtilization []types.FleetUtilization

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetUtilizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetUtilization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetUtilization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFleetUtilization"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetUtilization(options.Region), middleware.Before); err != nil {
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

// DescribeFleetUtilizationPaginatorOptions is the paginator options for
// DescribeFleetUtilization
type DescribeFleetUtilizationPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is ignored when the
	// request specifies one or a list of fleet IDs.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetUtilizationPaginator is a paginator for DescribeFleetUtilization
type DescribeFleetUtilizationPaginator struct {
	options   DescribeFleetUtilizationPaginatorOptions
	client    DescribeFleetUtilizationAPIClient
	params    *DescribeFleetUtilizationInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetUtilizationPaginator returns a new
// DescribeFleetUtilizationPaginator
func NewDescribeFleetUtilizationPaginator(client DescribeFleetUtilizationAPIClient, params *DescribeFleetUtilizationInput, optFns ...func(*DescribeFleetUtilizationPaginatorOptions)) *DescribeFleetUtilizationPaginator {
	if params == nil {
		params = &DescribeFleetUtilizationInput{}
	}

	options := DescribeFleetUtilizationPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetUtilizationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetUtilizationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleetUtilization page.
func (p *DescribeFleetUtilizationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetUtilizationOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeFleetUtilization(ctx, &params, optFns...)
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

// DescribeFleetUtilizationAPIClient is a client that implements the
// DescribeFleetUtilization operation.
type DescribeFleetUtilizationAPIClient interface {
	DescribeFleetUtilization(context.Context, *DescribeFleetUtilizationInput, ...func(*Options)) (*DescribeFleetUtilizationOutput, error)
}

var _ DescribeFleetUtilizationAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeFleetUtilization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFleetUtilization",
	}
}
