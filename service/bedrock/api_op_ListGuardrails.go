// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists details about all the guardrails in an account. To list the DRAFT version
// of all your guardrails, don't specify the guardrailIdentifier field. To list
// all versions of a guardrail, specify the ARN of the guardrail in the
// guardrailIdentifier field.
//
// You can set the maximum number of results to return in a response in the
// maxResults field. If there are more results than the number you set, the
// response returns a nextToken that you can send in another ListGuardrails
// request to see the next batch of results.
func (c *Client) ListGuardrails(ctx context.Context, params *ListGuardrailsInput, optFns ...func(*Options)) (*ListGuardrailsOutput, error) {
	if params == nil {
		params = &ListGuardrailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGuardrails", params, optFns, c.addOperationListGuardrailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGuardrailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGuardrailsInput struct {

	// The unique identifier of the guardrail.
	GuardrailIdentifier *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// If there are more results than were returned in the response, the response
	// returns a nextToken that you can send in another ListGuardrails request to see
	// the next batch of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGuardrailsOutput struct {

	// A list of objects, each of which contains details about a guardrail.
	//
	// This member is required.
	Guardrails []types.GuardrailSummary

	// If there are more results than were returned in the response, the response
	// returns a nextToken that you can send in another ListGuardrails request to see
	// the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGuardrailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGuardrails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGuardrails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGuardrails"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGuardrails(options.Region), middleware.Before); err != nil {
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

// ListGuardrailsPaginatorOptions is the paginator options for ListGuardrails
type ListGuardrailsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGuardrailsPaginator is a paginator for ListGuardrails
type ListGuardrailsPaginator struct {
	options   ListGuardrailsPaginatorOptions
	client    ListGuardrailsAPIClient
	params    *ListGuardrailsInput
	nextToken *string
	firstPage bool
}

// NewListGuardrailsPaginator returns a new ListGuardrailsPaginator
func NewListGuardrailsPaginator(client ListGuardrailsAPIClient, params *ListGuardrailsInput, optFns ...func(*ListGuardrailsPaginatorOptions)) *ListGuardrailsPaginator {
	if params == nil {
		params = &ListGuardrailsInput{}
	}

	options := ListGuardrailsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGuardrailsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGuardrailsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGuardrails page.
func (p *ListGuardrailsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGuardrailsOutput, error) {
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
	result, err := p.client.ListGuardrails(ctx, &params, optFns...)
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

// ListGuardrailsAPIClient is a client that implements the ListGuardrails
// operation.
type ListGuardrailsAPIClient interface {
	ListGuardrails(context.Context, *ListGuardrailsInput, ...func(*Options)) (*ListGuardrailsOutput, error)
}

var _ ListGuardrailsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListGuardrails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGuardrails",
	}
}
