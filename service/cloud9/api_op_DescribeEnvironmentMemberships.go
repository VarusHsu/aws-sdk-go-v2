// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about environment members for an Cloud9 development
// environment.
func (c *Client) DescribeEnvironmentMemberships(ctx context.Context, params *DescribeEnvironmentMembershipsInput, optFns ...func(*Options)) (*DescribeEnvironmentMembershipsOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentMembershipsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentMemberships", params, optFns, c.addOperationDescribeEnvironmentMembershipsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentMembershipsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEnvironmentMembershipsInput struct {

	// The ID of the environment to get environment member information about.
	EnvironmentId *string

	// The maximum number of environment members to get information about.
	MaxResults *int32

	// During a previous call, if there are more than 25 items in the list, only the
	// first 25 items are returned, along with a unique string called a next token. To
	// get the next batch of items in the list, call this operation again, adding the
	// next token to the call. To get all of the items in the list, keep calling this
	// operation with each subsequent next token that is returned, until no more next
	// tokens are returned.
	NextToken *string

	// The type of environment member permissions to get information about. Available
	// values include:
	//
	//   - owner : Owns the environment.
	//
	//   - read-only : Has read-only access to the environment.
	//
	//   - read-write : Has read-write access to the environment.
	//
	// If no value is specified, information about all environment members are
	// returned.
	Permissions []types.Permissions

	// The Amazon Resource Name (ARN) of an individual environment member to get
	// information about. If no value is specified, information about all environment
	// members are returned.
	UserArn *string

	noSmithyDocumentSerde
}

type DescribeEnvironmentMembershipsOutput struct {

	// Information about the environment members for the environment.
	Memberships []types.EnvironmentMember

	// If there are more than 25 items in the list, only the first 25 items are
	// returned, along with a unique string called a next token. To get the next batch
	// of items in the list, call this operation again, adding the next token to the
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEnvironmentMembershipsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEnvironmentMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEnvironmentMemberships{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEnvironmentMemberships"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentMemberships(options.Region), middleware.Before); err != nil {
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

// DescribeEnvironmentMembershipsPaginatorOptions is the paginator options for
// DescribeEnvironmentMemberships
type DescribeEnvironmentMembershipsPaginatorOptions struct {
	// The maximum number of environment members to get information about.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEnvironmentMembershipsPaginator is a paginator for
// DescribeEnvironmentMemberships
type DescribeEnvironmentMembershipsPaginator struct {
	options   DescribeEnvironmentMembershipsPaginatorOptions
	client    DescribeEnvironmentMembershipsAPIClient
	params    *DescribeEnvironmentMembershipsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEnvironmentMembershipsPaginator returns a new
// DescribeEnvironmentMembershipsPaginator
func NewDescribeEnvironmentMembershipsPaginator(client DescribeEnvironmentMembershipsAPIClient, params *DescribeEnvironmentMembershipsInput, optFns ...func(*DescribeEnvironmentMembershipsPaginatorOptions)) *DescribeEnvironmentMembershipsPaginator {
	if params == nil {
		params = &DescribeEnvironmentMembershipsInput{}
	}

	options := DescribeEnvironmentMembershipsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEnvironmentMembershipsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEnvironmentMembershipsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEnvironmentMemberships page.
func (p *DescribeEnvironmentMembershipsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEnvironmentMembershipsOutput, error) {
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
	result, err := p.client.DescribeEnvironmentMemberships(ctx, &params, optFns...)
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

// DescribeEnvironmentMembershipsAPIClient is a client that implements the
// DescribeEnvironmentMemberships operation.
type DescribeEnvironmentMembershipsAPIClient interface {
	DescribeEnvironmentMemberships(context.Context, *DescribeEnvironmentMembershipsInput, ...func(*Options)) (*DescribeEnvironmentMembershipsOutput, error)
}

var _ DescribeEnvironmentMembershipsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeEnvironmentMemberships(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEnvironmentMemberships",
	}
}
