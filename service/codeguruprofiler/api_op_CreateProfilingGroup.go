// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a profiling group.
func (c *Client) CreateProfilingGroup(ctx context.Context, params *CreateProfilingGroupInput, optFns ...func(*Options)) (*CreateProfilingGroupOutput, error) {
	stack := middleware.NewStack("CreateProfilingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateProfilingGroupMiddlewares(stack)
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
	addOpCreateProfilingGroupValidationMiddleware(stack)
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
			OperationName: "CreateProfilingGroup",
			Err:           err,
		}
	}
	out := result.(*CreateProfilingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the createProfiliingGroupRequest.
type CreateProfilingGroupInput struct {

	// The agent orchestration configuration.
	AgentOrchestrationConfig *types.AgentOrchestrationConfig

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. This parameter specifies a unique identifier for the new profiling
	// group that helps ensure idempotency.
	//
	// This member is required.
	ClientToken *string

	ComputePlatform types.ComputePlatform

	// The name of the profiling group.
	//
	// This member is required.
	ProfilingGroupName *string
}

// The structure representing the createProfilingGroupResponse.
type CreateProfilingGroupOutput struct {

	// Information about the new profiling group
	//
	// This member is required.
	ProfilingGroup *types.ProfilingGroupDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateProfilingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateProfilingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProfilingGroup{}, middleware.After)
}
