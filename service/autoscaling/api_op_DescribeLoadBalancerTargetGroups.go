// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the target groups for the specified Auto Scaling group.
func (c *Client) DescribeLoadBalancerTargetGroups(ctx context.Context, params *DescribeLoadBalancerTargetGroupsInput, optFns ...func(*Options)) (*DescribeLoadBalancerTargetGroupsOutput, error) {
	stack := middleware.NewStack("DescribeLoadBalancerTargetGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeLoadBalancerTargetGroupsMiddlewares(stack)
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
	addOpDescribeLoadBalancerTargetGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancerTargetGroups(options.Region), middleware.Before)
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
			OperationName: "DescribeLoadBalancerTargetGroups",
			Err:           err,
		}
	}
	out := result.(*DescribeLoadBalancerTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLoadBalancerTargetGroupsInput struct {

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100.
	MaxRecords *int32
}

type DescribeLoadBalancerTargetGroupsOutput struct {

	// Information about the target groups.
	LoadBalancerTargetGroups []*types.LoadBalancerTargetGroupState

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeLoadBalancerTargetGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancerTargetGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancerTargetGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLoadBalancerTargetGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeLoadBalancerTargetGroups",
	}
}
