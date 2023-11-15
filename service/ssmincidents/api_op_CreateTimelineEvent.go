// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmincidents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssmincidents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a custom timeline event on the incident details page of an incident
// record. Incident Manager automatically creates timeline events that mark key
// moments during an incident. You can create custom timeline events to mark
// important events that Incident Manager can detect automatically.
func (c *Client) CreateTimelineEvent(ctx context.Context, params *CreateTimelineEventInput, optFns ...func(*Options)) (*CreateTimelineEventOutput, error) {
	if params == nil {
		params = &CreateTimelineEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTimelineEvent", params, optFns, c.addOperationCreateTimelineEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTimelineEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTimelineEventInput struct {

	// A short description of the event.
	//
	// This member is required.
	EventData *string

	// The time that the event occurred.
	//
	// This member is required.
	EventTime *time.Time

	// The type of event. You can create timeline events of type Custom Event .
	//
	// This member is required.
	EventType *string

	// The Amazon Resource Name (ARN) of the incident record that the action adds the
	// incident to.
	//
	// This member is required.
	IncidentRecordArn *string

	// A token that ensures that a client calls the action only once with the
	// specified details.
	ClientToken *string

	// Adds one or more references to the TimelineEvent . A reference is an Amazon Web
	// Services resource involved or associated with the incident. To specify a
	// reference, enter its Amazon Resource Name (ARN). You can also specify a related
	// item associated with a resource. For example, to specify an Amazon DynamoDB
	// (DynamoDB) table as a resource, use the table's ARN. You can also specify an
	// Amazon CloudWatch metric associated with the DynamoDB table as a related item.
	EventReferences []types.EventReference

	noSmithyDocumentSerde
}

type CreateTimelineEventOutput struct {

	// The ID of the event for easy reference later.
	//
	// This member is required.
	EventId *string

	// The ARN of the incident record that you added the event to.
	//
	// This member is required.
	IncidentRecordArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTimelineEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTimelineEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTimelineEvent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTimelineEvent"); err != nil {
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
	if err = addIdempotencyToken_opCreateTimelineEventMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateTimelineEventValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTimelineEvent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateTimelineEvent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateTimelineEvent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateTimelineEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateTimelineEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateTimelineEventInput ")
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
func addIdempotencyToken_opCreateTimelineEventMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateTimelineEvent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateTimelineEvent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTimelineEvent",
	}
}
