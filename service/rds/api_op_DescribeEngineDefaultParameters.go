// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the default engine and system parameter information for the specified
// database engine.
func (c *Client) DescribeEngineDefaultParameters(ctx context.Context, params *DescribeEngineDefaultParametersInput, optFns ...func(*Options)) (*DescribeEngineDefaultParametersOutput, error) {
	if params == nil {
		params = &DescribeEngineDefaultParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEngineDefaultParameters", params, optFns, c.addOperationDescribeEngineDefaultParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEngineDefaultParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEngineDefaultParametersInput struct {

	// The name of the DB parameter group family. Valid Values:
	//   - aurora-mysql5.7
	//   - aurora-mysql8.0
	//   - aurora-postgresql10
	//   - aurora-postgresql11
	//   - aurora-postgresql12
	//   - aurora-postgresql13
	//   - aurora-postgresql14
	//   - custom-oracle-ee-19
	//   - mariadb10.2
	//   - mariadb10.3
	//   - mariadb10.4
	//   - mariadb10.5
	//   - mariadb10.6
	//   - mysql5.7
	//   - mysql8.0
	//   - oracle-ee-19
	//   - oracle-ee-cdb-19
	//   - oracle-ee-cdb-21
	//   - oracle-se2-19
	//   - oracle-se2-cdb-19
	//   - oracle-se2-cdb-21
	//   - postgres10
	//   - postgres11
	//   - postgres12
	//   - postgres13
	//   - postgres14
	//   - sqlserver-ee-11.0
	//   - sqlserver-ee-12.0
	//   - sqlserver-ee-13.0
	//   - sqlserver-ee-14.0
	//   - sqlserver-ee-15.0
	//   - sqlserver-ex-11.0
	//   - sqlserver-ex-12.0
	//   - sqlserver-ex-13.0
	//   - sqlserver-ex-14.0
	//   - sqlserver-ex-15.0
	//   - sqlserver-se-11.0
	//   - sqlserver-se-12.0
	//   - sqlserver-se-13.0
	//   - sqlserver-se-14.0
	//   - sqlserver-se-15.0
	//   - sqlserver-web-11.0
	//   - sqlserver-web-12.0
	//   - sqlserver-web-13.0
	//   - sqlserver-web-14.0
	//   - sqlserver-web-15.0
	//
	// This member is required.
	DBParameterGroupFamily *string

	// This parameter isn't currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous
	// DescribeEngineDefaultParameters request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeEngineDefaultParametersOutput struct {

	// Contains the result of a successful invocation of the
	// DescribeEngineDefaultParameters action.
	EngineDefaults *types.EngineDefaults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEngineDefaultParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEngineDefaultParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEngineDefaultParameters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEngineDefaultParameters"); err != nil {
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
	if err = addOpDescribeEngineDefaultParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEngineDefaultParameters(options.Region), middleware.Before); err != nil {
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

// DescribeEngineDefaultParametersAPIClient is a client that implements the
// DescribeEngineDefaultParameters operation.
type DescribeEngineDefaultParametersAPIClient interface {
	DescribeEngineDefaultParameters(context.Context, *DescribeEngineDefaultParametersInput, ...func(*Options)) (*DescribeEngineDefaultParametersOutput, error)
}

var _ DescribeEngineDefaultParametersAPIClient = (*Client)(nil)

// DescribeEngineDefaultParametersPaginatorOptions is the paginator options for
// DescribeEngineDefaultParameters
type DescribeEngineDefaultParametersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEngineDefaultParametersPaginator is a paginator for
// DescribeEngineDefaultParameters
type DescribeEngineDefaultParametersPaginator struct {
	options   DescribeEngineDefaultParametersPaginatorOptions
	client    DescribeEngineDefaultParametersAPIClient
	params    *DescribeEngineDefaultParametersInput
	nextToken *string
	firstPage bool
}

// NewDescribeEngineDefaultParametersPaginator returns a new
// DescribeEngineDefaultParametersPaginator
func NewDescribeEngineDefaultParametersPaginator(client DescribeEngineDefaultParametersAPIClient, params *DescribeEngineDefaultParametersInput, optFns ...func(*DescribeEngineDefaultParametersPaginatorOptions)) *DescribeEngineDefaultParametersPaginator {
	if params == nil {
		params = &DescribeEngineDefaultParametersInput{}
	}

	options := DescribeEngineDefaultParametersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEngineDefaultParametersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEngineDefaultParametersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEngineDefaultParameters page.
func (p *DescribeEngineDefaultParametersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEngineDefaultParametersOutput, error) {
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

	result, err := p.client.DescribeEngineDefaultParameters(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = nil
	if result.EngineDefaults != nil {
		p.nextToken = result.EngineDefaults.Marker
	}

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeEngineDefaultParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEngineDefaultParameters",
	}
}
