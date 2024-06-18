// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Create a new registration attachment to use for uploading a file or a URL to a
// file. The maximum file size is 1MiB and valid file extensions are PDF, JPEG and
// PNG. For example, many sender ID registrations require a signed “letter of
// authorization” (LOA) to be submitted.
func (c *Client) CreateRegistrationAttachment(ctx context.Context, params *CreateRegistrationAttachmentInput, optFns ...func(*Options)) (*CreateRegistrationAttachmentOutput, error) {
	if params == nil {
		params = &CreateRegistrationAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRegistrationAttachment", params, optFns, c.addOperationCreateRegistrationAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRegistrationAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRegistrationAttachmentInput struct {

	// The registration file to upload. The maximum file size is 1MiB and valid file
	// extensions are PDF, JPEG and PNG.
	AttachmentBody []byte

	// A URL to the required registration file. For example, you can provide the S3
	// object URL.
	AttachmentUrl *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don't specify a client token, a randomly generated token is
	// used for the request to ensure idempotency.
	ClientToken *string

	// An array of tags (key and value pairs) to associate with the registration
	// attachment.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateRegistrationAttachmentOutput struct {

	// The status of the registration attachment.
	//
	//   - UPLOAD_IN_PROGRESS The attachment is being uploaded.
	//
	//   - UPLOAD_COMPLETE The attachment has been uploaded.
	//
	//   - UPLOAD_FAILED The attachment failed to uploaded.
	//
	//   - DELETED The attachment has been deleted..
	//
	// This member is required.
	AttachmentStatus types.AttachmentStatus

	// The time when the registration attachment was created, in [UNIX epoch time] format.
	//
	// [UNIX epoch time]: https://www.epochconverter.com/
	//
	// This member is required.
	CreatedTimestamp *time.Time

	// The Amazon Resource Name (ARN) for the registration attachment.
	//
	// This member is required.
	RegistrationAttachmentArn *string

	// The unique identifier for the registration attachment.
	//
	// This member is required.
	RegistrationAttachmentId *string

	// An array of tags (key and value pairs) to associate with the registration
	// attachment.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRegistrationAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateRegistrationAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateRegistrationAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRegistrationAttachment"); err != nil {
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
	if err = addIdempotencyToken_opCreateRegistrationAttachmentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRegistrationAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRegistrationAttachment(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateRegistrationAttachment struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateRegistrationAttachment) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateRegistrationAttachment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateRegistrationAttachmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateRegistrationAttachmentInput ")
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
func addIdempotencyToken_opCreateRegistrationAttachmentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateRegistrationAttachment{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateRegistrationAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRegistrationAttachment",
	}
}
