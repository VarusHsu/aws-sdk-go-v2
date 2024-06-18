// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all customer managed policies attached to a specified PermissionSet.
func (c *Client) ListCustomerManagedPolicyReferencesInPermissionSet(ctx context.Context, params *ListCustomerManagedPolicyReferencesInPermissionSetInput, optFns ...func(*Options)) (*ListCustomerManagedPolicyReferencesInPermissionSetOutput, error) {
	if params == nil {
		params = &ListCustomerManagedPolicyReferencesInPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomerManagedPolicyReferencesInPermissionSet", params, optFns, c.addOperationListCustomerManagedPolicyReferencesInPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomerManagedPolicyReferencesInPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomerManagedPolicyReferencesInPermissionSetInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the PermissionSet .
	//
	// This member is required.
	PermissionSetArn *string

	// The maximum number of results to display for the list call.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomerManagedPolicyReferencesInPermissionSetOutput struct {

	// Specifies the names and paths of the customer managed policies that you have
	// attached to your permission set.
	CustomerManagedPolicyReferences []types.CustomerManagedPolicyReference

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomerManagedPolicyReferencesInPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCustomerManagedPolicyReferencesInPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCustomerManagedPolicyReferencesInPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCustomerManagedPolicyReferencesInPermissionSet"); err != nil {
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
	if err = addOpListCustomerManagedPolicyReferencesInPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomerManagedPolicyReferencesInPermissionSet(options.Region), middleware.Before); err != nil {
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

// ListCustomerManagedPolicyReferencesInPermissionSetPaginatorOptions is the
// paginator options for ListCustomerManagedPolicyReferencesInPermissionSet
type ListCustomerManagedPolicyReferencesInPermissionSetPaginatorOptions struct {
	// The maximum number of results to display for the list call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomerManagedPolicyReferencesInPermissionSetPaginator is a paginator for
// ListCustomerManagedPolicyReferencesInPermissionSet
type ListCustomerManagedPolicyReferencesInPermissionSetPaginator struct {
	options   ListCustomerManagedPolicyReferencesInPermissionSetPaginatorOptions
	client    ListCustomerManagedPolicyReferencesInPermissionSetAPIClient
	params    *ListCustomerManagedPolicyReferencesInPermissionSetInput
	nextToken *string
	firstPage bool
}

// NewListCustomerManagedPolicyReferencesInPermissionSetPaginator returns a new
// ListCustomerManagedPolicyReferencesInPermissionSetPaginator
func NewListCustomerManagedPolicyReferencesInPermissionSetPaginator(client ListCustomerManagedPolicyReferencesInPermissionSetAPIClient, params *ListCustomerManagedPolicyReferencesInPermissionSetInput, optFns ...func(*ListCustomerManagedPolicyReferencesInPermissionSetPaginatorOptions)) *ListCustomerManagedPolicyReferencesInPermissionSetPaginator {
	if params == nil {
		params = &ListCustomerManagedPolicyReferencesInPermissionSetInput{}
	}

	options := ListCustomerManagedPolicyReferencesInPermissionSetPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomerManagedPolicyReferencesInPermissionSetPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomerManagedPolicyReferencesInPermissionSetPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomerManagedPolicyReferencesInPermissionSet
// page.
func (p *ListCustomerManagedPolicyReferencesInPermissionSetPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomerManagedPolicyReferencesInPermissionSetOutput, error) {
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
	result, err := p.client.ListCustomerManagedPolicyReferencesInPermissionSet(ctx, &params, optFns...)
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

// ListCustomerManagedPolicyReferencesInPermissionSetAPIClient is a client that
// implements the ListCustomerManagedPolicyReferencesInPermissionSet operation.
type ListCustomerManagedPolicyReferencesInPermissionSetAPIClient interface {
	ListCustomerManagedPolicyReferencesInPermissionSet(context.Context, *ListCustomerManagedPolicyReferencesInPermissionSetInput, ...func(*Options)) (*ListCustomerManagedPolicyReferencesInPermissionSetOutput, error)
}

var _ ListCustomerManagedPolicyReferencesInPermissionSetAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListCustomerManagedPolicyReferencesInPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCustomerManagedPolicyReferencesInPermissionSet",
	}
}
