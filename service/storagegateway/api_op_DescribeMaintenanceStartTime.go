// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns your gateway's weekly maintenance start time including the day and time
// of the week. Note that values are in terms of the gateway's time zone.
func (c *Client) DescribeMaintenanceStartTime(ctx context.Context, params *DescribeMaintenanceStartTimeInput, optFns ...func(*Options)) (*DescribeMaintenanceStartTimeOutput, error) {
	stack := middleware.NewStack("DescribeMaintenanceStartTime", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeMaintenanceStartTimeMiddlewares(stack)
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
	addOpDescribeMaintenanceStartTimeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceStartTime(options.Region), middleware.Before)
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
			OperationName: "DescribeMaintenanceStartTime",
			Err:           err,
		}
	}
	out := result.(*DescribeMaintenanceStartTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway.
type DescribeMaintenanceStartTimeInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

// A JSON object containing the following fields:  <ul> <li> <p>
// <a>DescribeMaintenanceStartTimeOutput$DayOfMonth</a> </p> </li> <li> <p>
// <a>DescribeMaintenanceStartTimeOutput$DayOfWeek</a> </p> </li> <li> <p>
// <a>DescribeMaintenanceStartTimeOutput$HourOfDay</a> </p> </li> <li> <p>
// <a>DescribeMaintenanceStartTimeOutput$MinuteOfHour</a> </p> </li> <li> <p>
// <a>DescribeMaintenanceStartTimeOutput$Timezone</a> </p> </li> </ul>
type DescribeMaintenanceStartTimeOutput struct {

	// An ordinal number between 0 and 6 that represents the day of the week, where 0
	// represents Sunday and 6 represents Saturday. The day of week is in the time zone
	// of the gateway.
	DayOfWeek *int32

	// A value that indicates the time zone that is set for the gateway. The start time
	// and day of week specified should be in the time zone of the gateway.
	Timezone *string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// The hour component of the maintenance start time represented as hh, where hh is
	// the hour (0 to 23). The hour of the day is in the time zone of the gateway.
	HourOfDay *int32

	// The day of the month component of the maintenance start time represented as an
	// ordinal number from 1 to 28, where 1 represents the first day of the month and
	// 28 represents the last day of the month.
	DayOfMonth *int32

	// The minute component of the maintenance start time represented as mm, where mm
	// is the minute (0 to 59). The minute of the hour is in the time zone of the
	// gateway.
	MinuteOfHour *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeMaintenanceStartTimeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceStartTime{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceStartTime{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMaintenanceStartTime(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeMaintenanceStartTime",
	}
}
