// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns a description of a processing job.
func (c *Client) DescribeProcessingJob(ctx context.Context, params *DescribeProcessingJobInput, optFns ...func(*Options)) (*DescribeProcessingJobOutput, error) {
	if params == nil {
		params = &DescribeProcessingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProcessingJob", params, optFns, c.addOperationDescribeProcessingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProcessingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProcessingJobInput struct {

	// The name of the processing job. The name must be unique within an Amazon Web
	// Services Region in the Amazon Web Services account.
	//
	// This member is required.
	ProcessingJobName *string

	noSmithyDocumentSerde
}

type DescribeProcessingJobOutput struct {

	// Configures the processing job to run a specified container image.
	//
	// This member is required.
	AppSpecification *types.AppSpecification

	// The time at which the processing job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the processing job.
	//
	// This member is required.
	ProcessingJobArn *string

	// The name of the processing job. The name must be unique within an Amazon Web
	// Services Region in the Amazon Web Services account.
	//
	// This member is required.
	ProcessingJobName *string

	// Provides the status of a processing job.
	//
	// This member is required.
	ProcessingJobStatus types.ProcessingJobStatus

	// Identifies the resources, ML compute instances, and ML storage volumes to
	// deploy for a processing job. In distributed training, you specify more than one
	// instance.
	//
	// This member is required.
	ProcessingResources *types.ProcessingResources

	// The ARN of an AutoML job associated with this processing job.
	AutoMLJobArn *string

	// The environment variables set in the Docker container.
	Environment map[string]string

	// An optional string, up to one KB in size, that contains metadata from the
	// processing container when the processing job exits.
	ExitMessage *string

	// The configuration information used to create an experiment.
	ExperimentConfig *types.ExperimentConfig

	// A string, up to one KB in size, that contains the reason a processing job
	// failed, if it failed.
	FailureReason *string

	// The time at which the processing job was last modified.
	LastModifiedTime *time.Time

	// The ARN of a monitoring schedule for an endpoint associated with this
	// processing job.
	MonitoringScheduleArn *string

	// Networking options for a processing job.
	NetworkConfig *types.NetworkConfig

	// The time at which the processing job completed.
	ProcessingEndTime *time.Time

	// The inputs for a processing job.
	ProcessingInputs []types.ProcessingInput

	// Output configuration for the processing job.
	ProcessingOutputConfig *types.ProcessingOutputConfig

	// The time at which the processing job started.
	ProcessingStartTime *time.Time

	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume
	// to perform tasks on your behalf.
	RoleArn *string

	// The time limit for how long the processing job is allowed to run.
	StoppingCondition *types.ProcessingStoppingCondition

	// The ARN of a training job associated with this processing job.
	TrainingJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeProcessingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProcessingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProcessingJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeProcessingJob"); err != nil {
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
	if err = addOpDescribeProcessingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProcessingJob(options.Region), middleware.Before); err != nil {
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

// ProcessingJobCompletedOrStoppedWaiterOptions are waiter options for
// ProcessingJobCompletedOrStoppedWaiter
type ProcessingJobCompletedOrStoppedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	//
	// Passing options here is functionally equivalent to passing values to this
	// config's ClientOptions field that extend the inner client's APIOptions directly.
	APIOptions []func(*middleware.Stack) error

	// Functional options to be passed to all operations invoked by this client.
	//
	// Function values that modify the inner APIOptions are applied after the waiter
	// config's own APIOptions modifiers.
	ClientOptions []func(*Options)

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ProcessingJobCompletedOrStoppedWaiter will use default minimum delay of 60
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ProcessingJobCompletedOrStoppedWaiter will use default max delay of
	// 120 seconds. Note that MaxDelay must resolve to value greater than or equal to
	// the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state.
	//
	// By default service-modeled logic will populate this option. This option can
	// thus be used to define a custom waiter state with fall-back to service-modeled
	// waiter state mutators.The function returns an error in case of a failure state.
	// In case of retry state, this function returns a bool value of true and nil
	// error, while in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeProcessingJobInput, *DescribeProcessingJobOutput, error) (bool, error)
}

// ProcessingJobCompletedOrStoppedWaiter defines the waiters for
// ProcessingJobCompletedOrStopped
type ProcessingJobCompletedOrStoppedWaiter struct {
	client DescribeProcessingJobAPIClient

	options ProcessingJobCompletedOrStoppedWaiterOptions
}

// NewProcessingJobCompletedOrStoppedWaiter constructs a
// ProcessingJobCompletedOrStoppedWaiter.
func NewProcessingJobCompletedOrStoppedWaiter(client DescribeProcessingJobAPIClient, optFns ...func(*ProcessingJobCompletedOrStoppedWaiterOptions)) *ProcessingJobCompletedOrStoppedWaiter {
	options := ProcessingJobCompletedOrStoppedWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = processingJobCompletedOrStoppedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ProcessingJobCompletedOrStoppedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ProcessingJobCompletedOrStopped waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ProcessingJobCompletedOrStoppedWaiter) Wait(ctx context.Context, params *DescribeProcessingJobInput, maxWaitDur time.Duration, optFns ...func(*ProcessingJobCompletedOrStoppedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ProcessingJobCompletedOrStopped
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ProcessingJobCompletedOrStoppedWaiter) WaitForOutput(ctx context.Context, params *DescribeProcessingJobInput, maxWaitDur time.Duration, optFns ...func(*ProcessingJobCompletedOrStoppedWaiterOptions)) (*DescribeProcessingJobOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeProcessingJob(ctx, params, func(o *Options) {
			baseOpts := []func(*Options){
				addIsWaiterUserAgent,
			}
			o.APIOptions = append(o.APIOptions, apiOptions...)
			for _, opt := range baseOpts {
				opt(o)
			}
			for _, opt := range options.ClientOptions {
				opt(o)
			}
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ProcessingJobCompletedOrStopped waiter")
}

func processingJobCompletedOrStoppedStateRetryable(ctx context.Context, input *DescribeProcessingJobInput, output *DescribeProcessingJobOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ProcessingJobStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Completed"
		value, ok := pathValue.(types.ProcessingJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ProcessingJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ProcessingJobStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Stopped"
		value, ok := pathValue.(types.ProcessingJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ProcessingJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ProcessingJobStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.ProcessingJobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ProcessingJobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "ValidationException" == apiErr.ErrorCode() {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// DescribeProcessingJobAPIClient is a client that implements the
// DescribeProcessingJob operation.
type DescribeProcessingJobAPIClient interface {
	DescribeProcessingJob(context.Context, *DescribeProcessingJobInput, ...func(*Options)) (*DescribeProcessingJobOutput, error)
}

var _ DescribeProcessingJobAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeProcessingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeProcessingJob",
	}
}
