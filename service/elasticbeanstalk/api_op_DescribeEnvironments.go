// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns descriptions for existing environments.
func (c *Client) DescribeEnvironments(ctx context.Context, params *DescribeEnvironmentsInput, optFns ...func(*Options)) (*DescribeEnvironmentsOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironments", params, optFns, c.addOperationDescribeEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe one or more environments.
type DescribeEnvironmentsInput struct {

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that are associated with this application.
	ApplicationName *string

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that have the specified IDs.
	EnvironmentIds []string

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that have the specified names.
	EnvironmentNames []string

	// Indicates whether to include deleted environments:
	//
	// true : Environments that have been deleted after IncludedDeletedBackTo are
	// displayed.
	//
	// false : Do not include deleted environments.
	IncludeDeleted *bool

	//  If specified when IncludeDeleted is set to true , then environments deleted
	// after this date are displayed.
	IncludedDeletedBackTo *time.Time

	// For a paginated request. Specify a maximum number of environments to include in
	// each response.
	//
	// If no MaxRecords is specified, all available environments are retrieved in a
	// single response.
	MaxRecords *int32

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical to
	// the ones specified in the initial request.
	//
	// If no NextToken is specified, the first page is retrieved.
	NextToken *string

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// include only those that are associated with this application version.
	VersionLabel *string

	noSmithyDocumentSerde
}

// Result message containing a list of environment descriptions.
type DescribeEnvironmentsOutput struct {

	//  Returns an EnvironmentDescription list.
	Environments []types.EnvironmentDescription

	// In a paginated request, the token that you can pass in a subsequent request to
	// get the next response page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEnvironments"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironments(options.Region), middleware.Before); err != nil {
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

// EnvironmentExistsWaiterOptions are waiter options for EnvironmentExistsWaiter
type EnvironmentExistsWaiterOptions struct {

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
	// EnvironmentExistsWaiter will use default minimum delay of 20 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, EnvironmentExistsWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeEnvironmentsInput, *DescribeEnvironmentsOutput, error) (bool, error)
}

// EnvironmentExistsWaiter defines the waiters for EnvironmentExists
type EnvironmentExistsWaiter struct {
	client DescribeEnvironmentsAPIClient

	options EnvironmentExistsWaiterOptions
}

// NewEnvironmentExistsWaiter constructs a EnvironmentExistsWaiter.
func NewEnvironmentExistsWaiter(client DescribeEnvironmentsAPIClient, optFns ...func(*EnvironmentExistsWaiterOptions)) *EnvironmentExistsWaiter {
	options := EnvironmentExistsWaiterOptions{}
	options.MinDelay = 20 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = environmentExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &EnvironmentExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for EnvironmentExists waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *EnvironmentExistsWaiter) Wait(ctx context.Context, params *DescribeEnvironmentsInput, maxWaitDur time.Duration, optFns ...func(*EnvironmentExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for EnvironmentExists waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *EnvironmentExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeEnvironmentsInput, maxWaitDur time.Duration, optFns ...func(*EnvironmentExistsWaiterOptions)) (*DescribeEnvironmentsOutput, error) {
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

		out, err := w.client.DescribeEnvironments(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for EnvironmentExists waiter")
}

func environmentExistsStateRetryable(ctx context.Context, input *DescribeEnvironmentsInput, output *DescribeEnvironmentsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Environments[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Ready"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.EnvironmentStatus)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.EnvironmentStatus value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Environments[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Launching"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.EnvironmentStatus)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.EnvironmentStatus value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return true, nil
		}
	}

	return true, nil
}

// EnvironmentTerminatedWaiterOptions are waiter options for
// EnvironmentTerminatedWaiter
type EnvironmentTerminatedWaiterOptions struct {

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
	// EnvironmentTerminatedWaiter will use default minimum delay of 20 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, EnvironmentTerminatedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *DescribeEnvironmentsInput, *DescribeEnvironmentsOutput, error) (bool, error)
}

// EnvironmentTerminatedWaiter defines the waiters for EnvironmentTerminated
type EnvironmentTerminatedWaiter struct {
	client DescribeEnvironmentsAPIClient

	options EnvironmentTerminatedWaiterOptions
}

// NewEnvironmentTerminatedWaiter constructs a EnvironmentTerminatedWaiter.
func NewEnvironmentTerminatedWaiter(client DescribeEnvironmentsAPIClient, optFns ...func(*EnvironmentTerminatedWaiterOptions)) *EnvironmentTerminatedWaiter {
	options := EnvironmentTerminatedWaiterOptions{}
	options.MinDelay = 20 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = environmentTerminatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &EnvironmentTerminatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for EnvironmentTerminated waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *EnvironmentTerminatedWaiter) Wait(ctx context.Context, params *DescribeEnvironmentsInput, maxWaitDur time.Duration, optFns ...func(*EnvironmentTerminatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for EnvironmentTerminated waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *EnvironmentTerminatedWaiter) WaitForOutput(ctx context.Context, params *DescribeEnvironmentsInput, maxWaitDur time.Duration, optFns ...func(*EnvironmentTerminatedWaiterOptions)) (*DescribeEnvironmentsOutput, error) {
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

		out, err := w.client.DescribeEnvironments(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for EnvironmentTerminated waiter")
}

func environmentTerminatedStateRetryable(ctx context.Context, input *DescribeEnvironmentsInput, output *DescribeEnvironmentsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Environments[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Terminated"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.EnvironmentStatus)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.EnvironmentStatus value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Environments[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Terminating"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.EnvironmentStatus)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.EnvironmentStatus value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return true, nil
		}
	}

	return true, nil
}

// EnvironmentUpdatedWaiterOptions are waiter options for EnvironmentUpdatedWaiter
type EnvironmentUpdatedWaiterOptions struct {

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
	// EnvironmentUpdatedWaiter will use default minimum delay of 20 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, EnvironmentUpdatedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeEnvironmentsInput, *DescribeEnvironmentsOutput, error) (bool, error)
}

// EnvironmentUpdatedWaiter defines the waiters for EnvironmentUpdated
type EnvironmentUpdatedWaiter struct {
	client DescribeEnvironmentsAPIClient

	options EnvironmentUpdatedWaiterOptions
}

// NewEnvironmentUpdatedWaiter constructs a EnvironmentUpdatedWaiter.
func NewEnvironmentUpdatedWaiter(client DescribeEnvironmentsAPIClient, optFns ...func(*EnvironmentUpdatedWaiterOptions)) *EnvironmentUpdatedWaiter {
	options := EnvironmentUpdatedWaiterOptions{}
	options.MinDelay = 20 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = environmentUpdatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &EnvironmentUpdatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for EnvironmentUpdated waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *EnvironmentUpdatedWaiter) Wait(ctx context.Context, params *DescribeEnvironmentsInput, maxWaitDur time.Duration, optFns ...func(*EnvironmentUpdatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for EnvironmentUpdated waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *EnvironmentUpdatedWaiter) WaitForOutput(ctx context.Context, params *DescribeEnvironmentsInput, maxWaitDur time.Duration, optFns ...func(*EnvironmentUpdatedWaiterOptions)) (*DescribeEnvironmentsOutput, error) {
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

		out, err := w.client.DescribeEnvironments(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for EnvironmentUpdated waiter")
}

func environmentUpdatedStateRetryable(ctx context.Context, input *DescribeEnvironmentsInput, output *DescribeEnvironmentsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Environments[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Ready"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.EnvironmentStatus)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.EnvironmentStatus value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Environments[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Updating"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.EnvironmentStatus)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.EnvironmentStatus value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return true, nil
		}
	}

	return true, nil
}

// DescribeEnvironmentsAPIClient is a client that implements the
// DescribeEnvironments operation.
type DescribeEnvironmentsAPIClient interface {
	DescribeEnvironments(context.Context, *DescribeEnvironmentsInput, ...func(*Options)) (*DescribeEnvironmentsOutput, error)
}

var _ DescribeEnvironmentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeEnvironments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEnvironments",
	}
}
