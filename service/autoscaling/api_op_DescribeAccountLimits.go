// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the current Amazon EC2 Auto Scaling resource quotas for your AWS
// account. For information about requesting an increase, see Amazon EC2 Auto
// Scaling Service Quotas
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-account-limits.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DescribeAccountLimits(ctx context.Context, params *DescribeAccountLimitsInput, optFns ...func(*Options)) (*DescribeAccountLimitsOutput, error) {
	stack := middleware.NewStack("DescribeAccountLimits", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeAccountLimitsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountLimits(options.Region), middleware.Before)
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
			OperationName: "DescribeAccountLimits",
			Err:           err,
		}
	}
	out := result.(*DescribeAccountLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountLimitsInput struct {
}

type DescribeAccountLimitsOutput struct {

	// The current number of groups for your AWS account.
	NumberOfAutoScalingGroups *int32

	// The current number of launch configurations for your AWS account.
	NumberOfLaunchConfigurations *int32

	// The maximum number of groups allowed for your AWS account. The default is 200
	// groups per AWS Region.
	MaxNumberOfAutoScalingGroups *int32

	// The maximum number of launch configurations allowed for your AWS account. The
	// default is 200 launch configurations per AWS Region.
	MaxNumberOfLaunchConfigurations *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeAccountLimitsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAccountLimits{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAccountLimits{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAccountLimits(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeAccountLimits",
	}
}
