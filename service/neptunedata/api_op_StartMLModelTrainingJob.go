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

// Creates a new Neptune ML model training job. See [Model training using the modeltraining command]modeltraining .
//
// When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the [neptune-db:StartMLModelTrainingJob]IAM action in that cluster.
//
// [Model training using the modeltraining command]: https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-modeltraining.html
// [neptune-db:StartMLModelTrainingJob]: https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#startmlmodeltrainingjob
func (c *Client) StartMLModelTrainingJob(ctx context.Context, params *StartMLModelTrainingJobInput, optFns ...func(*Options)) (*StartMLModelTrainingJobOutput, error) {
	if params == nil {
		params = &StartMLModelTrainingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMLModelTrainingJob", params, optFns, c.addOperationStartMLModelTrainingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMLModelTrainingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMLModelTrainingJobInput struct {

	// The job ID of the completed data-processing job that has created the data that
	// the training will work with.
	//
	// This member is required.
	DataProcessingJobId *string

	// The location in Amazon S3 where the model artifacts are to be stored.
	//
	// This member is required.
	TrainModelS3Location *string

	// The type of ML instance used in preparing and managing training of ML models.
	// This is a CPU instance chosen based on memory requirements for processing the
	// training data and model.
	BaseProcessingInstanceType *string

	// The configuration for custom model training. This is a JSON object.
	CustomModelTrainingParameters *types.CustomModelTrainingParameters

	// Optimizes the cost of training machine-learning models by using Amazon Elastic
	// Compute Cloud spot instances. The default is False .
	EnableManagedSpotTraining *bool

	// A unique identifier for the new job. The default is An autogenerated UUID.
	Id *string

	// Maximum total number of training jobs to start for the hyperparameter tuning
	// job. The default is 2. Neptune ML automatically tunes the hyperparameters of the
	// machine learning model. To obtain a model that performs well, use at least 10
	// jobs (in other words, set maxHPONumberOfTrainingJobs to 10). In general, the
	// more tuning runs, the better the results.
	MaxHPONumberOfTrainingJobs *int32

	// Maximum number of parallel training jobs to start for the hyperparameter tuning
	// job. The default is 2. The number of parallel jobs you can run is limited by the
	// available resources on your training instance.
	MaxHPOParallelTrainingJobs *int32

	// The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3
	// resources. This must be listed in your DB cluster parameter group or an error
	// will occur.
	NeptuneIamRoleArn *string

	// The job ID of a completed model-training job that you want to update
	// incrementally based on updated data.
	PreviousModelTrainingJobId *string

	// The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt the
	// output of the processing job. The default is none.
	S3OutputEncryptionKMSKey *string

	// The ARN of an IAM role for SageMaker execution.This must be listed in your DB
	// cluster parameter group or an error will occur.
	SagemakerIamRoleArn *string

	// The VPC security group IDs. The default is None.
	SecurityGroupIds []string

	// The IDs of the subnets in the Neptune VPC. The default is None.
	Subnets []string

	// The type of ML instance used for model training. All Neptune ML models support
	// CPU, GPU, and multiGPU training. The default is ml.p3.2xlarge . Choosing the
	// right instance type for training depends on the task type, graph size, and your
	// budget.
	TrainingInstanceType *string

	// The disk volume size of the training instance. Both input data and the output
	// model are stored on disk, so the volume size must be large enough to hold both
	// data sets. The default is 0. If not specified or 0, Neptune ML selects a disk
	// volume size based on the recommendation generated in the data processing step.
	TrainingInstanceVolumeSizeInGB *int32

	// Timeout in seconds for the training job. The default is 86,400 (1 day).
	TrainingTimeOutInSeconds *int32

	// The Amazon Key Management Service (KMS) key that SageMaker uses to encrypt data
	// on the storage volume attached to the ML compute instances that run the training
	// job. The default is None.
	VolumeEncryptionKMSKey *string

	noSmithyDocumentSerde
}

type StartMLModelTrainingJobOutput struct {

	// The ARN of the new model training job.
	Arn *string

	// The model training job creation time, in milliseconds.
	CreationTimeInMillis *int64

	// The unique ID of the new model training job.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMLModelTrainingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMLModelTrainingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMLModelTrainingJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartMLModelTrainingJob"); err != nil {
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
	if err = addOpStartMLModelTrainingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMLModelTrainingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMLModelTrainingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartMLModelTrainingJob",
	}
}
