// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns descriptions of the available Amazon Redshift cluster versions. You can
// call this operation even before creating any clusters to learn more about the
// Amazon Redshift versions.
//
// For more information about managing clusters, go to [Amazon Redshift Clusters] in the Amazon Redshift
// Cluster Management Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
func (c *Client) DescribeClusterVersions(ctx context.Context, params *DescribeClusterVersionsInput, optFns ...func(*Options)) (*DescribeClusterVersionsOutput, error) {
	if params == nil {
		params = &DescribeClusterVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusterVersions", params, optFns, c.addOperationDescribeClusterVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClusterVersionsInput struct {

	// The name of a specific cluster parameter group family to return details for.
	//
	// Constraints:
	//
	//   - Must be 1 to 255 alphanumeric characters
	//
	//   - First character must be a letter
	//
	//   - Cannot end with a hyphen or contain two consecutive hyphens
	ClusterParameterGroupFamily *string

	// The specific cluster version to return.
	//
	// Example: 1.0
	ClusterVersion *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterVersionsrequest exceed the value specified in
	// MaxRecords , Amazon Web Services returns a value in the Marker field of the
	// response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// Contains the output from the DescribeClusterVersions action.
type DescribeClusterVersionsOutput struct {

	// A list of Version elements.
	ClusterVersions []types.ClusterVersion

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClusterVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeClusterVersions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterVersions(options.Region), middleware.Before); err != nil {
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

// DescribeClusterVersionsPaginatorOptions is the paginator options for
// DescribeClusterVersions
type DescribeClusterVersionsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClusterVersionsPaginator is a paginator for DescribeClusterVersions
type DescribeClusterVersionsPaginator struct {
	options   DescribeClusterVersionsPaginatorOptions
	client    DescribeClusterVersionsAPIClient
	params    *DescribeClusterVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeClusterVersionsPaginator returns a new
// DescribeClusterVersionsPaginator
func NewDescribeClusterVersionsPaginator(client DescribeClusterVersionsAPIClient, params *DescribeClusterVersionsInput, optFns ...func(*DescribeClusterVersionsPaginatorOptions)) *DescribeClusterVersionsPaginator {
	if params == nil {
		params = &DescribeClusterVersionsInput{}
	}

	options := DescribeClusterVersionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClusterVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClusterVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeClusterVersions page.
func (p *DescribeClusterVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClusterVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeClusterVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeClusterVersionsAPIClient is a client that implements the
// DescribeClusterVersions operation.
type DescribeClusterVersionsAPIClient interface {
	DescribeClusterVersions(context.Context, *DescribeClusterVersionsInput, ...func(*Options)) (*DescribeClusterVersionsOutput, error)
}

var _ DescribeClusterVersionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeClusterVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeClusterVersions",
	}
}
