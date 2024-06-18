// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a framework with one or more controls. A framework is a collection of
// controls that you can use to evaluate your backup practices. By using pre-built
// customizable controls to define your policies, you can evaluate whether your
// backup practices comply with your policies and which resources are not yet in
// compliance.
func (c *Client) CreateFramework(ctx context.Context, params *CreateFrameworkInput, optFns ...func(*Options)) (*CreateFrameworkOutput, error) {
	if params == nil {
		params = &CreateFrameworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFramework", params, optFns, c.addOperationCreateFrameworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFrameworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFrameworkInput struct {

	// A list of the controls that make up the framework. Each control in the list has
	// a name, input parameters, and scope.
	//
	// This member is required.
	FrameworkControls []types.FrameworkControl

	// The unique name of the framework. The name must be between 1 and 256
	// characters, starting with a letter, and consisting of letters (a-z, A-Z),
	// numbers (0-9), and underscores (_).
	//
	// This member is required.
	FrameworkName *string

	// An optional description of the framework with a maximum of 1,024 characters.
	FrameworkDescription *string

	// Metadata that you can assign to help organize the frameworks that you create.
	// Each tag is a key-value pair.
	FrameworkTags map[string]string

	// A customer-chosen string that you can use to distinguish between otherwise
	// identical calls to CreateFrameworkInput . Retrying a successful request with the
	// same idempotency token results in a success message with no action taken.
	IdempotencyToken *string

	noSmithyDocumentSerde
}

type CreateFrameworkOutput struct {

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format
	// of the ARN depends on the resource type.
	FrameworkArn *string

	// The unique name of the framework. The name must be between 1 and 256
	// characters, starting with a letter, and consisting of letters (a-z, A-Z),
	// numbers (0-9), and underscores (_).
	FrameworkName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFrameworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFramework{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFramework{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFramework"); err != nil {
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
	if err = addIdempotencyToken_opCreateFrameworkMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFrameworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFramework(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFramework struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFramework) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFramework) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFrameworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFrameworkInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFrameworkMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFramework{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFramework(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFramework",
	}
}
