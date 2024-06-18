// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns a list of DescribeEvaluations that match the search criteria in the
// request.
func (c *Client) DescribeEvaluations(ctx context.Context, params *DescribeEvaluationsInput, optFns ...func(*Options)) (*DescribeEvaluationsOutput, error) {
	if params == nil {
		params = &DescribeEvaluationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEvaluations", params, optFns, c.addOperationDescribeEvaluationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEvaluationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEvaluationsInput struct {

	// The equal to operator. The Evaluation results will have FilterVariable values
	// that exactly match the value specified with EQ .
	EQ *string

	// Use one of the following variable to filter a list of Evaluation objects:
	//
	//   - CreatedAt - Sets the search criteria to the Evaluation creation date.
	//
	//   - Status - Sets the search criteria to the Evaluation status.
	//
	//   - Name - Sets the search criteria to the contents of Evaluation Name .
	//
	//   - IAMUser - Sets the search criteria to the user account that invoked an
	//   Evaluation .
	//
	//   - MLModelId - Sets the search criteria to the MLModel that was evaluated.
	//
	//   - DataSourceId - Sets the search criteria to the DataSource used in Evaluation
	//   .
	//
	//   - DataUri - Sets the search criteria to the data file(s) used in Evaluation .
	//   The URL can identify either a file or an Amazon Simple Storage Solution (Amazon
	//   S3) bucket or directory.
	FilterVariable types.EvaluationFilterVariable

	// The greater than or equal to operator. The Evaluation results will have
	// FilterVariable values that are greater than or equal to the value specified with
	// GE .
	GE *string

	// The greater than operator. The Evaluation results will have FilterVariable
	// values that are greater than the value specified with GT .
	GT *string

	// The less than or equal to operator. The Evaluation results will have
	// FilterVariable values that are less than or equal to the value specified with LE
	// .
	LE *string

	// The less than operator. The Evaluation results will have FilterVariable values
	// that are less than the value specified with LT .
	LT *string

	//  The maximum number of Evaluation to include in the result.
	Limit *int32

	// The not equal to operator. The Evaluation results will have FilterVariable
	// values not equal to the value specified with NE .
	NE *string

	// The ID of the page in the paginated results.
	NextToken *string

	// A string that is found at the beginning of a variable, such as Name or Id .
	//
	// For example, an Evaluation could have the Name 2014-09-09-HolidayGiftMailer . To
	// search for this Evaluation , select Name for the FilterVariable and any of the
	// following strings for the Prefix :
	//
	//   - 2014-09
	//
	//   - 2014-09-09
	//
	//   - 2014-09-09-Holiday
	Prefix *string

	// A two-value parameter that determines the sequence of the resulting list of
	// Evaluation .
	//
	//   - asc - Arranges the list in ascending order (A-Z, 0-9).
	//
	//   - dsc - Arranges the list in descending order (Z-A, 9-0).
	//
	// Results are sorted by FilterVariable .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

// Represents the query results from a DescribeEvaluations operation. The content
// is essentially a list of Evaluation .
type DescribeEvaluationsOutput struct {

	// The ID of the next page in the paginated results that indicates at least one
	// more page follows.
	NextToken *string

	// A list of Evaluation that meet the search criteria.
	Results []types.Evaluation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEvaluationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEvaluations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEvaluations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEvaluations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEvaluations(options.Region), middleware.Before); err != nil {
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

// EvaluationAvailableWaiterOptions are waiter options for
// EvaluationAvailableWaiter
type EvaluationAvailableWaiterOptions struct {

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
	// EvaluationAvailableWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, EvaluationAvailableWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeEvaluationsInput, *DescribeEvaluationsOutput, error) (bool, error)
}

// EvaluationAvailableWaiter defines the waiters for EvaluationAvailable
type EvaluationAvailableWaiter struct {
	client DescribeEvaluationsAPIClient

	options EvaluationAvailableWaiterOptions
}

// NewEvaluationAvailableWaiter constructs a EvaluationAvailableWaiter.
func NewEvaluationAvailableWaiter(client DescribeEvaluationsAPIClient, optFns ...func(*EvaluationAvailableWaiterOptions)) *EvaluationAvailableWaiter {
	options := EvaluationAvailableWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = evaluationAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &EvaluationAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for EvaluationAvailable waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *EvaluationAvailableWaiter) Wait(ctx context.Context, params *DescribeEvaluationsInput, maxWaitDur time.Duration, optFns ...func(*EvaluationAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for EvaluationAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *EvaluationAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeEvaluationsInput, maxWaitDur time.Duration, optFns ...func(*EvaluationAvailableWaiterOptions)) (*DescribeEvaluationsOutput, error) {
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

		out, err := w.client.DescribeEvaluations(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for EvaluationAvailable waiter")
}

func evaluationAvailableStateRetryable(ctx context.Context, input *DescribeEvaluationsInput, output *DescribeEvaluationsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Results[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "COMPLETED"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.EntityStatus)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.EntityStatus value, got %T", pathValue)
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
		pathValue, err := jmespath.Search("Results[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.EntityStatus)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.EntityStatus value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// DescribeEvaluationsPaginatorOptions is the paginator options for
// DescribeEvaluations
type DescribeEvaluationsPaginatorOptions struct {
	//  The maximum number of Evaluation to include in the result.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEvaluationsPaginator is a paginator for DescribeEvaluations
type DescribeEvaluationsPaginator struct {
	options   DescribeEvaluationsPaginatorOptions
	client    DescribeEvaluationsAPIClient
	params    *DescribeEvaluationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEvaluationsPaginator returns a new DescribeEvaluationsPaginator
func NewDescribeEvaluationsPaginator(client DescribeEvaluationsAPIClient, params *DescribeEvaluationsInput, optFns ...func(*DescribeEvaluationsPaginatorOptions)) *DescribeEvaluationsPaginator {
	if params == nil {
		params = &DescribeEvaluationsInput{}
	}

	options := DescribeEvaluationsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEvaluationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEvaluationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEvaluations page.
func (p *DescribeEvaluationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEvaluationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeEvaluations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeEvaluationsAPIClient is a client that implements the
// DescribeEvaluations operation.
type DescribeEvaluationsAPIClient interface {
	DescribeEvaluations(context.Context, *DescribeEvaluationsInput, ...func(*Options)) (*DescribeEvaluationsOutput, error)
}

var _ DescribeEvaluationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeEvaluations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEvaluations",
	}
}
