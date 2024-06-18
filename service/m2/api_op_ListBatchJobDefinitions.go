// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the available batch job definitions based on the batch job resources
// uploaded during the application creation. You can use the batch job definitions
// in the list to start a batch job.
func (c *Client) ListBatchJobDefinitions(ctx context.Context, params *ListBatchJobDefinitionsInput, optFns ...func(*Options)) (*ListBatchJobDefinitionsOutput, error) {
	if params == nil {
		params = &ListBatchJobDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBatchJobDefinitions", params, optFns, c.addOperationListBatchJobDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBatchJobDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBatchJobDefinitionsInput struct {

	// The identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The maximum number of batch job definitions to return.
	MaxResults *int32

	// A pagination token returned from a previous call to this operation. This
	// specifies the next item to return. To return to the beginning of the list,
	// exclude this parameter.
	NextToken *string

	// If the batch job definition is a FileBatchJobDefinition, the prefix allows you
	// to search on the file names of FileBatchJobDefinitions.
	Prefix *string

	noSmithyDocumentSerde
}

type ListBatchJobDefinitionsOutput struct {

	// The list of batch job definitions.
	//
	// This member is required.
	BatchJobDefinitions []types.BatchJobDefinition

	// If there are more items to return, this contains a token that is passed to a
	// subsequent call to this operation to retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBatchJobDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBatchJobDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBatchJobDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBatchJobDefinitions"); err != nil {
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
	if err = addOpListBatchJobDefinitionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBatchJobDefinitions(options.Region), middleware.Before); err != nil {
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

// ListBatchJobDefinitionsPaginatorOptions is the paginator options for
// ListBatchJobDefinitions
type ListBatchJobDefinitionsPaginatorOptions struct {
	// The maximum number of batch job definitions to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBatchJobDefinitionsPaginator is a paginator for ListBatchJobDefinitions
type ListBatchJobDefinitionsPaginator struct {
	options   ListBatchJobDefinitionsPaginatorOptions
	client    ListBatchJobDefinitionsAPIClient
	params    *ListBatchJobDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListBatchJobDefinitionsPaginator returns a new
// ListBatchJobDefinitionsPaginator
func NewListBatchJobDefinitionsPaginator(client ListBatchJobDefinitionsAPIClient, params *ListBatchJobDefinitionsInput, optFns ...func(*ListBatchJobDefinitionsPaginatorOptions)) *ListBatchJobDefinitionsPaginator {
	if params == nil {
		params = &ListBatchJobDefinitionsInput{}
	}

	options := ListBatchJobDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBatchJobDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBatchJobDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBatchJobDefinitions page.
func (p *ListBatchJobDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBatchJobDefinitionsOutput, error) {
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
	result, err := p.client.ListBatchJobDefinitions(ctx, &params, optFns...)
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

// ListBatchJobDefinitionsAPIClient is a client that implements the
// ListBatchJobDefinitions operation.
type ListBatchJobDefinitionsAPIClient interface {
	ListBatchJobDefinitions(context.Context, *ListBatchJobDefinitionsInput, ...func(*Options)) (*ListBatchJobDefinitionsOutput, error)
}

var _ ListBatchJobDefinitionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListBatchJobDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBatchJobDefinitions",
	}
}
