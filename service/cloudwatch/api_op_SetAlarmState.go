// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Temporarily sets the state of an alarm for testing purposes. When the updated
// state differs from the previous value, the action configured for the appropriate
// state is invoked. For example, if your alarm is configured to send an Amazon SNS
// message when an alarm is triggered, temporarily changing the alarm state to
// ALARM sends an SNS message. Metric alarms returns to their actual state quickly,
// often within seconds. Because the metric alarm state change happens quickly, it
// is typically only visible in the alarm's History tab in the Amazon CloudWatch
// console or through DescribeAlarmHistory
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmHistory.html).
// If you use SetAlarmState on a composite alarm, the composite alarm is not
// guaranteed to return to its actual state. It returns to its actual state only
// once any of its children alarms change state. It is also reevaluated if you
// update its configuration. If an alarm triggers EC2 Auto Scaling policies or
// application Auto Scaling policies, you must include information in the
// StateReasonData parameter to enable the policy to take the correct action.
func (c *Client) SetAlarmState(ctx context.Context, params *SetAlarmStateInput, optFns ...func(*Options)) (*SetAlarmStateOutput, error) {
	stack := middleware.NewStack("SetAlarmState", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetAlarmStateMiddlewares(stack)
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
	addOpSetAlarmStateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetAlarmState(options.Region), middleware.Before)
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
			OperationName: "SetAlarmState",
			Err:           err,
		}
	}
	out := result.(*SetAlarmStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetAlarmStateInput struct {

	// The reason that this alarm is set to this specific state, in text format.
	//
	// This member is required.
	StateReason *string

	// The reason that this alarm is set to this specific state, in JSON format. For
	// SNS or EC2 alarm actions, this is just informational. But for EC2 Auto Scaling
	// or application Auto Scaling alarm actions, the Auto Scaling policy uses the
	// information in this field to take the correct action.
	StateReasonData *string

	// The value of the state.
	//
	// This member is required.
	StateValue types.StateValue

	// The name of the alarm.
	//
	// This member is required.
	AlarmName *string
}

type SetAlarmStateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetAlarmStateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetAlarmState{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetAlarmState{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetAlarmState(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "SetAlarmState",
	}
}
