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

// Lists the assignee of the specified Amazon Web Services account with the
// specified permission set.
func (c *Client) ListAccountAssignments(ctx context.Context, params *ListAccountAssignmentsInput, optFns ...func(*Options)) (*ListAccountAssignmentsOutput, error) {
	if params == nil {
		params = &ListAccountAssignmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountAssignments", params, optFns, c.addOperationListAccountAssignmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountAssignmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountAssignmentsInput struct {

	// The identifier of the Amazon Web Services account from which to list the
	// assignments.
	//
	// This member is required.
	AccountId *string

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services Service Namespacesin the Amazon Web Services
	// General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the permission set from which to list assignments.
	//
	// This member is required.
	PermissionSetArn *string

	// The maximum number of results to display for the assignment.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAccountAssignmentsOutput struct {

	// The list of assignments that match the input Amazon Web Services account and
	// permission set.
	AccountAssignments []types.AccountAssignment

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountAssignmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAccountAssignments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAccountAssignments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccountAssignments"); err != nil {
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
	if err = addOpListAccountAssignmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountAssignments(options.Region), middleware.Before); err != nil {
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

// ListAccountAssignmentsPaginatorOptions is the paginator options for
// ListAccountAssignments
type ListAccountAssignmentsPaginatorOptions struct {
	// The maximum number of results to display for the assignment.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountAssignmentsPaginator is a paginator for ListAccountAssignments
type ListAccountAssignmentsPaginator struct {
	options   ListAccountAssignmentsPaginatorOptions
	client    ListAccountAssignmentsAPIClient
	params    *ListAccountAssignmentsInput
	nextToken *string
	firstPage bool
}

// NewListAccountAssignmentsPaginator returns a new ListAccountAssignmentsPaginator
func NewListAccountAssignmentsPaginator(client ListAccountAssignmentsAPIClient, params *ListAccountAssignmentsInput, optFns ...func(*ListAccountAssignmentsPaginatorOptions)) *ListAccountAssignmentsPaginator {
	if params == nil {
		params = &ListAccountAssignmentsInput{}
	}

	options := ListAccountAssignmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountAssignmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountAssignmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccountAssignments page.
func (p *ListAccountAssignmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountAssignmentsOutput, error) {
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
	result, err := p.client.ListAccountAssignments(ctx, &params, optFns...)
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

// ListAccountAssignmentsAPIClient is a client that implements the
// ListAccountAssignments operation.
type ListAccountAssignmentsAPIClient interface {
	ListAccountAssignments(context.Context, *ListAccountAssignmentsInput, ...func(*Options)) (*ListAccountAssignmentsOutput, error)
}

var _ ListAccountAssignmentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAccountAssignments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccountAssignments",
	}
}
