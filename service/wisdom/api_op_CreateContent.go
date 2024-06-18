// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates Wisdom content. Before to calling this API, use [StartContentUpload] to upload an asset.
//
// [StartContentUpload]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_StartContentUpload.html
func (c *Client) CreateContent(ctx context.Context, params *CreateContentInput, optFns ...func(*Options)) (*CreateContentOutput, error) {
	if params == nil {
		params = &CreateContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContent", params, optFns, c.addOperationCreateContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContentInput struct {

	// The identifier of the knowledge base. This should not be a QUICK_RESPONSES type
	// knowledge base if you're storing Wisdom Content resource to it. Can be either
	// the ID or the ARN. URLs cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The name of the content. Each piece of content in a knowledge base must have a
	// unique name. You can retrieve a piece of content using only its knowledge base
	// and its name with the [SearchContent]API.
	//
	// [SearchContent]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_SearchContent.html
	//
	// This member is required.
	Name *string

	// A pointer to the uploaded asset. This value is returned by [StartContentUpload].
	//
	// [StartContentUpload]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_StartContentUpload.html
	//
	// This member is required.
	UploadId *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see [Making retries safe with idempotent APIs].
	//
	// [Making retries safe with idempotent APIs]: https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/
	ClientToken *string

	// A key/value map to store attributes without affecting tagging or
	// recommendations. For example, when synchronizing data between an external system
	// and Wisdom, you can store an external version identifier as metadata to utilize
	// for determining drift.
	Metadata map[string]string

	// The URI you want to use for the article. If the knowledge base has a
	// templateUri, setting this argument overrides it for this piece of content.
	OverrideLinkOutUri *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// The title of the content. If not set, the title is equal to the name.
	Title *string

	noSmithyDocumentSerde
}

type CreateContentOutput struct {

	// The content.
	Content *types.ContentData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateContent"); err != nil {
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
	if err = addIdempotencyToken_opCreateContentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateContent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateContent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateContent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateContentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateContentInput ")
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
func addIdempotencyToken_opCreateContentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateContent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateContent",
	}
}
