// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This action is only used by the Amazon ECS agent, and it is not intended for use
// outside of the agent. Sent to acknowledge that a container changed states.
func (c *Client) SubmitContainerStateChange(ctx context.Context, params *SubmitContainerStateChangeInput, optFns ...func(*Options)) (*SubmitContainerStateChangeOutput, error) {
	stack := middleware.NewStack("SubmitContainerStateChange", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSubmitContainerStateChangeMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitContainerStateChange(options.Region), middleware.Before)
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
			OperationName: "SubmitContainerStateChange",
			Err:           err,
		}
	}
	out := result.(*SubmitContainerStateChangeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubmitContainerStateChangeInput struct {

	// The network bindings of the container.
	NetworkBindings []*types.NetworkBinding

	// The name of the container.
	ContainerName *string

	// The reason for the state change request.
	Reason *string

	// The task ID or full Amazon Resource Name (ARN) of the task that hosts the
	// container.
	Task *string

	// The short name or full ARN of the cluster that hosts the container.
	Cluster *string

	// The exit code returned for the state change request.
	ExitCode *int32

	// The status of the state change request.
	Status *string

	// The ID of the Docker container.
	RuntimeId *string
}

type SubmitContainerStateChangeOutput struct {

	// Acknowledgement of the state change.
	Acknowledgment *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSubmitContainerStateChangeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSubmitContainerStateChange{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSubmitContainerStateChange{}, middleware.After)
}

func newServiceMetadataMiddleware_opSubmitContainerStateChange(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "SubmitContainerStateChange",
	}
}
