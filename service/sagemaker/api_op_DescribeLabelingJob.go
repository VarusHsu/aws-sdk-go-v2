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
	"time"
)

// Gets information about a labeling job.
func (c *Client) DescribeLabelingJob(ctx context.Context, params *DescribeLabelingJobInput, optFns ...func(*Options)) (*DescribeLabelingJobOutput, error) {
	stack := middleware.NewStack("DescribeLabelingJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeLabelingJobMiddlewares(stack)
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
	addOpDescribeLabelingJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLabelingJob(options.Region), middleware.Before)
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
			OperationName: "DescribeLabelingJob",
			Err:           err,
		}
	}
	out := result.(*DescribeLabelingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLabelingJobInput struct {

	// The name of the labeling job to return information for.
	//
	// This member is required.
	LabelingJobName *string
}

type DescribeLabelingJobOutput struct {

	// The date and time that the labeling job was last updated.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The date and time that the labeling job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Input configuration information for the labeling job, such as the Amazon S3
	// location of the data objects and the location of the manifest file that
	// describes the data objects.
	//
	// This member is required.
	InputConfig *types.LabelingJobInputConfig

	// Configuration information required for human workers to complete a labeling
	// task.
	//
	// This member is required.
	HumanTaskConfig *types.HumanTaskConfig

	// A set of conditions for stopping a labeling job. If any of the conditions are
	// met, the job is automatically stopped.
	StoppingConditions *types.LabelingJobStoppingConditions

	// Configuration information for automated data labeling.
	LabelingJobAlgorithmsConfig *types.LabelingJobAlgorithmsConfig

	// The processing status of the labeling job.
	//
	// This member is required.
	LabelingJobStatus types.LabelingJobStatus

	// The Amazon Resource Name (ARN) of the labeling job.
	//
	// This member is required.
	LabelingJobArn *string

	// If the job failed, the reason that it failed.
	FailureReason *string

	// The name assigned to the labeling job when it was created.
	//
	// This member is required.
	LabelingJobName *string

	// The S3 location of the JSON file that defines the categories used to label data
	// objects. Please note the following label-category limits:
	//
	//     * Semantic
	// segmentation labeling jobs using automated labeling: 20 labels
	//
	//     * Box
	// bounding labeling jobs (all): 10 labels
	//
	// The file is a JSON structure in the
	// following format: {
	//     "document-version": "2018-11-28"
	//
	//     "labels": [
	//
	//
	// {
	//
	//     "label": "label 1"
	//
	//     },
	//
	//     {
	//
	//     "label": "label 2"
	//
	//     },
	//
	//
	// ...
	//
	//     {
	//
	//     "label": "label n"
	//
	//     }
	//
	//     ]
	//
	//     }
	LabelCategoryConfigS3Uri *string

	// An array of key/value pairs. For more information, see Using Cost Allocation
	// Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)
	// in the AWS Billing and Cost Management User Guide.
	Tags []*types.Tag

	// The location of the output produced by the labeling job.
	LabelingJobOutput *types.LabelingJobOutput

	// The attribute used as the label in the output manifest file.
	LabelAttributeName *string

	// A unique identifier for work done as part of a labeling job.
	//
	// This member is required.
	JobReferenceCode *string

	// Provides a breakdown of the number of data objects labeled by humans, the number
	// of objects labeled by machine, the number of objects than couldn't be labeled,
	// and the total number of objects labeled.
	//
	// This member is required.
	LabelCounters *types.LabelCounters

	// The location of the job's output data and the AWS Key Management Service key ID
	// for the key used to encrypt the output data, if any.
	//
	// This member is required.
	OutputConfig *types.LabelingJobOutputConfig

	// The Amazon Resource Name (ARN) that Amazon SageMaker assumes to perform tasks on
	// your behalf during data labeling.
	//
	// This member is required.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeLabelingJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLabelingJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLabelingJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLabelingJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeLabelingJob",
	}
}
