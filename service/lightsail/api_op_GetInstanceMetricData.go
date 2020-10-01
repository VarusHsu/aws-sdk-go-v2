// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the data points for the specified Amazon Lightsail instance metric,
// given an instance name. Metrics report the utilization of your resources, and
// the error counts generated by them. Monitor and collect metric data regularly to
// maintain the reliability, availability, and performance of your resources.
func (c *Client) GetInstanceMetricData(ctx context.Context, params *GetInstanceMetricDataInput, optFns ...func(*Options)) (*GetInstanceMetricDataOutput, error) {
	stack := middleware.NewStack("GetInstanceMetricData", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetInstanceMetricDataMiddlewares(stack)
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
	addOpGetInstanceMetricDataValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstanceMetricData(options.Region), middleware.Before)
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
			OperationName: "GetInstanceMetricData",
			Err:           err,
		}
	}
	out := result.(*GetInstanceMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstanceMetricDataInput struct {

	// The start time of the time period.
	//
	// This member is required.
	StartTime *time.Time

	// The statistic for the metric. The following statistics are available:
	//
	//     *
	// Minimum - The lowest value observed during the specified period. Use this value
	// to determine low volumes of activity for your application.
	//
	//     * Maximum - The
	// highest value observed during the specified period. Use this value to determine
	// high volumes of activity for your application.
	//
	//     * Sum - All values submitted
	// for the matching metric added together. You can use this statistic to determine
	// the total volume of a metric.
	//
	//     * Average - The value of Sum / SampleCount
	// during the specified period. By comparing this statistic with the Minimum and
	// Maximum values, you can determine the full scope of a metric and how close the
	// average use is to the Minimum and Maximum values. This comparison helps you to
	// know when to increase or decrease your resources.
	//
	//     * SampleCount - The
	// count, or number, of data points used for the statistical calculation.
	//
	// This member is required.
	Statistics []types.MetricStatistic

	// The end time of the time period.
	//
	// This member is required.
	EndTime *time.Time

	// The unit for the metric data request. Valid units depend on the metric data
	// being requested. For the valid units to specify with each available metric, see
	// the metricName parameter.
	//
	// This member is required.
	Unit types.MetricUnit

	// The name of the instance for which you want to get metrics data.
	//
	// This member is required.
	InstanceName *string

	// The granularity, in seconds, of the returned data points. The StatusCheckFailed,
	// StatusCheckFailed_Instance, and StatusCheckFailed_System instance metric data is
	// available in 1-minute (60 seconds) granularity. All other instance metric data
	// is available in 5-minute (300 seconds) granularity.
	//
	// This member is required.
	Period *int32

	// The metric for which you want to return information. Valid instance metric names
	// are listed below, along with the most useful statistics to include in your
	// request, and the published unit value.
	//
	//     * BurstCapacityPercentage - The
	// percentage of CPU performance available for your instance to burst above its
	// baseline. Your instance continuously accrues and consumes burst capacity. Burst
	// capacity stops accruing when your instance's BurstCapacityPercentage reaches
	// 100%. For more information, see Viewing instance burst capacity in Amazon
	// Lightsail
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-viewing-instance-burst-capacity).
	// Statistics: The most useful statistics are Maximum and Average. Unit: The
	// published unit is Percent.
	//
	//     * BurstCapacityTime - The available amount of
	// time for your instance to burst at 100% CPU utilization. Your instance
	// continuously accrues and consumes burst capacity. Burst capacity time stops
	// accruing when your instance's BurstCapacityPercentage metric reaches 100%. Burst
	// capacity time is consumed at the full rate only when your instance operates at
	// 100% CPU utilization. For example, if your instance operates at 50% CPU
	// utilization in the burstable zone for a 5-minute period, then it consumes CPU
	// burst capacity minutes at a 50% rate in that period. Your instance consumed 2
	// minutes and 30 seconds of CPU burst capacity minutes in the 5-minute period. For
	// more information, see Viewing instance burst capacity in Amazon Lightsail
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-viewing-instance-burst-capacity).
	// Statistics: The most useful statistics are Maximum and Average. Unit: The
	// published unit is Seconds.
	//
	//     * CPUUtilization - The percentage of allocated
	// compute units that are currently in use on the instance. This metric identifies
	// the processing power to run the applications on the instance. Tools in your
	// operating system can show a lower percentage than Lightsail when the instance is
	// not allocated a full processor core. Statistics: The most useful statistics are
	// Maximum and Average. Unit: The published unit is Percent.
	//
	//     * NetworkIn - The
	// number of bytes received on all network interfaces by the instance. This metric
	// identifies the volume of incoming network traffic to the instance. The number
	// reported is the number of bytes received during the period. Because this metric
	// is reported in 5-minute intervals, divide the reported number by 300 to find
	// Bytes/second. Statistics: The most useful statistic is Sum. Unit: The published
	// unit is Bytes.
	//
	//     * NetworkOut - The number of bytes sent out on all network
	// interfaces by the instance. This metric identifies the volume of outgoing
	// network traffic from the instance. The number reported is the number of bytes
	// sent during the period. Because this metric is reported in 5-minute intervals,
	// divide the reported number by 300 to find Bytes/second. Statistics: The most
	// useful statistic is Sum. Unit: The published unit is Bytes.
	//
	//     *
	// StatusCheckFailed - Reports whether the instance passed or failed both the
	// instance status check and the system status check. This metric can be either 0
	// (passed) or 1 (failed). This metric data is available in 1-minute (60 seconds)
	// granularity. Statistics: The most useful statistic is Sum. Unit: The published
	// unit is Count.
	//
	//     * StatusCheckFailed_Instance - Reports whether the instance
	// passed or failed the instance status check. This metric can be either 0 (passed)
	// or 1 (failed). This metric data is available in 1-minute (60 seconds)
	// granularity. Statistics: The most useful statistic is Sum. Unit: The published
	// unit is Count.
	//
	//     * StatusCheckFailed_System - Reports whether the instance
	// passed or failed the system status check. This metric can be either 0 (passed)
	// or 1 (failed). This metric data is available in 1-minute (60 seconds)
	// granularity. Statistics: The most useful statistic is Sum. Unit: The published
	// unit is Count.
	//
	// This member is required.
	MetricName types.InstanceMetricName
}

type GetInstanceMetricDataOutput struct {

	// An array of objects that describe the metric data returned.
	MetricData []*types.MetricDatapoint

	// The name of the metric returned.
	MetricName types.InstanceMetricName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetInstanceMetricDataMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetInstanceMetricData{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInstanceMetricData{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetInstanceMetricData(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetInstanceMetricData",
	}
}
