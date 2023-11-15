// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the phone numbers for the specified Amazon Chime SDK account, Amazon
// Chime SDK user, Amazon Chime SDK Voice Connector, or Amazon Chime SDK Voice
// Connector group.
func (c *Client) ListPhoneNumbers(ctx context.Context, params *ListPhoneNumbersInput, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) {
	if params == nil {
		params = &ListPhoneNumbersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPhoneNumbers", params, optFns, c.addOperationListPhoneNumbersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPhoneNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPhoneNumbersInput struct {

	// The filter to limit the number of results.
	FilterName types.PhoneNumberAssociationName

	// The filter value.
	FilterValue *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token used to return the next page of results.
	NextToken *string

	// The phone number product types.
	ProductType types.PhoneNumberProductType

	// The status of your organization's phone numbers.
	Status *string

	noSmithyDocumentSerde
}

type ListPhoneNumbersOutput struct {

	// The token used to return the next page of results.
	NextToken *string

	// The phone number details.
	PhoneNumbers []types.PhoneNumber

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPhoneNumbersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPhoneNumbers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPhoneNumbers(options.Region), middleware.Before); err != nil {
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

// ListPhoneNumbersAPIClient is a client that implements the ListPhoneNumbers
// operation.
type ListPhoneNumbersAPIClient interface {
	ListPhoneNumbers(context.Context, *ListPhoneNumbersInput, ...func(*Options)) (*ListPhoneNumbersOutput, error)
}

var _ ListPhoneNumbersAPIClient = (*Client)(nil)

// ListPhoneNumbersPaginatorOptions is the paginator options for ListPhoneNumbers
type ListPhoneNumbersPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPhoneNumbersPaginator is a paginator for ListPhoneNumbers
type ListPhoneNumbersPaginator struct {
	options   ListPhoneNumbersPaginatorOptions
	client    ListPhoneNumbersAPIClient
	params    *ListPhoneNumbersInput
	nextToken *string
	firstPage bool
}

// NewListPhoneNumbersPaginator returns a new ListPhoneNumbersPaginator
func NewListPhoneNumbersPaginator(client ListPhoneNumbersAPIClient, params *ListPhoneNumbersInput, optFns ...func(*ListPhoneNumbersPaginatorOptions)) *ListPhoneNumbersPaginator {
	if params == nil {
		params = &ListPhoneNumbersInput{}
	}

	options := ListPhoneNumbersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPhoneNumbersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPhoneNumbersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPhoneNumbers page.
func (p *ListPhoneNumbersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPhoneNumbersOutput, error) {
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

	result, err := p.client.ListPhoneNumbers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPhoneNumbers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPhoneNumbers",
	}
}
