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

// Describes the notification types that are supported by Amazon EC2 Auto Scaling.
func (c *Client) DescribeAutoScalingNotificationTypes(ctx context.Context, params *DescribeAutoScalingNotificationTypesInput, optFns ...func(*Options)) (*DescribeAutoScalingNotificationTypesOutput, error) {
	stack := middleware.NewStack("DescribeAutoScalingNotificationTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeAutoScalingNotificationTypesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutoScalingNotificationTypes(options.Region), middleware.Before)
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
			OperationName: "DescribeAutoScalingNotificationTypes",
			Err:           err,
		}
	}
	out := result.(*DescribeAutoScalingNotificationTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutoScalingNotificationTypesInput struct {
}

type DescribeAutoScalingNotificationTypesOutput struct {

	// The notification types.
	AutoScalingNotificationTypes []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeAutoScalingNotificationTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAutoScalingNotificationTypes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAutoScalingNotificationTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAutoScalingNotificationTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeAutoScalingNotificationTypes",
	}
}
