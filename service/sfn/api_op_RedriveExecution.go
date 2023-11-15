// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restarts unsuccessful executions of Standard workflows that didn't complete
// successfully in the last 14 days. These include failed, aborted, or timed out
// executions. When you redrive (https://docs.aws.amazon.com/step-functions/latest/dg/redrive-executions.html)
// an execution, it continues the failed execution from the unsuccessful step and
// uses the same input. Step Functions preserves the results and execution history
// of the successful steps, and doesn't rerun these steps when you redrive an
// execution. Redriven executions use the same state machine definition and
// execution ARN as the original execution attempt. For workflows that include an
// Inline Map (https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-map-state.html)
// or Parallel (https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-parallel-state.html)
// state, RedriveExecution API action reschedules and redrives only the iterations
// and branches that failed or aborted. To redrive a workflow that includes a
// Distributed Map state with failed child workflow executions, you must redrive
// the parent workflow (https://docs.aws.amazon.com/step-functions/latest/dg/use-dist-map-orchestrate-large-scale-parallel-workloads.html#dist-map-orchestrate-parallel-workloads-key-terms)
// . The parent workflow redrives all the unsuccessful states, including
// Distributed Map. This API action is not supported by EXPRESS state machines.
// However, you can restart the unsuccessful executions of Express child workflows
// in a Distributed Map by redriving its Map Run. When you redrive a Map Run, the
// Express child workflows are rerun using the StartExecution API action. For more
// information, see Redriving Map Runs (https://docs.aws.amazon.com/step-functions/latest/dg/redrive-map-run.html)
// . You can redrive executions if your original execution meets the following
// conditions:
//   - The execution status isn't SUCCEEDED .
//   - Your workflow execution has not exceeded the redrivable period of 14 days.
//     Redrivable period refers to the time during which you can redrive a given
//     execution. This period starts from the day a state machine completes its
//     execution.
//   - The workflow execution has not exceeded the maximum open time of one year.
//     For more information about state machine quotas, see Quotas related to state
//     machine executions (https://docs.aws.amazon.com/step-functions/latest/dg/limits-overview.html#service-limits-state-machine-executions)
//     .
//   - The execution event history count is less than 24,999. Redriven executions
//     append their event history to the existing event history. Make sure your
//     workflow execution contains less than 24,999 events to accommodate the
//     ExecutionRedriven history event and at least one other history event.
func (c *Client) RedriveExecution(ctx context.Context, params *RedriveExecutionInput, optFns ...func(*Options)) (*RedriveExecutionOutput, error) {
	if params == nil {
		params = &RedriveExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RedriveExecution", params, optFns, c.addOperationRedriveExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RedriveExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RedriveExecutionInput struct {

	// The Amazon Resource Name (ARN) of the execution to be redriven.
	//
	// This member is required.
	ExecutionArn *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If you don’t specify a client token, the Amazon Web Services SDK
	// automatically generates a client token and uses it for the request to ensure
	// idempotency. The API uses one of the last 10 client tokens provided.
	ClientToken *string

	noSmithyDocumentSerde
}

type RedriveExecutionOutput struct {

	// The date the execution was last redriven.
	//
	// This member is required.
	RedriveDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRedriveExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpRedriveExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpRedriveExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RedriveExecution"); err != nil {
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
	if err = addIdempotencyToken_opRedriveExecutionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRedriveExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRedriveExecution(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRedriveExecution struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRedriveExecution) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRedriveExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RedriveExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RedriveExecutionInput ")
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
func addIdempotencyToken_opRedriveExecutionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRedriveExecution{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRedriveExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RedriveExecution",
	}
}
