// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the user activities in a specified time period.
func (c *Client) DescribeActivities(ctx context.Context, params *DescribeActivitiesInput, optFns ...func(*Options)) (*DescribeActivitiesOutput, error) {
	if params == nil {
		params = &DescribeActivitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeActivities", params, optFns, c.addOperationDescribeActivitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeActivitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeActivitiesInput struct {

	// Specifies which activity types to include in the response. If this field is
	// left empty, all activity types are returned.
	ActivityTypes *string

	// Amazon WorkDocs authentication token. Not required when using Amazon Web
	// Services administrator credentials to access the API.
	AuthenticationToken *string

	// The timestamp that determines the end time of the activities. The response
	// includes the activities performed before the specified timestamp.
	EndTime *time.Time

	// Includes indirect activities. An indirect activity results from a direct
	// activity performed on a parent resource. For example, sharing a parent folder
	// (the direct activity) shares all of the subfolders and documents within the
	// parent folder (the indirect activity).
	IncludeIndirectActivities bool

	// The maximum number of items to return.
	Limit *int32

	// The marker for the next set of results.
	Marker *string

	// The ID of the organization. This is a mandatory parameter when using
	// administrative API (SigV4) requests.
	OrganizationId *string

	// The document or folder ID for which to describe activity types.
	ResourceId *string

	// The timestamp that determines the starting time of the activities. The response
	// includes the activities performed after the specified timestamp.
	StartTime *time.Time

	// The ID of the user who performed the action. The response includes activities
	// pertaining to this user. This is an optional parameter and is only applicable
	// for administrative API (SigV4) requests.
	UserId *string

	noSmithyDocumentSerde
}

type DescribeActivitiesOutput struct {

	// The marker for the next set of results.
	Marker *string

	// The list of activities for the specified user and time period.
	UserActivities []types.Activity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeActivitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeActivities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeActivities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeActivities"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActivities(options.Region), middleware.Before); err != nil {
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

// DescribeActivitiesPaginatorOptions is the paginator options for
// DescribeActivities
type DescribeActivitiesPaginatorOptions struct {
	// The maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeActivitiesPaginator is a paginator for DescribeActivities
type DescribeActivitiesPaginator struct {
	options   DescribeActivitiesPaginatorOptions
	client    DescribeActivitiesAPIClient
	params    *DescribeActivitiesInput
	nextToken *string
	firstPage bool
}

// NewDescribeActivitiesPaginator returns a new DescribeActivitiesPaginator
func NewDescribeActivitiesPaginator(client DescribeActivitiesAPIClient, params *DescribeActivitiesInput, optFns ...func(*DescribeActivitiesPaginatorOptions)) *DescribeActivitiesPaginator {
	if params == nil {
		params = &DescribeActivitiesInput{}
	}

	options := DescribeActivitiesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeActivitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeActivitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeActivities page.
func (p *DescribeActivitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeActivitiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeActivities(ctx, &params, optFns...)
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

// DescribeActivitiesAPIClient is a client that implements the DescribeActivities
// operation.
type DescribeActivitiesAPIClient interface {
	DescribeActivities(context.Context, *DescribeActivitiesInput, ...func(*Options)) (*DescribeActivitiesOutput, error)
}

var _ DescribeActivitiesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeActivities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeActivities",
	}
}
