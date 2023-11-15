// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List of Trusted Advisor checks summarized for all accounts related to the
// workload.
func (c *Client) ListCheckSummaries(ctx context.Context, params *ListCheckSummariesInput, optFns ...func(*Options)) (*ListCheckSummariesOutput, error) {
	if params == nil {
		params = &ListCheckSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCheckSummaries", params, optFns, c.addOperationListCheckSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCheckSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCheckSummariesInput struct {

	// The ID of a choice.
	//
	// This member is required.
	ChoiceId *string

	// Well-Architected Lens ARN.
	//
	// This member is required.
	LensArn *string

	// The ID used to identify a pillar, for example, security . A pillar is identified
	// by its PillarReviewSummary$PillarId .
	//
	// This member is required.
	PillarId *string

	// The ID of the question.
	//
	// This member is required.
	QuestionId *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	//
	// This member is required.
	WorkloadId *string

	// The maximum number of results to return for this request.
	MaxResults *int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCheckSummariesOutput struct {

	// List of Trusted Advisor summaries related to the Well-Architected best practice.
	CheckSummaries []types.CheckSummary

	// The token to use to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCheckSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCheckSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCheckSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCheckSummaries"); err != nil {
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
	if err = addOpListCheckSummariesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCheckSummaries(options.Region), middleware.Before); err != nil {
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

// ListCheckSummariesAPIClient is a client that implements the ListCheckSummaries
// operation.
type ListCheckSummariesAPIClient interface {
	ListCheckSummaries(context.Context, *ListCheckSummariesInput, ...func(*Options)) (*ListCheckSummariesOutput, error)
}

var _ ListCheckSummariesAPIClient = (*Client)(nil)

// ListCheckSummariesPaginatorOptions is the paginator options for
// ListCheckSummaries
type ListCheckSummariesPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCheckSummariesPaginator is a paginator for ListCheckSummaries
type ListCheckSummariesPaginator struct {
	options   ListCheckSummariesPaginatorOptions
	client    ListCheckSummariesAPIClient
	params    *ListCheckSummariesInput
	nextToken *string
	firstPage bool
}

// NewListCheckSummariesPaginator returns a new ListCheckSummariesPaginator
func NewListCheckSummariesPaginator(client ListCheckSummariesAPIClient, params *ListCheckSummariesInput, optFns ...func(*ListCheckSummariesPaginatorOptions)) *ListCheckSummariesPaginator {
	if params == nil {
		params = &ListCheckSummariesInput{}
	}

	options := ListCheckSummariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCheckSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCheckSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCheckSummaries page.
func (p *ListCheckSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCheckSummariesOutput, error) {
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

	result, err := p.client.ListCheckSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCheckSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCheckSummaries",
	}
}
