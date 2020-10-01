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

// Gets a description of a hyperparameter tuning job.
func (c *Client) DescribeHyperParameterTuningJob(ctx context.Context, params *DescribeHyperParameterTuningJobInput, optFns ...func(*Options)) (*DescribeHyperParameterTuningJobOutput, error) {
	stack := middleware.NewStack("DescribeHyperParameterTuningJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeHyperParameterTuningJobMiddlewares(stack)
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
	addOpDescribeHyperParameterTuningJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHyperParameterTuningJob(options.Region), middleware.Before)
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
			OperationName: "DescribeHyperParameterTuningJob",
			Err:           err,
		}
	}
	out := result.(*DescribeHyperParameterTuningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHyperParameterTuningJobInput struct {

	// The name of the tuning job.
	//
	// This member is required.
	HyperParameterTuningJobName *string
}

type DescribeHyperParameterTuningJobOutput struct {

	// The HyperParameterTuningJobConfig () object that specifies the configuration of
	// the tuning job.
	//
	// This member is required.
	HyperParameterTuningJobConfig *types.HyperParameterTuningJobConfig

	// If the hyperparameter tuning job is an warm start tuning job with a
	// WarmStartType of IDENTICAL_DATA_AND_ALGORITHM, this is the TrainingJobSummary ()
	// for the training job with the best objective metric value of all training jobs
	// launched by this tuning job and all parent jobs specified for the warm start
	// tuning job.
	OverallBestTrainingJob *types.HyperParameterTrainingJobSummary

	// The TrainingJobStatusCounters () object that specifies the number of training
	// jobs, categorized by status, that this tuning job launched.
	//
	// This member is required.
	TrainingJobStatusCounters *types.TrainingJobStatusCounters

	// The date and time that the tuning job ended.
	HyperParameterTuningEndTime *time.Time

	// The HyperParameterTrainingJobDefinition () object that specifies the definition
	// of the training jobs that this tuning job launches.
	TrainingJobDefinition *types.HyperParameterTrainingJobDefinition

	// The configuration for starting the hyperparameter parameter tuning job using one
	// or more previous tuning jobs as a starting point. The results of previous tuning
	// jobs are used to inform which combinations of hyperparameters to search over in
	// the new tuning job.
	WarmStartConfig *types.HyperParameterTuningJobWarmStartConfig

	// The ObjectiveStatusCounters () object that specifies the number of training
	// jobs, categorized by the status of their final objective metric, that this
	// tuning job launched.
	//
	// This member is required.
	ObjectiveStatusCounters *types.ObjectiveStatusCounters

	// The status of the tuning job: InProgress, Completed, Failed, Stopping, or
	// Stopped.
	//
	// This member is required.
	HyperParameterTuningJobStatus types.HyperParameterTuningJobStatus

	// A list of the HyperParameterTrainingJobDefinition () objects launched for this
	// tuning job.
	TrainingJobDefinitions []*types.HyperParameterTrainingJobDefinition

	// The date and time that the tuning job started.
	//
	// This member is required.
	CreationTime *time.Time

	// The date and time that the status of the tuning job was modified.
	LastModifiedTime *time.Time

	// The name of the tuning job.
	//
	// This member is required.
	HyperParameterTuningJobName *string

	// A TrainingJobSummary () object that describes the training job that completed
	// with the best current HyperParameterTuningJobObjective ().
	BestTrainingJob *types.HyperParameterTrainingJobSummary

	// The Amazon Resource Name (ARN) of the tuning job.
	//
	// This member is required.
	HyperParameterTuningJobArn *string

	// If the tuning job failed, the reason it failed.
	FailureReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeHyperParameterTuningJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHyperParameterTuningJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHyperParameterTuningJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeHyperParameterTuningJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeHyperParameterTuningJob",
	}
}
