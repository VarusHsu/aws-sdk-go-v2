// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the Compute Optimizer enrollment (opt-in) status of organization member
// accounts, if your account is an organization management account. To get the
// enrollment status of standalone accounts, use the GetEnrollmentStatus action.
func (c *Client) GetEnrollmentStatusesForOrganization(ctx context.Context, params *GetEnrollmentStatusesForOrganizationInput, optFns ...func(*Options)) (*GetEnrollmentStatusesForOrganizationOutput, error) {
	if params == nil {
		params = &GetEnrollmentStatusesForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnrollmentStatusesForOrganization", params, optFns, c.addOperationGetEnrollmentStatusesForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnrollmentStatusesForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnrollmentStatusesForOrganizationInput struct {

	// An array of objects to specify a filter that returns a more specific list of
	// account enrollment statuses.
	Filters []types.EnrollmentFilter

	// The maximum number of account enrollment statuses to return with a single
	// request. You can specify up to 100 statuses to return with each request. To
	// retrieve the remaining results, make another request with the returned nextToken
	// value.
	MaxResults *int32

	// The token to advance to the next page of account enrollment statuses.
	NextToken *string

	noSmithyDocumentSerde
}

type GetEnrollmentStatusesForOrganizationOutput struct {

	// An array of objects that describe the enrollment statuses of organization
	// member accounts.
	AccountEnrollmentStatuses []types.AccountEnrollmentStatus

	// The token to use to advance to the next page of account enrollment statuses.
	// This value is null when there are no more pages of account enrollment statuses
	// to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnrollmentStatusesForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEnrollmentStatusesForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEnrollmentStatusesForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEnrollmentStatusesForOrganization"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnrollmentStatusesForOrganization(options.Region), middleware.Before); err != nil {
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

// GetEnrollmentStatusesForOrganizationAPIClient is a client that implements the
// GetEnrollmentStatusesForOrganization operation.
type GetEnrollmentStatusesForOrganizationAPIClient interface {
	GetEnrollmentStatusesForOrganization(context.Context, *GetEnrollmentStatusesForOrganizationInput, ...func(*Options)) (*GetEnrollmentStatusesForOrganizationOutput, error)
}

var _ GetEnrollmentStatusesForOrganizationAPIClient = (*Client)(nil)

// GetEnrollmentStatusesForOrganizationPaginatorOptions is the paginator options
// for GetEnrollmentStatusesForOrganization
type GetEnrollmentStatusesForOrganizationPaginatorOptions struct {
	// The maximum number of account enrollment statuses to return with a single
	// request. You can specify up to 100 statuses to return with each request. To
	// retrieve the remaining results, make another request with the returned nextToken
	// value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetEnrollmentStatusesForOrganizationPaginator is a paginator for
// GetEnrollmentStatusesForOrganization
type GetEnrollmentStatusesForOrganizationPaginator struct {
	options   GetEnrollmentStatusesForOrganizationPaginatorOptions
	client    GetEnrollmentStatusesForOrganizationAPIClient
	params    *GetEnrollmentStatusesForOrganizationInput
	nextToken *string
	firstPage bool
}

// NewGetEnrollmentStatusesForOrganizationPaginator returns a new
// GetEnrollmentStatusesForOrganizationPaginator
func NewGetEnrollmentStatusesForOrganizationPaginator(client GetEnrollmentStatusesForOrganizationAPIClient, params *GetEnrollmentStatusesForOrganizationInput, optFns ...func(*GetEnrollmentStatusesForOrganizationPaginatorOptions)) *GetEnrollmentStatusesForOrganizationPaginator {
	if params == nil {
		params = &GetEnrollmentStatusesForOrganizationInput{}
	}

	options := GetEnrollmentStatusesForOrganizationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetEnrollmentStatusesForOrganizationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetEnrollmentStatusesForOrganizationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetEnrollmentStatusesForOrganization page.
func (p *GetEnrollmentStatusesForOrganizationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetEnrollmentStatusesForOrganizationOutput, error) {
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

	result, err := p.client.GetEnrollmentStatusesForOrganization(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetEnrollmentStatusesForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEnrollmentStatusesForOrganization",
	}
}
