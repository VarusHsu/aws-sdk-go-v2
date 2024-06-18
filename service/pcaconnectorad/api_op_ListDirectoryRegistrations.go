// Code generated by smithy-go-codegen DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorad/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the directory registrations that you created by using the [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration] action.
//
// [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration]: https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration
func (c *Client) ListDirectoryRegistrations(ctx context.Context, params *ListDirectoryRegistrationsInput, optFns ...func(*Options)) (*ListDirectoryRegistrationsOutput, error) {
	if params == nil {
		params = &ListDirectoryRegistrationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDirectoryRegistrations", params, optFns, c.addOperationListDirectoryRegistrationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDirectoryRegistrationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDirectoryRegistrationsInput struct {

	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response on each page. If additional items exist beyond
	// the number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	MaxResults *int32

	// Use this parameter when paginating results in a subsequent request after you
	// receive a response with truncated results. Set it to the value of the NextToken
	// parameter from the response you just received.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDirectoryRegistrationsOutput struct {

	// Summary information about each directory registration you have created.
	DirectoryRegistrations []types.DirectoryRegistrationSummary

	// Use this parameter when paginating results in a subsequent request after you
	// receive a response with truncated results. Set it to the value of the NextToken
	// parameter from the response you just received.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDirectoryRegistrationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDirectoryRegistrations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDirectoryRegistrations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDirectoryRegistrations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDirectoryRegistrations(options.Region), middleware.Before); err != nil {
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

// ListDirectoryRegistrationsPaginatorOptions is the paginator options for
// ListDirectoryRegistrations
type ListDirectoryRegistrationsPaginatorOptions struct {
	// Use this parameter when paginating results to specify the maximum number of
	// items to return in the response on each page. If additional items exist beyond
	// the number you specify, the NextToken element is sent in the response. Use this
	// NextToken value in a subsequent request to retrieve additional items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDirectoryRegistrationsPaginator is a paginator for
// ListDirectoryRegistrations
type ListDirectoryRegistrationsPaginator struct {
	options   ListDirectoryRegistrationsPaginatorOptions
	client    ListDirectoryRegistrationsAPIClient
	params    *ListDirectoryRegistrationsInput
	nextToken *string
	firstPage bool
}

// NewListDirectoryRegistrationsPaginator returns a new
// ListDirectoryRegistrationsPaginator
func NewListDirectoryRegistrationsPaginator(client ListDirectoryRegistrationsAPIClient, params *ListDirectoryRegistrationsInput, optFns ...func(*ListDirectoryRegistrationsPaginatorOptions)) *ListDirectoryRegistrationsPaginator {
	if params == nil {
		params = &ListDirectoryRegistrationsInput{}
	}

	options := ListDirectoryRegistrationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDirectoryRegistrationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDirectoryRegistrationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDirectoryRegistrations page.
func (p *ListDirectoryRegistrationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDirectoryRegistrationsOutput, error) {
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
	result, err := p.client.ListDirectoryRegistrations(ctx, &params, optFns...)
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

// ListDirectoryRegistrationsAPIClient is a client that implements the
// ListDirectoryRegistrations operation.
type ListDirectoryRegistrationsAPIClient interface {
	ListDirectoryRegistrations(context.Context, *ListDirectoryRegistrationsInput, ...func(*Options)) (*ListDirectoryRegistrationsOutput, error)
}

var _ ListDirectoryRegistrationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDirectoryRegistrations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDirectoryRegistrations",
	}
}
