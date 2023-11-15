// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a model package that you can use to create SageMaker models or list on
// Amazon Web Services Marketplace, or a versioned model that is part of a model
// group. Buyers can subscribe to model packages listed on Amazon Web Services
// Marketplace to create models in SageMaker. To create a model package by
// specifying a Docker container that contains your inference code and the Amazon
// S3 location of your model artifacts, provide values for InferenceSpecification .
// To create a model from an algorithm resource that you created or subscribed to
// in Amazon Web Services Marketplace, provide a value for
// SourceAlgorithmSpecification . There are two types of model packages:
//   - Versioned - a model that is part of a model group in the model registry.
//   - Unversioned - a model package that is not part of a model group.
func (c *Client) CreateModelPackage(ctx context.Context, params *CreateModelPackageInput, optFns ...func(*Options)) (*CreateModelPackageOutput, error) {
	if params == nil {
		params = &CreateModelPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModelPackage", params, optFns, c.addOperationCreateModelPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelPackageInput struct {

	// An array of additional Inference Specification objects. Each additional
	// Inference Specification specifies artifacts based on this model package that can
	// be used on inference endpoints. Generally used with SageMaker Neo to store the
	// compiled artifacts.
	AdditionalInferenceSpecifications []types.AdditionalInferenceSpecificationDefinition

	// Whether to certify the model package for listing on Amazon Web Services
	// Marketplace. This parameter is optional for unversioned models, and does not
	// apply to versioned models.
	CertifyForMarketplace *bool

	// A unique token that guarantees that the call to this API is idempotent.
	ClientToken *string

	// The metadata properties associated with the model package versions.
	CustomerMetadataProperties map[string]string

	// The machine learning domain of your model package and its components. Common
	// machine learning domains include computer vision and natural language
	// processing.
	Domain *string

	// Represents the drift check baselines that can be used when the model monitor is
	// set using the model package. For more information, see the topic on Drift
	// Detection against Previous Baselines in SageMaker Pipelines (https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines-quality-clarify-baseline-lifecycle.html#pipelines-quality-clarify-baseline-drift-detection)
	// in the Amazon SageMaker Developer Guide.
	DriftCheckBaselines *types.DriftCheckBaselines

	// Specifies details about inference jobs that can be run with models based on
	// this model package, including the following:
	//   - The Amazon ECR paths of containers that contain the inference code and
	//   model artifacts.
	//   - The instance types that the model package supports for transform jobs and
	//   real-time endpoints used for inference.
	//   - The input and output content formats that the model package supports for
	//   inference.
	InferenceSpecification *types.InferenceSpecification

	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties *types.MetadataProperties

	// Whether the model is approved for deployment. This parameter is optional for
	// versioned models, and does not apply to unversioned models. For versioned
	// models, the value of this parameter must be set to Approved to deploy the model.
	ModelApprovalStatus types.ModelApprovalStatus

	// A structure that contains model metrics reports.
	ModelMetrics *types.ModelMetrics

	// A description of the model package.
	ModelPackageDescription *string

	// The name or Amazon Resource Name (ARN) of the model package group that this
	// model version belongs to. This parameter is required for versioned models, and
	// does not apply to unversioned models.
	ModelPackageGroupName *string

	// The name of the model package. The name must have 1 to 63 characters. Valid
	// characters are a-z, A-Z, 0-9, and - (hyphen). This parameter is required for
	// unversioned models. It is not applicable to versioned models.
	ModelPackageName *string

	// The Amazon Simple Storage Service (Amazon S3) path where the sample payload is
	// stored. This path must point to a single gzip compressed tar archive (.tar.gz
	// suffix). This archive can hold multiple files that are all equally used in the
	// load test. Each file in the archive must satisfy the size constraints of the
	// InvokeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html#API_runtime_InvokeEndpoint_RequestSyntax)
	// call.
	SamplePayloadUrl *string

	// Indicates if you want to skip model validation.
	SkipModelValidation types.SkipModelValidation

	// Details about the algorithm that was used to create the model package.
	SourceAlgorithmSpecification *types.SourceAlgorithmSpecification

	// A list of key value pairs associated with the model. For more information, see
	// Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the Amazon Web Services General Reference Guide. If you supply
	// ModelPackageGroupName , your model package belongs to the model group you
	// specify and uses the tags associated with the model group. In this case, you
	// cannot supply a tag argument.
	Tags []types.Tag

	// The machine learning task your model package accomplishes. Common machine
	// learning tasks include object detection and image classification. The following
	// tasks are supported by Inference Recommender: "IMAGE_CLASSIFICATION" |
	// "OBJECT_DETECTION" | "TEXT_GENERATION" | "IMAGE_SEGMENTATION" | "FILL_MASK" |
	// "CLASSIFICATION" | "REGRESSION" | "OTHER" . Specify "OTHER" if none of the tasks
	// listed fit your use case.
	Task *string

	// Specifies configurations for one or more transform jobs that SageMaker runs to
	// test the model package.
	ValidationSpecification *types.ModelPackageValidationSpecification

	noSmithyDocumentSerde
}

type CreateModelPackageOutput struct {

	// The Amazon Resource Name (ARN) of the new model package.
	//
	// This member is required.
	ModelPackageArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateModelPackage"); err != nil {
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
	if err = addIdempotencyToken_opCreateModelPackageMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateModelPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModelPackage(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateModelPackage struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateModelPackage) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateModelPackage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateModelPackageInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateModelPackageInput ")
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
func addIdempotencyToken_opCreateModelPackageMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateModelPackage{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateModelPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateModelPackage",
	}
}
