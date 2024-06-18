// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the quota increase requests for the specified quota.
func (c *Client) ListRequestedServiceQuotaChangeHistoryByQuota(ctx context.Context, params *ListRequestedServiceQuotaChangeHistoryByQuotaInput, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	if params == nil {
		params = &ListRequestedServiceQuotaChangeHistoryByQuotaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRequestedServiceQuotaChangeHistoryByQuota", params, optFns, c.addOperationListRequestedServiceQuotaChangeHistoryByQuotaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRequestedServiceQuotaChangeHistoryByQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRequestedServiceQuotaChangeHistoryByQuotaInput struct {

	// Specifies the quota identifier. To find the quota code for a specific quota,
	// use the ListServiceQuotasoperation, and look for the QuotaCode response in the output for the
	// quota you want.
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServicesoperation.
	//
	// This member is required.
	ServiceCode *string

	// Specifies the maximum number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results.
	//
	// An API operation can return fewer results than the maximum even when there are
	// more results available. You should check NextToken after every operation to
	// ensure that you receive all of the results.
	MaxResults *int32

	// Specifies a value for receiving additional results after you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	// Specifies at which level within the Amazon Web Services account the quota
	// request applies to.
	QuotaRequestedAtLevel types.AppliedLevelEnum

	// Specifies that you want to filter the results to only the requests with the
	// matching status.
	Status types.RequestStatus

	noSmithyDocumentSerde
}

type ListRequestedServiceQuotaChangeHistoryByQuotaOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string

	// Information about the quota increase requests.
	RequestedQuotas []types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRequestedServiceQuotaChangeHistoryByQuotaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRequestedServiceQuotaChangeHistoryByQuota{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRequestedServiceQuotaChangeHistoryByQuota{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRequestedServiceQuotaChangeHistoryByQuota"); err != nil {
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
	if err = addOpListRequestedServiceQuotaChangeHistoryByQuotaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistoryByQuota(options.Region), middleware.Before); err != nil {
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

// ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions is the paginator
// options for ListRequestedServiceQuotaChangeHistoryByQuota
type ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions struct {
	// Specifies the maximum number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results.
	//
	// An API operation can return fewer results than the maximum even when there are
	// more results available. You should check NextToken after every operation to
	// ensure that you receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRequestedServiceQuotaChangeHistoryByQuotaPaginator is a paginator for
// ListRequestedServiceQuotaChangeHistoryByQuota
type ListRequestedServiceQuotaChangeHistoryByQuotaPaginator struct {
	options   ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions
	client    ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient
	params    *ListRequestedServiceQuotaChangeHistoryByQuotaInput
	nextToken *string
	firstPage bool
}

// NewListRequestedServiceQuotaChangeHistoryByQuotaPaginator returns a new
// ListRequestedServiceQuotaChangeHistoryByQuotaPaginator
func NewListRequestedServiceQuotaChangeHistoryByQuotaPaginator(client ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient, params *ListRequestedServiceQuotaChangeHistoryByQuotaInput, optFns ...func(*ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions)) *ListRequestedServiceQuotaChangeHistoryByQuotaPaginator {
	if params == nil {
		params = &ListRequestedServiceQuotaChangeHistoryByQuotaInput{}
	}

	options := ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRequestedServiceQuotaChangeHistoryByQuotaPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRequestedServiceQuotaChangeHistoryByQuotaPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRequestedServiceQuotaChangeHistoryByQuota page.
func (p *ListRequestedServiceQuotaChangeHistoryByQuotaPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
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
	result, err := p.client.ListRequestedServiceQuotaChangeHistoryByQuota(ctx, &params, optFns...)
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

// ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient is a client that
// implements the ListRequestedServiceQuotaChangeHistoryByQuota operation.
type ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient interface {
	ListRequestedServiceQuotaChangeHistoryByQuota(context.Context, *ListRequestedServiceQuotaChangeHistoryByQuotaInput, ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error)
}

var _ ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistoryByQuota(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRequestedServiceQuotaChangeHistoryByQuota",
	}
}
