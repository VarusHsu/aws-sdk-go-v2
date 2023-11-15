// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of entities that have been affected by one or more events for
// one or more accounts in your organization in Organizations, based on the filter
// criteria. Entities can refer to individual customer resources, groups of
// customer resources, or any other construct, depending on the Amazon Web Service.
// At least one event Amazon Resource Name (ARN) and account ID are required.
// Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the EnableHealthServiceAccessForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html)
// operation from your organization's management account.
//   - This API operation uses pagination. Specify the nextToken parameter in the
//     next request to return more results.
//   - This operation doesn't support resource-level permissions. You can't use
//     this operation to allow or deny access to specific Health events. For more
//     information, see Resource- and action-based conditions (https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions)
//     in the Health User Guide.
func (c *Client) DescribeAffectedEntitiesForOrganization(ctx context.Context, params *DescribeAffectedEntitiesForOrganizationInput, optFns ...func(*Options)) (*DescribeAffectedEntitiesForOrganizationOutput, error) {
	if params == nil {
		params = &DescribeAffectedEntitiesForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAffectedEntitiesForOrganization", params, optFns, c.addOperationDescribeAffectedEntitiesForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAffectedEntitiesForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAffectedEntitiesForOrganizationInput struct {

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string

	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	MaxResults *int32

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	// A JSON set of elements including the awsAccountId , eventArn and a set of
	// statusCodes .
	OrganizationEntityAccountFilters []types.EntityAccountFilter

	// A JSON set of elements including the awsAccountId and the eventArn .
	//
	// Deprecated: This property is deprecated, use organizationEntityAccountFilters
	// instead.
	OrganizationEntityFilters []types.EventAccountFilter

	noSmithyDocumentSerde
}

type DescribeAffectedEntitiesForOrganizationOutput struct {

	// A JSON set of elements including the awsAccountId and its entityArn ,
	// entityValue and its entityArn , lastUpdatedTime , and statusCode .
	Entities []types.AffectedEntity

	// A JSON set of elements of the failed response, including the awsAccountId ,
	// errorMessage , errorName , and eventArn .
	FailedSet []types.OrganizationAffectedEntitiesErrorItem

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAffectedEntitiesForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAffectedEntitiesForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAffectedEntitiesForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAffectedEntitiesForOrganization"); err != nil {
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
	if err = addOpDescribeAffectedEntitiesForOrganizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAffectedEntitiesForOrganization(options.Region), middleware.Before); err != nil {
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

// DescribeAffectedEntitiesForOrganizationAPIClient is a client that implements
// the DescribeAffectedEntitiesForOrganization operation.
type DescribeAffectedEntitiesForOrganizationAPIClient interface {
	DescribeAffectedEntitiesForOrganization(context.Context, *DescribeAffectedEntitiesForOrganizationInput, ...func(*Options)) (*DescribeAffectedEntitiesForOrganizationOutput, error)
}

var _ DescribeAffectedEntitiesForOrganizationAPIClient = (*Client)(nil)

// DescribeAffectedEntitiesForOrganizationPaginatorOptions is the paginator
// options for DescribeAffectedEntitiesForOrganization
type DescribeAffectedEntitiesForOrganizationPaginatorOptions struct {
	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAffectedEntitiesForOrganizationPaginator is a paginator for
// DescribeAffectedEntitiesForOrganization
type DescribeAffectedEntitiesForOrganizationPaginator struct {
	options   DescribeAffectedEntitiesForOrganizationPaginatorOptions
	client    DescribeAffectedEntitiesForOrganizationAPIClient
	params    *DescribeAffectedEntitiesForOrganizationInput
	nextToken *string
	firstPage bool
}

// NewDescribeAffectedEntitiesForOrganizationPaginator returns a new
// DescribeAffectedEntitiesForOrganizationPaginator
func NewDescribeAffectedEntitiesForOrganizationPaginator(client DescribeAffectedEntitiesForOrganizationAPIClient, params *DescribeAffectedEntitiesForOrganizationInput, optFns ...func(*DescribeAffectedEntitiesForOrganizationPaginatorOptions)) *DescribeAffectedEntitiesForOrganizationPaginator {
	if params == nil {
		params = &DescribeAffectedEntitiesForOrganizationInput{}
	}

	options := DescribeAffectedEntitiesForOrganizationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAffectedEntitiesForOrganizationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAffectedEntitiesForOrganizationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAffectedEntitiesForOrganization page.
func (p *DescribeAffectedEntitiesForOrganizationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAffectedEntitiesForOrganizationOutput, error) {
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

	result, err := p.client.DescribeAffectedEntitiesForOrganization(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAffectedEntitiesForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAffectedEntitiesForOrganization",
	}
}
