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

// This operation is not supported by directory buckets.
//
// Returns a list of the access points that are owned by the current account
// that's associated with the specified bucket. You can retrieve up to 1000 access
// points per call. If the specified bucket has more than 1,000 access points (or
// the number specified in maxResults , whichever is less), the response will
// include a continuation token that you can use to list the additional access
// points.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to ListAccessPoints :
//
// [CreateAccessPoint]
//
// [DeleteAccessPoint]
//
// [GetAccessPoint]
//
// [CreateAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html
// [GetAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html
// [DeleteAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html#API_control_GetAccessPoint_Examples
func (c *Client) ListAccessPoints(ctx context.Context, params *ListAccessPointsInput, optFns ...func(*Options)) (*ListAccessPointsOutput, error) {
	if params == nil {
		params = &ListAccessPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessPoints", params, optFns, c.addOperationListAccessPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessPointsInput struct {

	// The Amazon Web Services account ID for the account that owns the specified
	// access points.
	//
	// This member is required.
	AccountId *string

	// The name of the bucket whose associated access points you want to list.
	//
	// For using this parameter with Amazon S3 on Outposts with the REST API, you must
	// specify the name and the x-amz-outpost-id as well.
	//
	// For using this parameter with S3 on Outposts with the Amazon Web Services SDK
	// and CLI, you must specify the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/ . For example, to access the bucket
	// reports through Outpost my-outpost owned by account 123456789012 in Region
	// us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	Bucket *string

	// The maximum number of access points that you want to include in the list. If
	// the specified bucket has more than this number of access points, then the
	// response will include a continuation token in the NextToken field that you can
	// use to retrieve the next page of access points.
	MaxResults int32

	// A continuation token. If a previous call to ListAccessPoints returned a
	// continuation token in the NextToken field, then providing that value here
	// causes Amazon S3 to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

func (in *ListAccessPointsInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.Bucket = in.Bucket
	p.RequiresAccountId = ptr.Bool(true)
}

type ListAccessPointsOutput struct {

	// Contains identification and configuration information for one or more access
	// points associated with the specified bucket.
	AccessPointList []types.AccessPoint

	// If the specified bucket has more access points than can be returned in one call
	// to this API, this field contains a continuation token that you can provide in
	// subsequent calls to this API to retrieve additional access points.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccessPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListAccessPoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListAccessPoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccessPoints"); err != nil {
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
	if err = addEndpointPrefix_opListAccessPointsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAccessPointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessPoints(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListAccessPointsUpdateEndpoint(stack, options); err != nil {
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

// ListAccessPointsPaginatorOptions is the paginator options for ListAccessPoints
type ListAccessPointsPaginatorOptions struct {
	// The maximum number of access points that you want to include in the list. If
	// the specified bucket has more than this number of access points, then the
	// response will include a continuation token in the NextToken field that you can
	// use to retrieve the next page of access points.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessPointsPaginator is a paginator for ListAccessPoints
type ListAccessPointsPaginator struct {
	options   ListAccessPointsPaginatorOptions
	client    ListAccessPointsAPIClient
	params    *ListAccessPointsInput
	nextToken *string
	firstPage bool
}

// NewListAccessPointsPaginator returns a new ListAccessPointsPaginator
func NewListAccessPointsPaginator(client ListAccessPointsAPIClient, params *ListAccessPointsInput, optFns ...func(*ListAccessPointsPaginatorOptions)) *ListAccessPointsPaginator {
	if params == nil {
		params = &ListAccessPointsInput{}
	}

	options := ListAccessPointsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessPointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessPointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccessPoints page.
func (p *ListAccessPointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessPointsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListAccessPoints(ctx, &params, optFns...)
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

func (m *ListAccessPointsInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *ListAccessPointsInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opListAccessPointsMiddleware struct {
}

func (*endpointPrefix_opListAccessPointsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAccessPointsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
	input, ok := opaqueInput.(*ListAccessPointsInput)
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
func addEndpointPrefix_opListAccessPointsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListAccessPointsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListAccessPointsAPIClient is a client that implements the ListAccessPoints
// operation.
type ListAccessPointsAPIClient interface {
	ListAccessPoints(context.Context, *ListAccessPointsInput, ...func(*Options)) (*ListAccessPointsOutput, error)
}

var _ ListAccessPointsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAccessPoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccessPoints",
	}
}

func copyListAccessPointsInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*ListAccessPointsInput)
	if !ok {
		return nil, fmt.Errorf("expect *ListAccessPointsInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *ListAccessPointsInput) copy() interface{} {
	v := *in
	return &v
}
func getListAccessPointsARNMember(input interface{}) (*string, bool) {
	in := input.(*ListAccessPointsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setListAccessPointsARNMember(input interface{}, v string) error {
	in := input.(*ListAccessPointsInput)
	in.Bucket = &v
	return nil
}
func backFillListAccessPointsAccountID(input interface{}, v string) error {
	in := input.(*ListAccessPointsInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addListAccessPointsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getListAccessPointsARNMember,
			BackfillAccountID: backFillListAccessPointsAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setListAccessPointsARNMember,
			CopyInput:         copyListAccessPointsInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
