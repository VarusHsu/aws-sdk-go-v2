// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a lens share.
//
// The owner of a lens can share it with other Amazon Web Services accounts,
// users, an organization, and organizational units (OUs) in the same Amazon Web
// Services Region. Lenses provided by Amazon Web Services (Amazon Web Services
// Official Content) cannot be shared.
//
// Shared access to a lens is not removed until the lens invitation is deleted.
//
// If you share a lens with an organization or OU, all accounts in the
// organization or OU are granted access to the lens.
//
// For more information, see [Sharing a custom lens] in the Well-Architected Tool User Guide.
//
// # Disclaimer
//
// By sharing your custom lenses with other Amazon Web Services accounts, you
// acknowledge that Amazon Web Services will make your custom lenses available to
// those other accounts. Those other accounts may continue to access and use your
// shared custom lenses even if you delete the custom lenses from your own Amazon
// Web Services account or terminate your Amazon Web Services account.
//
// [Sharing a custom lens]: https://docs.aws.amazon.com/wellarchitected/latest/userguide/lenses-sharing.html
func (c *Client) CreateLensShare(ctx context.Context, params *CreateLensShareInput, optFns ...func(*Options)) (*CreateLensShareOutput, error) {
	if params == nil {
		params = &CreateLensShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLensShare", params, optFns, c.addOperationCreateLensShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLensShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLensShareInput struct {

	// A unique case-sensitive string used to ensure that this request is idempotent
	// (executes only once).
	//
	// You should not reuse the same token for other requests. If you retry a request
	// with the same client request token and the same parameters after the original
	// request has completed successfully, the result of the original request is
	// returned.
	//
	// This token is listed as required, however, if you do not specify it, the Amazon
	// Web Services SDKs automatically generate one for you. If you are not using the
	// Amazon Web Services SDK or the CLI, you must provide this token or the request
	// will fail.
	//
	// This member is required.
	ClientRequestToken *string

	// The alias of the lens.
	//
	// For Amazon Web Services official lenses, this is either the lens alias, such as
	// serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses.
	//
	// For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// .
	//
	// Each lens is identified by its LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The Amazon Web Services account ID, organization ID, or organizational unit
	// (OU) ID with which the workload, lens, profile, or review template is shared.
	//
	// This member is required.
	SharedWith *string

	noSmithyDocumentSerde
}

type CreateLensShareOutput struct {

	// The ID associated with the share.
	ShareId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLensShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLensShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLensShare{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLensShare"); err != nil {
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
	if err = addIdempotencyToken_opCreateLensShareMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateLensShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLensShare(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateLensShare struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateLensShare) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateLensShare) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateLensShareInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateLensShareInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateLensShareMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateLensShare{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateLensShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLensShare",
	}
}
