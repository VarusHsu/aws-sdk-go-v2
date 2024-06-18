// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates worlds using the specified template.
func (c *Client) CreateWorldGenerationJob(ctx context.Context, params *CreateWorldGenerationJobInput, optFns ...func(*Options)) (*CreateWorldGenerationJobOutput, error) {
	if params == nil {
		params = &CreateWorldGenerationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorldGenerationJob", params, optFns, c.addOperationCreateWorldGenerationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorldGenerationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorldGenerationJobInput struct {

	// The Amazon Resource Name (arn) of the world template describing the worlds you
	// want to create.
	//
	// This member is required.
	Template *string

	// Information about the world count.
	//
	// This member is required.
	WorldCount *types.WorldCount

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// A map that contains tag keys and tag values that are attached to the world
	// generator job.
	Tags map[string]string

	// A map that contains tag keys and tag values that are attached to the generated
	// worlds.
	WorldTags map[string]string

	noSmithyDocumentSerde
}

type CreateWorldGenerationJobOutput struct {

	// The Amazon Resource Name (ARN) of the world generator job.
	Arn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The time, in milliseconds since the epoch, when the world generator job was
	// created.
	CreatedAt *time.Time

	// The failure code of the world generator job if it failed:
	//
	// InternalServiceError Internal service error.
	//
	// LimitExceeded The requested resource exceeds the maximum number allowed, or the
	// number of concurrent stream requests exceeds the maximum number allowed.
	//
	// ResourceNotFound The specified resource could not be found.
	//
	// RequestThrottled The request was throttled.
	//
	// InvalidInput An input parameter in the request is not valid.
	FailureCode types.WorldGenerationJobErrorCode

	// The status of the world generator job.
	//
	// Pending The world generator job request is pending.
	//
	// Running The world generator job is running.
	//
	// Completed The world generator job completed.
	//
	// Failed The world generator job failed. See failureCode for more information.
	//
	// PartialFailed Some worlds did not generate.
	//
	// Canceled The world generator job was cancelled.
	//
	// Canceling The world generator job is being cancelled.
	Status types.WorldGenerationJobStatus

	// A map that contains tag keys and tag values that are attached to the world
	// generator job.
	Tags map[string]string

	// The Amazon Resource Name (arn) of the world template.
	Template *string

	// Information about the world count.
	WorldCount *types.WorldCount

	// A map that contains tag keys and tag values that are attached to the generated
	// worlds.
	WorldTags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorldGenerationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorldGenerationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorldGenerationJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWorldGenerationJob"); err != nil {
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
	if err = addIdempotencyToken_opCreateWorldGenerationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWorldGenerationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorldGenerationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWorldGenerationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWorldGenerationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWorldGenerationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateWorldGenerationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWorldGenerationJobInput ")
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
func addIdempotencyToken_opCreateWorldGenerationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWorldGenerationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWorldGenerationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWorldGenerationJob",
	}
}
