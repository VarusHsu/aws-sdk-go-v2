// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a speaker search task.
//
// Before starting any speaker search tasks, you must provide all notices and
// obtain all consents from the speaker as required under applicable privacy and
// biometrics laws, and as required under the [AWS service terms]for the Amazon Chime SDK.
//
// [AWS service terms]: https://aws.amazon.com/service-terms/
func (c *Client) StartSpeakerSearchTask(ctx context.Context, params *StartSpeakerSearchTaskInput, optFns ...func(*Options)) (*StartSpeakerSearchTaskOutput, error) {
	if params == nil {
		params = &StartSpeakerSearchTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSpeakerSearchTask", params, optFns, c.addOperationStartSpeakerSearchTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSpeakerSearchTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSpeakerSearchTaskInput struct {

	// The unique identifier of the resource to be updated. Valid values include the
	// ID and ARN of the media insights pipeline.
	//
	// This member is required.
	Identifier *string

	// The ARN of the voice profile domain that will store the voice profile.
	//
	// This member is required.
	VoiceProfileDomainArn *string

	// The unique identifier for the client request. Use a different token for
	// different speaker search tasks.
	ClientRequestToken *string

	// The task configuration for the Kinesis video stream source of the media
	// insights pipeline.
	KinesisVideoStreamSourceTaskConfiguration *types.KinesisVideoStreamSourceTaskConfiguration

	noSmithyDocumentSerde
}

type StartSpeakerSearchTaskOutput struct {

	// The details of the speaker search task.
	SpeakerSearchTask *types.SpeakerSearchTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSpeakerSearchTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSpeakerSearchTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSpeakerSearchTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartSpeakerSearchTask"); err != nil {
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
	if err = addIdempotencyToken_opStartSpeakerSearchTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartSpeakerSearchTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSpeakerSearchTask(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartSpeakerSearchTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSpeakerSearchTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSpeakerSearchTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSpeakerSearchTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSpeakerSearchTaskInput ")
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
func addIdempotencyToken_opStartSpeakerSearchTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartSpeakerSearchTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSpeakerSearchTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartSpeakerSearchTask",
	}
}
