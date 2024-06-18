// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates a new media analysis job. Accepts a manifest file in an Amazon S3
// bucket. The output is a manifest file and a summary of the manifest stored in
// the Amazon S3 bucket.
func (c *Client) StartMediaAnalysisJob(ctx context.Context, params *StartMediaAnalysisJobInput, optFns ...func(*Options)) (*StartMediaAnalysisJobOutput, error) {
	if params == nil {
		params = &StartMediaAnalysisJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMediaAnalysisJob", params, optFns, c.addOperationStartMediaAnalysisJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMediaAnalysisJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMediaAnalysisJobInput struct {

	// Input data to be analyzed by the job.
	//
	// This member is required.
	Input *types.MediaAnalysisInput

	// Configuration options for the media analysis job to be created.
	//
	// This member is required.
	OperationsConfig *types.MediaAnalysisOperationsConfig

	// The Amazon S3 bucket location to store the results.
	//
	// This member is required.
	OutputConfig *types.MediaAnalysisOutputConfig

	// Idempotency token used to prevent the accidental creation of duplicate
	// versions. If you use the same token with multiple StartMediaAnalysisJobRequest
	// requests, the same response is returned. Use ClientRequestToken to prevent the
	// same request from being processed more than once.
	ClientRequestToken *string

	// The name of the job. Does not have to be unique.
	JobName *string

	// The identifier of customer managed AWS KMS key (name or ARN). The key is used
	// to encrypt images copied into the service. The key is also used to encrypt
	// results and manifest files written to the output Amazon S3 bucket.
	KmsKeyId *string

	noSmithyDocumentSerde
}

type StartMediaAnalysisJobOutput struct {

	// Identifier for the created job.
	//
	// This member is required.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMediaAnalysisJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMediaAnalysisJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMediaAnalysisJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMediaAnalysisJob"); err != nil {
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
	if err = addIdempotencyToken_opStartMediaAnalysisJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartMediaAnalysisJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMediaAnalysisJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartMediaAnalysisJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartMediaAnalysisJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartMediaAnalysisJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartMediaAnalysisJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartMediaAnalysisJobInput ")
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
func addIdempotencyToken_opStartMediaAnalysisJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartMediaAnalysisJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartMediaAnalysisJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMediaAnalysisJob",
	}
}
