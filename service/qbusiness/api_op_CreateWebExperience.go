// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Q Business web experience.
func (c *Client) CreateWebExperience(ctx context.Context, params *CreateWebExperienceInput, optFns ...func(*Options)) (*CreateWebExperienceOutput, error) {
	if params == nil {
		params = &CreateWebExperienceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWebExperience", params, optFns, c.addOperationCreateWebExperienceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWebExperienceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWebExperienceInput struct {

	// The identifier of the Amazon Q Business web experience.
	//
	// This member is required.
	ApplicationId *string

	// A token you provide to identify a request to create an Amazon Q Business web
	// experience.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the service role attached to your web
	// experience.
	RoleArn *string

	// Determines whether sample prompts are enabled in the web experience for an end
	// user.
	SamplePromptsControlMode types.WebExperienceSamplePromptsControlMode

	// A subtitle to personalize your Amazon Q Business web experience.
	Subtitle *string

	// A list of key-value pairs that identify or categorize your Amazon Q Business
	// web experience. You can also use tags to help control access to the web
	// experience. Tag keys and values can consist of Unicode letters, digits, white
	// space, and any of the following symbols: _ . : / = + - @.
	Tags []types.Tag

	// The title for your Amazon Q Business web experience.
	Title *string

	// The customized welcome message for end users of an Amazon Q Business web
	// experience.
	WelcomeMessage *string

	noSmithyDocumentSerde
}

type CreateWebExperienceOutput struct {

	//  The Amazon Resource Name (ARN) of an Amazon Q Business web experience.
	WebExperienceArn *string

	// The identifier of the Amazon Q Business web experience.
	WebExperienceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWebExperienceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWebExperience{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWebExperience{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWebExperience"); err != nil {
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
	if err = addIdempotencyToken_opCreateWebExperienceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWebExperienceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWebExperience(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWebExperience struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWebExperience) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWebExperience) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateWebExperienceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWebExperienceInput ")
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
func addIdempotencyToken_opCreateWebExperienceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWebExperience{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWebExperience(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWebExperience",
	}
}
