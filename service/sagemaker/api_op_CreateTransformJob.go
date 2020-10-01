// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a transform job. A transform job uses a trained model to get inferences
// on a dataset and saves these results to an Amazon S3 location that you specify.
// To perform batch transformations, you create a transform job and use the data
// that you have readily available. In the request body, you provide the
// following:
//
//     * TransformJobName - Identifies the transform job. The name must
// be unique within an AWS Region in an AWS account.
//
//     * ModelName - Identifies
// the model to use. ModelName must be the name of an existing Amazon SageMaker
// model in the same AWS Region and AWS account. For information on creating a
// model, see CreateModel ().
//
//     * TransformInput - Describes the dataset to be
// transformed and the Amazon S3 location where it is stored.
//
//     *
// TransformOutput - Identifies the Amazon S3 location where you want Amazon
// SageMaker to save the results from the transform job.
//
//     * TransformResources
// - Identifies the ML compute instances for the transform job.
//
// For more
// information about how batch transformation works, see Batch Transform
// (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform.html).
func (c *Client) CreateTransformJob(ctx context.Context, params *CreateTransformJobInput, optFns ...func(*Options)) (*CreateTransformJobOutput, error) {
	stack := middleware.NewStack("CreateTransformJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTransformJobMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateTransformJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransformJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateTransformJob",
			Err:           err,
		}
	}
	out := result.(*CreateTransformJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransformJobInput struct {

	// The name of the model that you want to use for the transform job. ModelName must
	// be the name of an existing Amazon SageMaker model within an AWS Region in an AWS
	// account.
	//
	// This member is required.
	ModelName *string

	// Describes the results of the transform job.
	//
	// This member is required.
	TransformOutput *types.TransformOutput

	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//
	//     * CreateProcessingJob ()
	//
	//     *
	// CreateTrainingJob ()
	//
	//     * CreateTransformJob ()
	ExperimentConfig *types.ExperimentConfig

	// The name of the transform job. The name must be unique within an AWS Region in
	// an AWS account.
	//
	// This member is required.
	TransformJobName *string

	// The maximum allowed size of the payload, in MB. A payload is the data portion of
	// a record (without metadata). The value in MaxPayloadInMB must be greater than,
	// or equal to, the size of a single record. To estimate the size of a record in
	// MB, divide the size of your dataset by the number of records. To ensure that the
	// records fit within the maximum payload size, we recommend using a slightly
	// larger value. The default value is 6 MB. For cases where the payload might be
	// arbitrarily large and is transmitted using HTTP chunked encoding, set the value
	// to 0. This feature works only in supported algorithms. Currently, Amazon
	// SageMaker built-in algorithms do not support HTTP chunked encoding.
	MaxPayloadInMB *int32

	// Specifies the number of records to include in a mini-batch for an HTTP inference
	// request. A record is a single unit of input data that inference can be made on.
	// For example, a single line in a CSV file is a record. To enable the batch
	// strategy, you must set the SplitType property to Line, RecordIO, or TFRecord. To
	// use only one record when making an HTTP invocation request to a container, set
	// BatchStrategy to SingleRecord and SplitType to Line. To fit as many records in a
	// mini-batch as can fit within the MaxPayloadInMB limit, set BatchStrategy to
	// MultiRecord and SplitType to Line.
	BatchStrategy types.BatchStrategy

	// (Optional) An array of key-value pairs. For more information, see Using Cost
	// Allocation Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the AWS Billing and Cost Management User Guide.
	Tags []*types.Tag

	// Describes the resources, including ML instance types and ML instance count, to
	// use for the transform job.
	//
	// This member is required.
	TransformResources *types.TransformResources

	// The data structure used to specify the data to be used for inference in a batch
	// transform job and to associate the data that is relevant to the prediction
	// results in the output. The input filter provided allows you to exclude input
	// data that is not needed for inference in a batch transform job. The output
	// filter provided allows you to include input data relevant to interpreting the
	// predictions in the output from the job. For more information, see Associate
	// Prediction Results with their Corresponding Input Records
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html).
	DataProcessing *types.DataProcessing

	// The environment variables to set in the Docker container. We support up to 16
	// key and values entries in the map.
	Environment map[string]*string

	// Configures the timeout and maximum number of retries for processing a transform
	// job invocation.
	ModelClientConfig *types.ModelClientConfig

	// Describes the input source and the way the transform job consumes it.
	//
	// This member is required.
	TransformInput *types.TransformInput

	// The maximum number of parallel requests that can be sent to each instance in a
	// transform job. If MaxConcurrentTransforms is set to 0 or left unset, Amazon
	// SageMaker checks the optional execution-parameters to determine the settings for
	// your chosen algorithm. If the execution-parameters endpoint is not enabled, the
	// default value is 1. For more information on execution-parameters, see How
	// Containers Serve Requests
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-batch-code.html#your-algorithms-batch-code-how-containe-serves-requests).
	// For built-in algorithms, you don't need to set a value for
	// MaxConcurrentTransforms.
	MaxConcurrentTransforms *int32
}

type CreateTransformJobOutput struct {

	// The Amazon Resource Name (ARN) of the transform job.
	//
	// This member is required.
	TransformJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTransformJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTransformJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTransformJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTransformJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateTransformJob",
	}
}
