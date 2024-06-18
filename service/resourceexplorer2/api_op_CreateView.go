// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a view that users can query by using the Search operation. Results from
// queries that you make using this view include only resources that match the
// view's Filters . For more information about Amazon Web Services Resource
// Explorer views, see [Managing views]in the Amazon Web Services Resource Explorer User Guide.
//
// Only the principals with an IAM identity-based policy that grants Allow to the
// Search action on a Resource with the [Amazon resource name (ARN)] of this view can Search using views you create
// with this operation.
//
// [Managing views]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views.html
// [Amazon resource name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
func (c *Client) CreateView(ctx context.Context, params *CreateViewInput, optFns ...func(*Options)) (*CreateViewOutput, error) {
	if params == nil {
		params = &CreateViewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateView", params, optFns, c.addOperationCreateViewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateViewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateViewInput struct {

	// The name of the new view. This name appears in the list of views in Resource
	// Explorer.
	//
	// The name must be no more than 64 characters long, and can include letters,
	// digits, and the dash (-) character. The name must be unique within its Amazon
	// Web Services Region.
	//
	// This member is required.
	ViewName *string

	// This value helps ensure idempotency. Resource Explorer uses this value to
	// prevent the accidental creation of duplicate versions. We recommend that you
	// generate a [UUID-type value]to ensure the uniqueness of your views.
	//
	// [UUID-type value]: https://wikipedia.org/wiki/Universally_unique_identifier
	ClientToken *string

	// An array of strings that specify which resources are included in the results of
	// queries made using this view. When you use this view in a Searchoperation, the filter
	// string is combined with the search's QueryString parameter using a logical AND
	// operator.
	//
	// For information about the supported syntax, see [Search query reference for Resource Explorer] in the Amazon Web Services
	// Resource Explorer User Guide.
	//
	// This query string in the context of this operation supports only [filter prefixes] with optional [operators]
	// . It doesn't support free-form text. For example, the string region:us*
	// service:ec2 -tag:stage=prod includes all Amazon EC2 resources in any Amazon Web
	// Services Region that begins with the letters us and is not tagged with a key
	// Stage that has the value prod .
	//
	// [filter prefixes]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-filters
	// [Search query reference for Resource Explorer]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html
	// [operators]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-operators
	Filters *types.SearchFilter

	// Specifies optional fields that you want included in search results from this
	// view. It is a list of objects that each describe a field to include.
	//
	// The default is an empty list, with no optional fields included in the results.
	IncludedProperties []types.IncludedProperty

	// The root ARN of the account, an organizational unit (OU), or an organization
	// ARN. If left empty, the default is account.
	Scope *string

	// Tag key and value pairs that are attached to the view.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateViewOutput struct {

	// A structure that contains the details about the new view.
	View *types.View

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateViewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateView{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateView{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateView"); err != nil {
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
	if err = addIdempotencyToken_opCreateViewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateViewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateView(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateView struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateView) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateView) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateViewInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateViewInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateViewMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateView{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateView(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateView",
	}
}
