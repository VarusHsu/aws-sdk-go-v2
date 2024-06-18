// Code generated by smithy-go-codegen DO NOT EDIT.

package costandusagereportservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Amazon Web Services Cost and Usage Report available to this account.
func (c *Client) DescribeReportDefinitions(ctx context.Context, params *DescribeReportDefinitionsInput, optFns ...func(*Options)) (*DescribeReportDefinitionsOutput, error) {
	if params == nil {
		params = &DescribeReportDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReportDefinitions", params, optFns, c.addOperationDescribeReportDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReportDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests a Amazon Web Services Cost and Usage Report list owned by the account.
type DescribeReportDefinitionsInput struct {

	// The maximum number of results that Amazon Web Services returns for the
	// operation.
	MaxResults *int32

	// A generic string.
	NextToken *string

	noSmithyDocumentSerde
}

// If the action is successful, the service sends back an HTTP 200 response.
type DescribeReportDefinitionsOutput struct {

	// A generic string.
	NextToken *string

	// An Amazon Web Services Cost and Usage Report list owned by the account.
	ReportDefinitions []types.ReportDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReportDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReportDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReportDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeReportDefinitions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReportDefinitions(options.Region), middleware.Before); err != nil {
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

// DescribeReportDefinitionsPaginatorOptions is the paginator options for
// DescribeReportDefinitions
type DescribeReportDefinitionsPaginatorOptions struct {
	// The maximum number of results that Amazon Web Services returns for the
	// operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReportDefinitionsPaginator is a paginator for DescribeReportDefinitions
type DescribeReportDefinitionsPaginator struct {
	options   DescribeReportDefinitionsPaginatorOptions
	client    DescribeReportDefinitionsAPIClient
	params    *DescribeReportDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReportDefinitionsPaginator returns a new
// DescribeReportDefinitionsPaginator
func NewDescribeReportDefinitionsPaginator(client DescribeReportDefinitionsAPIClient, params *DescribeReportDefinitionsInput, optFns ...func(*DescribeReportDefinitionsPaginatorOptions)) *DescribeReportDefinitionsPaginator {
	if params == nil {
		params = &DescribeReportDefinitionsInput{}
	}

	options := DescribeReportDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReportDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReportDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReportDefinitions page.
func (p *DescribeReportDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReportDefinitionsOutput, error) {
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
	result, err := p.client.DescribeReportDefinitions(ctx, &params, optFns...)
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

// DescribeReportDefinitionsAPIClient is a client that implements the
// DescribeReportDefinitions operation.
type DescribeReportDefinitionsAPIClient interface {
	DescribeReportDefinitions(context.Context, *DescribeReportDefinitionsInput, ...func(*Options)) (*DescribeReportDefinitionsOutput, error)
}

var _ DescribeReportDefinitionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeReportDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeReportDefinitions",
	}
}
