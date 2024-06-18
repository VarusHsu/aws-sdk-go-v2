// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the account creation requests that match the specified status that is
// currently being tracked for the organization.
//
// Always check the NextToken response parameter for a null value when calling a
// List* operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// This operation can be called only from the organization's management account or
// by a member account that is a delegated administrator for an Amazon Web Services
// service.
func (c *Client) ListCreateAccountStatus(ctx context.Context, params *ListCreateAccountStatusInput, optFns ...func(*Options)) (*ListCreateAccountStatusOutput, error) {
	if params == nil {
		params = &ListCreateAccountStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCreateAccountStatus", params, optFns, c.addOperationListCreateAccountStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCreateAccountStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCreateAccountStatusInput struct {

	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	// A list of one or more states that you want included in the response. If this
	// parameter isn't present, all requests are included in the response.
	States []types.CreateAccountState

	noSmithyDocumentSerde
}

type ListCreateAccountStatusOutput struct {

	// A list of objects with details about the requests. Certain elements, such as
	// the accountId number, are present in the output only after the account has been
	// successfully created.
	CreateAccountStatuses []types.CreateAccountStatus

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCreateAccountStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCreateAccountStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCreateAccountStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCreateAccountStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCreateAccountStatus(options.Region), middleware.Before); err != nil {
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

// ListCreateAccountStatusPaginatorOptions is the paginator options for
// ListCreateAccountStatus
type ListCreateAccountStatusPaginatorOptions struct {
	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCreateAccountStatusPaginator is a paginator for ListCreateAccountStatus
type ListCreateAccountStatusPaginator struct {
	options   ListCreateAccountStatusPaginatorOptions
	client    ListCreateAccountStatusAPIClient
	params    *ListCreateAccountStatusInput
	nextToken *string
	firstPage bool
}

// NewListCreateAccountStatusPaginator returns a new
// ListCreateAccountStatusPaginator
func NewListCreateAccountStatusPaginator(client ListCreateAccountStatusAPIClient, params *ListCreateAccountStatusInput, optFns ...func(*ListCreateAccountStatusPaginatorOptions)) *ListCreateAccountStatusPaginator {
	if params == nil {
		params = &ListCreateAccountStatusInput{}
	}

	options := ListCreateAccountStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCreateAccountStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCreateAccountStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCreateAccountStatus page.
func (p *ListCreateAccountStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCreateAccountStatusOutput, error) {
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
	result, err := p.client.ListCreateAccountStatus(ctx, &params, optFns...)
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

// ListCreateAccountStatusAPIClient is a client that implements the
// ListCreateAccountStatus operation.
type ListCreateAccountStatusAPIClient interface {
	ListCreateAccountStatus(context.Context, *ListCreateAccountStatusInput, ...func(*Options)) (*ListCreateAccountStatusOutput, error)
}

var _ ListCreateAccountStatusAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCreateAccountStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCreateAccountStatus",
	}
}
