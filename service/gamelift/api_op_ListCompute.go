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

//	This operation has been expanded to use with the Amazon GameLift containers
//
// feature, which is currently in public preview.
//
// Retrieves information on the compute resources in an Amazon GameLift fleet.
//
// To request a list of computes, specify the fleet ID. Use the pagination
// parameters to retrieve results in a set of sequential pages.
//
// You can filter the result set by location.
//
// If successful, this operation returns information on all computes in the
// requested fleet. Depending on the fleet's compute type, the result includes the
// following information:
//
//   - For EC2 fleets, this operation returns information about the EC2 instance.
//     Compute names are instance IDs.
//
//   - For ANYWHERE fleets, this operation returns the compute names and details
//     provided when the compute was registered with RegisterCompute . The
//     GameLiftServiceSdkEndpoint or GameLiftAgentEndpoint is included.
//
//   - For CONTAINER fleets, this operation returns information about containers
//     that are registered as computes, and the instances they're running on. Compute
//     names are container names.
func (c *Client) ListCompute(ctx context.Context, params *ListComputeInput, optFns ...func(*Options)) (*ListComputeOutput, error) {
	if params == nil {
		params = &ListComputeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCompute", params, optFns, c.addOperationListComputeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComputeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComputeInput struct {

	// A unique identifier for the fleet to retrieve compute resources for.
	//
	// This member is required.
	FleetId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// The name of a location to retrieve compute resources for. For an Amazon
	// GameLift Anywhere fleet, use a custom location. For a multi-location EC2 or
	// container fleet, provide a Amazon Web Services Region or Local Zone code (for
	// example: us-west-2 or us-west-2-lax-1 ).
	Location *string

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComputeOutput struct {

	// A list of compute resources in the specified fleet.
	ComputeList []types.Compute

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComputeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCompute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCompute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCompute"); err != nil {
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
	if err = addOpListComputeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCompute(options.Region), middleware.Before); err != nil {
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

// ListComputePaginatorOptions is the paginator options for ListCompute
type ListComputePaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComputePaginator is a paginator for ListCompute
type ListComputePaginator struct {
	options   ListComputePaginatorOptions
	client    ListComputeAPIClient
	params    *ListComputeInput
	nextToken *string
	firstPage bool
}

// NewListComputePaginator returns a new ListComputePaginator
func NewListComputePaginator(client ListComputeAPIClient, params *ListComputeInput, optFns ...func(*ListComputePaginatorOptions)) *ListComputePaginator {
	if params == nil {
		params = &ListComputeInput{}
	}

	options := ListComputePaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComputePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComputePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCompute page.
func (p *ListComputePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComputeOutput, error) {
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
	result, err := p.client.ListCompute(ctx, &params, optFns...)
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

// ListComputeAPIClient is a client that implements the ListCompute operation.
type ListComputeAPIClient interface {
	ListCompute(context.Context, *ListComputeInput, ...func(*Options)) (*ListComputeOutput, error)
}

var _ ListComputeAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCompute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCompute",
	}
}
