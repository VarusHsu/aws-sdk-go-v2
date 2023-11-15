// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a notification rule for a resource. The rule specifies the events you
// want notifications about and the targets (such as Chatbot topics or Chatbot
// clients configured for Slack) where you want to receive them.
func (c *Client) CreateNotificationRule(ctx context.Context, params *CreateNotificationRuleInput, optFns ...func(*Options)) (*CreateNotificationRuleOutput, error) {
	if params == nil {
		params = &CreateNotificationRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNotificationRule", params, optFns, c.addOperationCreateNotificationRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNotificationRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNotificationRuleInput struct {

	// The level of detail to include in the notifications for this resource. BASIC
	// will include only the contents of the event as it would appear in Amazon
	// CloudWatch. FULL will include any supplemental information provided by AWS
	// CodeStar Notifications and/or the service for the resource for which the
	// notification is created.
	//
	// This member is required.
	DetailType types.DetailType

	// A list of event types associated with this notification rule. For a list of
	// allowed events, see EventTypeSummary .
	//
	// This member is required.
	EventTypeIds []string

	// The name for the notification rule. Notification rule names must be unique in
	// your Amazon Web Services account.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the resource to associate with the
	// notification rule. Supported resources include pipelines in CodePipeline,
	// repositories in CodeCommit, and build projects in CodeBuild.
	//
	// This member is required.
	Resource *string

	// A list of Amazon Resource Names (ARNs) of Amazon Simple Notification Service
	// topics and Chatbot clients to associate with the notification rule.
	//
	// This member is required.
	Targets []types.Target

	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request
	// with the same parameters is received and a token is included, the request
	// returns information about the initial request that used that token. The Amazon
	// Web Services SDKs prepopulate client request tokens. If you are using an Amazon
	// Web Services SDK, an idempotency token is created for you.
	ClientRequestToken *string

	// The status of the notification rule. The default value is ENABLED . If the
	// status is set to DISABLED , notifications aren't sent for the notification rule.
	Status types.NotificationRuleStatus

	// A list of tags to apply to this notification rule. Key names cannot start with "
	// aws ".
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateNotificationRuleOutput struct {

	// The Amazon Resource Name (ARN) of the notification rule.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNotificationRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateNotificationRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNotificationRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNotificationRule"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addIdempotencyToken_opCreateNotificationRuleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNotificationRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNotificationRule(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

type idempotencyToken_initializeOpCreateNotificationRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNotificationRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNotificationRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNotificationRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNotificationRuleInput ")
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
func addIdempotencyToken_opCreateNotificationRuleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateNotificationRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNotificationRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNotificationRule",
	}
}
