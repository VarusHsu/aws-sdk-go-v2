// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the margin summary report, which includes the Amazon Web Services
// cost and charged amount (pro forma cost) by Amazon Web Service for a specific
// billing group.
func (c *Client) GetBillingGroupCostReport(ctx context.Context, params *GetBillingGroupCostReportInput, optFns ...func(*Options)) (*GetBillingGroupCostReportOutput, error) {
	if params == nil {
		params = &GetBillingGroupCostReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBillingGroupCostReport", params, optFns, c.addOperationGetBillingGroupCostReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBillingGroupCostReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBillingGroupCostReportInput struct {

	// The Amazon Resource Number (ARN) that uniquely identifies the billing group.
	//
	// This member is required.
	Arn *string

	// A time range for which the margin summary is effective. You can specify up to
	// 12 months.
	BillingPeriodRange *types.BillingPeriodRange

	// A list of strings that specify the attributes that are used to break down costs
	// in the margin summary reports for the billing group. For example, you can view
	// your costs by the Amazon Web Service name or the billing period.
	GroupBy []types.GroupByAttributeName

	// The maximum number of margin summary reports to retrieve.
	MaxResults *int32

	// The pagination token used on subsequent calls to get reports.
	NextToken *string

	noSmithyDocumentSerde
}

type GetBillingGroupCostReportOutput struct {

	// The list of margin summary reports.
	BillingGroupCostReportResults []types.BillingGroupCostReportResultElement

	// The pagination token used on subsequent calls to get reports.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBillingGroupCostReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBillingGroupCostReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBillingGroupCostReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBillingGroupCostReport"); err != nil {
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
	if err = addOpGetBillingGroupCostReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBillingGroupCostReport(options.Region), middleware.Before); err != nil {
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

// GetBillingGroupCostReportPaginatorOptions is the paginator options for
// GetBillingGroupCostReport
type GetBillingGroupCostReportPaginatorOptions struct {
	// The maximum number of margin summary reports to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBillingGroupCostReportPaginator is a paginator for GetBillingGroupCostReport
type GetBillingGroupCostReportPaginator struct {
	options   GetBillingGroupCostReportPaginatorOptions
	client    GetBillingGroupCostReportAPIClient
	params    *GetBillingGroupCostReportInput
	nextToken *string
	firstPage bool
}

// NewGetBillingGroupCostReportPaginator returns a new
// GetBillingGroupCostReportPaginator
func NewGetBillingGroupCostReportPaginator(client GetBillingGroupCostReportAPIClient, params *GetBillingGroupCostReportInput, optFns ...func(*GetBillingGroupCostReportPaginatorOptions)) *GetBillingGroupCostReportPaginator {
	if params == nil {
		params = &GetBillingGroupCostReportInput{}
	}

	options := GetBillingGroupCostReportPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBillingGroupCostReportPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBillingGroupCostReportPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBillingGroupCostReport page.
func (p *GetBillingGroupCostReportPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBillingGroupCostReportOutput, error) {
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
	result, err := p.client.GetBillingGroupCostReport(ctx, &params, optFns...)
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

// GetBillingGroupCostReportAPIClient is a client that implements the
// GetBillingGroupCostReport operation.
type GetBillingGroupCostReportAPIClient interface {
	GetBillingGroupCostReport(context.Context, *GetBillingGroupCostReportInput, ...func(*Options)) (*GetBillingGroupCostReportOutput, error)
}

var _ GetBillingGroupCostReportAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetBillingGroupCostReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBillingGroupCostReport",
	}
}
