// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the notifications that are associated with a budget.
func (c *Client) DescribeNotificationsForBudget(ctx context.Context, params *DescribeNotificationsForBudgetInput, optFns ...func(*Options)) (*DescribeNotificationsForBudgetOutput, error) {
	if params == nil {
		params = &DescribeNotificationsForBudgetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNotificationsForBudget", params, optFns, c.addOperationDescribeNotificationsForBudgetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNotificationsForBudgetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of DescribeNotificationsForBudget
type DescribeNotificationsForBudgetInput struct {

	// The accountId that is associated with the budget whose notifications you want
	// descriptions of.
	//
	// This member is required.
	AccountId *string

	// The name of the budget whose notifications you want descriptions of.
	//
	// This member is required.
	BudgetName *string

	// An optional integer that represents how many entries a paginated response
	// contains.
	MaxResults *int32

	// The pagination token that you include in your request to indicate the next set
	// of results that you want to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

// Response of GetNotificationsForBudget
type DescribeNotificationsForBudgetOutput struct {

	// The pagination token in the service response that indicates the next set of
	// results that you can retrieve.
	NextToken *string

	// A list of notifications that are associated with a budget.
	Notifications []types.Notification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNotificationsForBudgetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNotificationsForBudget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNotificationsForBudget{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeNotificationsForBudget"); err != nil {
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
	if err = addOpDescribeNotificationsForBudgetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotificationsForBudget(options.Region), middleware.Before); err != nil {
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

// DescribeNotificationsForBudgetPaginatorOptions is the paginator options for
// DescribeNotificationsForBudget
type DescribeNotificationsForBudgetPaginatorOptions struct {
	// An optional integer that represents how many entries a paginated response
	// contains.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNotificationsForBudgetPaginator is a paginator for
// DescribeNotificationsForBudget
type DescribeNotificationsForBudgetPaginator struct {
	options   DescribeNotificationsForBudgetPaginatorOptions
	client    DescribeNotificationsForBudgetAPIClient
	params    *DescribeNotificationsForBudgetInput
	nextToken *string
	firstPage bool
}

// NewDescribeNotificationsForBudgetPaginator returns a new
// DescribeNotificationsForBudgetPaginator
func NewDescribeNotificationsForBudgetPaginator(client DescribeNotificationsForBudgetAPIClient, params *DescribeNotificationsForBudgetInput, optFns ...func(*DescribeNotificationsForBudgetPaginatorOptions)) *DescribeNotificationsForBudgetPaginator {
	if params == nil {
		params = &DescribeNotificationsForBudgetInput{}
	}

	options := DescribeNotificationsForBudgetPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeNotificationsForBudgetPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNotificationsForBudgetPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeNotificationsForBudget page.
func (p *DescribeNotificationsForBudgetPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNotificationsForBudgetOutput, error) {
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
	result, err := p.client.DescribeNotificationsForBudget(ctx, &params, optFns...)
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

// DescribeNotificationsForBudgetAPIClient is a client that implements the
// DescribeNotificationsForBudget operation.
type DescribeNotificationsForBudgetAPIClient interface {
	DescribeNotificationsForBudget(context.Context, *DescribeNotificationsForBudgetInput, ...func(*Options)) (*DescribeNotificationsForBudgetOutput, error)
}

var _ DescribeNotificationsForBudgetAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeNotificationsForBudget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeNotificationsForBudget",
	}
}
