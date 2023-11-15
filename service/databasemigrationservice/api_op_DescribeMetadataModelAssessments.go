// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of metadata model assessments for your account in the
// current region.
func (c *Client) DescribeMetadataModelAssessments(ctx context.Context, params *DescribeMetadataModelAssessmentsInput, optFns ...func(*Options)) (*DescribeMetadataModelAssessmentsOutput, error) {
	if params == nil {
		params = &DescribeMetadataModelAssessmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMetadataModelAssessments", params, optFns, c.addOperationDescribeMetadataModelAssessmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMetadataModelAssessmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMetadataModelAssessmentsInput struct {

	// The name or Amazon Resource Name (ARN) of the migration project.
	//
	// This member is required.
	MigrationProjectIdentifier *string

	// Filters applied to the metadata model assessments described in the form of
	// key-value pairs.
	Filters []types.Filter

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords . If Marker
	// is returned by a previous response, there are more results available. The value
	// of Marker is a unique pagination token for each page. To retrieve the next
	// page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, DMS includes a pagination token in the
	// response so that you can retrieve the remaining results.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeMetadataModelAssessmentsOutput struct {

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords . If Marker
	// is returned by a previous response, there are more results available. The value
	// of Marker is a unique pagination token for each page. To retrieve the next
	// page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// A paginated list of metadata model assessments for the specified migration
	// project.
	Requests []types.SchemaConversionRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMetadataModelAssessmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMetadataModelAssessments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMetadataModelAssessments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMetadataModelAssessments"); err != nil {
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
	if err = addOpDescribeMetadataModelAssessmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMetadataModelAssessments(options.Region), middleware.Before); err != nil {
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

// DescribeMetadataModelAssessmentsAPIClient is a client that implements the
// DescribeMetadataModelAssessments operation.
type DescribeMetadataModelAssessmentsAPIClient interface {
	DescribeMetadataModelAssessments(context.Context, *DescribeMetadataModelAssessmentsInput, ...func(*Options)) (*DescribeMetadataModelAssessmentsOutput, error)
}

var _ DescribeMetadataModelAssessmentsAPIClient = (*Client)(nil)

// DescribeMetadataModelAssessmentsPaginatorOptions is the paginator options for
// DescribeMetadataModelAssessments
type DescribeMetadataModelAssessmentsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, DMS includes a pagination token in the
	// response so that you can retrieve the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMetadataModelAssessmentsPaginator is a paginator for
// DescribeMetadataModelAssessments
type DescribeMetadataModelAssessmentsPaginator struct {
	options   DescribeMetadataModelAssessmentsPaginatorOptions
	client    DescribeMetadataModelAssessmentsAPIClient
	params    *DescribeMetadataModelAssessmentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMetadataModelAssessmentsPaginator returns a new
// DescribeMetadataModelAssessmentsPaginator
func NewDescribeMetadataModelAssessmentsPaginator(client DescribeMetadataModelAssessmentsAPIClient, params *DescribeMetadataModelAssessmentsInput, optFns ...func(*DescribeMetadataModelAssessmentsPaginatorOptions)) *DescribeMetadataModelAssessmentsPaginator {
	if params == nil {
		params = &DescribeMetadataModelAssessmentsInput{}
	}

	options := DescribeMetadataModelAssessmentsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMetadataModelAssessmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMetadataModelAssessmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMetadataModelAssessments page.
func (p *DescribeMetadataModelAssessmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMetadataModelAssessmentsOutput, error) {
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

	result, err := p.client.DescribeMetadataModelAssessments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMetadataModelAssessments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMetadataModelAssessments",
	}
}
