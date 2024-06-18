// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Returns the list of access grants in your S3 Access Grants instance.
//
// Permissions You must have the s3:ListAccessGrants permission to use this
// operation.
func (c *Client) ListAccessGrants(ctx context.Context, params *ListAccessGrantsInput, optFns ...func(*Options)) (*ListAccessGrantsOutput, error) {
	if params == nil {
		params = &ListAccessGrantsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessGrants", params, optFns, c.addOperationListAccessGrantsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessGrantsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessGrantsInput struct {

	// The ID of the Amazon Web Services account that is making this request.
	//
	// This member is required.
	AccountId *string

	// The Amazon Resource Name (ARN) of an Amazon Web Services IAM Identity Center
	// application associated with your Identity Center instance. If the grant includes
	// an application ARN, the grantee can only access the S3 data through this
	// application.
	ApplicationArn *string

	// The S3 path of the data to which you are granting access. It is the result of
	// appending the Subprefix to the location scope.
	GrantScope *string

	// The unique identifer of the Grantee . If the grantee type is IAM , the
	// identifier is the IAM Amazon Resource Name (ARN) of the user or role. If the
	// grantee type is a directory user or group, the identifier is 128-bit universally
	// unique identifier (UUID) in the format a1b2c3d4-5678-90ab-cdef-EXAMPLE11111 .
	// You can obtain this UUID from your Amazon Web Services IAM Identity Center
	// instance.
	GranteeIdentifier *string

	// The type of the grantee to which access has been granted. It can be one of the
	// following values:
	//
	//   - IAM - An IAM user or role.
	//
	//   - DIRECTORY_USER - Your corporate directory user. You can use this option if
	//   you have added your corporate identity directory to IAM Identity Center and
	//   associated the IAM Identity Center instance with your S3 Access Grants instance.
	//
	//   - DIRECTORY_GROUP - Your corporate directory group. You can use this option if
	//   you have added your corporate identity directory to IAM Identity Center and
	//   associated the IAM Identity Center instance with your S3 Access Grants instance.
	GranteeType types.GranteeType

	// The maximum number of access grants that you would like returned in the List
	// Access Grants response. If the results include the pagination token NextToken ,
	// make another call using the NextToken to determine if there are more results.
	MaxResults int32

	// A pagination token to request the next page of results. Pass this value into a
	// subsequent List Access Grants request in order to retrieve the next page of
	// results.
	NextToken *string

	// The type of permission granted to your S3 data, which can be set to one of the
	// following values:
	//
	//   - READ – Grant read-only access to the S3 data.
	//
	//   - WRITE – Grant write-only access to the S3 data.
	//
	//   - READWRITE – Grant both read and write access to the S3 data.
	Permission types.Permission

	noSmithyDocumentSerde
}

func (in *ListAccessGrantsInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type ListAccessGrantsOutput struct {

	// A container for a list of grants in an S3 Access Grants instance.
	AccessGrantsList []types.ListAccessGrantEntry

	// A pagination token to request the next page of results. Pass this value into a
	// subsequent List Access Grants request in order to retrieve the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccessGrantsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListAccessGrants{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListAccessGrants{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccessGrants"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opListAccessGrantsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAccessGrantsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessGrants(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListAccessGrantsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

// ListAccessGrantsPaginatorOptions is the paginator options for ListAccessGrants
type ListAccessGrantsPaginatorOptions struct {
	// The maximum number of access grants that you would like returned in the List
	// Access Grants response. If the results include the pagination token NextToken ,
	// make another call using the NextToken to determine if there are more results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessGrantsPaginator is a paginator for ListAccessGrants
type ListAccessGrantsPaginator struct {
	options   ListAccessGrantsPaginatorOptions
	client    ListAccessGrantsAPIClient
	params    *ListAccessGrantsInput
	nextToken *string
	firstPage bool
}

// NewListAccessGrantsPaginator returns a new ListAccessGrantsPaginator
func NewListAccessGrantsPaginator(client ListAccessGrantsAPIClient, params *ListAccessGrantsInput, optFns ...func(*ListAccessGrantsPaginatorOptions)) *ListAccessGrantsPaginator {
	if params == nil {
		params = &ListAccessGrantsInput{}
	}

	options := ListAccessGrantsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessGrantsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessGrantsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccessGrants page.
func (p *ListAccessGrantsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessGrantsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListAccessGrants(ctx, &params, optFns...)
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

type endpointPrefix_opListAccessGrantsMiddleware struct {
}

func (*endpointPrefix_opListAccessGrantsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAccessGrantsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*ListAccessGrantsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListAccessGrantsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListAccessGrantsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListAccessGrantsAPIClient is a client that implements the ListAccessGrants
// operation.
type ListAccessGrantsAPIClient interface {
	ListAccessGrants(context.Context, *ListAccessGrantsInput, ...func(*Options)) (*ListAccessGrantsOutput, error)
}

var _ ListAccessGrantsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAccessGrants(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccessGrants",
	}
}

func copyListAccessGrantsInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*ListAccessGrantsInput)
	if !ok {
		return nil, fmt.Errorf("expect *ListAccessGrantsInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *ListAccessGrantsInput) copy() interface{} {
	v := *in
	return &v
}
func backFillListAccessGrantsAccountID(input interface{}, v string) error {
	in := input.(*ListAccessGrantsInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addListAccessGrantsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyListAccessGrantsInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
