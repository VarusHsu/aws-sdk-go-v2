// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists one or more Amazon Q Business Web Experiences.
func (c *Client) ListWebExperiences(ctx context.Context, params *ListWebExperiencesInput, optFns ...func(*Options)) (*ListWebExperiencesOutput, error) {
	if params == nil {
		params = &ListWebExperiencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWebExperiences", params, optFns, c.addOperationListWebExperiencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWebExperiencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWebExperiencesInput struct {

	// The identifier of the Amazon Q Business application linked to the listed web
	// experiences.
	//
	// This member is required.
	ApplicationId *string

	// The maximum number of Amazon Q Business Web Experiences to return.
	MaxResults *int32

	// If the maxResults response was incomplete because there is more data to
	// retrieve, Amazon Q Business returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of Amazon Q Business
	// conversations.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWebExperiencesOutput struct {

	// If the response is truncated, Amazon Q Business returns this token, which you
	// can use in a later request to list the next set of messages.
	NextToken *string

	// An array of summary information for one or more Amazon Q Business experiences.
	WebExperiences []types.WebExperience

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWebExperiencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWebExperiences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWebExperiences{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWebExperiences"); err != nil {
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
	if err = addOpListWebExperiencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWebExperiences(options.Region), middleware.Before); err != nil {
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

// ListWebExperiencesAPIClient is a client that implements the ListWebExperiences
// operation.
type ListWebExperiencesAPIClient interface {
	ListWebExperiences(context.Context, *ListWebExperiencesInput, ...func(*Options)) (*ListWebExperiencesOutput, error)
}

var _ ListWebExperiencesAPIClient = (*Client)(nil)

// ListWebExperiencesPaginatorOptions is the paginator options for
// ListWebExperiences
type ListWebExperiencesPaginatorOptions struct {
	// The maximum number of Amazon Q Business Web Experiences to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWebExperiencesPaginator is a paginator for ListWebExperiences
type ListWebExperiencesPaginator struct {
	options   ListWebExperiencesPaginatorOptions
	client    ListWebExperiencesAPIClient
	params    *ListWebExperiencesInput
	nextToken *string
	firstPage bool
}

// NewListWebExperiencesPaginator returns a new ListWebExperiencesPaginator
func NewListWebExperiencesPaginator(client ListWebExperiencesAPIClient, params *ListWebExperiencesInput, optFns ...func(*ListWebExperiencesPaginatorOptions)) *ListWebExperiencesPaginator {
	if params == nil {
		params = &ListWebExperiencesInput{}
	}

	options := ListWebExperiencesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWebExperiencesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWebExperiencesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWebExperiences page.
func (p *ListWebExperiencesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWebExperiencesOutput, error) {
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

	result, err := p.client.ListWebExperiences(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWebExperiences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWebExperiences",
	}
}
