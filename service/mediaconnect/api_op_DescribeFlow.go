// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Displays the details of a flow. The response includes the flow ARN, name, and
// Availability Zone, as well as details about the source, outputs, and
// entitlements.
func (c *Client) DescribeFlow(ctx context.Context, params *DescribeFlowInput, optFns ...func(*Options)) (*DescribeFlowOutput, error) {
	if params == nil {
		params = &DescribeFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFlow", params, optFns, c.addOperationDescribeFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFlowInput struct {

	// The ARN of the flow that you want to describe.
	//
	// This member is required.
	FlowArn *string

	noSmithyDocumentSerde
}

type DescribeFlowOutput struct {

	// The settings for a flow, including its source, outputs, and entitlements.
	Flow *types.Flow

	// Messages that provide the state of the flow.
	Messages *types.Messages

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFlow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFlow"); err != nil {
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
	if err = addOpDescribeFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFlow(options.Region), middleware.Before); err != nil {
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

// FlowActiveWaiterOptions are waiter options for FlowActiveWaiter
type FlowActiveWaiterOptions struct {

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
	// FlowActiveWaiter will use default minimum delay of 3 seconds. Note that MinDelay
	// must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, FlowActiveWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeFlowInput, *DescribeFlowOutput, error) (bool, error)
}

// FlowActiveWaiter defines the waiters for FlowActive
type FlowActiveWaiter struct {
	client DescribeFlowAPIClient

	options FlowActiveWaiterOptions
}

// NewFlowActiveWaiter constructs a FlowActiveWaiter.
func NewFlowActiveWaiter(client DescribeFlowAPIClient, optFns ...func(*FlowActiveWaiterOptions)) *FlowActiveWaiter {
	options := FlowActiveWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = flowActiveStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &FlowActiveWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for FlowActive waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *FlowActiveWaiter) Wait(ctx context.Context, params *DescribeFlowInput, maxWaitDur time.Duration, optFns ...func(*FlowActiveWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for FlowActive waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *FlowActiveWaiter) WaitForOutput(ctx context.Context, params *DescribeFlowInput, maxWaitDur time.Duration, optFns ...func(*FlowActiveWaiterOptions)) (*DescribeFlowOutput, error) {
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

		out, err := w.client.DescribeFlow(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for FlowActive waiter")
}

func flowActiveStateRetryable(ctx context.Context, input *DescribeFlowInput, output *DescribeFlowOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Flow.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Flow.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STARTING"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Flow.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "UPDATING"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.ServiceUnavailableException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Flow.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ERROR"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// FlowDeletedWaiterOptions are waiter options for FlowDeletedWaiter
type FlowDeletedWaiterOptions struct {

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
	// FlowDeletedWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, FlowDeletedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeFlowInput, *DescribeFlowOutput, error) (bool, error)
}

// FlowDeletedWaiter defines the waiters for FlowDeleted
type FlowDeletedWaiter struct {
	client DescribeFlowAPIClient

	options FlowDeletedWaiterOptions
}

// NewFlowDeletedWaiter constructs a FlowDeletedWaiter.
func NewFlowDeletedWaiter(client DescribeFlowAPIClient, optFns ...func(*FlowDeletedWaiterOptions)) *FlowDeletedWaiter {
	options := FlowDeletedWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = flowDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &FlowDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for FlowDeleted waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *FlowDeletedWaiter) Wait(ctx context.Context, params *DescribeFlowInput, maxWaitDur time.Duration, optFns ...func(*FlowDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for FlowDeleted waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *FlowDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeFlowInput, maxWaitDur time.Duration, optFns ...func(*FlowDeletedWaiterOptions)) (*DescribeFlowOutput, error) {
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

		out, err := w.client.DescribeFlow(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for FlowDeleted waiter")
}

func flowDeletedStateRetryable(ctx context.Context, input *DescribeFlowInput, output *DescribeFlowOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.NotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Flow.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETING"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.ServiceUnavailableException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Flow.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ERROR"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// FlowStandbyWaiterOptions are waiter options for FlowStandbyWaiter
type FlowStandbyWaiterOptions struct {

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
	// FlowStandbyWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, FlowStandbyWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeFlowInput, *DescribeFlowOutput, error) (bool, error)
}

// FlowStandbyWaiter defines the waiters for FlowStandby
type FlowStandbyWaiter struct {
	client DescribeFlowAPIClient

	options FlowStandbyWaiterOptions
}

// NewFlowStandbyWaiter constructs a FlowStandbyWaiter.
func NewFlowStandbyWaiter(client DescribeFlowAPIClient, optFns ...func(*FlowStandbyWaiterOptions)) *FlowStandbyWaiter {
	options := FlowStandbyWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = flowStandbyStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &FlowStandbyWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for FlowStandby waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *FlowStandbyWaiter) Wait(ctx context.Context, params *DescribeFlowInput, maxWaitDur time.Duration, optFns ...func(*FlowStandbyWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for FlowStandby waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *FlowStandbyWaiter) WaitForOutput(ctx context.Context, params *DescribeFlowInput, maxWaitDur time.Duration, optFns ...func(*FlowStandbyWaiterOptions)) (*DescribeFlowOutput, error) {
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

		out, err := w.client.DescribeFlow(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for FlowStandby waiter")
}

func flowStandbyStateRetryable(ctx context.Context, input *DescribeFlowInput, output *DescribeFlowOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Flow.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STANDBY"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Flow.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STOPPING"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.ServiceUnavailableException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Flow.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ERROR"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// DescribeFlowAPIClient is a client that implements the DescribeFlow operation.
type DescribeFlowAPIClient interface {
	DescribeFlow(context.Context, *DescribeFlowInput, ...func(*Options)) (*DescribeFlowOutput, error)
}

var _ DescribeFlowAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFlow",
	}
}
