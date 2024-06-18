// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets event data results of a query. You must specify the QueryID value returned
// by the StartQuery operation.
func (c *Client) GetQueryResults(ctx context.Context, params *GetQueryResultsInput, optFns ...func(*Options)) (*GetQueryResultsOutput, error) {
	if params == nil {
		params = &GetQueryResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueryResults", params, optFns, c.addOperationGetQueryResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueryResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryResultsInput struct {

	// The ID of the query for which you want to get results.
	//
	// This member is required.
	QueryId *string

	// The ARN (or ID suffix of the ARN) of the event data store against which the
	// query was run.
	//
	// Deprecated: EventDataStore is no longer required by GetQueryResultsRequest
	EventDataStore *string

	// The maximum number of query results to display on a single page.
	MaxQueryResults *int32

	// A token you can use to get the next page of query results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetQueryResultsOutput struct {

	// The error message returned if a query failed.
	ErrorMessage *string

	// A token you can use to get the next page of query results.
	NextToken *string

	// Contains the individual event results of the query.
	QueryResultRows [][]map[string]string

	// Shows the count of query results.
	QueryStatistics *types.QueryStatistics

	// The status of the query. Values include QUEUED , RUNNING , FINISHED , FAILED ,
	// TIMED_OUT , or CANCELLED .
	QueryStatus types.QueryStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQueryResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetQueryResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetQueryResults{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetQueryResults"); err != nil {
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
	if err = addOpGetQueryResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueryResults(options.Region), middleware.Before); err != nil {
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

// GetQueryResultsPaginatorOptions is the paginator options for GetQueryResults
type GetQueryResultsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetQueryResultsPaginator is a paginator for GetQueryResults
type GetQueryResultsPaginator struct {
	options   GetQueryResultsPaginatorOptions
	client    GetQueryResultsAPIClient
	params    *GetQueryResultsInput
	nextToken *string
	firstPage bool
}

// NewGetQueryResultsPaginator returns a new GetQueryResultsPaginator
func NewGetQueryResultsPaginator(client GetQueryResultsAPIClient, params *GetQueryResultsInput, optFns ...func(*GetQueryResultsPaginatorOptions)) *GetQueryResultsPaginator {
	if params == nil {
		params = &GetQueryResultsInput{}
	}

	options := GetQueryResultsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetQueryResultsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetQueryResultsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetQueryResults page.
func (p *GetQueryResultsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetQueryResultsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetQueryResults(ctx, &params, optFns...)
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

// GetQueryResultsAPIClient is a client that implements the GetQueryResults
// operation.
type GetQueryResultsAPIClient interface {
	GetQueryResults(context.Context, *GetQueryResultsInput, ...func(*Options)) (*GetQueryResultsOutput, error)
}

var _ GetQueryResultsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetQueryResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetQueryResults",
	}
}
