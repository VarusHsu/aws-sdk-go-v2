// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Q in Connect quick response.
func (c *Client) CreateQuickResponse(ctx context.Context, params *CreateQuickResponseInput, optFns ...func(*Options)) (*CreateQuickResponseOutput, error) {
	if params == nil {
		params = &CreateQuickResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateQuickResponse", params, optFns, c.addOperationCreateQuickResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateQuickResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateQuickResponseInput struct {

	// The content of the quick response.
	//
	// This member is required.
	Content types.QuickResponseDataProvider

	// The identifier of the knowledge base. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The name of the quick response.
	//
	// This member is required.
	Name *string

	// The Amazon Connect channels this quick response applies to.
	Channels []string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see [Making retries safe with idempotent APIs].
	//
	// [Making retries safe with idempotent APIs]: https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/
	ClientToken *string

	// The media type of the quick response content.
	//
	//   - Use application/x.quickresponse;format=plain for a quick response written in
	//   plain text.
	//
	//   - Use application/x.quickresponse;format=markdown for a quick response written
	//   in richtext.
	ContentType *string

	// The description of the quick response.
	Description *string

	// The configuration information of the user groups that the quick response is
	// accessible to.
	GroupingConfiguration *types.GroupingConfiguration

	// Whether the quick response is active.
	IsActive *bool

	// The language code value for the language in which the quick response is
	// written. The supported language codes include de_DE , en_US , es_ES , fr_FR ,
	// id_ID , it_IT , ja_JP , ko_KR , pt_BR , zh_CN , zh_TW
	Language *string

	// The shortcut key of the quick response. The value should be unique across the
	// knowledge base.
	ShortcutKey *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateQuickResponseOutput struct {

	// The quick response.
	QuickResponse *types.QuickResponseData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateQuickResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateQuickResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateQuickResponse{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateQuickResponse"); err != nil {
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
	if err = addIdempotencyToken_opCreateQuickResponseMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateQuickResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateQuickResponse(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateQuickResponse struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateQuickResponse) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateQuickResponse) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateQuickResponseInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateQuickResponseInput ")
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
func addIdempotencyToken_opCreateQuickResponseMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateQuickResponse{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateQuickResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateQuickResponse",
	}
}
