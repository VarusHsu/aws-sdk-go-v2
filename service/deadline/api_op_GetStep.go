// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a step.
func (c *Client) GetStep(ctx context.Context, params *GetStepInput, optFns ...func(*Options)) (*GetStepOutput, error) {
	if params == nil {
		params = &GetStepInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStep", params, optFns, c.addOperationGetStepMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStepOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStepInput struct {

	// The farm ID for the step.
	//
	// This member is required.
	FarmId *string

	// The job ID for the step.
	//
	// This member is required.
	JobId *string

	// The queue ID for the step.
	//
	// This member is required.
	QueueId *string

	// The step ID.
	//
	// This member is required.
	StepId *string

	noSmithyDocumentSerde
}

type GetStepOutput struct {

	// The date and time the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user or system that created this resource.
	//
	// This member is required.
	CreatedBy *string

	// The life cycle status of the step.
	//
	// This member is required.
	LifecycleStatus types.StepLifecycleStatus

	// The name of the step.
	//
	// This member is required.
	Name *string

	// The step ID.
	//
	// This member is required.
	StepId *string

	// The task run status for the job.
	//
	// This member is required.
	TaskRunStatus types.TaskRunStatus

	// The number of tasks running on the job.
	//
	// This member is required.
	TaskRunStatusCounts map[string]int32

	// The number of dependencies in the step.
	DependencyCounts *types.DependencyCounts

	// The description of the step.
	Description *string

	// The date and time the resource ended running.
	EndedAt *time.Time

	// A message that describes the lifecycle status of the step.
	LifecycleStatusMessage *string

	// A list of step parameters and the combination expression for the step.
	ParameterSpace *types.ParameterSpace

	// The required capabilities of the step.
	RequiredCapabilities *types.StepRequiredCapabilities

	// The date and time the resource started running.
	StartedAt *time.Time

	// The task status with which the job started.
	TargetTaskRunStatus types.StepTargetTaskRunStatus

	// The date and time the resource was updated.
	UpdatedAt *time.Time

	// The user or system that updated this resource.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStepMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetStep{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStep{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetStep"); err != nil {
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
	if err = addEndpointPrefix_opGetStepMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetStepValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStep(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetStepMiddleware struct {
}

func (*endpointPrefix_opGetStepMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetStepMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetStepMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetStepMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetStep(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetStep",
	}
}
