// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rum/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the raw performance events that RUM has collected from your web
// application, so that you can do your own processing or analysis of this data.
func (c *Client) GetAppMonitorData(ctx context.Context, params *GetAppMonitorDataInput, optFns ...func(*Options)) (*GetAppMonitorDataOutput, error) {
	if params == nil {
		params = &GetAppMonitorDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAppMonitorData", params, optFns, c.addOperationGetAppMonitorDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAppMonitorDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAppMonitorDataInput struct {

	// The name of the app monitor that collected the data that you want to retrieve.
	//
	// This member is required.
	Name *string

	// A structure that defines the time range that you want to retrieve results from.
	//
	// This member is required.
	TimeRange *types.TimeRange

	// An array of structures that you can use to filter the results to those that
	// match one or more sets of key-value pairs that you specify.
	Filters []types.QueryFilter

	// The maximum number of results to return in one operation.
	MaxResults int32

	// Use the token returned by the previous operation to request the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetAppMonitorDataOutput struct {

	// The events that RUM collected that match your request.
	Events []string

	// A token that you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAppMonitorDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAppMonitorData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAppMonitorData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAppMonitorData"); err != nil {
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
	if err = addOpGetAppMonitorDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAppMonitorData(options.Region), middleware.Before); err != nil {
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

// GetAppMonitorDataAPIClient is a client that implements the GetAppMonitorData
// operation.
type GetAppMonitorDataAPIClient interface {
	GetAppMonitorData(context.Context, *GetAppMonitorDataInput, ...func(*Options)) (*GetAppMonitorDataOutput, error)
}

var _ GetAppMonitorDataAPIClient = (*Client)(nil)

// GetAppMonitorDataPaginatorOptions is the paginator options for GetAppMonitorData
type GetAppMonitorDataPaginatorOptions struct {
	// The maximum number of results to return in one operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetAppMonitorDataPaginator is a paginator for GetAppMonitorData
type GetAppMonitorDataPaginator struct {
	options   GetAppMonitorDataPaginatorOptions
	client    GetAppMonitorDataAPIClient
	params    *GetAppMonitorDataInput
	nextToken *string
	firstPage bool
}

// NewGetAppMonitorDataPaginator returns a new GetAppMonitorDataPaginator
func NewGetAppMonitorDataPaginator(client GetAppMonitorDataAPIClient, params *GetAppMonitorDataInput, optFns ...func(*GetAppMonitorDataPaginatorOptions)) *GetAppMonitorDataPaginator {
	if params == nil {
		params = &GetAppMonitorDataInput{}
	}

	options := GetAppMonitorDataPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetAppMonitorDataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetAppMonitorDataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetAppMonitorData page.
func (p *GetAppMonitorDataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetAppMonitorDataOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetAppMonitorData(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetAppMonitorData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAppMonitorData",
	}
}
