// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	This is a paginated call to list linked accounts that are linked to the payer
//
// account for the specified time period. If no information is provided, the
// current billing period is used. The response will optionally include the billing
// group that's associated with the linked account.
func (c *Client) ListAccountAssociations(ctx context.Context, params *ListAccountAssociationsInput, optFns ...func(*Options)) (*ListAccountAssociationsOutput, error) {
	if params == nil {
		params = &ListAccountAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountAssociations", params, optFns, c.addOperationListAccountAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountAssociationsInput struct {

	//  The preferred billing period to get account associations.
	BillingPeriod *string

	// The filter on the account ID of the linked account, or any of the following:
	//
	// MONITORED : linked accounts that are associated to billing groups.
	//
	// UNMONITORED : linked accounts that aren't associated to billing groups.
	//
	// Billing Group Arn : linked accounts that are associated to the provided billing
	// group Arn.
	Filters *types.ListAccountAssociationsFilter

	//  The pagination token that's used on subsequent calls to retrieve accounts.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAccountAssociationsOutput struct {

	//  The list of linked accounts in the payer account.
	LinkedAccounts []types.AccountAssociationsListElement

	//  The pagination token that's used on subsequent calls to get accounts.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAccountAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAccountAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccountAssociations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountAssociations(options.Region), middleware.Before); err != nil {
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

// ListAccountAssociationsPaginatorOptions is the paginator options for
// ListAccountAssociations
type ListAccountAssociationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountAssociationsPaginator is a paginator for ListAccountAssociations
type ListAccountAssociationsPaginator struct {
	options   ListAccountAssociationsPaginatorOptions
	client    ListAccountAssociationsAPIClient
	params    *ListAccountAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListAccountAssociationsPaginator returns a new
// ListAccountAssociationsPaginator
func NewListAccountAssociationsPaginator(client ListAccountAssociationsAPIClient, params *ListAccountAssociationsInput, optFns ...func(*ListAccountAssociationsPaginatorOptions)) *ListAccountAssociationsPaginator {
	if params == nil {
		params = &ListAccountAssociationsInput{}
	}

	options := ListAccountAssociationsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccountAssociations page.
func (p *ListAccountAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListAccountAssociations(ctx, &params, optFns...)
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

// ListAccountAssociationsAPIClient is a client that implements the
// ListAccountAssociations operation.
type ListAccountAssociationsAPIClient interface {
	ListAccountAssociations(context.Context, *ListAccountAssociationsInput, ...func(*Options)) (*ListAccountAssociationsOutput, error)
}

var _ ListAccountAssociationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAccountAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccountAssociations",
	}
}
