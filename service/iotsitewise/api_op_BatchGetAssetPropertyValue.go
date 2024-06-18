// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the current value for one or more asset properties. For more information,
// see [Querying current values]in the IoT SiteWise User Guide.
//
// [Querying current values]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#current-values
func (c *Client) BatchGetAssetPropertyValue(ctx context.Context, params *BatchGetAssetPropertyValueInput, optFns ...func(*Options)) (*BatchGetAssetPropertyValueOutput, error) {
	if params == nil {
		params = &BatchGetAssetPropertyValueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetAssetPropertyValue", params, optFns, c.addOperationBatchGetAssetPropertyValueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetAssetPropertyValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetAssetPropertyValueInput struct {

	// The list of asset property value entries for the batch get request. You can
	// specify up to 128 entries per request.
	//
	// This member is required.
	Entries []types.BatchGetAssetPropertyValueEntry

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type BatchGetAssetPropertyValueOutput struct {

	// A list of the errors (if any) associated with the batch request. Each error
	// entry contains the entryId of the entry that failed.
	//
	// This member is required.
	ErrorEntries []types.BatchGetAssetPropertyValueErrorEntry

	// A list of entries that were not processed by this batch request. because these
	// entries had been completely processed by previous paginated requests. Each
	// skipped entry contains the entryId of the entry that skipped.
	//
	// This member is required.
	SkippedEntries []types.BatchGetAssetPropertyValueSkippedEntry

	// A list of entries that were processed successfully by this batch request. Each
	// success entry contains the entryId of the entry that succeeded and the latest
	// query result.
	//
	// This member is required.
	SuccessEntries []types.BatchGetAssetPropertyValueSuccessEntry

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetAssetPropertyValueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetAssetPropertyValue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetAssetPropertyValue{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetAssetPropertyValue"); err != nil {
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
	if err = addEndpointPrefix_opBatchGetAssetPropertyValueMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchGetAssetPropertyValueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetAssetPropertyValue(options.Region), middleware.Before); err != nil {
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

// BatchGetAssetPropertyValuePaginatorOptions is the paginator options for
// BatchGetAssetPropertyValue
type BatchGetAssetPropertyValuePaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// BatchGetAssetPropertyValuePaginator is a paginator for
// BatchGetAssetPropertyValue
type BatchGetAssetPropertyValuePaginator struct {
	options   BatchGetAssetPropertyValuePaginatorOptions
	client    BatchGetAssetPropertyValueAPIClient
	params    *BatchGetAssetPropertyValueInput
	nextToken *string
	firstPage bool
}

// NewBatchGetAssetPropertyValuePaginator returns a new
// BatchGetAssetPropertyValuePaginator
func NewBatchGetAssetPropertyValuePaginator(client BatchGetAssetPropertyValueAPIClient, params *BatchGetAssetPropertyValueInput, optFns ...func(*BatchGetAssetPropertyValuePaginatorOptions)) *BatchGetAssetPropertyValuePaginator {
	if params == nil {
		params = &BatchGetAssetPropertyValueInput{}
	}

	options := BatchGetAssetPropertyValuePaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &BatchGetAssetPropertyValuePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *BatchGetAssetPropertyValuePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next BatchGetAssetPropertyValue page.
func (p *BatchGetAssetPropertyValuePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*BatchGetAssetPropertyValueOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.BatchGetAssetPropertyValue(ctx, &params, optFns...)
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

type endpointPrefix_opBatchGetAssetPropertyValueMiddleware struct {
}

func (*endpointPrefix_opBatchGetAssetPropertyValueMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchGetAssetPropertyValueMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opBatchGetAssetPropertyValueMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opBatchGetAssetPropertyValueMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// BatchGetAssetPropertyValueAPIClient is a client that implements the
// BatchGetAssetPropertyValue operation.
type BatchGetAssetPropertyValueAPIClient interface {
	BatchGetAssetPropertyValue(context.Context, *BatchGetAssetPropertyValueInput, ...func(*Options)) (*BatchGetAssetPropertyValueOutput, error)
}

var _ BatchGetAssetPropertyValueAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opBatchGetAssetPropertyValue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetAssetPropertyValue",
	}
}
