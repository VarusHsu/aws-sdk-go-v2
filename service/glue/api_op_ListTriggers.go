// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the names of all trigger resources in this Amazon Web Services
// account, or the resources with the specified tag. This operation allows you to
// see which resources are available in your account, and their names.
//
// This operation takes the optional Tags field, which you can use as a filter on
// the response so that tagged resources can be retrieved as a group. If you choose
// to use tags filtering, only resources with the tag are retrieved.
func (c *Client) ListTriggers(ctx context.Context, params *ListTriggersInput, optFns ...func(*Options)) (*ListTriggersOutput, error) {
	if params == nil {
		params = &ListTriggersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTriggers", params, optFns, c.addOperationListTriggersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTriggersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTriggersInput struct {

	//  The name of the job for which to retrieve triggers. The trigger that can start
	// this job is returned. If there is no such trigger, all triggers are returned.
	DependentJobName *string

	// The maximum size of a list to return.
	MaxResults *int32

	// A continuation token, if this is a continuation request.
	NextToken *string

	// Specifies to return only these tagged resources.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ListTriggersOutput struct {

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string

	// The names of all triggers in the account, or the triggers with the specified
	// tags.
	TriggerNames []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTriggersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTriggers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTriggers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTriggers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTriggers(options.Region), middleware.Before); err != nil {
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

// ListTriggersPaginatorOptions is the paginator options for ListTriggers
type ListTriggersPaginatorOptions struct {
	// The maximum size of a list to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTriggersPaginator is a paginator for ListTriggers
type ListTriggersPaginator struct {
	options   ListTriggersPaginatorOptions
	client    ListTriggersAPIClient
	params    *ListTriggersInput
	nextToken *string
	firstPage bool
}

// NewListTriggersPaginator returns a new ListTriggersPaginator
func NewListTriggersPaginator(client ListTriggersAPIClient, params *ListTriggersInput, optFns ...func(*ListTriggersPaginatorOptions)) *ListTriggersPaginator {
	if params == nil {
		params = &ListTriggersInput{}
	}

	options := ListTriggersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTriggersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTriggersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTriggers page.
func (p *ListTriggersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTriggersOutput, error) {
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
	result, err := p.client.ListTriggers(ctx, &params, optFns...)
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

// ListTriggersAPIClient is a client that implements the ListTriggers operation.
type ListTriggersAPIClient interface {
	ListTriggers(context.Context, *ListTriggersInput, ...func(*Options)) (*ListTriggersOutput, error)
}

var _ ListTriggersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListTriggers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTriggers",
	}
}
