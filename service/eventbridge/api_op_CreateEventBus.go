// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new event bus within your account. This can be a custom event bus
// which you can use to receive events from your custom applications and services,
// or it can be a partner event bus which can be matched to a partner event source.
func (c *Client) CreateEventBus(ctx context.Context, params *CreateEventBusInput, optFns ...func(*Options)) (*CreateEventBusOutput, error) {
	if params == nil {
		params = &CreateEventBusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventBus", params, optFns, c.addOperationCreateEventBusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventBusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventBusInput struct {

	// The name of the new event bus.
	//
	// Custom event bus names can't contain the / character, but you can use the /
	// character in partner event bus names. In addition, for partner event buses, the
	// name must exactly match the name of the partner event source that this event bus
	// is matched to.
	//
	// You can't use the name default for a custom event bus, as this name is already
	// used for your account's default event bus.
	//
	// This member is required.
	Name *string

	// Configuration details of the Amazon SQS queue for EventBridge to use as a
	// dead-letter queue (DLQ).
	//
	// For more information, see Event retry policy and using dead-letter queues in the EventBridge User Guide.
	DeadLetterConfig *types.DeadLetterConfig

	// The event bus description.
	Description *string

	// If you are creating a partner event bus, this specifies the partner event
	// source that the new event bus will be matched with.
	EventSourceName *string

	// The identifier of the KMS customer managed key for EventBridge to use, if you
	// choose to use a customer managed key to encrypt events on this event bus. The
	// identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key
	// alias ARN.
	//
	// If you do not specify a customer managed key identifier, EventBridge uses an
	// Amazon Web Services owned key to encrypt events on the event bus.
	//
	// For more information, see [Managing keys] in the Key Management Service Developer Guide.
	//
	// Archives and schema discovery are not supported for event buses encrypted using
	// a customer managed key. EventBridge returns an error if:
	//
	//   - You call [CreateArchive]on an event bus set to use a customer managed key for encryption.
	//
	//   - You call [CreateDiscoverer]on an event bus set to use a customer managed key for encryption.
	//
	//   - You call [UpdatedEventBus]to set a customer managed key on an event bus with an archives or
	//   schema discovery enabled.
	//
	// To enable archives or schema discovery on an event bus, choose to use an Amazon
	// Web Services owned key. For more information, see [Data encryption in EventBridge]in the Amazon EventBridge
	// User Guide.
	//
	// [UpdatedEventBus]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_UpdatedEventBus.html
	// [Data encryption in EventBridge]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-encryption.html
	// [Managing keys]: https://docs.aws.amazon.com/kms/latest/developerguide/getting-started.html
	// [CreateArchive]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateArchive.html
	// [CreateDiscoverer]: https://docs.aws.amazon.com/eventbridge/latest/schema-reference/v1-discoverers.html#CreateDiscoverer
	KmsKeyIdentifier *string

	// Tags to associate with the event bus.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEventBusOutput struct {

	// Configuration details of the Amazon SQS queue for EventBridge to use as a
	// dead-letter queue (DLQ).
	//
	// For more information, see Event retry policy and using dead-letter queues in the EventBridge User Guide.
	DeadLetterConfig *types.DeadLetterConfig

	// The event bus description.
	Description *string

	// The ARN of the new event bus.
	EventBusArn *string

	// The identifier of the KMS customer managed key for EventBridge to use to
	// encrypt events on this event bus, if one has been specified.
	//
	// For more information, see [Data encryption in EventBridge] in the Amazon EventBridge User Guide.
	//
	// [Data encryption in EventBridge]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-encryption.html
	KmsKeyIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventBusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEventBus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEventBus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEventBus"); err != nil {
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
	if err = addOpCreateEventBusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventBus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEventBus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEventBus",
	}
}
