// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifyuibuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of code generation jobs for a specified Amplify app and
// backend environment.
func (c *Client) ListCodegenJobs(ctx context.Context, params *ListCodegenJobsInput, optFns ...func(*Options)) (*ListCodegenJobsOutput, error) {
	if params == nil {
		params = &ListCodegenJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCodegenJobs", params, optFns, c.addOperationListCodegenJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCodegenJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCodegenJobsInput struct {

	// The unique ID for the Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The maximum number of jobs to retrieve.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCodegenJobsOutput struct {

	// The list of code generation jobs for the Amplify app.
	//
	// This member is required.
	Entities []types.CodegenJobSummary

	// The pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCodegenJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCodegenJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCodegenJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCodegenJobs"); err != nil {
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
	if err = addOpListCodegenJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCodegenJobs(options.Region), middleware.Before); err != nil {
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

// ListCodegenJobsPaginatorOptions is the paginator options for ListCodegenJobs
type ListCodegenJobsPaginatorOptions struct {
	// The maximum number of jobs to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCodegenJobsPaginator is a paginator for ListCodegenJobs
type ListCodegenJobsPaginator struct {
	options   ListCodegenJobsPaginatorOptions
	client    ListCodegenJobsAPIClient
	params    *ListCodegenJobsInput
	nextToken *string
	firstPage bool
}

// NewListCodegenJobsPaginator returns a new ListCodegenJobsPaginator
func NewListCodegenJobsPaginator(client ListCodegenJobsAPIClient, params *ListCodegenJobsInput, optFns ...func(*ListCodegenJobsPaginatorOptions)) *ListCodegenJobsPaginator {
	if params == nil {
		params = &ListCodegenJobsInput{}
	}

	options := ListCodegenJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCodegenJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCodegenJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCodegenJobs page.
func (p *ListCodegenJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCodegenJobsOutput, error) {
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
	result, err := p.client.ListCodegenJobs(ctx, &params, optFns...)
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

// ListCodegenJobsAPIClient is a client that implements the ListCodegenJobs
// operation.
type ListCodegenJobsAPIClient interface {
	ListCodegenJobs(context.Context, *ListCodegenJobsInput, ...func(*Options)) (*ListCodegenJobsOutput, error)
}

var _ ListCodegenJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCodegenJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCodegenJobs",
	}
}
