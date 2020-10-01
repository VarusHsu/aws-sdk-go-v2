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

// Disables group metrics for the specified Auto Scaling group.
func (c *Client) DisableMetricsCollection(ctx context.Context, params *DisableMetricsCollectionInput, optFns ...func(*Options)) (*DisableMetricsCollectionOutput, error) {
	stack := middleware.NewStack("DisableMetricsCollection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDisableMetricsCollectionMiddlewares(stack)
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
	addOpDisableMetricsCollectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableMetricsCollection(options.Region), middleware.Before)
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
			OperationName: "DisableMetricsCollection",
			Err:           err,
		}
	}
	out := result.(*DisableMetricsCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableMetricsCollectionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Specifies one or more of the following metrics:
	//
	//     * GroupMinSize
	//
	//     *
	// GroupMaxSize
	//
	//     * GroupDesiredCapacity
	//
	//     * GroupInServiceInstances
	//
	//     *
	// GroupPendingInstances
	//
	//     * GroupStandbyInstances
	//
	//     *
	// GroupTerminatingInstances
	//
	//     * GroupTotalInstances
	//
	//     *
	// GroupInServiceCapacity
	//
	//     * GroupPendingCapacity
	//
	//     * GroupStandbyCapacity
	//
	//
	// * GroupTerminatingCapacity
	//
	//     * GroupTotalCapacity
	//
	// If you omit this
	// parameter, all metrics are disabled.
	Metrics []*string
}

type DisableMetricsCollectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDisableMetricsCollectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDisableMetricsCollection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableMetricsCollection{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableMetricsCollection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DisableMetricsCollection",
	}
}
