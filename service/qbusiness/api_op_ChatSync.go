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

// Starts or continues a non-streaming Amazon Q Business conversation.
func (c *Client) ChatSync(ctx context.Context, params *ChatSyncInput, optFns ...func(*Options)) (*ChatSyncOutput, error) {
	if params == nil {
		params = &ChatSyncInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChatSync", params, optFns, c.addOperationChatSyncMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChatSyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ChatSyncInput struct {

	// The identifier of the Amazon Q Business application linked to the Amazon Q
	// Business conversation.
	//
	// This member is required.
	ApplicationId *string

	// A request from an end user to perform an Amazon Q Business plugin action.
	ActionExecution *types.ActionExecution

	// A list of files uploaded directly during chat. You can upload a maximum of 5
	// files of upto 10 MB each.
	Attachments []types.AttachmentInput

	// Enables filtering of Amazon Q Business web experience responses based on
	// document attributes or metadata fields.
	AttributeFilter *types.AttributeFilter

	// The chat modes available in an Amazon Q Business web experience.
	//   - RETRIEVAL_MODE - The default chat mode for an Amazon Q Business application.
	//   When this mode is enabled, Amazon Q Business generates responses only from data
	//   sources connected to an Amazon Q Business application.
	//   - CREATOR_MODE - By selecting this mode, users can choose to generate
	//   responses only from the LLM knowledge, without consulting connected data
	//   sources, for a chat request.
	//   - PLUGIN_MODE - By selecting this mode, users can choose to use plugins in
	//   chat.
	// For more information, see Admin controls and guardrails (https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/guardrails.html)
	// , Plugins (https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/plugins.html)
	// , and Conversation settings (https://docs.aws.amazon.com/amazonq/latest/business-use-dg/using-web-experience.html#chat-source-scope)
	// .
	ChatMode types.ChatMode

	// The chat mode configuration for an Amazon Q Business application.
	ChatModeConfiguration types.ChatModeConfiguration

	// A token that you provide to identify a chat request.
	ClientToken *string

	// The identifier of the Amazon Q Business conversation.
	ConversationId *string

	// The identifier of the previous end user text input message in a conversation.
	ParentMessageId *string

	// The groups that a user associated with the chat input belongs to.
	UserGroups []string

	// The identifier of the user attached to the chat input.
	UserId *string

	// A end user message in a conversation.
	UserMessage *string

	noSmithyDocumentSerde
}

type ChatSyncOutput struct {

	// A request from Amazon Q Business to the end user for information Amazon Q
	// Business needs to successfully complete a requested plugin action.
	ActionReview *types.ActionReview

	// The identifier of the Amazon Q Business conversation.
	ConversationId *string

	// A list of files which failed to upload during chat.
	FailedAttachments []types.AttachmentOutput

	// The source documents used to generate the conversation response.
	SourceAttributions []*types.SourceAttribution

	// An AI-generated message in a conversation.
	SystemMessage *string

	// The identifier of an Amazon Q Business AI generated message within the
	// conversation.
	SystemMessageId *string

	// The identifier of an Amazon Q Business end user text input message within the
	// conversation.
	UserMessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationChatSyncMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpChatSync{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpChatSync{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ChatSync"); err != nil {
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
	if err = addIdempotencyToken_opChatSyncMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpChatSyncValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChatSync(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpChatSync struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpChatSync) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpChatSync) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ChatSyncInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ChatSyncInput ")
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
func addIdempotencyToken_opChatSyncMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpChatSync{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opChatSync(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ChatSync",
	}
}
