// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified scaling plan. Deleting a scaling plan deletes the
// underlying ScalingInstruction () for all of the scalable resources that are
// covered by the plan. If the plan has launched resources or has scaling
// activities in progress, you must delete those resources separately.
func (c *Client) DeleteScalingPlan(ctx context.Context, params *DeleteScalingPlanInput, optFns ...func(*Options)) (*DeleteScalingPlanOutput, error) {
	stack := middleware.NewStack("DeleteScalingPlan", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteScalingPlanMiddlewares(stack)
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
	addOpDeleteScalingPlanValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteScalingPlan(options.Region), middleware.Before)
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
			OperationName: "DeleteScalingPlan",
			Err:           err,
		}
	}
	out := result.(*DeleteScalingPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteScalingPlanInput struct {

	// The name of the scaling plan.
	//
	// This member is required.
	ScalingPlanName *string

	// The version number of the scaling plan.
	//
	// This member is required.
	ScalingPlanVersion *int64
}

type DeleteScalingPlanOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteScalingPlanMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteScalingPlan{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteScalingPlan{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteScalingPlan(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling-plans",
		OperationName: "DeleteScalingPlan",
	}
}
