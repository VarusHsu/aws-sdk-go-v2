// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of metadata model imports.
func (c *Client) DescribeMetadataModelImports(ctx context.Context, params *DescribeMetadataModelImportsInput, optFns ...func(*Options)) (*DescribeMetadataModelImportsOutput, error) {
	if params == nil {
		params = &DescribeMetadataModelImportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMetadataModelImports", params, optFns, c.addOperationDescribeMetadataModelImportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMetadataModelImportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMetadataModelImportsInput struct {

	// The migration project name or Amazon Resource Name (ARN).
	//
	// This member is required.
	MigrationProjectIdentifier *string

	// Filters applied to the metadata model imports described in the form of
	// key-value pairs.
	Filters []types.Filter

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords .
	//
	// If Marker is returned by a previous response, there are more results available.
	// The value of Marker is a unique pagination token for each page. To retrieve the
	// next page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// A paginated list of metadata model imports.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeMetadataModelImportsOutput struct {

	// Specifies the unique pagination token that makes it possible to display the
	// next page of results. If this parameter is specified, the response includes only
	// records beyond the marker, up to the value specified by MaxRecords .
	//
	// If Marker is returned by a previous response, there are more results available.
	// The value of Marker is a unique pagination token for each page. To retrieve the
	// next page, make the call again using the returned token and keeping all other
	// arguments unchanged.
	Marker *string

	// A paginated list of metadata model imports.
	Requests []types.SchemaConversionRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMetadataModelImportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMetadataModelImports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMetadataModelImports{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeMetadataModelImports"); err != nil {
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
	if err = addOpDescribeMetadataModelImportsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMetadataModelImports(options.Region), middleware.Before); err != nil {
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

// DescribeMetadataModelImportsPaginatorOptions is the paginator options for
// DescribeMetadataModelImports
type DescribeMetadataModelImportsPaginatorOptions struct {
	// A paginated list of metadata model imports.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMetadataModelImportsPaginator is a paginator for
// DescribeMetadataModelImports
type DescribeMetadataModelImportsPaginator struct {
	options   DescribeMetadataModelImportsPaginatorOptions
	client    DescribeMetadataModelImportsAPIClient
	params    *DescribeMetadataModelImportsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMetadataModelImportsPaginator returns a new
// DescribeMetadataModelImportsPaginator
func NewDescribeMetadataModelImportsPaginator(client DescribeMetadataModelImportsAPIClient, params *DescribeMetadataModelImportsInput, optFns ...func(*DescribeMetadataModelImportsPaginatorOptions)) *DescribeMetadataModelImportsPaginator {
	if params == nil {
		params = &DescribeMetadataModelImportsInput{}
	}

	options := DescribeMetadataModelImportsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMetadataModelImportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMetadataModelImportsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMetadataModelImports page.
func (p *DescribeMetadataModelImportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMetadataModelImportsOutput, error) {
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
	result, err := p.client.DescribeMetadataModelImports(ctx, &params, optFns...)
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

// DescribeMetadataModelImportsAPIClient is a client that implements the
// DescribeMetadataModelImports operation.
type DescribeMetadataModelImportsAPIClient interface {
	DescribeMetadataModelImports(context.Context, *DescribeMetadataModelImportsInput, ...func(*Options)) (*DescribeMetadataModelImportsOutput, error)
}

var _ DescribeMetadataModelImportsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeMetadataModelImports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeMetadataModelImports",
	}
}
