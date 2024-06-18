// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a version of the SageMaker image specified by ImageName . The version
// represents the Amazon ECR container image specified by BaseImage .
func (c *Client) CreateImageVersion(ctx context.Context, params *CreateImageVersionInput, optFns ...func(*Options)) (*CreateImageVersionOutput, error) {
	if params == nil {
		params = &CreateImageVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateImageVersion", params, optFns, c.addOperationCreateImageVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateImageVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateImageVersionInput struct {

	// The registry path of the container image to use as the starting point for this
	// version. The path is an Amazon ECR URI in the following format:
	//
	//     .dkr.ecr..amazonaws.com/
	//
	// This member is required.
	BaseImage *string

	// A unique ID. If not specified, the Amazon Web Services CLI and Amazon Web
	// Services SDKs, such as the SDK for Python (Boto3), add a unique value to the
	// call.
	//
	// This member is required.
	ClientToken *string

	// The ImageName of the Image to create a version of.
	//
	// This member is required.
	ImageName *string

	// A list of aliases created with the image version.
	Aliases []string

	// Indicates Horovod compatibility.
	Horovod *bool

	// Indicates SageMaker job type compatibility.
	//
	//   - TRAINING : The image version is compatible with SageMaker training jobs.
	//
	//   - INFERENCE : The image version is compatible with SageMaker inference jobs.
	//
	//   - NOTEBOOK_KERNEL : The image version is compatible with SageMaker notebook
	//   kernels.
	JobType types.JobType

	// The machine learning framework vended in the image version.
	MLFramework *string

	// Indicates CPU or GPU compatibility.
	//
	//   - CPU : The image version is compatible with CPU.
	//
	//   - GPU : The image version is compatible with GPU.
	Processor types.Processor

	// The supported programming language and its version.
	ProgrammingLang *string

	// The maintainer description of the image version.
	ReleaseNotes *string

	// The stability of the image version, specified by the maintainer.
	//
	//   - NOT_PROVIDED : The maintainers did not provide a status for image version
	//   stability.
	//
	//   - STABLE : The image version is stable.
	//
	//   - TO_BE_ARCHIVED : The image version is set to be archived. Custom image
	//   versions that are set to be archived are automatically archived after three
	//   months.
	//
	//   - ARCHIVED : The image version is archived. Archived image versions are not
	//   searchable and are no longer actively supported.
	VendorGuidance types.VendorGuidance

	noSmithyDocumentSerde
}

type CreateImageVersionOutput struct {

	// The ARN of the image version.
	ImageVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateImageVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateImageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateImageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateImageVersion"); err != nil {
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
	if err = addIdempotencyToken_opCreateImageVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateImageVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateImageVersion(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateImageVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateImageVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateImageVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateImageVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateImageVersionInput ")
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
func addIdempotencyToken_opCreateImageVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateImageVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateImageVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateImageVersion",
	}
}
