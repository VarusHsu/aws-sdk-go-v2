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

// Describes the task sets in the specified cluster and service. This is used when
// a service uses the EXTERNAL deployment controller type. For more information,
// see Amazon ECS Deployment Types
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) DescribeTaskSets(ctx context.Context, params *DescribeTaskSetsInput, optFns ...func(*Options)) (*DescribeTaskSetsOutput, error) {
	stack := middleware.NewStack("DescribeTaskSets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTaskSetsMiddlewares(stack)
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
	addOpDescribeTaskSetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTaskSets(options.Region), middleware.Before)
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
			OperationName: "DescribeTaskSets",
			Err:           err,
		}
	}
	out := result.(*DescribeTaskSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTaskSetsInput struct {

	// The short name or full Amazon Resource Name (ARN) of the service that the task
	// sets exist in.
	//
	// This member is required.
	Service *string

	// Specifies whether to see the resource tags for the task set. If TAGS is
	// specified, the tags are included in the response. If this field is omitted, tags
	// are not included in the response.
	Include []types.TaskSetField

	// The ID or full Amazon Resource Name (ARN) of task sets to describe.
	TaskSets []*string

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service that the task sets exist in.
	//
	// This member is required.
	Cluster *string
}

type DescribeTaskSetsOutput struct {

	// The list of task sets described.
	TaskSets []*types.TaskSet

	// Any failures associated with the call.
	Failures []*types.Failure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTaskSetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTaskSets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTaskSets{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTaskSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DescribeTaskSets",
	}
}
