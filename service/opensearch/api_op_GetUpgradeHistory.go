// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the complete history of the last 10 upgrades performed on an Amazon
// OpenSearch Service domain.
func (c *Client) GetUpgradeHistory(ctx context.Context, params *GetUpgradeHistoryInput, optFns ...func(*Options)) (*GetUpgradeHistoryOutput, error) {
	if params == nil {
		params = &GetUpgradeHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUpgradeHistory", params, optFns, c.addOperationGetUpgradeHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUpgradeHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to the GetUpgradeHistory operation.
type GetUpgradeHistoryInput struct {

	// The name of an existing domain.
	//
	// This member is required.
	DomainName *string

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults int32

	// If your initial GetUpgradeHistory operation returns a nextToken , you can
	// include the returned nextToken in subsequent GetUpgradeHistory operations,
	// which returns results in the next page.
	NextToken *string

	noSmithyDocumentSerde
}

// Container for the response returned by the GetUpgradeHistory operation.
type GetUpgradeHistoryOutput struct {

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// A list of objects corresponding to each upgrade or upgrade eligibility check
	// performed on a domain.
	UpgradeHistories []types.UpgradeHistory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUpgradeHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUpgradeHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUpgradeHistory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUpgradeHistory"); err != nil {
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
	if err = addOpGetUpgradeHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUpgradeHistory(options.Region), middleware.Before); err != nil {
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

// GetUpgradeHistoryAPIClient is a client that implements the GetUpgradeHistory
// operation.
type GetUpgradeHistoryAPIClient interface {
	GetUpgradeHistory(context.Context, *GetUpgradeHistoryInput, ...func(*Options)) (*GetUpgradeHistoryOutput, error)
}

var _ GetUpgradeHistoryAPIClient = (*Client)(nil)

// GetUpgradeHistoryPaginatorOptions is the paginator options for GetUpgradeHistory
type GetUpgradeHistoryPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetUpgradeHistoryPaginator is a paginator for GetUpgradeHistory
type GetUpgradeHistoryPaginator struct {
	options   GetUpgradeHistoryPaginatorOptions
	client    GetUpgradeHistoryAPIClient
	params    *GetUpgradeHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetUpgradeHistoryPaginator returns a new GetUpgradeHistoryPaginator
func NewGetUpgradeHistoryPaginator(client GetUpgradeHistoryAPIClient, params *GetUpgradeHistoryInput, optFns ...func(*GetUpgradeHistoryPaginatorOptions)) *GetUpgradeHistoryPaginator {
	if params == nil {
		params = &GetUpgradeHistoryInput{}
	}

	options := GetUpgradeHistoryPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetUpgradeHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetUpgradeHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetUpgradeHistory page.
func (p *GetUpgradeHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetUpgradeHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetUpgradeHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetUpgradeHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUpgradeHistory",
	}
}
