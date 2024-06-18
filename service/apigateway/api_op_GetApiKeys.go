// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the current ApiKeys resource.
func (c *Client) GetApiKeys(ctx context.Context, params *GetApiKeysInput, optFns ...func(*Options)) (*GetApiKeysOutput, error) {
	if params == nil {
		params = &GetApiKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApiKeys", params, optFns, c.addOperationGetApiKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApiKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get information about the current ApiKeys resource.
type GetApiKeysInput struct {

	// The identifier of a customer in Amazon Web Services Marketplace or an external
	// system, such as a developer portal.
	CustomerId *string

	// A boolean flag to specify whether ( true ) or not ( false ) the result contains
	// key values.
	IncludeValues *bool

	// The maximum number of returned results per page. The default value is 25 and
	// the maximum value is 500.
	Limit *int32

	// The name of queried API keys.
	NameQuery *string

	// The current pagination position in the paged result set.
	Position *string

	noSmithyDocumentSerde
}

// Represents a collection of API keys as represented by an ApiKeys resource.
type GetApiKeysOutput struct {

	// The current page of elements from this collection.
	Items []types.ApiKey

	// The current pagination position in the paged result set.
	Position *string

	// A list of warning messages logged during the import of API keys when the
	// failOnWarnings option is set to true.
	Warnings []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApiKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApiKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApiKeys{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetApiKeys"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApiKeys(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

// GetApiKeysPaginatorOptions is the paginator options for GetApiKeys
type GetApiKeysPaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and
	// the maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetApiKeysPaginator is a paginator for GetApiKeys
type GetApiKeysPaginator struct {
	options   GetApiKeysPaginatorOptions
	client    GetApiKeysAPIClient
	params    *GetApiKeysInput
	nextToken *string
	firstPage bool
}

// NewGetApiKeysPaginator returns a new GetApiKeysPaginator
func NewGetApiKeysPaginator(client GetApiKeysAPIClient, params *GetApiKeysInput, optFns ...func(*GetApiKeysPaginatorOptions)) *GetApiKeysPaginator {
	if params == nil {
		params = &GetApiKeysInput{}
	}

	options := GetApiKeysPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetApiKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Position,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetApiKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetApiKeys page.
func (p *GetApiKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetApiKeysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Position = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetApiKeys(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Position

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// GetApiKeysAPIClient is a client that implements the GetApiKeys operation.
type GetApiKeysAPIClient interface {
	GetApiKeys(context.Context, *GetApiKeysInput, ...func(*Options)) (*GetApiKeysOutput, error)
}

var _ GetApiKeysAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetApiKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetApiKeys",
	}
}
