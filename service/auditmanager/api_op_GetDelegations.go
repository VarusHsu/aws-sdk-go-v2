// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of delegations from an audit owner to a delegate.
func (c *Client) GetDelegations(ctx context.Context, params *GetDelegationsInput, optFns ...func(*Options)) (*GetDelegationsOutput, error) {
	if params == nil {
		params = &GetDelegationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDelegations", params, optFns, c.addOperationGetDelegationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDelegationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDelegationsInput struct {

	//  Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	//  The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetDelegationsOutput struct {

	//  The list of delegations that the GetDelegations API returned.
	Delegations []types.DelegationMetadata

	//  The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDelegationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDelegations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDelegations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDelegations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDelegations(options.Region), middleware.Before); err != nil {
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

// GetDelegationsPaginatorOptions is the paginator options for GetDelegations
type GetDelegationsPaginatorOptions struct {
	//  Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDelegationsPaginator is a paginator for GetDelegations
type GetDelegationsPaginator struct {
	options   GetDelegationsPaginatorOptions
	client    GetDelegationsAPIClient
	params    *GetDelegationsInput
	nextToken *string
	firstPage bool
}

// NewGetDelegationsPaginator returns a new GetDelegationsPaginator
func NewGetDelegationsPaginator(client GetDelegationsAPIClient, params *GetDelegationsInput, optFns ...func(*GetDelegationsPaginatorOptions)) *GetDelegationsPaginator {
	if params == nil {
		params = &GetDelegationsInput{}
	}

	options := GetDelegationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetDelegationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDelegationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetDelegations page.
func (p *GetDelegationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDelegationsOutput, error) {
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
	result, err := p.client.GetDelegations(ctx, &params, optFns...)
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

// GetDelegationsAPIClient is a client that implements the GetDelegations
// operation.
type GetDelegationsAPIClient interface {
	GetDelegations(context.Context, *GetDelegationsInput, ...func(*Options)) (*GetDelegationsOutput, error)
}

var _ GetDelegationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetDelegations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDelegations",
	}
}
