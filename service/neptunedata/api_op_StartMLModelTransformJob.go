// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new model transform job. See [Use a trained model to generate new model artifacts].
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:StartMLModelTransformJob]IAM action in that cluster.
//
// [Use a trained model to generate new model artifacts]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-model-transform.html
// [neptune-db:StartMLModelTransformJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startmlmodeltransformjob
func (c *Client) StartMLModelTransformJob(ctx context.Context, params *StartMLModelTransformJobInput, optFns ...func(*Options)) (*StartMLModelTransformJobOutput, error) {
	if params == nil {
		params = &StartMLModelTransformJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMLModelTransformJob", params, optFns, c.addOperationStartMLModelTransformJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMLModelTransformJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMLModelTransformJobInput struct {

	// The location in Amazon S3 where the model artifacts are to be stored.
	//
	// This member is required.
	ModelTransformOutputS3Location *string

	// The type of ML instance used in preparing and managing training of ML models.
	// This is an ML compute instance chosen based on memory requirements for
	// processing the training data and model.
	BaseProcessingInstanceType *string

	// The disk volume size of the training instance in gigabytes. The default is 0.
	// Both input data and the output model are stored on disk, so the volume size must
	// be large enough to hold both data sets. If not specified or 0, Neptune ML
	// selects a disk volume size based on the recommendation generated in the data
	// processing step.
	BaseProcessingInstanceVolumeSizeInGB *int32

	// Configuration information for a model transform using a custom model. The
	// customModelTransformParameters object contains the following fields, which must
	// have values compatible with the saved model parameters from the training job:
	CustomModelTransformParameters *types.CustomModelTransformParameters

	// The job ID of a completed data-processing job. You must include either
	// dataProcessingJobId and a mlModelTrainingJobId , or a trainingJobName .
	DataProcessingJobId *string

	// A unique identifier for the new job. The default is an autogenerated UUID.
	Id *string

	// The job ID of a completed model-training job. You must include either
	// dataProcessingJobId and a mlModelTrainingJobId , or a trainingJobName .
	MlModelTrainingJobId *string

	// The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3
	// resources. This must be listed in your DB cluster parameter group or an error
	// will occur.
	NeptuneIamRoleArn *string

	// The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt the
	// output of the processing job. The default is none.
	S3OutputEncryptionKMSKey *string

	// The ARN of an IAM role for SageMaker execution. This must be listed in your DB
	// cluster parameter group or an error will occur.
	SagemakerIamRoleArn *string

	// The VPC security group IDs. The default is None.
	SecurityGroupIds []string

	// The IDs of the subnets in the Neptune VPC. The default is None.
	Subnets []string

	// The name of a completed SageMaker training job. You must include either
	// dataProcessingJobId and a mlModelTrainingJobId , or a trainingJobName .
	TrainingJobName *string

	// The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt data
	// on the storage volume attached to the ML compute instances that run the training
	// job. The default is None.
	VolumeEncryptionKMSKey *string

	noSmithyDocumentSerde
}

type StartMLModelTransformJobOutput struct {

	// The ARN of the model transform job.
	Arn *string

	// The creation time of the model transform job, in milliseconds.
	CreationTimeInMillis *int64

	// The unique ID of the new model transform job.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMLModelTransformJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMLModelTransformJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMLModelTransformJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMLModelTransformJob"); err != nil {
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
	if err = addOpStartMLModelTransformJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMLModelTransformJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMLModelTransformJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMLModelTransformJob",
	}
}
